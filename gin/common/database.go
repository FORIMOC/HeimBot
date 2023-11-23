package common

import (
	"fmt"
	"forimoc.com/Heimbot/model/ModelAdmin"
	"forimoc.com/Heimbot/model/ModelFraudDetect"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/model/ModelKeyUser"
	"forimoc.com/Heimbot/model/ModelKeyword"
	"forimoc.com/Heimbot/model/ModelTimeSeries"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 保存数据库实例的全局变量
var DB *gorm.DB

// InitDB 数据库配置初始化
// =>
func InitDB() {
	// 从配置文件 application.yml 获得配置信息
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	// loc=Local 设置本地时区
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 自动将模型迁移至数据库表
	err = db.AutoMigrate(
		&ModelAdmin.Admin{},
		&ModelGroup.Group{},
		&ModelGroup.GroupCacheMessage{},
		&ModelGroup.GroupGraph{},
		&ModelGroup.GroupKeyword{},
		&ModelGroup.GroupMessage{},
		&ModelGroup.GroupTime{},
		&ModelGroup.GroupUserTime{},
		&ModelGroup.GroupUser{},
		&ModelKeyword.KeywordList{},
		&ModelKeyUser.KeyUser{},
		&ModelKeyUser.KeyUserTime{},
		&ModelKeyUser.KeyUserCacheMessage{},
		&ModelKeyUser.KeyUserKeyword{},
		&ModelKeyUser.KeyUserGroup{},
		&ModelFraudDetect.FDCacheMessage{},
		&ModelFraudDetect.FDMessage{},
		&ModelFraudDetect.FDGroup{},
		&ModelTimeSeries.TSCacheTime{},
		&ModelTimeSeries.TSCacheResult{},
	)
	if err != nil {
		panic(err)
	}
	DB = db
}

// GetDB 返回数据库实例
// => 数据库实例
func GetDB() *gorm.DB {
	return DB
}
