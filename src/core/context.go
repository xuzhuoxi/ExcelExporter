package core

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
)

type TitleContext struct {
	// 使用的字段索引
	FieldType FieldType
	// 使用的编程语言
	ProgramLanguage string
}

func (o TitleContext) String() string {
	return fmt.Sprintf("TitleContent(FieldType=%v, ProgramLanguage=%s)", o.FieldType, o.ProgramLanguage)
}

type DataContext struct {
	// 使用的字段索引
	FieldType FieldType
	// 输出的文件类型
	DataFileFormat string
}

func (o DataContext) String() string {
	return fmt.Sprintf("DataContext(FieldTypeIndex=%v, DataFileFormat=%s)", o.FieldType, o.DataFileFormat)
}

type ConstContext struct {
}

//--------------------------------

type TempDataProxy struct {
	Sheet     *excel.ExcelSheet
	Excel     *excel.ExcelProxy
	Index     []int
	TitleName string
	Language  string
}

func (o *TempDataProxy) GetTitleName(index int) string {
	nameRowIndex := Setting.Excel.Title.NameRow - 1
	value, err := o.Sheet.GetRowAt(nameRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldName Error At %d", index))
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleRemark(index int) string {
	remarkRowIndex := Setting.Excel.Title.RemarkRow - 1
	value, err := o.Sheet.GetRowAt(remarkRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldRemark Error At %d", index))
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleValueFormat(index int) string {
	formatRowIndex := Setting.Excel.Title.FieldFormatRow - 1
	value, err := o.Sheet.GetRowAt(formatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error At %d", index))
		return ""
	}
	ls, ok := Setting.System.FindProgramLanguage(o.Language)
	if !ok {
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error At %d", index))
		return ""
	}
	format, ok := ls.Setting.GetFormat(value)
	if !ok {
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error At %d", index))
		return ""
	}
	return format.Name
}

func (o *TempDataProxy) GetFieldName(index int) string {
	fmt.Println("TempDataProxy.GetFieldName:", index)
	langFormatRowIndex := Setting.Excel.Title.FieldNameRows.GetRowNum(o.Language) - 1
	value, err := o.Sheet.GetRowAt(langFormatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldName Error At %d", index))
		return ""
	}
	return value
}
