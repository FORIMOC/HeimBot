package HeimboardController

import (
	"encoding/json"
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelKeyUser"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"unsafe"
)

type GroupCard struct {
	Nickname string
	Card     map[string]string
}

// UpdateKeyUserStatus 更新关键用户状态控制器
// user_id(用户ID), status(要更改为的状态) =>
func UpdateKeyUserStatus(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	status := ctx.PostForm("status")
	if len(status) == 0 || len(userID) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "参数不足")
		return
	}
	var keyUser ModelKeyUser.KeyUser
	err := DB.Where("user_id = ?", userID).First(&keyUser).Error
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在该关键用户")
		return
	}
	keyUser.Status = status
	DB.Save(&keyUser)
	common.Response(ctx, http.StatusOK, 200, nil, "状态更新成功")
}

// SearchUsername 搜寻用户ID对于的昵称控制器
// user_id => []
func SearchUsername(ctx *gin.Context) {
	// 获取参数
	userID := ctx.PostForm("user_id")
	resp, err := http.Get("http://121.5.167.91/util/search_username.php?user_id=" + userID)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器错误")
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器错误")
		return
	}
	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	var groupCard GroupCard
	err = json.Unmarshal([]byte(*str), &groupCard)
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器错误")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{
		"nickname": groupCard.Nickname,
		"card":     groupCard.Card,
	}, "搜索关键用户昵称成功")
}

// AddKeyUser 添加关键用户控制器
// user_id(用户ID) =>
func AddKeyUser(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	if len(userID) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "关键用户QQ不能为空")
		return
	}
	var keyUser ModelKeyUser.KeyUser
	if err := DB.Where("user_id = ?", userID).First(&keyUser).Error; err == nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "已存在此关键用户")
		return
	}
	newKeyUser := ModelKeyUser.KeyUser{
		UserID: userID,
	}
	DB.Create(&newKeyUser)
	common.Response(ctx, http.StatusOK, 200, nil, "关键用户添加成功")
}

// DeleteKeyUser 删除关键用户控制器
// user_id(用户ID) =>
func DeleteKeyUser(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	var keyUser ModelKeyUser.KeyUser
	if err := DB.Where("user_id = ?", userID).First(&keyUser).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键用户")
		return
	}
	DB.Unscoped().Delete(&keyUser)
	common.Response(ctx, http.StatusOK, 200, nil, "关键用户删除成功")
}

// SetKeyUserInfo 设置关键用户备注控制器
// user_id(用户ID) =>
func SetKeyUserInfo(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	username := ctx.PostForm("username")
	remark := ctx.PostForm("remark")
	// 参数合法性检验
	if len(username) == 0 && len(remark) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名和备注不能都为空")
		return
	}

	var keyUser ModelKeyUser.KeyUser
	if err := DB.Where("user_id = ?", userID).First(&keyUser).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键用户")
		return
	}
	// 设置新用户名和备注
	if len(username) > 0 {
		keyUser.Username = username
	}
	if len(remark) > 0 {
		keyUser.Remark = remark
	}
	if err := DB.Save(&keyUser).Error; err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "设置备注发生错误")
		return
	}
	common.Response(ctx, http.StatusOK, 200, nil, "关键用户备注设置成功")
}

// GetKeyUserList 获取关键用户列表控制器
// => key_user_list([]ModelKeyUser.KeyUser)
func GetKeyUserList(ctx *gin.Context) {
	DB := common.GetDB()
	// keywords 数组
	var keyUserList []ModelKeyUser.KeyUser
	if err := DB.Find(&keyUserList).Error; err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "关键用户表查询失败")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"key_user_list": keyUserList}, "关键用户表查询成功")
}

// GetKeyUserInfo 获取关键用户信息控制器
// user_id(用户ID) => key_user(ModelKeyUser.KeyUser)
func GetKeyUserInfo(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	var keyUser ModelKeyUser.KeyUser
	if err := DB.Where("user_id = ?", userID).First(&keyUser).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键用户")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"key_user": keyUser}, "关键用户信息查询成功")
}

// GetKeyUserGroup 获取关键用户发言群聊控制器
// user_id(用户ID) => key_user_groups([]ModelKeyUser.KeyUserGroup)
func GetKeyUserGroup(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	groupUserLimit := viper.GetInt("data_limit.groupUserLimit")
	var keyUserGroups []ModelKeyUser.KeyUserGroup
	var count int64
	if err := DB.Where("user_id = ?", userID).Order("freq DESC").Limit(groupUserLimit).Find(&keyUserGroups).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键用户")
		return
	}
	DB.Table("key_user_groups").Where("user_id = ?", userID).Count(&count)
	common.Response(ctx, http.StatusOK, 200, gin.H{"key_user_groups": keyUserGroups, "count": count}, "关键用户发言群聊查询成功")
}

// GetKeyUserTime 获取关键用户发言时间分布控制器
func GetKeyUserTime(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	var keyUserTimes []ModelKeyUser.KeyUserTime
	if err := DB.Where("user_id = ?", userID).Find(&keyUserTimes).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键用户")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"key_user_times": keyUserTimes}, "关键用户发言时间分布查询成功")
}

// GetKeyUserKeywords 获取群聊关键词信息控制器
// user_id(群ID) => (指定user_id的[]ModelKeyUser.KeyUserKeyword)
func GetKeyUserKeywords(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	var keywords []ModelKeyUser.KeyUserKeyword
	var count int64
	err := DB.Where("user_id = ?", userID).Order("freq DESC").Limit(50).Find(&keywords).Error
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键用户")
		return
	}
	DB.Table("key_user_keywords").Where("user_id = ?", userID).Count(&count)
	common.Response(ctx, http.StatusOK, 200, gin.H{"keywords": keywords, "count": count}, "关键用户关键词数据获取成功")
}

// GetKeyUserCacheMessages 获取关键用户缓存消息控制器
// user_id(用户ID) => key_user_cache_messages()
func GetKeyUserCacheMessages(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	userID := ctx.PostForm("user_id")
	var keyUserCacheMessages []ModelKeyUser.KeyUserCacheMessage
	if err := DB.Where("user_id = ?", userID).Find(&keyUserCacheMessages).Error; err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键用户")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"key_user_cache_messages": keyUserCacheMessages}, "关键用户缓存消息查询成功")
}
