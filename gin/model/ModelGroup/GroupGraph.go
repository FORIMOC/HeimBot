/* GraphGroup.go 社交网络模型
------------------------
GroupID: 群ID
------------------------
UserID1: 用户1ID
------------------------
Username1: 用户1名称
------------------------
UserID2: 用户2ID
------------------------
Username2: 用户2名称
------------------------
Value1: 用户1倾向于用户2的值
------------------------
Value2: 用户2倾向于用户1的值
------------------------
*/

package ModelGroup

import "gorm.io/gorm"

type GroupGraph struct {
	gorm.Model
	GroupID   string  `gorm:"varchar(255);not null;"`
	UserID1   string  `gorm:"varchar(255);not null;uniqueIndex:UserID1_UserID2;size:255;"`
	Username1 string  `gorm:"varchar(255);not null;"`
	UserID2   string  `gorm:"varchar(255);not null;uniqueIndex:UserID1_UserID2;size:255;"`
	Username2 string  `gorm:"varchar(255);not null;"`
	Value1    float64 `gorm:"not null;default:0;"`
	Value2    float64 `gorm:"not null;default:0;"`
}
