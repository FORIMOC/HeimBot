/* KeyUserCacheMessage.go 群聊缓存消息模型
------------------------
GroupID: 群ID
------------------------
UserID: 用户ID
------------------------
Username: 用户名
------------------------
GroupName: 群聊名
------------------------
Content: 消息内容
------------------------
*/

package ModelKeyUser

import "gorm.io/gorm"

type KeyUserCacheMessage struct {
	gorm.Model
	UserID    string `gorm:"varchar(255);not null;"`
	Username  string `gorm:"varchar(255);not null;"`
	GroupID   string `gorm:"varchar(255);not null;"`
	GroupName string `gorm:"varchar(255);not null;"`
	Content   string `gorm:"size:2047;not null;default:0"`
}
