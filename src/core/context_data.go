package core

import (
	"fmt"
)

// 数据导出上下文
type DataContext struct {
	RangeName      string         // 使用的字段索引名称
	RangeType      FieldRangeType // 使用的字段索引
	DataFileFormat string         // 输出的文件类型
}

func (o DataContext) String() string {
	return fmt.Sprintf("DataContext(RangeName=%s, RangeType=%v, ProgramLanguage=%s)",
		o.RangeName, o.RangeType, o.DataFileFormat)
}
