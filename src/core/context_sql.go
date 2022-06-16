package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"strings"
)

// Sql上下文
type SqlContext struct {
	// 使用的字段索引名称
	RangeName string
	// 使用的字段索引
	RangeType FieldRangeType
	// 表定义Sql是否启用
	TitleOn bool
	// 数据Sql是否启用
	DataOn bool
	// 合并Sql文件
	SqlMerge bool
}

func (o SqlContext) String() string {
	return fmt.Sprintf("SqlContext(TitleOn=%v, DataOn=%v, SqlMerge=%v)", o.TitleOn, o.DataOn, o.SqlMerge)
}

// 单条数据项
type SqlItem struct {
	excelRow   *excel.ExcelRow
	selects    []int
	fieldTypes []string
}

func (i *SqlItem) FieldLen() int {
	return len(i.selects)
}

func (i *SqlItem) GetValues() []string {
	rs := make([]string, 0, len(i.selects))
	for index := range i.selects {
		value := i.GetSqlValue(index)
		rs = append(rs, value)
	}
	return rs
}

func (i *SqlItem) GetValue(selectIndex int) string {
	index := i.selects[selectIndex]
	value, _ := i.excelRow.ValueAtIndex(index)
	return value
}

func (i *SqlItem) GetSqlValue(selectIndex int) string {
	fieldType := i.fieldTypes[selectIndex]
	db, _ := Setting.System.GetDatabase()
	types, _ := db.GetDataTypes()
	value := i.GetValue(selectIndex)
	sqlDataType, _ := types.GetType(fieldType)
	//fmt.Println("GetSqlValue:", selectIndex, sqlDataType.IsNumber, value)
	if sqlDataType.IsNumber {
		return value
	}
	value = strings.ReplaceAll(value, `"`, `\"`)
	value = fmt.Sprintf("\"%s\"", value)
	//fmt.Println("更新:", selectIndex, sqlDataType.IsNumber, value)
	return value
}

// Sql模板代理
type TempSqlProxy struct {
	Sheet      *excel.ExcelSheet
	Excel      *excel.ExcelProxy
	SqlCtx     *SqlContext
	TableName  string
	FieldIndex []int
	StartRow   int
	EndRow     int

	fieldValue []string
}

func (o *TempSqlProxy) FieldLen() int {
	return len(o.FieldIndex)
}

func (o *TempSqlProxy) ItemLen() int {
	return o.EndRow - o.StartRow
}

func (o *TempSqlProxy) ValueAtAxis(axis string) string {
	value, err := o.Sheet.ValueAtAxis(axis)
	if nil != err {
		return ""
	}
	return value
}

func (o *TempSqlProxy) GetFieldName(index int) string {
	nameRowIndex := Setting.Excel.TitleData.GetFileKeyRow(setting.FileNameSql) - 1
	value, err := o.Sheet.GetRowAt(nameRowIndex).ValueAtIndex(index)
	if err != nil {
		Logger.Error(fmt.Sprintf("GetFieldName Error At %d", index))
		return ""
	}
	return value
}

func (o *TempSqlProxy) GetItems() []SqlItem {
	capRow := o.ItemLen()
	if capRow <= 0 {
		return nil
	}
	rs := make([]SqlItem, 0, capRow)
	for row := o.StartRow; row <= o.EndRow; row += 1 {
		item, err := o.GetItem(row)
		if nil != err {
			continue
		}
		rs = append(rs, item)
	}
	return rs
}

func (o *TempSqlProxy) GetItem(row int) (item SqlItem, err error) {
	if !o.checkItemRow(row) {
		err = errors.New(fmt.Sprintf("Row[%d] out of range. ", row))
		return
	}
	excelRow := o.Sheet.GetRowAt(row - 1)
	selects := o.FieldIndex
	fieldTypes := o.getFieldTypes()
	return SqlItem{excelRow: excelRow, selects: selects, fieldTypes: fieldTypes}, nil
}

func (o *TempSqlProxy) checkItemRow(row int) bool {
	return row >= o.StartRow && row < o.EndRow
}

func (o *TempSqlProxy) getFieldTypes() []string {
	if nil != o.fieldValue {
		return o.fieldValue
	}

	row := o.Sheet.GetRowAt(Setting.Excel.TitleData.FieldFormatRow - 1)
	rs := make([]string, len(o.FieldIndex))
	for index := range rs {
		fieldType, _ := row.ValueAtIndex(o.FieldIndex[index])
		rs[index] = o.formatFieldType(fieldType)
	}
	o.fieldValue = rs
	return rs
}

func (o *TempSqlProxy) formatFieldType(fieldType string) string {
	if !setting.RegFixedString.MatchString(fieldType) {
		return fieldType
	}
	return setting.RegFixedString.ReplaceAllString(fieldType, "string(*)")
}
