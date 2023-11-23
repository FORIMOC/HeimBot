package HeimboardController

import (
	"forimoc.com/Heimbot/common"
	"forimoc.com/Heimbot/model/ModelTimeSeries"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// GetTimeSeriesResult 获取时间序列检测结果控制器
// => []ModelTimeSeries.TSCacheResult
func GetTimeSeriesResult(ctx *gin.Context) {
	DB := common.GetDB()
	var timeSeriesResult []ModelTimeSeries.TSCacheResult
	// 获取参数
	resultLimit := viper.GetInt("api.timeSeriesDetect.resultLimit")

	if err := DB.Find(&timeSeriesResult).Limit(resultLimit).Error; err != nil {
		log.Println(len(timeSeriesResult))
		common.Response(ctx, http.StatusInternalServerError, 500, nil, "未找到时间序列检测结果")
		return
	}
	common.Response(ctx, http.StatusOK, 200, gin.H{"result": timeSeriesResult}, "时间序列检测结果查询成功")
}
