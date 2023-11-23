/* GroupUser.go 群聊用户模型
------------------------
GroupID: 群ID
------------------------
UserID: 用户ID
------------------------
Username: 用户名
------------------------
Freq: 发言频数
------------------------
Value: 社交网络中用户点的大小(参照社交网络相关逻辑以及Heimborad相关部分)
	30
	40
	50
	60
------------------------
KeywordNum: 发言内容中捕获的关键词频数
------------------------
*/

package ModelGroup

import "gorm.io/gorm"

type GroupUser struct {
	gorm.Model
	GroupID    string  `gorm:"varchar(255);not null;uniqueIndex:GroupID_UserID;size:255;"`
	UserID     string  `gorm:"varchar(255);not null;uniqueIndex:GroupID_UserID;size:255;"`
	Username   string  `gorm:"varchar(255);not null;"`
	Freq       int     `gorm:"not null;default:0;"`
	Value      float64 `gorm:"not null;default:0;"`
	KeywordNum int     `gorm:"not null;default:0;"`
}
