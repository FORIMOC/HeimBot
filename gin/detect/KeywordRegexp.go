package detect

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/model/ModelKeyword"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

// GetKeywordRegexp 对消息使用关键词词表进行正则表达式匹配
// => 匹配的关键词列表
func GetKeywordRegexp(DB *gorm.DB, data string) []string {
	msg := gjson.Get(data, "message").String()
	var keywordList []ModelKeyword.KeywordList
	var keywords []string
	var record []string

	DB.Find(&keywordList)
	for i := 0; i < len(keywordList); i++ {
		keywords = append(keywords, keywordList[i].Keyword)
	}

	// 对于每一个关键词进行正则表达式匹配
	for i := 0; i < len(keywords); i++ {
		re := regexp.MustCompile("(?i)" + keywords[i])
		if len(re.FindStringIndex(msg)) > 0 {
			record = append(record, keywords[i])
		}
	}
	return record
}

// GroupKeywordWarning 对基群进行报警
func GroupKeywordWarning(DB *gorm.DB, data string) {
	var group ModelGroup.Group
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	DB.Where("group_id = ?", groupID).First(&group)

	record := GetKeywordRegexp(DB, data)

	if len(record) > 0 {
		resp := "===捕获到关键词===\n" +
			"**[" + group.GroupName + "]-(" + groupID + ")**\n" +
			"用户" + username + "(" + userID + ") 的消息内容捕捉到 " + strings.Join(record, "|")
		common.BotDefaultResponse(resp)
	}
}
