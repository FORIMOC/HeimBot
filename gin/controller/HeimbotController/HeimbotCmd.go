package HeimbotController

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelAdmin"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"regexp"
)

func ownerCmd(DB *gorm.DB, data string) {
	msg := gjson.Get(data, "message").String()

	// 代理发送
	re := regexp.MustCompile(`^send (\d{4}) (\d+) (.+)$`)
	res := re.FindAllStringSubmatch(msg, -1)
	if res != nil {
		common.BotResponse(res[0][1], res[0][2], res[0][3])
	}
}

func cmd(DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	groupID := gjson.Get(data, "group_id").String()
	msg := gjson.Get(data, "message").String()
	botPort := viper.GetString("bot.port")
	// 激活
	if msg == "activate" {
		var admin ModelAdmin.Admin
		if err := DB.Where("user_id = ?", userID).First(&admin).Error; err != nil {
			common.BotResponse(botPort, groupID, "权限不足, 许可拒绝")
			return
		}
		if admin.Level == 1 {
			common.BotResponse(botPort, groupID, "已经激活, 许可拒绝")
			return
		}
		if admin.Level == 0 {
			common.BotResponse(botPort, groupID, "您已是最高权限, 无需激活")
			return
		}
		admin.Level = 1
		DB.Save(&admin)
		common.BotResponse(botPort, groupID, "激活成功")
	}
}
