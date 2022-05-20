package model

import "github.com/jinzhu/gorm"

type Keyword struct {
	gorm.Model
	GroupID string `gorm:"varchar(255);not null;"`
	Keyword string `gorm:"varchar(255);not null;"`
	Freq    int    `gorm:"not null;default:0;"`
}
