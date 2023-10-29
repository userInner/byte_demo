package utils

import "time"

const (
	OrmatTimeStr = "2006-01-02 15:04:05"
)

// 日期时间 转化为时间戳
func GetTimeInt64(t string) int64 {
	local, _ := time.LoadLocation("Local")
	thetime, _ := time.ParseInLocation(OrmatTimeStr, t, local)
	return thetime.Unix()
}

// 获取当前时间
func GetTime() time.Time {
	return time.Now().Local()
}
