package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupID   string  `gorm:"varchar(255);not null;unique;"`
	GroupName string  `gorm:"varchar(255);not null;"`
	Status    string  `gorm:"not null;default:0;"` //0:不监听 1:摘要监听 2:控局监听 3:全局监听
	SumMsg    int     `gorm:"not null;default:0;"`
	AvgMsg    float64 `gorm:"not null;default:0;"`
	Member    int     `gorm:"not null;default:0;"`
	MaxMember int     `gorm:"not null;default:0;"`
}
