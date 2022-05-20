package src

import (
	"fmt"
	"forimoc.com/Hoshino/model"
	"forimoc.com/Hoshino/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tidwall/gjson"
	"regexp"
	"strings"
)

func Warning(ctx *gin.Context, DB *gorm.DB, data string) []string {
	var group model.Group
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	DB.Where("group_id = ?", groupID).First(&group)
	msg := gjson.Get(data, "message").String()
	var keywordList []model.KeywordList
	var keywords []string
	DB.Find(&keywordList)
	for i := 0; i < len(keywordList); i++ {
		keywords = append(keywords, keywordList[i].Keyword)
	}
	var record []string
	for i := 0; i < len(keywords); i++ {
		re := regexp.MustCompile(keywords[i])
		fmt.Println(re.FindStringIndex(msg))
		if len(re.FindStringIndex(msg)) > 0 {
			record = append(record, keywords[i])
		}
	}
	if len(record) > 0 {
		resp := "**[" + group.GroupName + "]-(" + groupID + ")**\n" +
			strings.Join(record, "|") + "\n" +
			username + "(" + userID + ") > " + msg
		util.Send("Komari", "781815666", resp)
	}
	return record
}
