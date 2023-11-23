/* TSCacheResult.go 时间序列检测结果模型
------------------------
GroupID: 群ID
------------------------
AmplitudeX: 时间序列第X强子序列的振幅
------------------------
FreqX: 时间序列第X强子序列的振幅频率
------------------------
PhaseX: 时间序列第X强子序列的相位差
------------------------
ACFX: 时间序列第X强子序列的自相关系数
------------------------
MSEMean: 所有用户和
------------------------
*/

package ModelTimeSeries

import "gorm.io/gorm"

type TSCacheResult struct {
	gorm.Model
	GroupID     string  `gorm:"varchar(255);not null;size:255;"`
	Amplitude1  float64 `gorm:"not null;default:0;"`
	Amplitude2  float64 `gorm:"not null;default:0;"`
	Amplitude3  float64 `gorm:"not null;default:0;"`
	Freq1       float64 `gorm:"not null;default:0;"`
	Freq2       float64 `gorm:"not null;default:0;"`
	Freq3       float64 `gorm:"not null;default:0;"`
	Phase1      float64 `gorm:"not null;default:0;"`
	Phase2      float64 `gorm:"not null;default:0;"`
	Phase3      float64 `gorm:"not null;default:0;"`
	ACF1        float64 `gorm:"not null;default:0;"`
	ACF2        float64 `gorm:"not null;default:0;"`
	ACF3        float64 `gorm:"not null;default:0;"`
	MSEMean     float64 `gorm:"not null;default:0;"`
	MSEVariance float64 `gorm:"not null;default:0;"`
	Outliers    string  `gorm:"varchar(255);not null;size:255;"`
	OutliersMSE float64 `gorm:"not null;default:0;"`
}
