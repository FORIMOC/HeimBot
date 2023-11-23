/* Admin.go 管理员用户模型
------------------------
UserID :用户ID
------------------------
Password: 密码
------------------------
Level: 管理员等级
	0: 全权
		允许关闭/重启机器人，允许修改任意配置，修改后不会发出提示消息
	1: 协助
		允许修改监听模式、关键词、关键用户等，且修改后会在基群发出提示消息，不允许修改管理员组
	2: 只读
		只允许浏览信息，不允许修改任何配置
------------------------
*/

package ModelAdmin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserID   string `gorm:"varchar(255);not null;unique;"`
	Password string `gorm:"varchar(255);not null;"`
	Level    int    `gorm:"not null;"`
}
