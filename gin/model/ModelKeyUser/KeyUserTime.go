/* KeyUserTime.go 关键用户发言时间分布模型
------------------------
UserID: 用户ID
------------------------
Month: 月份
------------------------
Day: 日
------------------------
HourX: 某时间段的消息数
	将一天24小时平均划分为12个时间段
------------------------
*/

package ModelKeyUser

import "gorm.io/gorm"

type KeyUserTime struct {
	gorm.Model
	UserID string `gorm:"varchar(255);not null;"`
	Month  int    `gorm:"not null;default:0;"`
	Day    int    `gorm:"not null;default:0;"`
	Hour1  int    `gorm:"not null;default:0;"`
	Hour2  int    `gorm:"not null;default:0;"`
	Hour3  int    `gorm:"not null;default:0;"`
	Hour4  int    `gorm:"not null;default:0;"`
	Hour5  int    `gorm:"not null;default:0;"`
	Hour6  int    `gorm:"not null;default:0;"`
	Hour7  int    `gorm:"not null;default:0;"`
	Hour8  int    `gorm:"not null;default:0;"`
	Hour9  int    `gorm:"not null;default:0;"`
	Hour10 int    `gorm:"not null;default:0;"`
	Hour11 int    `gorm:"not null;default:0;"`
	Hour12 int    `gorm:"not null;default:0;"`
}
