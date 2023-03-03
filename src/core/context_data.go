package core

import (
	"fmt"
)

// 数据导出上下文
type DataContext struct {
	EnablePrefix   string         // 开启前缀
	RangeName      string         // 使用的字段索引名称
	RangeType      FieldRangeType // 使用的字段索引
	DataFileFormat string         // 输出的文件类型
	StartRowNum    int            // 数据开始行号
	StartColIndex  int            // 数据开始列索引
}

func (o DataContext) String() string {
	return fmt.Sprintf("DataContext(Prefix=%s, RangeName=%s, RangeType=%v, ProgramLanguage=%s, StartRowNum=%d, StartColIndex=%d)",
		o.EnablePrefix, o.RangeName, o.RangeType, o.DataFileFormat, o.StartRowNum, o.StartColIndex)
}
