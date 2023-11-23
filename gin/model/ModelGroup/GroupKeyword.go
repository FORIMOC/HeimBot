/* GroupKeyword.go 群聊关键词模型
------------------------
GroupID: 群ID
------------------------
Keyword: 关键词
------------------------
Freq: 出现的频数
------------------------
*/

package ModelGroup

type GroupKeyword struct {
	GroupID string `gorm:"varchar(255);not null;uniqueIndex:GroupID_Keyword;size:255;"`
	Keyword string `gorm:"varchar(255);not null;uniqueIndex:GroupID_Keyword;size:255;"`
	Freq    int    `gorm:"not null;default:0;"`
}
