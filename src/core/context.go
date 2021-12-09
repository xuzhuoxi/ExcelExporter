package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
)

type TitleContext struct {
	// 使用的字段索引名称
	RangeName string
	// 使用的字段索引
	RangeType FieldRangeType
	// 使用的编程语言
	ProgramLanguage string
}

func (o TitleContext) String() string {
	return fmt.Sprintf("TitleContent(RangeName=%s, RangeType=%v, ProgramLanguage=%s)",
		o.RangeName, o.RangeType, o.ProgramLanguage)
}

type DataContext struct {
	// 使用的字段索引名称
	RangeName string
	// 使用的字段索引
	RangeType FieldRangeType
	// 输出的文件类型
	DataFileFormat string
}

func (o DataContext) String() string {
	return fmt.Sprintf("DataContext(RangeName=%s, RangeType=%v, ProgramLanguage=%s)",
		o.RangeName, o.RangeType, o.DataFileFormat)
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
	//nameRowIndex := Setting.Excel.Title.NameRow - 1
	nameRowIndex := Setting.Excel.TitleData.NameRow - 1
	value, err := o.Sheet.GetRowAt(nameRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldName Error At %d", index))
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleRemark(index int) string {
	//remarkRowIndex := Setting.Excel.Title.RemarkRow - 1
	remarkRowIndex := Setting.Excel.TitleData.RemarkRow - 1
	value, err := o.Sheet.GetRowAt(remarkRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldRemark Error At %d", index))
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleLangDefine(index int) setting.FieldOperate {
	//formatRowIndex := Setting.Excel.Title.FieldFormatRow - 1
	formatRowIndex := Setting.Excel.TitleData.FieldFormatRow - 1
	value, err := o.Sheet.GetRowAt(formatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error At %d: %v", index, err))
		return setting.FieldOperate{}
	}
	ls, ok := Setting.System.FindProgramLanguage(o.Language)
	if !ok {
		err = errors.New(fmt.Sprintf("Find Program Language Fail At %d ", index))
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error At %d: %v", index, err))
		return setting.FieldOperate{}
	}
	value = setting.FormatStringField(value)
	format, ok := ls.Setting.GetLangDefine(value)
	if !ok {
		err = errors.New(fmt.Sprintf("Get Lang Define Fail At %d, %s ", index, value))
		Logger.Error(fmt.Sprintf("GetFieldValueFormat Error At %d: %v", index, err))
		return setting.FieldOperate{}
	}
	return format
}

func (o *TempDataProxy) GetFieldName(index int) string {
	return o.GetTitleLangKey(index, o.Language)
}

func (o *TempDataProxy) GetTitleLangKey(index int, langName string) string {
	//fmt.Println("TempDataProxy.GetFieldName:", index)
	//langFormatRowIndex := Setting.Excel.Title.LangKeyRows.GetRowNum(langName) - 1
	langFormatRowIndex := Setting.Excel.TitleData.GetFieldNameRow(langName) - 1
	value, err := o.Sheet.GetRowAt(langFormatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetTitleLangKey Error At [%s,%d]", langName, index))
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleFileKey(index int, fileType string) string {
	//fmt.Println("TempDataProxy.GetTitleFileKey:", fileType, index)
	//langFormatRowIndex := Setting.Excel.Title.FileKeyRows.GetRowNum(fileType) - 1
	langFormatRowIndex := Setting.Excel.TitleData.GetFileKeyRow(fileType) - 1
	value, err := o.Sheet.GetRowAt(langFormatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetTitleFileKey Error At [%s,%d]", fileType, index))
		return ""
	}
	return value
}
