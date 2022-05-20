package model

import "github.com/jinzhu/gorm"

type Graph struct {
	gorm.Model
	GroupID   string  `gorm:"varchar(255);not null;"`
	UserID1   string  `gorm:"varchar(255);not null;"`
	Username1 string  `gorm:"varchar(255);not null;"`
	UserID2   string  `gorm:"varchar(255);not null;"`
	Username2 string  `gorm:"varchar(255);not null;"`
	Value1    float64 `gorm:"not null;default:0;"`
	Value2    float64 `gorm:"not null;default:0;"`
}
