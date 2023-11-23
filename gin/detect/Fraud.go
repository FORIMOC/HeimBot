package detect

import (
	"encoding/json"
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelFraudDetect"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/model/ModelKeyUser"
	"forimoc.com/Heimbot/util"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

// GetFraudDetect 将消息记录并定时交由模型处理, 将其中识别为诈骗的消息存入 FDMessage
func GetFraudDetect(DB *gorm.DB, data string) {
	// 获取参数
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	content := gjson.Get(data, "message").String()
	dataFlushLimit := viper.GetInt64("api.fraudDetect.dataFlushLimit")
	host := viper.GetString("api.fraudDetect.host")
	port := viper.GetString("api.fraudDetect.port")
	route := viper.GetString("api.fraudDetect.route")
	url := "http://" + host + ":" + port + route

	var count int64
	fraudCacheMessages := make([]ModelFraudDetect.FDCacheMessage, 0)
	tobeWarning := make([]ModelFraudDetect.FDMessage, 0)
	texts := make([]string, 0)

	var result util.FDModelResult

	// 增 1
	newCacheMessage := ModelGroup.GroupCacheMessage{
		UserID:   userID,
		Username: username,
		GroupID:  groupID,
		Content:  content,
	}
	DB.Table("fd_cache_messages").Create(&newCacheMessage)

	// 如果数据条数大于等于 dataFlushLimit 条, 则发由模型处理, 并清空缓存
	DB.Table("fd_cache_messages").Count(&count)
	if count >= dataFlushLimit {
		// 准备要发送的json数据
		DB.Table("fd_cache_messages").Order("created_at asc").Find(&fraudCacheMessages)
		for i := 0; i < int(dataFlushLimit); i++ {
			texts = append(texts, fraudCacheMessages[i].Content)
		}
		jsonData := map[string]interface{}{
			"texts": texts,
		}
		// 发送请求给模型
		resp := util.SendJSON(url, jsonData)
		json.Unmarshal(resp, &result)
		// 确认为诈骗消息的数据存入 FDMessage, 和 tobeWarning
		for i := 0; i < int(dataFlushLimit); i++ {
			if result.Result[i] == 1 {
				// 存入 fraud_messages
				newFraudMessage := ModelFraudDetect.FDMessage{
					GroupID:  fraudCacheMessages[i].GroupID,
					UserID:   fraudCacheMessages[i].UserID,
					Username: fraudCacheMessages[i].Username,
					Content:  fraudCacheMessages[i].Content,
				}
				DB.Table("fd_messages").Create(&newFraudMessage)
				// 匹配每条涉诈消息中的群聊ID并存入 fraud_groups
				re := regexp.MustCompile("\\d{8,10}")
				matches := re.FindAllString(fraudCacheMessages[i].Content, -1)
				for _, match := range matches {
					var fraudGroup ModelFraudDetect.FDGroup
					if err := DB.Where("group_id = ? AND sender_id = ?", match, fraudCacheMessages[i].UserID).First(&fraudGroup).Error; err != nil {
						newFraudGroup := ModelFraudDetect.FDGroup{
							GroupID:  match,
							SenderID: fraudCacheMessages[i].UserID,
						}
						DB.Create(&newFraudGroup)
					}
				}
				// 发送涉诈信息的用户存入 key_user
				var newKeyUser ModelKeyUser.KeyUser
				if err := DB.Table("key_users").Where("user_id = ?", fraudCacheMessages[i].UserID).First(&newKeyUser).Error; err != nil {
					newKeyUser = ModelKeyUser.KeyUser{
						UserID:    fraudCacheMessages[i].UserID,
						Username:  fraudCacheMessages[i].Username,
						Remark:    "",
						SumMsgNum: 1,
						AvgMsgLen: float64(len(fraudCacheMessages[i].Content)),
						Status:    "1",
					}
					DB.Create(&newKeyUser)
				}
				// 需要报警的内容
				tobeWarning = append(tobeWarning, newFraudMessage)
			}
		}
		// 清空缓存
		DB.Exec("DELETE FROM fd_cache_messages")
		if len(tobeWarning) > 0 {
			common.BotDefaultResponse("===模型检测到涉诈消息===\n详细内容于Heimboard查看")
		}
	}
}

// GetFraudGroupDetect 将消息记录进行QQ群黑名单筛选, 并返回结果
func GetFraudGroupDetect(DB *gorm.DB, data string) []string {
	// 获取参数
	msg := gjson.Get(data, "message").String()
	groupID := gjson.Get(data, "group_id").String()
	userID := gjson.Get(data, "user_id").String()
	var tobeWarning []string

	// 捕获群号
	re := regexp.MustCompile("\\d{8,10}")
	matches := re.FindAllString(msg, -1)
	for _, match := range matches {
		var fraudGroup ModelFraudDetect.FDGroup
		if err := DB.Where("group_id = ?", match).First(&fraudGroup).Error; err == nil {
			// 添加到返回列表
			tobeWarning = append(tobeWarning, fraudGroup.GroupID)
			// 如果发送人不一样, 则添加到 FDGroup 里
			var fraudGroup2 ModelFraudDetect.FDGroup
			if err2 := DB.Where("group_id = ? AND sender_id", groupID, userID).First(&fraudGroup2).Error; err2 != nil {
				newFraudGroup := ModelFraudDetect.FDGroup{
					GroupID:  match,
					SenderID: userID,
				}
				DB.Create(&newFraudGroup)
			}
		}
	}
	return tobeWarning
}

// GroupFraudGroupWarning 对基群就QQ群黑名单进行报警
func GroupFraudGroupWarning(DB *gorm.DB, data string) {
	var group ModelGroup.Group
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	DB.Where("group_id = ?", groupID).First(&group)

	tobeWarning := GetFraudGroupDetect(DB, data)

	if len(tobeWarning) > 0 {
		resp := "===检测到恶意引流QQ群===\n" +
			"**[" + group.GroupName + "]-(" + groupID + ")**\n" +
			"用户" + username + "(" + userID + ") 引流 群聊 " + strings.Join(tobeWarning, "|")
		common.BotDefaultResponse(resp)
	}
}
