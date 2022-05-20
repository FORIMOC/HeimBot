package common

import (
	"fmt"
	"forimoc.com/Hoshino/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	// loc=Local 设置本地时区
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.Graph{})
	db.AutoMigrate(&model.Group{})
	db.AutoMigrate(&model.Keyword{})
	db.AutoMigrate(&model.KeywordList{})
	db.AutoMigrate(&model.Message{})
	db.AutoMigrate(&model.Premessage{})
	db.AutoMigrate(&model.Time{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.WhiteList{})
	db.AutoMigrate(&model.Login{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
