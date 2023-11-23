/* GroupUserTime.go 群聊用户消息时间分布模型
------------------------
GroupID: 群ID
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

package ModelGroup

import "gorm.io/gorm"

type GroupUserTime struct {
	gorm.Model
	GroupID string `gorm:"varchar(255);not null;uniqueIndex:GroupID_UserID_Month_Day;size:255;"`
	UserID  string `gorm:"varchar(255);not null;uniqueIndex:GroupID_UserID_Month_Day;size:255;"`
	Month   int    `gorm:"not null;default:0;uniqueIndex:GroupID_UserID_Month_Day;"`
	Day     int    `gorm:"not null;default:0;uniqueIndex:GroupID_UserID_Month_Day;"`
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
