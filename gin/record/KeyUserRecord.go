package record

import (
	"forimoc.com/Heimbot/detect"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/model/ModelKeyUser"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	time2 "time"
)

// KeyUserUpdate 关键用户自身基本信息更新
func KeyUserUpdate(DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	msg := gjson.Get(data, "message").String()
	var keyUser ModelKeyUser.KeyUser
	DB.Where("user_id = ?", userID).First(&keyUser)
	keyUser.SumMsgNum += 1
	keyUser.AvgMsgLen = (keyUser.AvgMsgLen*float64(keyUser.SumMsgNum-1) + float64(len(msg))) / float64(keyUser.SumMsgNum)
	DB.Save(&keyUser)
	return
}

// KeyUserGroupUpdate 发言群聊数据更新
func KeyUserGroupUpdate(DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	groupID := gjson.Get(data, "group_id").String()
	var group ModelGroup.Group
	DB.Where("group_id = ?", groupID).First(&group)
	groupName := group.GroupName

	var keyUserGroup ModelKeyUser.KeyUserGroup
	err := DB.Where("user_id = ?", userID).Where("group_id = ?", groupID).First(&keyUserGroup).Error
	if err != nil {
		newKeyUserGroup := ModelKeyUser.KeyUserGroup{
			UserID:    userID,
			GroupID:   groupID,
			GroupName: groupName,
			Freq:      1,
		}
		DB.Create(&newKeyUserGroup)
		return
	}
	keyUserGroup.Freq += 1
	keyUserGroup.GroupName = groupName
	DB.Save(&keyUserGroup)
	return
}

// KeyUserTimeUpdate 发言时间分布数据更新
func KeyUserTimeUpdate(DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	groupTimeLimit := viper.GetInt("data_limit.groupTimeLimit")
	var time ModelKeyUser.KeyUserTime
	var times []ModelKeyUser.KeyUserTime
	var count int64

	// 判断是否为关键用户
	if err := DB.Table("key_users").Where("user_id = ?", userID).Error; err != nil {
		return
	}

	// 获取时间参数
	hour := time2.Now().Hour()
	month := time2.Now().Month()
	day := time2.Now().Day()

	// 没有此天时间分布记录则新建
	if err := DB.Where("user_id = ? AND month = ? AND day = ?", userID, month, day).First(&time).Error; err != nil {
		newTime := ModelKeyUser.KeyUserTime{
			UserID: userID,
			Month:  int(month),
			Day:    day,
		}
		DB.Create(&newTime)
	}

	// 记录更新
	DB.Where("user_id = ? AND month = ? AND day = ?", userID, month, day).First(&time)
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

	// 限制记录条数在最近 5 天
	DB.Table("key_user_times").Where("user_id = ?", userID).Count(&count)
	if count > int64(groupTimeLimit) {
		DB.Where("user_id = ?", userID).Order("created_at desc").Find(&times)
		DB.Unscoped().Delete(&times[groupTimeLimit])
	}
	return
}

// KeyUserKeywordUpdate 关键词记录数据更新
func KeyUserKeywordUpdate(DB *gorm.DB, data string) {
	userID := gjson.Get(data, "user_id").String()
	var keyword ModelKeyUser.KeyUserKeyword
	record := detect.GetKeywordRegexp(DB, data)
	for i := 0; i < len(record); i++ {
		err := DB.Where("user_id = ?", userID).Where("keyword = ?", record[i]).First(&keyword).Error
		if err != nil {
			newKeyword := ModelKeyUser.KeyUserKeyword{
				UserID:  userID,
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

// KeyUserCacheMessageUpdate 缓存消息更新
func KeyUserCacheMessageUpdate(DB *gorm.DB, data string) {
	// 获取参数
	userID := gjson.Get(data, "user_id").String()
	username := gjson.Get(data, "sender.nickname").String()
	groupID := gjson.Get(data, "group_id").String()
	content := gjson.Get(data, "message").String()
	groupCacheMessageLimit := viper.GetInt64("data_limit.groupCacheMessageLimit")
	groupCacheMessageFlushLimit := viper.GetInt("data_limit.groupCacheMessageFlushLimit")
	var count int64
	var group ModelGroup.Group
	DB.Where("group_id = ?", groupID).First(&group)
	groupName := group.GroupName
	// 如果数据条数大于 groupCacheMessageLimit 条, 则按时间排序删除 groupCacheMessageFlushLimit 条
	DB.Table("key_user_cache_messages").Where("user_id = ?", userID).Count(&count)
	if count >= groupCacheMessageLimit {
		DB.Table("key_user_cache_messages").Where("user_id = ?", userID).Order("created_at ASC").Limit(groupCacheMessageFlushLimit).Unscoped().Delete(ModelKeyUser.KeyUserCacheMessage{})
	}
	// 增 1
	newCacheMessage := ModelKeyUser.KeyUserCacheMessage{
		UserID:    userID,
		Username:  username,
		GroupID:   groupID,
		GroupName: groupName,
		Content:   content,
	}
	DB.Create(&newCacheMessage)
	return
}
