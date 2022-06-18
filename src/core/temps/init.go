package temps

import (
	"github.com/xuzhuoxi/ExcelExporter/src/core/tools"
)

// 初始化游戏到模板中的函数
func init() {
	RegisterFunc("ToLowerCamelCase", tools.ToLowerCamelCase)
	RegisterFunc("ToUpperCamelCase", tools.ToUpperCamelCase)

	RegisterFunc("Add", tools.Add)
	RegisterFunc("Sub", tools.Sub)

	RegisterFunc("NowTime", tools.NowTime)
	RegisterFunc("NowTimeStr", tools.NowTimeStr)
	RegisterFunc("NowTimeFormat", tools.NowTimeFormat)
	RegisterFunc("NowYear", tools.NowYear)
	RegisterFunc("NowMonth", tools.NowMonth)
	RegisterFunc("NowDay", tools.NowDay)
	RegisterFunc("NowWeekday", tools.NowWeekday)
	RegisterFunc("NowHour", tools.NowHour)
	RegisterFunc("NowMinute", tools.NowMinute)
	RegisterFunc("NowSecond", tools.NowSecond)
	RegisterFunc("NowUnix", tools.NowUnix)
	RegisterFunc("NowUnixNano", tools.NowUnixNano)
}
