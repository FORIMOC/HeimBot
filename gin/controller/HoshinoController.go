package controller

import (
	"forimoc.com/Hoshino/common"
	"forimoc.com/Hoshino/model"
	"forimoc.com/Hoshino/src"
	"forimoc.com/Hoshino/util"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func MainEntry(ctx *gin.Context) {
	DB := common.GetDB()
	rawData, _ := ioutil.ReadAll(ctx.Request.Body)
	data := string(rawData)
	postType := gjson.Get(data, "post_type").String()
	if postType != "message" {
		return
	}

	// 获取参数
	groupID := gjson.Get(data, "group_id").String()

	// 群聊存在控制
	var group model.Group
	err := DB.Where("group_id = ?", groupID).First(&group).Error
	if err != nil {
		groupName, member, maxMember := util.GetGroupInfo(groupID)
		newGroup := model.Group{
			GroupID:   groupID,
			GroupName: groupName,
			Status:    "1",
			SumMsg:    0,
			AvgMsg:    0,
			Member:    member,
			MaxMember: maxMember,
		}
		DB.Create(&newGroup)
	}

	// 状态控制(顺序不能改)
	switch group.Status {
	case "3": // 全局监听(记录型数据库支持)
		src.KeywordUpdate(ctx, DB, data)
		src.TimeUpdate(ctx, DB, data)

		src.UserUpdate(ctx, DB, data)
		src.GroupUpdate(ctx, DB, data)
		src.GraphUpdate(ctx, DB, data)
		src.PremessageUpdate(ctx, DB, data)
		break
	case "2": // 控局监听(缓存数据库支持)
		src.UserUpdate(ctx, DB, data)
		src.GroupUpdate(ctx, DB, data)
		src.GraphUpdate(ctx, DB, data)
		src.PremessageUpdate(ctx, DB, data)
		src.Warning(ctx, DB, data)
		break
	case "1": // 摘要监听(无数据库支持)
		src.Warning(ctx, DB, data)
		break
	default: // 无监听(特殊群聊不显示，比如情报部)
	case "0":
		return
	}
}
