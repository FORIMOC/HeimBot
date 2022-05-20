package src

import (
	"forimoc.com/Hoshino/model"
	"forimoc.com/Hoshino/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tidwall/gjson"
	time2 "time"
)

func GraphUpdate(ctx *gin.Context, DB *gorm.DB, data string) {
	userID1 := gjson.Get(data, "user_id").String()
	username1 := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	var err error
	var value float64
	var graph, newGraph model.Graph
	var premessages []model.Premessage
	var user model.User
	DB.Where("group_id = ?", groupID).Order("mark ASC").Find(&premessages)
	for i := 0; i < len(premessages); i++ {
		if userID1 == premessages[i].UserID {
			continue
		}
		if i < 2 {
			value = 1
		} else if i < 5 {
			value = 0.5
		} else {
			value = 0.2
		}
		err = DB.Raw("SELECT * FROM graphs WHERE group_id = ? AND user_id1 = ? AND user_id2 = ?", groupID, userID1, premessages[i].UserID).Scan(&graph).Error
		if err != nil {
			err = DB.Raw("SELECT * FROM graphs WHERE group_id = ? AND user_id1 = ? AND user_id2 = ?", groupID, premessages[i].UserID, userID1).Scan(&graph).Error
			if err != nil {
				newGraph = model.Graph{
					GroupID:   groupID,
					UserID1:   userID1,
					Username1: username1,
					UserID2:   premessages[i].UserID,
					Username2: premessages[i].Username,
					Value1:    value,
					Value2:    0,
				}
				DB.Create(&newGraph)
			} else {
				graph.Value2 += value
				DB.Save(&graph)
			}
		} else {
			graph.Value1 += value
			DB.Save(&graph)
		}
		err = DB.Raw("SELECT * FROM users WHERE group_id = ? AND user_id = ?", groupID, userID1).Scan(&user).Error
		if err == nil {
			user.Value += value
			DB.Save(&user)
		}
		err = DB.Raw("SELECT * FROM users WHERE group_id = ? AND user_id = ?", groupID, premessages[i].UserID).Scan(&user).Error
		if err == nil {
			user.Value += value
			DB.Save(&user)
		}
	}
	return
}

func MessageUpdate(ctx *gin.Context, DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	content := gjson.Get(data, "message").String()

	message := model.Message{
		UserID:   userID,
		Username: username,
		GroupID:  groupID,
		Content:  content,
	}
	DB.Create(&message)
	return
}

func PremessageUpdate(ctx *gin.Context, DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	content := gjson.Get(data, "message").String()
	// 删 10
	DB.Where("group_id = ?", groupID).Where("mark = ?", 10).Unscoped().Delete(model.Premessage{})
	// 1-9 -> 2-10
	var premessages []model.Premessage
	DB.Where("group_id = ?", groupID).Order("mark ASC").Find(&premessages)
	for i := len(premessages) - 1; i >= 0; i-- {
		premessages[i].Mark = i + 2
		DB.Save(&premessages[i])
	}
	// 增 1
	newPremessage := model.Premessage{
		UserID:   userID,
		Username: username,
		GroupID:  groupID,
		Content:  content,
		Mark:     1,
	}
	DB.Create(&newPremessage)
	return
}

func KeywordUpdate(ctx *gin.Context, DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	var keyword model.Keyword
	record := Warning(ctx, DB, data)
	for i := 0; i < len(record); i++ {
		err := DB.Where("group_id = ?", groupID).Where("keyword = ?", record[i]).First(&keyword).Error
		if err != nil {
			newKeyword := model.Keyword{
				GroupID: groupID,
				Keyword: record[i],
				Freq:    1,
			}
			DB.Create(&newKeyword)
		} else {
			keyword.Freq += 1
			DB.Save(&keyword)
		}
	}
	return
}

func TimeUpdate(ctx *gin.Context, DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	var time model.Time
	var count int

	hour := time2.Now().Hour()
	month := time2.Now().Month()
	day := time2.Now().Day()

	err := DB.Where("group_id = ? AND month = ? AND day = ?", groupID, month, day).First(&time).Error
	if err != nil {
		newTime := model.Time{
			GroupID: groupID,
			Month:   int(month),
			Day:     day,
		}
		DB.Create(&newTime)
	}
	DB.Where("group_id = ? AND month = ? AND day = ?", groupID, month, day).First(&time)
	if hour < 2 {
		time.Hour1 += 1
	} else if hour < 4 {
		time.Hour2 += 1
	} else if hour < 6 {
		time.Hour3 += 1
	} else if hour < 8 {
		time.Hour4 += 1
	} else if hour < 10 {
		time.Hour5 += 1
	} else if hour < 12 {
		time.Hour6 += 1
	} else if hour < 14 {
		time.Hour7 += 1
	} else if hour < 16 {
		time.Hour8 += 1
	} else if hour < 18 {
		time.Hour9 += 1
	} else if hour < 20 {
		time.Hour10 += 1
	} else if hour < 22 {
		time.Hour11 += 1
	} else {
		time.Hour12 += 1
	}
	DB.Save(&time)

	DB.Table("times").Where("group_id = ?", groupID).Count(&count)
	if count > 3 {
		DB.Where("group_id = ?", groupID).First(&time)
		DB.Where("group_id = ?", groupID).Unscoped().Delete(&time)
	}
	return
}

func UserUpdate(ctx *gin.Context, DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	var user model.User
	err := DB.Where("group_id = ?", groupID).Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		newUser := model.User{
			UserID:   userID,
			Username: username,
			GroupID:  groupID,
			Freq:     1,
			Value:    0,
		}
		DB.Create(&newUser)
		return
	}
	user.Freq += 1
	user.Username = username
	DB.Save(&user)
	return
}

func GroupUpdate(ctx *gin.Context, DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	msg := gjson.Get(data, "message").String()
	var group model.Group
	DB.Where("group_id = ?", groupID).First(&group)
	group.SumMsg += 1
	group.AvgMsg = (group.AvgMsg*float64(group.SumMsg-1) + float64(len(msg))) / float64(group.SumMsg)
	groupName, member, maxMember := util.GetGroupInfo(groupID)
	if len(groupName) > 0 {
		group.GroupName = groupName
	}
	if member != 0 {
		group.Member = member
	}
	if maxMember != 0 {
		group.MaxMember = maxMember
	}
	DB.Save(&group)
	return
}
