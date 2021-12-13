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
	// 使用的字段索引名称
	RangeName string
	// 使用的字段索引
	RangeType FieldRangeType
	// 使用的编程语言
	ProgramLanguage string
}

func (o ConstContext) String() string {
	return fmt.Sprintf("ConstContext(RangeName=%s, RangeType=%v, ProgramLanguage=%s)",
		o.RangeName, o.RangeType, o.ProgramLanguage)
}

//--------------------------------

type ConstItem struct {
	Name   string
	Value  string
	Type   string
	Remark string
}

type TempConstProxy struct {
	Sheet     *excel.ExcelSheet
	Excel     *excel.ExcelProxy
	TitleName string
	Language  string
	StartRow  int
	EndRow    int
}

func (o *TempConstProxy) ValueAtAxis(axis string) string {
	value, err := o.Sheet.ValueAtAxis(axis)
	fmt.Println("TempConstProxy.ValueAtAxis", value)
	if nil != err {
		return ""
	}
	return value
}

func (o *TempConstProxy) GetItems() []ConstItem {
	if o.EndRow-o.StartRow <= 0 {
		return nil
	}
	rs := make([]ConstItem, 0, o.EndRow-o.StartRow+1)
	for row := o.StartRow; row <= o.EndRow; row += 1 {
		rs = append(rs, o.GetItem(row))
	}
	return rs
}

func (o *TempConstProxy) GetItem(row int) (item ConstItem) {
	excelRow := o.Sheet.GetRowAt(row - 1)
	name, _ := excelRow.ValueAtAxis(Setting.Excel.Const.NameCol)
	remark, _ := excelRow.ValueAtAxis(Setting.Excel.Const.RemarkCol)
	tp, _ := excelRow.ValueAtAxis(Setting.Excel.Const.TypeCol)
	ld, ok := Setting.System.FindProgramLanguage(o.Language)
	err := errors.New(fmt.Sprintf("Const Item Type Error At Row %d ", row))
	if !ok {
		Logger.Error(err)
		return
	}
	typeFormat, ok := ld.Setting.GetLangDefine(tp)
	if !ok {
		Logger.Error(err)
		return
	}
	value, _ := excelRow.ValueAtAxis(Setting.Excel.Const.ValueCol)
	return ConstItem{Name: name, Type: typeFormat.Name, Value: value, Remark: remark}
}

//--------------------------------

type TempDataProxy struct {
	Sheet     *excel.ExcelSheet
	Excel     *excel.ExcelProxy
	Index     []int
	TitleName string
	Language  string
}

func (o *TempDataProxy) ValueAtAxis(axis string) string {
	value, err := o.Sheet.ValueAtAxis(axis)
	if nil != err {
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleName(index int) string {
	nameRowIndex := Setting.Excel.TitleData.NameRow - 1
	value, err := o.Sheet.GetRowAt(nameRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldName Error At %d", index))
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleRemark(index int) string {
	remarkRowIndex := Setting.Excel.TitleData.RemarkRow - 1
	value, err := o.Sheet.GetRowAt(remarkRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldRemark Error At %d", index))
		return ""
	}
	return value
}

func (o *TempDataProxy) GetTitleLangDefine(index int) setting.FieldOperate {
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
	langFormatRowIndex := Setting.Excel.TitleData.GetFileKeyRow(fileType) - 1
	value, err := o.Sheet.GetRowAt(langFormatRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetTitleFileKey Error At [%s,%d]", fileType, index))
		return ""
	}
	return value
}
