package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string  `gorm:"varchar(255);not null;"`
	Username string  `gorm:"varchar(255);not null;"`
	GroupID  string  `gorm:"varchar(255);not null;"`
	Freq     int     `gorm:"not null;default:0;"`
	Value    float64 `gorm:"not null;default:0;"`
}
