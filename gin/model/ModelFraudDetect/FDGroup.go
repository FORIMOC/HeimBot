/* FDGroup.go 出现的引流目标群聊模型
------------------------
GroupID: 群ID
------------------------
SenderID: 发送引流消息的用户ID
------------------------
*/

package ModelFraudDetect

import "gorm.io/gorm"

type FDGroup struct {
	gorm.Model
	GroupID  string `gorm:"varchar(255);not null;uniqueIndex:GroupID_SenderID;size:255;"`
	SenderID string `gorm:"varchar(255);not null;uniqueIndex:GroupID_SenderID;size:255;"`
}
