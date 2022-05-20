package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BotResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"reply": msg})
}

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
