/* KeyUserKeyword.go 关键用户关键词模型
------------------------
UserID: 用户ID
------------------------
Keyword: 关键词
------------------------
Freq: 出现的频数
------------------------
*/

package ModelKeyUser

import "gorm.io/gorm"

type KeyUserKeyword struct {
	gorm.Model
	UserID  string `gorm:"varchar(255);not null;"`
	Keyword string `gorm:"varchar(255);not null;"`
	Freq    int    `gorm:"not null;default:0;"`
}
