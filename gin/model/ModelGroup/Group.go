/* GroupGraph.go 群聊模型
------------------------
GroupID: 群ID
------------------------
GroupName: 群名
------------------------
Status: 状态
	0: 不监听
	1: 摘要监听
	2: 控局监听
	3: 全局监听
------------------------
SumMsgNum: 消息总数
------------------------
AvgMsgLen: 平均消息长度
------------------------
MemberNum: 成员数
------------------------
MaxMemberNum: 最大成员数
------------------------
*/

package ModelGroup

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupID      string  `gorm:"varchar(255);not null;unique;"`
	GroupName    string  `gorm:"varchar(255);not null;"`
	Status       string  `gorm:"not null;default:0;"` //0:不监听 1:摘要监听 2:控局监听 3:全局监听
	SumMsgNum    int     `gorm:"not null;default:0;"`
	AvgMsgLen    float64 `gorm:"not null;default:0;"`
	MemberNum    int     `gorm:"not null;default:0;"`
	MaxMemberNum int     `gorm:"not null;default:0;"`
}
