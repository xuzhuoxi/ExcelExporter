package core

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
)

type TitleContext struct {
	// 使用的字段索引
	FieldTypeIndex int
	// 使用的编程语言
	ProgramLanguage string
}

func (o TitleContext) String() string {
	return fmt.Sprintf("TitleContent(FieldTypeIndex=%d, ProgramLanguage=%s)", o.FieldTypeIndex, o.ProgramLanguage)
}

type DataContext struct {
	// 使用的字段索引
	FieldTypeIndex int
	// 输出的文件类型
	DataFileFormat string
}

func (o DataContext) String() string {
	return fmt.Sprintf("DataContext(FieldTypeIndex=%d, DataFileFormat=%s)", o.FieldTypeIndex, o.DataFileFormat)
}

type ConstContext struct {
}

type GoTempData struct {
	Sheet *excel.ExcelSheet
	Excel *excel.ExcelProxy
	Index []int
}
