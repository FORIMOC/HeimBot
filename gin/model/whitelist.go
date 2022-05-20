package model

import "github.com/jinzhu/gorm"

type WhiteList struct {
	gorm.Model
	UserID      string `gorm:"varchar(255);not null;unique;"`
	Username    string `gorm:"varchar(255);not null;"`
	isMonitored bool   `gorm:"not null;default:false;"`
	canCallBot  bool   `gorm:"not null;default:true;"`
}
