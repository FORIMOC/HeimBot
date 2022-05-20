package model

import "github.com/jinzhu/gorm"

type KeywordList struct {
	gorm.Model
	Keyword string `gorm:"varchar(255);not null;unique;"`
}
