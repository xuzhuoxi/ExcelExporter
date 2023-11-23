package tools

import "time"

const (
	defaultFormat = "2006-01-02 15:04:05"
)

// NowTime 当前时间
func NowTime() time.Time {
	return time.Now()
}

// NowTimeStr 当前时间
func NowTimeStr() string {
	return NowTimeFormat(defaultFormat)
}

// NowTimeFormat 当前时间
// 2006-01-02 15:04:05 PM Mon Jan
// 2006-01-_2 15:04:05 PM Mon Jan
func NowTimeFormat(format string) string {
	return time.Now().Format(format)
}

// NowYear 当前时间年份
func NowYear() int {
	return NowTime().Year()
}

// NowMonth 当前时间月份
// 一月: 1
func NowMonth() int {
	return int(NowTime().Month())
}

// NowDay 当前时间日期
func NowDay() int {
	return NowTime().Day()
}

// NowWeekday 当前时间星期几
// 星期日： 0
func NowWeekday() int {
	return int(NowTime().Weekday())
}

// NowHour 当前时间小时
func NowHour() int {
	return NowTime().Hour()
}

// NowMinute 当前时间分钟
func NowMinute() int {
	return NowTime().Minute()
}

// NowSecond 当前时间秒
func NowSecond() int {
	return NowTime().Second()
}

// NowUnix 当前时间秒戳（s）
func NowUnix() int64 {
	return NowTime().Unix()
}

// NowUnixNano 当前时间秒戳（ns）
func NowUnixNano() int64 {
	return NowTime().UnixNano()
}
