package middleware

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelAdmin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth 管理员验证中间件
// 管理员等级 =>
func Auth(adminLevel int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 检验票据格式
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 检验票据格式
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 验证通过后获取claim中的userID
		ID := claims.UserID
		DB := common.GetDB()
		var admin ModelAdmin.Admin
		DB.First(&admin, ID)

		// 用户不存在
		if admin.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 管理员等级过低
		if admin.Level > adminLevel {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 用户存在且等级足够 将admin信息写入上下文
		ctx.Set("admin", admin)

		ctx.Next()
	}
}
