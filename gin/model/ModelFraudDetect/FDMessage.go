/* FDMessage.go 模型确定为诈骗消息的模型
------------------------
GroupID: 群ID
------------------------
UserID: 用户ID
------------------------
Username: 用户名
------------------------
Content: 消息内容
------------------------
*/

package ModelFraudDetect

import "gorm.io/gorm"

type FDMessage struct {
	gorm.Model
	GroupID  string `gorm:"varchar(255);not null;"`
	UserID   string `gorm:"varchar(255);not null;"`
	Username string `gorm:"varchar(255);not null;"`
	Content  string `gorm:"size:2047;not null;default:0"`
}
