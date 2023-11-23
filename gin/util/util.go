package util

import (
	"forimoc.com/Heimbot/model/ModelAdmin"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"os/exec"
	"strconv"
	"strings"
)

// IsGroupInConfigItem 群聊是否在配置文件中的某个群聊集合中
// 目标群聊, 配置项路径 => 是/否
func IsGroupInConfigItem(groupID string, configItem string) bool {
	configString := viper.GetString(configItem)
	configGroups := strings.Split(configString, ",")
	for i := 0; i < len(configGroups); i++ {
		if groupID == configGroups[i] {
			return true
		}
	}
	return false
}

// IsHighLevelAdmin 判断是否为高等级管理员(0, 1)
// 数据库实例, 用户ID => 是/否
func IsHighLevelAdmin(DB *gorm.DB, userID string) bool {
	var admin ModelAdmin.Admin
	if err := DB.Where("user_id = ?", userID).First(&admin).Error; err == nil {
		if admin.Level < 2 {
			return true
		}
	}
	return false
}

// GetGroupInfo 获取群聊基本信息
// 群ID => 群名称, 群成员数, 群最大成员数
func GetGroupInfo(groupID string) (string, int, int) {
	botPort := viper.GetString("bot.port")
	u := "http://127.0.0.1:" + botPort + "/get_group_info?group_id=" + groupID
	json, _ := exec.Command("curl", u).CombinedOutput()
	groupName := gjson.Get(string(json), "data.group_name").String()
	member := gjson.Get(string(json), "data.member_count").Int()
	maxMember := gjson.Get(string(json), "data.max_member_count").Int()
	return groupName, int(member), int(maxMember)
}

// GetGroupList 获取群聊ID列表
// => 群聊ID字符串数组
func GetGroupList() []string {
	var groupList []string
	botPort := viper.GetString("bot.port")
	u := "http://127.0.0.1:" + botPort + "/get_group_list"
	json, _ := exec.Command("curl", u).CombinedOutput()
	for i := 0; i < int(gjson.Get(string(json), "data.#").Int()); i++ {
		groupList = append(groupList, gjson.Get(string(json), "data."+strconv.Itoa(i)+".group_id").String())
	}
	return groupList
}
