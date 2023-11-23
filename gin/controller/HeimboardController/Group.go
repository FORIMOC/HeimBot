package HeimboardController

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelGroup"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// UpdateGroupStatus 更新群聊状态控制器
// group_id(群ID), status(要更改为的状态) =>
func UpdateGroupStatus(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("group_id")
	status := ctx.PostForm("status")
	if len(status) == 0 || len(groupID) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "参数不足")
		return
	}
	var group ModelGroup.Group
	err := DB.Where("group_id = ?", groupID).First(&group).Error
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在该群聊")
		return
	}
	group.Status = status
	DB.Save(&group)
	common.Response(ctx, http.StatusOK, 200, nil, "状态更新成功")
}

// GetGroupList 获取群聊列表控制器
// => groups([]ModelGroup.Group)
func GetGroupList(ctx *gin.Context) {
	DB := common.GetDB()
	var groups []ModelGroup.Group
	DB.Where("status != 0").Find(&groups)
	common.Response(ctx, http.StatusOK, 200, gin.H{"groups": groups}, "群聊列表查询成功")
}

// GetGroupInfo 获取群聊基本信息控制器
// group_id(群ID) => group(ModelGroup.Group)
func GetGroupInfo(ctx *gin.Context) {
	DB := common.GetDB()
	var group ModelGroup.Group
	// 获取数据
	groupID := ctx.PostForm("group_id")
	if err := DB.Where("group_id = ?", groupID).First(&group).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"group": group}, "群聊数据获取成功")
}

// GetGroupGraph 获取群聊社交网络控制器
// group_id(群ID) => (指定group_id的[]ModelGroup.GroupGraph)
func GetGroupGraph(ctx *gin.Context) {
	DB := common.GetDB()
	// 防止 graphs 为 nil
	graphs := make([]ModelGroup.GroupGraph, 0)
	var users []ModelGroup.GroupUser
	// 获取参数
	groupID := ctx.PostForm("group_id")
	groupGraphPointLimit := viper.GetInt("data_limit.groupGraphPointLimit")

	// 获取 groupGraphPointLimit 个最大的点
	if err := DB.Where("group_id = ?", groupID).Order("value DESC").Limit(groupGraphPointLimit).Find(&users).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}

	// 根据 groupGraphPointLimit 个最大的点获取对应的图关系
	for _, user := range users {
		var tempGraphs []ModelGroup.GroupGraph
		if err := DB.Where("group_id = ? AND user_id1 = ?", groupID, user.UserID).Find(&tempGraphs).Error; err == nil {
			for _, graph := range tempGraphs {
				graphs = append(graphs, graph)
			}
		}
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"graphs": graphs, "users": users}, "群聊数据获取成功")
}

// GetUserList 获取群聊用户信息控制器
// group_id(群ID) => (指定group_id的[]ModelGroup.GroupUser)
func GetUserList(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("group_id")
	groupUserLimit := viper.GetInt("data_limit.groupUserLimit")

	var users []ModelGroup.GroupUser
	var count int64
	if err := DB.Where("group_id = ?", groupID).Order("freq DESC").Limit(groupUserLimit).Find(&users).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	DB.Table("group_users").Where("group_id = ?", groupID).Count(&count)
	common.Response(ctx, http.StatusOK, 200, gin.H{"users": users, "count": count}, "用户数据获取成功")
}

// GetGroupCacheMessages 获取群聊用户信息控制器
// group_id(群ID) => (指定group_id的[]ModelGroup.GroupCacheMessage)
func GetGroupCacheMessages(ctx *gin.Context) {
	DB := common.GetDB()
	groupID := ctx.PostForm("group_id")
	var cacheMessages []ModelGroup.GroupCacheMessage
	if err := DB.Where("group_id = ?", groupID).Find(&cacheMessages).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"cache_messages": cacheMessages}, "群聊数据获取成功")
}

// GetGroupKeywords 获取群聊关键词信息控制器
// group_id(群ID) => (指定group_id的[]ModelGroup.GroupKeyword)
func GetGroupKeywords(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("group_id")
	limit := viper.GetInt("data_limit.groupKeywordLimit")

	var keywords []ModelGroup.GroupKeyword
	var count int64
	if err := DB.Where("group_id = ?", groupID).Order("freq DESC").Limit(limit).Find(&keywords).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	DB.Table("group_keywords").Where("group_id = ?", groupID).Count(&count)
	common.Response(ctx, http.StatusOK, 200, gin.H{"keywords": keywords, "count": count}, "群聊关键词数据获取成功")
}

// GetGroupTimes 获取群聊消息时间分布控制器
// group_id(群ID) => (指定group_id的[]ModelGroup.GroupTime)
func GetGroupTimes(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("group_id")
	var times []ModelGroup.GroupTime
	err := DB.Where("group_id = ?", groupID).Find(&times).Error
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"times": times}, "群聊数据获取成功")
}

// DeleteGroupData
// 删除一个群聊的所有附属数据
func DeleteGroupData(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("group_id")
	var group ModelGroup.Group
	var groupUser ModelGroup.GroupUser
	var groupGraph ModelGroup.GroupGraph
	var groupKeyword ModelGroup.GroupKeyword
	var groupTime ModelGroup.GroupTime
	var groupCacheMessage ModelGroup.GroupCacheMessage
	var groupMessage ModelGroup.GroupMessage

	// 是否存在此群聊
	if err := DB.Where("group_id = ?", groupID).First(&group).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}

	// 开启事务, 如果任意一个删除不成功则回滚
	tx := DB.Begin()

	// 删除Group表中指定group_id的数据
	if err := tx.Where("group_id = ?", groupID).Unscoped().Delete(&group).Error; err != nil {
		tx.Rollback()
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "删除Group时发生错误")
		return
	}

	// 删除GroupUser表中指定group_id的数据
	if err := tx.Where("group_id = ?", groupID).Unscoped().Delete(&groupUser).Error; err != nil {
		tx.Rollback()
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "删除GroupUser时发生错误")
		return
	}

	// 删除GroupGraph表中指定group_id的数据
	if err := tx.Where("group_id = ?", groupID).Unscoped().Delete(&groupGraph).Error; err != nil {
		tx.Rollback()
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "删除GroupGraph时发生错误")
		return
	}

	// 删除GroupKeyword表中指定group_id的数据
	if err := tx.Where("group_id = ?", groupID).Unscoped().Delete(&groupKeyword).Error; err != nil {
		tx.Rollback()
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "删除GroupKeyword时发生错误")
		return
	}

	// 删除GroupTime表中指定group_id的数据
	if err := tx.Where("group_id = ?", groupID).Unscoped().Delete(&groupTime).Error; err != nil {
		tx.Rollback()
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "删除GroupTime时发生错误")
		return
	}

	// 删除GroupCacheMessage表中指定group_id的数据
	if err := tx.Where("group_id = ?", groupID).Unscoped().Delete(&groupCacheMessage).Error; err != nil {
		tx.Rollback()
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "删除GroupCacheMessage时发生错误")
		return
	}

	// 删除GroupMessage表中指定group_id的数据
	if err := tx.Where("group_id = ?", groupID).Unscoped().Delete(&groupMessage).Error; err != nil {
		tx.Rollback()
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "删除GroupMessage时发生错误")
		return
	}

	// 全部删除成功, 提交事务
	tx.Commit()
	common.Response(ctx, http.StatusOK, 200, nil, "删除群聊成功")
}
