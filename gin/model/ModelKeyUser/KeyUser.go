/* KeyUser.go 关键用户模型
------------------------
UserID: 用户ID
------------------------
Username: 用户名
------------------------
Remark: 备注
------------------------
SumMsgNum: 总发言数
------------------------
AvgMsgLen: 平均发言数
------------------------
Status: 监听状态
------------------------
*/

package ModelKeyUser

import "gorm.io/gorm"

type KeyUser struct {
	gorm.Model
	UserID    string  `gorm:"varchar(255);not null;unique;"`
	Username  string  `gorm:"varchar(255);"`
	Remark    string  `gorm:"varchar(255);"`
	SumMsgNum int     `gorm:"not null;default:0;"`
	AvgMsgLen float64 `gorm:"not null;default:0;"`
	Status    string  `gorm:"not null;default:1;"` //0:不监听 1:开启监听
}
