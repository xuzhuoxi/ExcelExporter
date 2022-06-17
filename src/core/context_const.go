package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"strings"
)

// 常量表上下文
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

// 常量数据
type ConstItem struct {
	Name   string
	Value  string
	Type   string
	Remark string
}

// 常量表模板代理
type TempConstProxy struct {
	Sheet     *excel.ExcelSheet
	Excel     *excel.ExcelProxy
	ConstCtx  *ConstContext
	FileName  string
	ClassName string
	Language  string
	StartRow  int
	EndRow    int
}

func (o *TempConstProxy) ValueAtAxis(axis string) string {
	value, err := o.Sheet.ValueAtAxis(axis)
	if nil != err {
		return ""
	}
	return value
}

func (o *TempConstProxy) GetItems() []ConstItem {
	capRow := o.EndRow - o.StartRow
	if capRow <= 0 {
		return nil
	}
	rs := make([]ConstItem, 0, capRow)
	for row := o.StartRow; row < o.EndRow; row += 1 {
		item, err := o.GetItem(row)
		if nil != err {
			continue
		}
		rs = append(rs, item)
	}
	return rs
}

func (o *TempConstProxy) GetItem(row int) (item ConstItem, err error) {
	//fmt.Println("GetItem:", row)
	if !o.CheckItemRow(row) {
		err = errors.New(fmt.Sprintf("Row[%d] out of range. ", row))
		return
	}
	excelRow := o.Sheet.GetRowAt(row - 1)
	name, err2 := excelRow.ValueAtAxis(Setting.Excel.Const.NameCol)
	if nil != err2 {
		err = err2
		return
	}
	if len(strings.TrimSpace(name)) == 0 {
		err = errors.New(fmt.Sprintf("Empty Row At %d. ", row))
		return
	}
	remark, _ := excelRow.ValueAtAxis(Setting.Excel.Const.RemarkCol)
	tp, _ := excelRow.ValueAtAxis(Setting.Excel.Const.TypeCol)
	ld, ok := Setting.System.FindProgramLanguage(o.Language)
	//fmt.Println("行:", o.Language, tp, ld, excelRow.Cell)
	err = errors.New(fmt.Sprintf("Const Item Type Error At Row %d ", row))
	if !ok {
		return
	}
	typeFormat, ok := ld.Setting.GetLangDefine(tp)
	if !ok {
		return
	}
	value, _ := excelRow.ValueAtAxis(Setting.Excel.Const.ValueCol)
	if tp == setting.FieldString {
		value = fmt.Sprintf("\"%s\"", value)
	}
	return ConstItem{Name: name, Type: typeFormat.LangTypeName, Value: value, Remark: remark}, nil
}

func (o *TempConstProxy) CheckItemRow(row int) bool {
	return row >= o.StartRow && row < o.EndRow
}
