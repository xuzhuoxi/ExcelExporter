package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/slicex"
	"github.com/xuzhuoxi/infra-go/stringx"
	"math"
	"strconv"
	"strings"
)

// Sql上下文
type SqlContext struct {
	EnablePrefix  string         // 开启前缀
	RangeName     string         // 使用的字段索引名称
	RangeType     FieldRangeType // 使用的字段索引
	TitleOn       bool           // 表定义Sql是否启用
	DataOn        bool           // 数据Sql是否启用
	SqlMerge      bool           // 合并Sql文件
	StartRowNum   int            // 数据开始行号
	StartColIndex int            // 数据开始列索引
}

func (o SqlContext) String() string {
	return fmt.Sprintf("SqlContext(Prefix=%s, TitleOn=%v, DataOn=%v, SqlMerge=%v, StartRowNum=%d, StartColIndex=%d)",
		o.EnablePrefix, o.TitleOn, o.DataOn, o.SqlMerge, o.StartRowNum, o.StartColIndex)
}

// 字段定义
type FieldItem struct {
	FieldName       string                  // 字段名称
	FieldType       string                  // 字段类型(原始未标准化)
	CustomFieldType string                  // 定制字段类型
	DatabaseExtend  *setting.DatabaseExtend // 数据库扩展定义
	DbField         setting.DbFieldType     // 数据库字段描述
	MaxByteSize     int                     // 最大utf-8字节数
	MaxRuneSize     int                     // 最大字符数
}

// 是否为定制类型
func (i FieldItem) IsCustomFieldType() bool {
	return "" != i.CustomFieldType
}

// 是否为数据库动态类型
func (i FieldItem) IsDynamicDbFieldType() bool {
	return i.DbField.IsDynamicLen()
}

// 是否为字段固定类型
func (i FieldItem) IsFixedFieldType() bool {
	return setting.RegFixedString.MatchString(i.FieldType)
}

func (i FieldItem) SqlFieldType() string {
	// 有定制类型
	if i.IsCustomFieldType() {
		return i.CustomFieldType
	}
	// 数据库类型无动态长度"*"，全部数值类型包含在内
	if !i.IsDynamicDbFieldType() {
		return i.DbField.FieldType
	}
	// 字段类型为string(n)
	if i.IsFixedFieldType() {
		size, _ := i.getFixedFieldTypeSize()
		return strings.ReplaceAll(i.DbField.FieldType, "*", strconv.Itoa(size))
	}

	// 动态计算
	size := i.getDynamicStatSize()
	return strings.ReplaceAll(i.DbField.FieldType, "*", strconv.Itoa(size))
}

func (i FieldItem) getFixedFieldTypeSize() (size int, err error) {
	start := strings.LastIndex(i.FieldType, "(")
	end := strings.LastIndex(i.FieldType, ")")
	size, err = strconv.Atoi(i.FieldType[start+1 : end])
	if nil != err {
		Logger.Warnln("FieldItem.getDynamicFieldTypeSize:", err)
	}
	if i.DbField.IsDynamicChar() {
		size = int(math.Ceil(i.DatabaseExtend.ScaleChar * float64(size)))
	} else if i.DbField.IsDynamicVarchar() {
		size = int(math.Ceil(i.DatabaseExtend.ScaleVarchar * float64(size)))
	}
	return
}

func (i FieldItem) getDynamicStatSize() int {
	size := i.MaxByteSize
	if i.DbField.IsDynamicChar() {
		size = int(math.Ceil(i.DatabaseExtend.ScaleChar * float64(i.MaxRuneSize)))
	} else if i.DbField.IsDynamicVarchar() {
		size = int(math.Ceil(i.DatabaseExtend.ScaleVarchar * float64(i.MaxRuneSize)))
	}
	return size
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
	Sheet         *excel.ExcelSheet // 当前执行的Sheet数据对象
	Excel         *excel.ExcelProxy // 当前执行的Excel数据代理对象
	SqlCtx        *SqlContext       // 当前执行的Sql上下文
	TableName     string            // 数据库表名
	FieldIndex    []int             // 字段选择索引
	StartRow      int               // 开始行号
	EndRow        int               // 结束行号
	StartColIndex int               //开始列索引

	fieldItems []FieldItem // 字段选择索引对应的字段定义
	primaryKey []FieldItem // 主键信息
}

func (o *TempSqlProxy) MergeOn() bool {
	return o.SqlCtx.SqlMerge
}

func (o *TempSqlProxy) NeedTruncateData() bool {
	return !o.SqlCtx.SqlMerge || !o.SqlCtx.TitleOn
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
	o.initFieldItems()
	return o.fieldItems
}

func (o *TempSqlProxy) GetFieldItem(realIndex int) FieldItem {
	localIndex, _ := slicex.IndexInt(o.FieldIndex, realIndex)
	return o.GetFieldItemLocal(localIndex)
}

func (o *TempSqlProxy) GetFieldItemLocal(localIndex int) FieldItem {
	o.initFieldItems()
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

func (o *TempSqlProxy) PrimaryKeyLen() int {
	o.initFieldItems()
	return len(o.primaryKey)
}

func (o *TempSqlProxy) GetPrimaryKeys() []FieldItem {
	o.initFieldItems()
	return o.primaryKey
}

func (o *TempSqlProxy) GetItem(row int) (item SqlItem, err error) {
	if !o.checkItemRow(row) {
		err = errors.New(fmt.Sprintf("Row[%d] out of range. ", row))
		return
	}
	o.initFieldItems()
	excelRow := o.Sheet.GetRowAt(row - 1)
	selects := o.FieldIndex
	return SqlItem{excelRow: excelRow, selects: selects, fieldItems: o.fieldItems}, nil
}

func (o *TempSqlProxy) checkItemRow(row int) bool {
	return row >= o.StartRow && row < o.EndRow
}

func (o *TempSqlProxy) initFieldItems() {
	if nil != o.fieldItems {
		return
	}
	sqlNameRow := o.Sheet.GetRowAt(Setting.Excel.TitleData.GetFileKeyRow(setting.FileNameSql) - 1)
	formatRow := o.Sheet.GetRowAt(Setting.Excel.TitleData.FieldFormatRow - 1)
	isCustomSqlField := Setting.Excel.TitleData.IsCustomSqlFieldType()
	o.fieldItems = make([]FieldItem, len(o.FieldIndex))
	db, _ := Setting.System.GetDatabase()
	extend, _ := db.GetDatabaseExtend()
	for index, realIndex := range o.FieldIndex {
		fieldName, _ := sqlNameRow.ValueAtIndex(realIndex)
		fieldType, _ := formatRow.ValueAtIndex(realIndex)
		formattedType := o.formatFieldType(fieldType)
		sqlDataType, _ := extend.GetFieldType(formattedType)

		if isCustomSqlField { // 定制类型处理
			customFieldType, _ := o.Sheet.ValueAtIndex(realIndex, Setting.Excel.TitleData.SqlFieldFormatRow-1)
			o.fieldItems[index] = FieldItem{
				FieldName: fieldName, FieldType: fieldType, CustomFieldType: strings.ToUpper(customFieldType),
				DatabaseExtend: extend, DbField: sqlDataType}
			continue
		}

		if sqlDataType.IsDynamicLen() { // 动态长度处理
			byteSize, runeSize := o.getMaxSize(realIndex)
			o.fieldItems[index] = FieldItem{
				FieldName: fieldName, FieldType: fieldType,
				DatabaseExtend: extend, DbField: sqlDataType,
				MaxByteSize: byteSize, MaxRuneSize: runeSize}
			continue
		}
		// 固定长度处理
		o.fieldItems[index] = FieldItem{
			FieldName: fieldName, FieldType: fieldType,
			DatabaseExtend: extend, DbField: sqlDataType}
	}
	o.initPrimaryKey()
}

func (o *TempSqlProxy) getMaxSize(colRealIndex int) (byteSize int, runeSize int) {
	for rowIndex := o.StartRow - 1; rowIndex < o.EndRow-1; rowIndex += 1 {
		value, err := o.Sheet.ValueAtIndex(colRealIndex, rowIndex)
		if nil != err {
			continue
		}
		byteSize = mathx.MaxInt(byteSize, len(value))
		runeSize = mathx.MaxInt(runeSize, stringx.GetRuneCount(value))
	}
	return
}

func (o *TempSqlProxy) initPrimaryKey() {
	axis := Setting.Excel.TitleData.Sql.PrimaryKeyAxis
	value, err := o.Sheet.ValueAtAxis(axis)
	if nil != err {
		return
	}
	value = strings.TrimSpace(value)
	if "" == value {
		return
	}
	arr := strings.Split(value, ",")
	key := make([]FieldItem, len(arr))
	for index := range arr {
		fieldCol := mathx.System26To10(arr[index])
		localIndex, _ := slicex.IndexInt(o.FieldIndex, fieldCol-1)
		if -1 == localIndex {
			Logger.Warnln(fmt.Sprintf("TempSqlProxy.initPrimaryKey: PrimaryKey Config Error At [%s:%s]", axis, value))
			continue
		}
		fieldItem := o.GetFieldItemLocal(localIndex)
		key[index] = fieldItem
	}
	o.primaryKey = key
}

func (o *TempSqlProxy) formatFieldType(fieldType string) string {
	if !setting.RegFixedString.MatchString(fieldType) {
		return fieldType
	}
	return setting.RegFixedString.ReplaceAllString(fieldType, "string(*)") // string(*) || []string(*)
}
