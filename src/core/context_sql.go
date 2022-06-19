package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strings"
)

// Sql上下文
type SqlContext struct {
	RangeName string         // 使用的字段索引名称
	RangeType FieldRangeType // 使用的字段索引
	TitleOn   bool           // 表定义Sql是否启用
	DataOn    bool           // 数据Sql是否启用
	SqlMerge  bool           // 合并Sql文件
}

func (o SqlContext) String() string {
	return fmt.Sprintf("SqlContext(TitleOn=%v, DataOn=%v, SqlMerge=%v)", o.TitleOn, o.DataOn, o.SqlMerge)
}

// 字段定义
type FieldItem struct {
	FieldName string              // 字段名称
	FieldType string              // 字段类型(原始未标准化)
	DbField   setting.DbFieldType // 数据库字段描述
}

func (i FieldItem) String() string {
	return fmt.Sprintf("FieldItem{Name=%v, Type=%v, FieldType=%v}", i.FieldName, i.FieldType, i.DbField)
}

// 单条数据项
type SqlItem struct {
	excelRow   *excel.ExcelRow // 数据行
	selects    []int           // 数据行选择索引
	fieldItems []FieldItem     // 数据行选择索引对应的字段定义
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

func (i *SqlItem) GetValue(localIndex int) string {
	index := i.selects[localIndex]
	value, _ := i.excelRow.ValueAtIndex(index)
	return value
}

func (i *SqlItem) GetSqlValue(localIndex int) string {
	value := i.GetValue(localIndex)
	fieldItem := i.fieldItems[localIndex]
	if fieldItem.DbField.IsNumber {
		return value
	}
	value = strings.ReplaceAll(value, `'`, `''`)
	value = fmt.Sprintf(`'%s'`, value)
	//fmt.Println("更新:", localIndex, sqlDataType.IsNumber, value)
	return value
}

// Sql模板代理
type TempSqlProxy struct {
	Sheet      *excel.ExcelSheet // 当前Sheet实例
	Excel      *excel.ExcelProxy // Excel实例代理
	SqlCtx     *SqlContext       // Sql上下文
	TableName  string            // 数据库表名
	FieldIndex []int             // 字段选择索引
	StartRow   int               // 开始行号
	EndRow     int               // 结束行号

	fieldItems []FieldItem // 字段选择索引对应的字段定义
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

func (o *TempSqlProxy) GetFieldItems() []FieldItem {
	o.updateFieldItems()
	return o.fieldItems
}

func (o *TempSqlProxy) GetFieldItem(realIndex int) FieldItem {
	localIndex, _ := slicex.IndexInt(o.FieldIndex, realIndex)
	return o.GetFieldItemLocal(localIndex)
}

func (o *TempSqlProxy) GetFieldItemLocal(localIndex int) FieldItem {
	o.updateFieldItems()
	return o.fieldItems[localIndex]
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
	o.updateFieldItems()
	excelRow := o.Sheet.GetRowAt(row - 1)
	selects := o.FieldIndex
	return SqlItem{excelRow: excelRow, selects: selects, fieldItems: o.fieldItems}, nil
}

func (o *TempSqlProxy) checkItemRow(row int) bool {
	return row >= o.StartRow && row < o.EndRow
}

func (o *TempSqlProxy) updateFieldItems() {
	if nil != o.fieldItems {
		return
	}
	sqlNameRow := o.Sheet.GetRowAt(Setting.Excel.TitleData.GetFileKeyRow(setting.FileNameSql) - 1)
	formatRow := o.Sheet.GetRowAt(Setting.Excel.TitleData.FieldFormatRow - 1)
	o.fieldItems = make([]FieldItem, len(o.FieldIndex))
	db, _ := Setting.System.GetDatabase()
	types, _ := db.GetDataTypes()
	for index, realIndex := range o.FieldIndex {
		fieldName, _ := sqlNameRow.ValueAtIndex(realIndex)
		fieldType, _ := formatRow.ValueAtIndex(realIndex)
		formattedType := o.formatFieldType(fieldType)
		sqlDataType, _ := types.GetFieldType(formattedType)
		o.fieldItems[index] = FieldItem{FieldName: fieldName, FieldType: fieldType, DbField: sqlDataType}
	}
}

func (o *TempSqlProxy) formatFieldType(fieldType string) string {
	if !setting.RegFixedString.MatchString(fieldType) {
		return fieldType
	}
	return setting.RegFixedString.ReplaceAllString(fieldType, "string(*)")
}
