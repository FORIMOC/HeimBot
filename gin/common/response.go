/* response.go 系统响应

 */

package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/url"
	"os/exec"
)

// BotResponse 指定机器人 向 指定群聊 发送 指定消息
// 机器人端口号, 目标群ID, 要发送的消息 =>
func BotResponse(port string, group string, content string) {
	u := "http://127.0.0.1:" + port + "/send_group_msg?group_id=" + group + "&message=" + url.QueryEscape(content)
	exec.Command("curl", u).Run()
}

// BotDefaultResponse Hoshino 向 基群 发送 指定消息
// 要发送的消息 =>
func BotDefaultResponse(content string) {
	port := viper.GetString("bot.port")
	baseGroupID := viper.GetString("bot.baseGroupID")
	isWarningEscape := viper.GetBool("bot.isWarningEscape")
	u := "http://127.0.0.1:" + port + "/send_group_msg?group_id=" + baseGroupID + "&message=" + url.QueryEscape(content)
	if isWarningEscape {
		u += "&auto_escape=true"
	}
	exec.Command("curl", u).Run()
}

// Response Heimboard前端通用响应包
// gin上下文, http状态码, 状态码, 响应数据, 响应消息
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
