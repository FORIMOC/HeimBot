package detect

import (
	"encoding/json"
	"forimoc.com/Heimbot/model/ModelGroup"
	"forimoc.com/Heimbot/model/ModelTimeSeries"
	"forimoc.com/Heimbot/util"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"gorm.io/gorm"
	"log"
	"math"
	"sort"
	time2 "time"
)

type sampleUser struct {
	userID     string
	timeSeries []float64
}

// GetFraudAndNormalUser 根据涉诈关键词个数选一定涉诈用户, 以及根据发言次数选一定普通用户
func GetFraudAndNormalUser(DB *gorm.DB, data string) []sampleUser {
	groupID := gjson.Get(data, "group_id").String()
	var count int64
	// 从配置文件获得相关变量
	sampleRatio := viper.GetFloat64("api.timeSeriesDetect.sampleRatio")
	sampleMinimum := viper.GetInt("api.timeSeriesDetect.sampleMinimum")
	sampleFraudRatio := viper.GetFloat64("api.timeSeriesDetect.sampleFraudRatio")
	dataFlushLimit := viper.GetInt("api.timeSeriesDetect.dataFlushLimit")

	// 确定样本数
	sampleNum := sampleMinimum
	if err := DB.Table("group_user").Where("group_id = ?", groupID).Count(&count).Error; err == nil {
		if int(float64(count)*sampleRatio) > sampleMinimum {
			sampleNum = int(float64(count) * sampleRatio)
		}
	}
	sampleUsers := make([]sampleUser, 0)

	// 检索关键词主序用户
	var fraudUsers []ModelGroup.GroupUser
	fraudNum := int(float64(sampleNum) * sampleFraudRatio)
	DB.Order("keyword_num desc").Limit(fraudNum).Find(&fraudUsers)
	for _, user := range fraudUsers {
		var userTimes []ModelGroup.GroupUserTime
		if err := DB.Where("group_id = ? AND user_id = ?", groupID, user.UserID).Order("month asc, day asc").Limit(dataFlushLimit).Find(&userTimes).Error; err == nil {
			timeSeries := make([]float64, 0)
			for _, userTime := range userTimes {
				timeSeries = append(timeSeries, float64(userTime.Hour1))
				timeSeries = append(timeSeries, float64(userTime.Hour2))
				timeSeries = append(timeSeries, float64(userTime.Hour3))
				timeSeries = append(timeSeries, float64(userTime.Hour4))
				timeSeries = append(timeSeries, float64(userTime.Hour5))
				timeSeries = append(timeSeries, float64(userTime.Hour6))
				timeSeries = append(timeSeries, float64(userTime.Hour7))
				timeSeries = append(timeSeries, float64(userTime.Hour8))
				timeSeries = append(timeSeries, float64(userTime.Hour9))
				timeSeries = append(timeSeries, float64(userTime.Hour10))
				timeSeries = append(timeSeries, float64(userTime.Hour11))
				timeSeries = append(timeSeries, float64(userTime.Hour12))
			}
			newSampleUser := sampleUser{
				userID:     user.UserID,
				timeSeries: timeSeries,
			}
			sampleUsers = append(sampleUsers, newSampleUser)
		}
	}

	// 填充发言频次主序用户
	var normalUsers []ModelGroup.GroupUser
	normalNum := sampleNum - fraudNum
	DB.Order("freq desc").Limit(normalNum).Find(&normalUsers)
	for _, user := range normalUsers {
		var userTimes []ModelGroup.GroupUserTime
		if err := DB.Where("group_id = ? AND user_id = ?", groupID, user.UserID).Order("created_at asc").Limit(dataFlushLimit).Find(&userTimes).Error; err == nil {
			timeSeries := make([]float64, 0)
			for _, userTime := range userTimes {
				timeSeries = append(timeSeries, float64(userTime.Hour1))
				timeSeries = append(timeSeries, float64(userTime.Hour2))
				timeSeries = append(timeSeries, float64(userTime.Hour3))
				timeSeries = append(timeSeries, float64(userTime.Hour4))
				timeSeries = append(timeSeries, float64(userTime.Hour5))
				timeSeries = append(timeSeries, float64(userTime.Hour6))
				timeSeries = append(timeSeries, float64(userTime.Hour7))
				timeSeries = append(timeSeries, float64(userTime.Hour8))
				timeSeries = append(timeSeries, float64(userTime.Hour9))
				timeSeries = append(timeSeries, float64(userTime.Hour10))
				timeSeries = append(timeSeries, float64(userTime.Hour11))
				timeSeries = append(timeSeries, float64(userTime.Hour12))
			}
			newSampleUser := sampleUser{
				userID:     user.UserID,
				timeSeries: timeSeries,
			}
			log.Println(len(timeSeries))
			sampleUsers = append(sampleUsers, newSampleUser)
		}
	}
	return sampleUsers
}

func GroupTimeSeriesWarning(DB *gorm.DB, data string) {
	groupID := gjson.Get(data, "group_id").String()
	dataFlushLimit := viper.GetInt("api.timeSeriesDetect.dataFlushLimit")
	var time ModelTimeSeries.TSCacheTime
	var times []ModelTimeSeries.TSCacheTime
	var count int64

	hour := time2.Now().Hour()
	month := time2.Now().Month()
	day := time2.Now().Day()

	err := DB.Where("group_id = ? AND month = ? AND day = ?", groupID, month, day).First(&time).Error
	if err != nil {
		newTime := ModelTimeSeries.TSCacheTime{
			GroupID: groupID,
			Month:   int(month),
			Day:     day,
		}
		DB.Create(&newTime)
	}
	DB.Where("group_id = ? AND month = ? AND day = ?", groupID, month, day).First(&time)
	if hour < 2 {
		time.Hour1 += 1
	} else if hour < 4 {
		time.Hour2 += 1
	} else if hour < 6 {
		time.Hour3 += 1
	} else if hour < 8 {
		time.Hour4 += 1
	} else if hour < 10 {
		time.Hour5 += 1
	} else if hour < 12 {
		time.Hour6 += 1
	} else if hour < 14 {
		time.Hour7 += 1
	} else if hour < 16 {
		time.Hour8 += 1
	} else if hour < 18 {
		time.Hour9 += 1
	} else if hour < 20 {
		time.Hour10 += 1
	} else if hour < 22 {
		time.Hour11 += 1
	} else {
		time.Hour12 += 1
	}
	DB.Save(&time)

	DB.Table("ts_cache_times").Where("group_id = ?", groupID).Count(&count)
	// 达到刷新条件, 将缓存发送算法端得到三角函数, 并进行后续检测
	if count > int64(dataFlushLimit) {
		DB.Where("group_id = ?", groupID).Order("created_at asc").Limit(dataFlushLimit).Find(&times)
		// 删除缓存
		DB.Exec("DELETE FROM ts_cache_times ORDER BY created_at ASC LIMIT 5")

		// 算法端 url
		host := viper.GetString("api.timeSeriesDetect.host")
		port := viper.GetString("api.timeSeriesDetect.port")
		route := viper.GetString("api.timeSeriesDetect.route")
		url := "http://" + host + ":" + port + route

		// 叠加 dataFlushLimit 天的时间序列
		timeSeries := make([]int, 0)
		for i := 0; i < len(times); i++ {
			timeSeries = append(timeSeries, times[i].Hour1)
			timeSeries = append(timeSeries, times[i].Hour2)
			timeSeries = append(timeSeries, times[i].Hour3)
			timeSeries = append(timeSeries, times[i].Hour4)
			timeSeries = append(timeSeries, times[i].Hour5)
			timeSeries = append(timeSeries, times[i].Hour6)
			timeSeries = append(timeSeries, times[i].Hour7)
			timeSeries = append(timeSeries, times[i].Hour8)
			timeSeries = append(timeSeries, times[i].Hour9)
			timeSeries = append(timeSeries, times[i].Hour10)
			timeSeries = append(timeSeries, times[i].Hour11)
			timeSeries = append(timeSeries, times[i].Hour12)
		}

		// 准备要发送的json数据
		jsonData := map[string]interface{}{
			"time_series": timeSeries,
			"top_k":       3,
		}

		// 发送请求给算法端
		resp := util.SendJSON(url, jsonData)
		var result util.TSModelResult
		json.Unmarshal(resp, &result)

		// 按 ACF 绝对值排序
		sort.Slice(result.Result, func(i, j int) bool {
			return math.Abs(result.Result[i].ACF) > math.Abs(result.Result[j].ACF)
		})

		// 对最大 ACF 的这条频率进行检测, 先求出还原的三角函数
		triFunc := func(x float64) float64 {
			return result.Result[0].Amplitude * math.Sin(2*math.Pi*result.Result[0].Freq*x+result.Result[0].Phase)
		}

		// 根据涉诈权值算法体系筛选待检测样本
		sampleUsers := GetFraudAndNormalUser(DB, data)

		// 先根据 triFunc 计算预测值
		timeLine := floats.Span(make([]float64, dataFlushLimit*12), 0, 1)
		preds := make([]float64, 0)
		for _, x := range timeLine {
			preds = append(preds, triFunc(x))
		}

		// 计算每个用户时间序列和 triFunc 函数预测值的 MSE
		MSEs := make([]float64, 0)
		for _, user := range sampleUsers {
			MSE := util.MeanSquaredError(user.timeSeries, preds)
			MSEs = append(MSEs, MSE)
		}

		// 计算 MSE 的均值和方差
		mean := stat.Mean(MSEs, nil)
		variance := stat.Variance(MSEs, nil)

		// 构建正态分布, 并检测哪些用户离群
		maxMSE := 0.0
		maxidx := 0
		for i, MSE := range MSEs {
			if MSE > maxMSE {
				maxMSE = MSE
				maxidx = i
			}
		}
		outliers := sampleUsers[maxidx].userID

		// 结果存入数据库
		newTSCacheResult := ModelTimeSeries.TSCacheResult{
			GroupID:     groupID,
			Amplitude1:  result.Result[0].Amplitude,
			Amplitude2:  result.Result[1].Amplitude,
			Amplitude3:  result.Result[2].Amplitude,
			Freq1:       result.Result[0].Freq,
			Freq2:       result.Result[1].Freq,
			Freq3:       result.Result[2].Freq,
			Phase1:      result.Result[0].Phase,
			Phase2:      result.Result[1].Phase,
			Phase3:      result.Result[2].Phase,
			ACF1:        result.Result[0].ACF,
			ACF2:        result.Result[1].ACF,
			ACF3:        result.Result[2].ACF,
			MSEMean:     mean,
			MSEVariance: variance,
			Outliers:    outliers,
			OutliersMSE: maxMSE,
		}
		DB.Create(&newTSCacheResult)

		// 如果超出限制, 删掉最早的记录
		resultLimit := viper.GetInt64("api.timeSeriesDetect.resultLimit")
		DB.Table("ts_cache_time").Count(&count)
		if count > resultLimit {
			DB.Unscoped().Order("created_at asc").Limit(1).Delete(&ModelTimeSeries.TSCacheResult{})
		}

		// 报警
		//var group ModelGroup.Group
		//DB.Where("group_id = ?", groupID).First(&group)
		//response := "**[" + group.GroupName + "]-(" + groupID + ")**\n" +
		//	"时间序列算法检测到如下离群用户, 可能包含涉诈用户\n" +
		//	outliers
		//common.BotDefaultResponse(response)
	}
}
