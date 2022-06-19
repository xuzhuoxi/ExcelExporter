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
	RangeName       string         // 使用的字段索引名称
	RangeType       FieldRangeType // 使用的字段索引
	ProgramLanguage string         // 使用的编程语言
}

func (o ConstContext) String() string {
	return fmt.Sprintf("ConstContext(RangeName=%s, RangeType=%v, ProgramLanguage=%s)",
		o.RangeName, o.RangeType, o.ProgramLanguage)
}

// 常量数据
type ConstItem struct {
	Name   string // Excel表格中常量名称
	Value  string // Excel表格中常量值
	Type   string // Excel表格中常量类型
	Remark string // Excel表格中常量备注内容
}

// 常量表模板代理
type TempConstProxy struct {
	Sheet     *excel.ExcelSheet // 当前执行的Sheet数据对象
	Excel     *excel.ExcelProxy // 当前Excel代理，可能包含多个Excel
	ConstCtx  *ConstContext     // 当前执行的上下文数据
	FileName  string            // 导出文件名
	ClassName string            // 导出常量类名
	Language  string            // 导出对应的编程语言
	StartRow  int               // 数据开始行号
	EndRow    int               // 数据结束行号
}

// 取当前Sheet中对应坐标的字符数据，若数据不存在，返回空字符串
func (o *TempConstProxy) ValueAtAxis(axis string) string {
	value, err := o.Sheet.ValueAtAxis(axis)
	if nil != err {
		return ""
	}
	return value
}

// 取当前Sheet全部常量数据列表(已经过滤中间的空行)
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

// 取当前Sheet指定行号数据，转换为常量项，格式非法则返回对应错误
func (o *TempConstProxy) GetItem(row int) (item ConstItem, err error) {
	//fmt.Println("GetItem:", row)
	if !o.checkItemRow(row) {
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

func (o *TempConstProxy) checkItemRow(row int) bool {
	return row >= o.StartRow && row < o.EndRow
}
