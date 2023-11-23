/* KeywordList.go 关键词表模型
------------------------
Keyword: 关键词
------------------------
*/

package ModelKeyword

import "gorm.io/gorm"

type KeywordList struct {
	gorm.Model
	Keyword string `gorm:"varchar(255);not null;unique;"`
}
