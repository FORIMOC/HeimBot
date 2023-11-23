package record

import (
	"forimoc.com/Heimbot/detect"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/util"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"math"
	time2 "time"
)

// GroupGraphUpdate 社交网络数据更新
func GroupGraphUpdate(DB *gorm.DB, data string) {
	userID1 := gjson.Get(data, "user_id").String()
	username1 := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()

	var err error
	var value float64
	var graph, newGraph ModelGroup.GroupGraph
	var cacheMessages []ModelGroup.GroupCacheMessage
	var user ModelGroup.GroupUser

	// 查询缓存消息
	DB.Where("group_id = ?", groupID).Order("created_at ASC").Find(&cacheMessages)

	// 最大能产生影响的消息条数
	maxAffectNum := viper.GetInt("data_limit.groupGraphMaxAffectMessageLimit")
	if len(cacheMessages) <= maxAffectNum {
		maxAffectNum = len(cacheMessages)
	}

	for i := 0; i < maxAffectNum; i++ {
		graph.ID = 0
		if userID1 == cacheMessages[i].UserID {
			continue
		}
		if i < 2 {
			value = 1
		} else if i < 5 {
			value = 0.5
		} else {
			value = 0.2
		}
		err = DB.Where("group_id = ? AND user_id1 = ? AND user_id2 = ?", groupID, userID1, cacheMessages[i].UserID).First(&graph).Error
		if err != nil {
			err = DB.Where("group_id = ? AND user_id1 = ? AND user_id2 = ?", groupID, cacheMessages[i].UserID, userID1).First(&graph).Error
			if err != nil {
				newGraph = ModelGroup.GroupGraph{
					GroupID:   groupID,
					UserID1:   userID1,
					Username1: username1,
					UserID2:   cacheMessages[i].UserID,
					Username2: cacheMessages[i].Username,
					Value1:    value,
					Value2:    0,
				}
				DB.Create(&newGraph)
			} else {
				graph.Value2 += value
				graph.Value2 = math.Round(graph.Value2*10) / 10
				DB.Save(&graph)
			}
		} else {
			graph.Value1 += value
			graph.Value1 = math.Round(graph.Value1*10) / 10
			DB.Save(&graph)
		}
		err = DB.Where("group_id = ? AND user_id = ?", groupID, userID1).First(&user).Error
		if err == nil {
			user.Value += value
			user.Value = math.Round(user.Value*10) / 10
			DB.Save(&user)
		}
		//err = DB.Raw("SELECT * FROM group_users WHERE group_id = ? AND user_id = ?", groupID, cacheMessages[i].UserID).Scan(&user).Error
		//if err == nil {
		//	user.Value += value
		//	DB.Save(&user)
		//}
	}
	return
}

// GroupMessageUpdate 消息数据更新
func GroupMessageUpdate(DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	content := gjson.Get(data, "message").String()

	message := ModelGroup.GroupMessage{
		UserID:   userID,
		Username: username,
		GroupID:  groupID,
		Content:  content,
	}
	DB.Create(&message)
	return
}

// GroupCacheMessageUpdate 缓存消息更新
func GroupCacheMessageUpdate(DB *gorm.DB, data string) {
	// 获取参数
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	content := gjson.Get(data, "message").String()
	groupCacheMessageLimit := viper.GetInt64("data_limit.groupCacheMessageLimit")
	groupCacheMessageFlushLimit := viper.GetInt("data_limit.groupCacheMessageFlushLimit")
	var count int64
	// 如果数据条数大于 groupCacheMessageLimit 条, 则按时间排序删除 groupCacheMessageFlushLimit 条
	DB.Table("group_cache_messages").Where("group_id = ?", groupID).Count(&count)
	if count >= groupCacheMessageLimit {
		DB.Table("group_cache_messages").Where("group_id = ?", groupID).Order("created_at ASC").Limit(groupCacheMessageFlushLimit).Unscoped().Delete(ModelGroup.GroupCacheMessage{})
	}
	// 增 1
	newCacheMessage := ModelGroup.GroupCacheMessage{
		UserID:   userID,
		Username: username,
		GroupID:  groupID,
		Content:  content,
	}
	DB.Create(&newCacheMessage)
	return
}

// GroupKeywordUpdate 关键词词频记录更新
func GroupKeywordUpdate(DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	userID := gjson.Get(data, "user_id").String()
	var keyword ModelGroup.GroupKeyword
	var user ModelGroup.GroupUser
	record := detect.GetKeywordRegexp(DB, data)
	// 更新 GroupKeyword
	for i := 0; i < len(record); i++ {
		if err := DB.Where("group_id = ? AND keyword = ?", groupID, record[i]).First(&keyword).Error; err != nil {
			newKeyword := ModelGroup.GroupKeyword{
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
	// 更新 GroupUser 中的 KeywordNum
	if err := DB.Where("group_id = ? AND user_id = ?", groupID, userID).First(&user).Error; err == nil {
		user.KeywordNum += len(record)
		DB.Save(&user)
	}
	return
}

// GroupTimeUpdate 发言时间分布数据更新
func GroupTimeUpdate(DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	groupTimeLimit := viper.GetInt("data_limit.groupTimeLimit")
	var time ModelGroup.GroupTime
	var times []ModelGroup.GroupTime
	var count int64

	hour := time2.Now().Hour()
	month := time2.Now().Month()
	day := time2.Now().Day()

	err := DB.Where("group_id = ? AND month = ? AND day = ?", groupID, month, day).First(&time).Error
	if err != nil {
		newTime := ModelGroup.GroupTime{
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

	DB.Table("group_times").Where("group_id = ?", groupID).Count(&count)
	if count > int64(groupTimeLimit) {
		DB.Where("group_id = ?", groupID).Order("created_at desc").Find(&times)
		DB.Unscoped().Delete(&times[groupTimeLimit])
	}
	return
}

// GroupUserTimeUpdate 单用户发言时间分布数据更新
func GroupUserTimeUpdate(DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	userID := gjson.Get(data, "user_id").String()
	groupUserTimeLimit := viper.GetInt("data_limit.groupUserTimeLimit")
	var time ModelGroup.GroupUserTime
	var times []ModelGroup.GroupUserTime
	var count int64

	hour := time2.Now().Hour()
	month := time2.Now().Month()
	day := time2.Now().Day()

	err := DB.Where("group_id = ? AND user_id = ? AND month = ? AND day = ?", groupID, userID, month, day).First(&time).Error
	if err != nil {
		newTime := ModelGroup.GroupUserTime{
			GroupID: groupID,
			UserID:  userID,
			Month:   int(month),
			Day:     day,
		}
		DB.Create(&newTime)
	}
	DB.Where("group_id = ? AND user_id = ? AND month = ? AND day = ?", groupID, userID, month, day).First(&time)
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

	DB.Where("group_id = ? AND user_id = ?", groupID, userID).Count(&count)
	if count > int64(groupUserTimeLimit) {
		DB.Where("group_id = ? AND user_id = ?", groupID, userID).Order("created_at desc").Find(&times)
		DB.Unscoped().Delete(&times[groupUserTimeLimit])
	}
	return
}

// GroupUserUpdate 用户数据更新
func GroupUserUpdate(DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	var user ModelGroup.GroupUser
	if err := DB.Where("group_id = ?", groupID).Where("user_id = ?", userID).First(&user).Error; err != nil {
		newUser := ModelGroup.GroupUser{
			UserID:     userID,
			Username:   username,
			GroupID:    groupID,
			Freq:       1,
			Value:      0,
			KeywordNum: 0,
		}
		DB.Create(&newUser)
		return
	}
	user.Freq += 1
	user.Username = username
	DB.Save(&user)
	return
}

// GroupUpdate 群聊自身基本信息更新
func GroupUpdate(DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	msg := gjson.Get(data, "message").String()
	var group ModelGroup.Group
	DB.Where("group_id = ?", groupID).First(&group)
	group.SumMsgNum += 1
	group.AvgMsgLen = (group.AvgMsgLen*float64(group.SumMsgNum-1) + float64(len(msg))) / float64(group.SumMsgNum)
	groupName, member, maxMember := util.GetGroupInfo(groupID)
	if len(groupName) > 0 {
		group.GroupName = groupName
	}
	if member > 0 {
		group.MemberNum = member
	}
	if maxMember > 0 {
		group.MaxMemberNum = maxMember
	}
	DB.Save(&group)
	return
}
