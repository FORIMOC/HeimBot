package model

import "github.com/jinzhu/gorm"

type Login struct {
	gorm.Model
	Token string `gorm:"varchar(255);not null;"`
}
