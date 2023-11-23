package HeimboardController

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/model/ModelKeyword"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddKeyword 增加关键词控制器
// keyword(关键词) =>
func AddKeyword(ctx *gin.Context) {
	DB := common.GetDB()
	keyword := ctx.PostForm("keyword")
	if len(keyword) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "关键词不能为空")
		return
	}
	var keywordList ModelKeyword.KeywordList
	err := DB.Where("keyword = ?", keyword).First(&keywordList).Error
	if err == nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "已存在此关键词")
		return
	}
	newKeyword := ModelKeyword.KeywordList{
		Keyword: keyword,
	}
	DB.Create(&newKeyword)
	//common.BotDefaultResponse("**Heimboard**\n添加了关键词: " + keyword)
	common.Response(ctx, http.StatusOK, 200, nil, "关键词添加成功")
}

// DeleteKeyword 删除关键词控制器
// keyword(关键词) =>
func DeleteKeyword(ctx *gin.Context) {
	DB := common.GetDB()
	keyword := ctx.PostForm("keyword")
	var keywordList ModelKeyword.KeywordList
	err := DB.Where("keyword = ?", keyword).First(&keywordList).Error
	if err != nil {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "不存在此关键词")
		return
	}
	DB.Unscoped().Delete(&keywordList)
	//common.BotDefaultResponse("**Heimboard**\n删除了关键词: " + keyword)
	common.Response(ctx, http.StatusOK, 200, nil, "关键词删除成功")
}

// GetKeywordList 获取关键词列表控制器
// => keywords
func GetKeywordList(ctx *gin.Context) {
	DB := common.GetDB()
	// keywords 数组
	var keywordList []ModelKeyword.KeywordList
	err := DB.Find(&keywordList).Error
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "关键词表查询失败")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"keyword_list": keywordList}, "关键词表查询成功")
}

// RegulateKeyword 重整关键词一致性控制器(删除词表中没有而群聊关键词记录中有的)
// =>
func RegulateKeyword(ctx *gin.Context) {
	DB := common.GetDB()
	var k = 0
	// keywordList 数组
	var keywordList []ModelKeyword.KeywordList
	// keyword 数组
	var keyword []ModelGroup.GroupKeyword
	err := DB.Find(&keywordList).Error
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "关键词表查询失败")
		return
	}
	err = DB.Find(&keyword).Error
	if err != nil {
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "关键词查询失败")
		return
	}

	for i := 0; i < len(keyword); i++ {
		k = 0
		for j := 0; j < len(keywordList); j++ {
			if keyword[i].Keyword == keywordList[j].Keyword {
				k = 1
				break
			}
		}
		if k == 0 {
			DB.Unscoped().Delete(keyword[i])
		}
	}

	//common.BotDefaultResponse("**Heimboard**\n重整了关键词一致性")
	common.Response(ctx, http.StatusOK, 200, nil, "关键词表一致性重整成功")
}
