package routes

import (
	"forimoc.com/Hoshino/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// HeimBot control
	r.POST("/", controller.MainEntry)
	// 登录
	r.POST("/api/auth/login", controller.Login)
	// group 数据
	r.POST("/api/group/update_status", controller.UpdateStatus)
	r.POST("/api/group/get_group_list", controller.GetGroupList)
	r.POST("/api/group/get_group_info", controller.GetGroupInfo)
	r.POST("/api/group/get_group_graphs", controller.GetGroupGraph)
	r.POST("/api/group/get_users_list", controller.GetUserList)
	r.POST("/api/group/get_group_premessages", controller.GetGroupPremessages)
	r.POST("/api/group/get_group_time", controller.GetGroupTime)
	r.POST("/api/group/get_group_keywords", controller.GetGroupKeywords)
	// user 数据

	// 关键词操作
	r.POST("/api/keyword/add_keyword", controller.AddKeyword)
	r.POST("/api/keyword/delete_keyword", controller.DeleteKeyword)
	r.POST("/api/keyword/list_keyword", controller.ListKeyword)

	return r
}
