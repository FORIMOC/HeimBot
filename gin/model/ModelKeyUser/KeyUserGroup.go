/* KeyUserGroup.go 关键用户群聊发言模型
------------------------
UserID: 用户ID
------------------------
GroupID: 群ID
------------------------
GroupName: 群聊名
------------------------
Freq: 发言频数
------------------------
*/

package ModelKeyUser

import "gorm.io/gorm"

type KeyUserGroup struct {
	gorm.Model
	UserID    string `gorm:"varchar(255);not null;"`
	GroupID   string `gorm:"varchar(255);not null;"`
	GroupName string `gorm:"varchar(255);not null;"`
	Freq      int    `gorm:"not null;default:0;"`
}
