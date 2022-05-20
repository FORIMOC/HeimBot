package model

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	UserID   string `gorm:"varchar(255);not null;"`
	Username string `gorm:"varchar(255);not null;"`
	GroupID  string `gorm:"varchar(255);not null;"`
	Content  string `gorm:"size:2047;not null;default:0;"`
}
