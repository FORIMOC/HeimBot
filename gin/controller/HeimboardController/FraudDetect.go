package HeimboardController

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelFraudDetect"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

// GetFraudMessages 获取诈骗消息控制器
// => fraud_messages()
func GetFraudMessages(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	fraudMessageLogLimit := viper.GetInt("api.fraudDetect.fraudMessageLogLimit")
	var fraudMessages []ModelFraudDetect.FDMessage
	DB.Table("fd_messages").Find(&fraudMessages).Limit(fraudMessageLogLimit)
	common.Response(ctx, http.StatusOK, 200, gin.H{"fraud_messages": fraudMessages}, "涉诈消息查询成功")
}

func GetFraudGroups(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	fraudGroupLogLimit := viper.GetInt("api.fraudDetect.fraudGroupLogLimit")
	var fraudGroup []ModelFraudDetect.FDGroup
	DB.Table("fd_groups").Find(&fraudGroup).Limit(fraudGroupLogLimit)
	common.Response(ctx, http.StatusOK, 200, gin.H{"fraud_groups": fraudGroup}, "涉诈群聊查询成功")
}
