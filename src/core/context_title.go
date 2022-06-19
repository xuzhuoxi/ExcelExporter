package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
)

// 表头导出上下文
type TitleContext struct {
	RangeName       string         // 使用的字段索引名称
	RangeType       FieldRangeType // 使用的字段索引
	ProgramLanguage string         // 使用的编程语言
}

func (o TitleContext) String() string {
	return fmt.Sprintf("TitleContent(RangeName=%s, RangeType=%v, ProgramLanguage=%s)",
		o.RangeName, o.RangeType, o.ProgramLanguage)
}

// 数据表代理
type TempTitleProxy struct {
	Sheet      *excel.ExcelSheet // 当前执行的Sheet数据对象
	Excel      *excel.ExcelProxy // 当前Excel代理，可能包含多个Excel
	TitleCtx   *TitleContext     // 当前执行的表头上下文数据
	FileName   string            // 表头导出类文件名
	ClassName  string            // 表头导出类名
	FieldIndex []int             // 当前选中的字段索引
	Language   string            // 当前的选择的编程语言
}

func (o *TempTitleProxy) ValueAtAxis(axis string) string {
	value, err := o.Sheet.ValueAtAxis(axis)
	if nil != err {
		return ""
	}
	return value
}

func (o *TempTitleProxy) FieldLen() int {
	return len(o.FieldIndex)
}

func (o *TempTitleProxy) GetTitleName(index int) string {
	nameRowIndex := Setting.Excel.TitleData.NameRow - 1
	value, err := o.Sheet.GetRowAt(nameRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldName Error At %d", index))
		return ""
	}
	return value
}

func (o *TempTitleProxy) GetTitleRemark(index int) string {
	remarkRowIndex := Setting.Excel.TitleData.RemarkRow - 1
	value, err := o.Sheet.GetRowAt(remarkRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldRemark Error At %d", index))
		return ""
	}
	return value
}

func (o *TempTitleProxy) GetTitleLangDefine(index int) setting.LangDataType {
	formatRowIndex := Setting.Excel.TitleData.FieldFormatRow - 1
	value, err := o.Sheet.GetRowAt(formatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error1 At %d: %v", index, err))
		return setting.LangDataType{}
	}
	ls, ok := Setting.System.FindProgramLanguage(o.Language)
	if !ok {
		err = errors.New(fmt.Sprintf("Find Program Language Fail At %d ", index))
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error2 At %d: %v", index, err))
		return setting.LangDataType{}
	}
	value = o.formatFieldType(value)
	format, ok := ls.Setting.GetLangDefine(value)
	if !ok {
		err = errors.New(fmt.Sprintf("Get Lang Define Fail At %d, %s ", index, value))
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error3 At %d: %v", index, err))
		return setting.LangDataType{}
	}
	return format
}

func (o *TempTitleProxy) GetFieldName(index int) string {
	return o.GetTitleLangKey(index, o.Language)
}

func (o *TempTitleProxy) GetTitleLangKey(index int, langName string) string {
	//fmt.Println("TempTitleProxy.GetFieldName:", index)
	langFormatRowIndex := Setting.Excel.TitleData.GetFieldNameRow(langName) - 1
	value, err := o.Sheet.GetRowAt(langFormatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetTitleLangKey Error At [%s,%d]", langName, index))
		return ""
	}
	return value
}

func (o *TempTitleProxy) GetTitleFileKey(index int, fileType string) string {
	//fmt.Println("TempTitleProxy.GetTitleFileKey:", fileType, index)
	langFormatRowIndex := Setting.Excel.TitleData.GetFileKeyRow(fileType) - 1
	value, err := o.Sheet.GetRowAt(langFormatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetTitleFileKey Error At [%s,%d]", fileType, index))
		return ""
	}
	return value
}

func (o *TempTitleProxy) formatFieldType(fieldValue string) string {
	if !setting.RegFixedString.MatchString(fieldValue) {
		return fieldValue
	}
	return setting.RegFixedString.ReplaceAllString(fieldValue, "string")
}
