package tools

import "time"

// 当前时间
func NowTime() time.Time {
	return time.Now()
}

// 当前时间
// 2006-01-02 15:04:05 PM Mon Jan
// 2006-01-_2 15:04:05 PM Mon Jan
func NowTimeFormat(format string) string {
	return time.Now().Format(format)
}

// 当前时间年份
func NowYear() int {
	return NowTime().Year()
}

// 当前时间月份
// 一月: 1
func NowMonth() int {
	return int(NowTime().Month())
}

// 当前时间日期
func NowDay() int {
	return NowTime().Day()
}

// 当前时间星期几
// 星期日： 0
func NowWeekday() int {
	return int(NowTime().Weekday())
}

// 当前时间小时
func NowHour() int {
	return NowTime().Hour()
}

// 当前时间分钟
func NowMinute() int {
	return NowTime().Minute()
}

// 当前时间秒
func NowSecond() int {
	return NowTime().Second()
}

// 当前时间秒戳（s）
func NowUnix() int64 {
	return NowTime().Unix()
}

// 当前时间秒戳（ns）
func NowUnixNano() int64 {
	return NowTime().UnixNano()
}
