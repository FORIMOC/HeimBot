package model

import "github.com/jinzhu/gorm"

type Time struct {
	gorm.Model
	GroupID string `gorm:"varchar(255);not null;"`
	Month   int    `gorm:"not null;default:0;"`
	Day     int    `gorm:"not null;default:0;"`
	Hour1   int    `gorm:"not null;default:0;"`
	Hour2   int    `gorm:"not null;default:0;"`
	Hour3   int    `gorm:"not null;default:0;"`
	Hour4   int    `gorm:"not null;default:0;"`
	Hour5   int    `gorm:"not null;default:0;"`
	Hour6   int    `gorm:"not null;default:0;"`
	Hour7   int    `gorm:"not null;default:0;"`
	Hour8   int    `gorm:"not null;default:0;"`
	Hour9   int    `gorm:"not null;default:0;"`
	Hour10  int    `gorm:"not null;default:0;"`
	Hour11  int    `gorm:"not null;default:0;"`
	Hour12  int    `gorm:"not null;default:0;"`
}
