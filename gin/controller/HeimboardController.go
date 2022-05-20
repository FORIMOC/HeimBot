package controller

import (
	"forimoc.com/Hoshino/common"
	"forimoc.com/Hoshino/model"
	"forimoc.com/Hoshino/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	token := ctx.PostForm("token")
	var login model.Login
	err := DB.Raw("SELECT * FROM logins WHERE id = 1").Scan(&login).Error
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务器错误")
		return
	}
	if token != login.Token {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "令牌错误")
		return
	}
	response.Response(ctx, http.StatusOK, 200, nil, "登录成功")
}

func GetGroupList(ctx *gin.Context) {
	DB := common.GetDB()
	var groups []model.Group
	DB.Find(&groups)
	response.Response(ctx, http.StatusOK, 200, gin.H{"groups": groups}, "群聊列表查询成功")
}

func GetGroupInfo(ctx *gin.Context) {
	DB := common.GetDB()
	var group model.Group
	// 获取数据
	groupID := ctx.PostForm("groupID")
	err := DB.Where("group_id = ?", groupID).First(&group).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{"group": group}, "群聊数据获取成功")
}

func GetGroupGraph(ctx *gin.Context) {
	DB := common.GetDB()
	var graphs []model.Graph
	// 获取参数
	groupID := ctx.PostForm("groupID")
	err := DB.Where("group_id = ?", groupID).Find(&graphs).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{"graphs": graphs}, "群聊数据获取成功")
}

func GetUserList(ctx *gin.Context) {
	DB := common.GetDB()
	groupID := ctx.PostForm("groupID")
	var users []model.User
	var count int
	err := DB.Where("group_id = ?", groupID).Order("freq DESC").Find(&users).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	DB.Table("users").Where("group_id = ?", groupID).Count(&count)
	response.Response(ctx, http.StatusOK, 200, gin.H{"users": users, "count": count}, "用户数据获取成功")
}

func GetGroupPremessages(ctx *gin.Context) {
	DB := common.GetDB()
	groupID := ctx.PostForm("groupID")
	var premessages []model.Premessage
	err := DB.Where("group_id = ?", groupID).Find(&premessages).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{"premessages": premessages}, "群聊数据获取成功")
}

func GetGroupKeywords(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("groupID")
	var keywords []model.Keyword
	var count int
	err := DB.Where("group_id = ?", groupID).Order("freq DESC").Limit(20).Find(&keywords).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	DB.Table("keywords").Where("group_id = ?", groupID).Count(&count)
	response.Response(ctx, http.StatusOK, 200, gin.H{"keywords": keywords, "count": count}, "群聊数据获取成功")
}

func GetGroupTime(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("groupID")
	var times []model.Time
	err := DB.Where("group_id = ?", groupID).Find(&times).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此群聊")
		return
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{"times": times}, "群聊数据获取成功")
}

func UpdateStatus(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	groupID := ctx.PostForm("groupID")
	status := ctx.PostForm("status")
	if len(status) == 0 || len(groupID) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "参数不足")
		return
	}
	var group model.Group
	err := DB.Where("group_id = ?", groupID).First(&group).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在该群聊")
		return
	}
	group.Status = status
	DB.Save(&group)
	response.Response(ctx, http.StatusOK, 200, nil, "状态更新成功")
}

func AddKeyword(ctx *gin.Context) {
	DB := common.GetDB()
	keyword := ctx.PostForm("keyword")
	var keywordList model.KeywordList
	err := DB.Where("keyword = ?", keyword).First(&keywordList).Error
	if err == nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "已存在此关键词")
		return
	}
	newKeyword := model.KeywordList{
		Keyword: keyword,
	}
	DB.Create(&newKeyword)
	response.Response(ctx, http.StatusOK, 200, nil, "关键词添加成功")
}

func DeleteKeyword(ctx *gin.Context) {
	DB := common.GetDB()
	keyword := ctx.PostForm("keyword")
	var keywordList model.KeywordList
	err := DB.Where("keyword = ?", keyword).First(&keywordList).Error
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键词")
		return
	}
	DB.Delete(&keywordList)
	response.Response(ctx, http.StatusOK, 200, nil, "关键词删除成功")
}

func ListKeyword(ctx *gin.Context) {
	DB := common.GetDB()
	// keywords 数组
	var keywordList []model.KeywordList
	err := DB.Find(&keywordList).Error
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "关键词表查询失败")
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{"keywords": keywordList}, "关键词表查询成功")
}
