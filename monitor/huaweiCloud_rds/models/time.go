package models

import "time"

// 转换查询时间，时间格式为"yyyy-mm-ddThh:mm:ssZ"，示例 "2023-08-17T08:00:00+0800"
func NowStr() (string, string) {
	queryEndTime := time.Now().Format("2006-01-02T15:04:05+0800")
	// queryStartTime := time.Now().Add(-1 * time.Hour).Format("2006-01-02T15:04:05+0800")
	queryStartTime := time.Now().Add(-2 * time.Minute).Format("2006-01-02T15:04:05+0800")
	return queryStartTime, queryEndTime
}
func Add8(intime string) (outtime string) {
	local, _ := time.LoadLocation("Asia/Shanghai")
	timestamp, _ := time.ParseInLocation("2006-01-02T15:04:05", intime, local)
	d3 := timestamp.Add(time.Hour * 8)
	return d3.String()
}
