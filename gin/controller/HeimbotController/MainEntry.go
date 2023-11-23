package HeimbotController

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/detect"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/model/ModelKeyUser"
	"forimoc.com/Heimbot/record"
	"forimoc.com/Heimbot/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"io"
)

func MainEntry(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取 go-cqhttp POST 转发信息
	rawData, _ := io.ReadAll(ctx.Request.Body)
	data := string(rawData)
	postType := gjson.Get(data, "post_type").String()
	if postType != "message" {
		return
	}

	// 获取参数
	groupID := gjson.Get(data, "group_id").String()
	userID := gjson.Get(data, "user_id").String()
	ownerID := viper.GetString("owner.ID")

	// 防止自检
	botSelf := viper.GetString("bot.selfID")
	if userID == botSelf {
		return
	}

	// 白名单过滤
	if util.IsGroupInConfigItem(groupID, "bot.whitelistGroupID") {
		return
	}

	// 群聊存在控制
	var group ModelGroup.Group
	defaultListenMode := viper.GetString("bot.defaultListenMode")
	if err := DB.Where("group_id = ?", groupID).First(&group).Error; err != nil {
		groupName, member, maxMember := util.GetGroupInfo(groupID)
		//if len(groupName) > 0 && len(groupID) > 0 {
		newGroup := ModelGroup.Group{
			GroupID:      groupID,
			GroupName:    groupName,
			Status:       defaultListenMode,
			SumMsgNum:    0,
			AvgMsgLen:    0,
			MemberNum:    member,
			MaxMemberNum: maxMember,
		}
		DB.Create(&newGroup)
		//}
	}

	// 群聊Cmd控制Heimbot
	isInteractable := viper.GetBool("bot.isInteractable")
	if isInteractable && util.IsGroupInConfigItem(groupID, "bot.authGroupID") {
		if userID == ownerID { // adminLevel == 0
			// Owner独有命令
			ownerCmd(DB, data)
		}
		cmd(DB, data)
	}

	// 群聊状态控制
	switch group.Status {
	case "3": // 全局监听
		go record.GroupUserUpdate(DB, data)
		go record.GroupUpdate(DB, data)

		go record.GroupTimeUpdate(DB, data)
		go record.GroupUserTimeUpdate(DB, data)
		go record.GroupCacheMessageUpdate(DB, data)
		go record.GroupMessageUpdate(DB, data)
		go record.GroupGraphUpdate(DB, data)
		go record.GroupKeywordUpdate(DB, data)

		go detect.GroupKeywordWarning(DB, data)
		go detect.GetFraudDetect(DB, data)
		go detect.GroupFraudGroupWarning(DB, data)
		go detect.GroupTimeSeriesWarning(DB, data)
		break
	case "2": // 控局监听
		go record.GroupUserUpdate(DB, data)
		go record.GroupUpdate(DB, data)

		go record.GroupTimeUpdate(DB, data)
		go record.GroupUserTimeUpdate(DB, data)
		go record.GroupCacheMessageUpdate(DB, data)
		go record.GroupGraphUpdate(DB, data)
		go record.GroupKeywordUpdate(DB, data)

		go detect.GroupKeywordWarning(DB, data)
		go detect.GetFraudDetect(DB, data)
		go detect.GroupFraudGroupWarning(DB, data)
		go detect.GroupTimeSeriesWarning(DB, data)
		break
	case "1": // 摘要监听
		go record.GroupUpdate(DB, data)

		go detect.GroupKeywordWarning(DB, data)
		go detect.GetFraudDetect(DB, data)
		go detect.GroupFraudGroupWarning(DB, data)
		go detect.GroupTimeSeriesWarning(DB, data)
		break
	default: // 无监听(不会发送给前端)
	case "0":
		break
	case "-1": // 仅调试
	}

	/* 关键用户 */
	// 是否属于关键用户
	var keyUser ModelKeyUser.KeyUser
	if err := DB.Table("key_users").Where("user_id = ?", userID).First(&keyUser).Error; err != nil {
		return
	}
	switch keyUser.Status {
	case "1": // 监听
		go record.KeyUserUpdate(DB, data)
		go record.KeyUserTimeUpdate(DB, data)
		go record.KeyUserCacheMessageUpdate(DB, data)
		go record.KeyUserGroupUpdate(DB, data)
		go record.KeyUserKeywordUpdate(DB, data)
	default: // 暂不监听
	case "0":
	}
}
