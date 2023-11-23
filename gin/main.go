package main

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelAdmin"
	"forimoc.com/Heimbot/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func main() {
	Init()                     // 配置文件和数据库的初始化
	r := gin.Default()         // gin实例
	r = routes.CollectRoute(r) // 初始化路由信息
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
}

// Init 读取配置文件并初始化
func Init() {
	// 初始化配置文件路径
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 数据库初始化
	common.InitDB()
	DB := common.GetDB()

	/*--- 初始账号与配置文件对齐 ---*/
	// 获取 ID 和 密码
	ownerID := viper.GetString("owner.ID")
	ownerPassword := viper.GetString("owner.password")

	// 获取加密的密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ownerPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	var admin ModelAdmin.Admin
	if err := DB.Where("user_id = ?", ownerID).First(&admin).Error; err != nil {
		// 添加全权用户
		newAdmin := &ModelAdmin.Admin{
			UserID:   ownerID,
			Password: string(hashedPassword),
			Level:    0,
		}
		if err := DB.Create(&newAdmin).Error; err != nil {
			panic(err)
		}
	} else {
		// 核对密码是否与配置文件一致
		if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(ownerPassword)); err != nil {
			// 更改密码
			admin.Password = string(hashedPassword)
			if err := DB.Save(&admin).Error; err != nil {
				panic(err)
			}
		}
	}
}
