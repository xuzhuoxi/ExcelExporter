// Package core
// Create on 2023/5/21
// @author xuzhuoxi
package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strconv"
	"strings"
)

// ProtoContext 协议表上下文
type ProtoContext struct {
	EnablePrefix string         // 开启前缀
	RangeName    string         // 使用的字段索引名称
	RangeType    FieldRangeType // 使用的字段索引
	Language     string         // 使用的编程语言
}

func (o ProtoContext) String() string {
	return fmt.Sprintf("ProtoContext(Prefix=%s, RangeName=%s, RangeType=%v, Language=%s)",
		o.EnablePrefix, o.RangeName, o.RangeType, o.Language)
}

// ProtoSheetTitle 协议表关定义
type ProtoSheetTitle struct {
	IdDataType   string   // Id数据类型
	RangeName    []string // 导出范围
	Namespace    string   // 命名空间(包名)
	ExportSubDir string   // 额外的导出子目录
}

func (o ProtoSheetTitle) MatchRange(rangeName string) bool {
	rangeName = strings.ToLower(rangeName)
	if len(o.RangeName) == 0 {
		return false
	}
	return slicex.ContainsString(o.RangeName, rangeName)
}

// ProtoFieldItem 协议属性
type ProtoFieldItem struct {
	Remark        string // 备注
	Name          string // 属性Key：Excel配置值
	Lang          string // 编程语言
	OriginalType  string // 属性数据类型：原始值
	FormattedType string // 属性数据类型：格式化值
	IsPointer     bool   // 属性数据类型：是否为指针类型
	IsArray       bool   // 属性数据类型：是否为自定义数组
	ArraySize     int    // 数组长度，-1 代表无固定谎称

	LangType       string               // 属性数据类型：编程语言值
	LangTypeDefine setting.LangDataType // 属性数据类型：编程语言定义
	IsCustomType   bool                 // 属性数据类型：是否为自定义类型
}

func (o ProtoFieldItem) TempLangType() string {
	if !o.IsCustomType {
		return o.LangType
	}
	langDefine, ok := Setting.System.FindProgramLanguage(o.Lang)
	if !ok {
		return o.LangType
	}
	custom := langDefine.Setting.Custom
	if !o.IsArray {
		return custom.ToLangType(o.LangType, o.IsPointer)
	}
	return custom.ToLangArrayType(o.LangType, o.IsPointer)
}

// ProtoItem 协议项
type ProtoItem struct {
	Id     string           // 协议Id
	Name   string           // 协议名称
	Remark string           // 协议备注
	File   string           // 协议文件名
	Fields []ProtoFieldItem // 协议属性
}

// TempProtoProxy 协议表模板代理
type TempProtoProxy struct {
	ProtoItem  ProtoItem        // 协议信息
	SheetProxy *ProtoSheetProxy // Sheet表信息代理
}

func (o *TempProtoProxy) ValueAtAxis(axis string) string {
	return o.SheetProxy.ValueAtAxis(axis)
}

func (o *TempProtoProxy) Namespace() string {
	return o.SheetProxy.Title.Namespace
}

func (o *TempProtoProxy) ProtoId() string {
	if setting.FieldString == o.SheetProxy.Title.IdDataType {
		return fmt.Sprintf("\"%s\"", o.ProtoItem.Id)
	}
	return o.ProtoItem.Id
}

func (o *TempProtoProxy) ProtoIdDataType() string {
	dtd, _ := o.SheetProxy.GetDataTypeDefine(o.SheetProxy.Title.IdDataType)
	return dtd.LangTypeName
}

func (o *TempProtoProxy) ClassName() string {
	return o.ProtoItem.Name
}

func (o *TempProtoProxy) ClassRemark() string {
	return o.ProtoItem.Remark
}

func (o *TempProtoProxy) GetFields() []ProtoFieldItem {
	return o.ProtoItem.Fields
}

type ProtoSheetProxy struct {
	Excel    *excel.ExcelProxy // 当前Excel代理，可能包含多个Excel
	Sheet    *excel.ExcelSheet // 当前执行的Sheet数据对象
	ProtoCtx *ProtoContext     // 当前执行的上下文数据
	Title    ProtoSheetTitle   // 表头信息

	protoNames []string // 协议定义列表
}

// ValueAtAxis 取当前Sheet中对应坐标的字符数据，若数据不存在，返回空字符串
func (o *ProtoSheetProxy) ValueAtAxis(axis string) string {
	value, err := o.Sheet.ValueAtAxis(axis)
	if nil != err {
		return ""
	}
	return value
}

func (o *ProtoSheetProxy) GetDataTypeDefine(dataType string) (define setting.LangDataType, ok bool) {
	ls, _ := Setting.System.FindProgramLanguage(o.ProtoCtx.Language)
	return ls.Setting.GetDataTypeDefine(dataType)
}

// GetItems 取当前Sheet全部协议数据列表(已经过滤中间的空行)
func (o *ProtoSheetProxy) GetItems() (items []ProtoItem, err error) {
	ps := Setting.Excel.Proto
	rowLen := o.Sheet.RowLength
	capRow := rowLen - ps.DataStartRow
	if capRow <= 0 {
		return nil, nil
	}
	rs := make([]ProtoItem, 0, capRow)
	for row := ps.DataStartRow; row <= rowLen; row += 1 {
		item, isItem, isBlank, err := o.GetItem(row)
		if nil != err {
			return rs, err
		}
		if isBlank && ps.BlankLineBreak {
			return rs, nil
		}
		if !isItem {
			continue
		}
		rs = append(rs, item)
	}
	return rs, nil
}

// GetItem 取当前Sheet指定行号数据，转换为协议项，格式非法则返回对应错误
func (o *ProtoSheetProxy) GetItem(row int) (item ProtoItem, isItem bool, isBlank bool, err error) {
	//fmt.Println("GetItem:", row)
	if !o.checkInRange(row) {
		err = errors.New(fmt.Sprintf("Row[%d] out of range. ", row))
		return
	}
	itemRow := o.Sheet.GetRowAt(row - 1)
	if itemRow.CellLength() == 0 {
		isBlank = true
		return
	}
	ps := Setting.Excel.Proto
	startFieldColIndex := ps.StartFieldColIndex()
	if startFieldColIndex >= itemRow.CellLength() {
		return
	}

	id := strings.TrimSpace(itemRow.Cell[ps.IdColIndex()])
	file := strings.TrimSpace(itemRow.Cell[ps.FileColIndex()])
	name := strings.TrimSpace(itemRow.Cell[ps.NameColIndex()])
	fieldSize := o.getFieldSize(itemRow, ps.StartFieldColIndex())
	if !o.checkIsItem(id, file, name) || fieldSize == 0 {
		return
	}

	remarkRow, remarkExist := o.getRemarkRow(row)
	fieldItems, err2 := o.getFieldItems(row, itemRow, remarkRow, fieldSize, remarkExist)
	if nil != err2 {
		err = err2
		return
	}
	nameRemark := ""
	if remarkExist {
		nameRemark = strings.TrimSpace(remarkRow.Cell[ps.NameColIndex()])
	}
	item = ProtoItem{Id: id, File: file, Name: name, Remark: nameRemark, Fields: fieldItems}
	suc := o.appendCustomProtoName(name)
	if !suc {
		err = errors.New(fmt.Sprintf("Duplicate ProtoName[%s] At Row[%d]. ", name, row))
		isItem = true
		return
	}
	return item, true, false, nil

}

func (o *ProtoSheetProxy) checkInRange(row int) bool {
	return row >= Setting.Excel.Proto.DataStartRow && row <= o.Sheet.RowLength
}

func (o *ProtoSheetProxy) checkIsItem(id string, file string, name string) bool {
	if len(file) == 0 || len(name) == 0 {
		return false
	}
	return true
}

// protoName不会是数组
func (o *ProtoSheetProxy) appendCustomProtoName(protoName string) bool {
	if o.containsCustomProtoName(protoName) {
		return false
	}
	o.protoNames = append(o.protoNames, protoName)
	return true
}

func (o *ProtoSheetProxy) containsCustomProtoName(protoName string) (ok bool) {
	if len(protoName) == 0 {
		return
	}
	ok = slicex.ContainsString(o.protoNames, protoName)
	return
}

func (o *ProtoSheetProxy) getFieldSize(excelRow *excel.ExcelRow, fieldStartIdx int) int {
	// len = 7 startIdx = 3	rs = 4
	celLen := excelRow.CellLength() //7
	if celLen <= 0 || celLen <= fieldStartIdx {
		return 0
	}
	for idx := fieldStartIdx; idx < celLen; idx += 1 {
		cellStr := strings.TrimSpace(excelRow.Cell[idx])
		if len(cellStr) == 0 {
			return idx - fieldStartIdx
		}
	}
	return celLen - fieldStartIdx
}

func (o *ProtoSheetProxy) getRemarkRow(row int) (remarkRow *excel.ExcelRow, exist bool) {
	ps := Setting.Excel.Proto
	remarkRowIdx := row + ps.RemarkOffset - 1
	if remarkRowIdx < 0 || remarkRowIdx >= o.Sheet.RowLength {
		return nil, false
	}
	remarkRow = o.Sheet.GetRowAt(remarkRowIdx)
	if nil == remarkRow {
		return nil, false
	}
	id := remarkRow.Cell[ps.IdColIndex()]
	file := remarkRow.Cell[ps.FileColIndex()]
	if len(id) != 0 || len(file) != 0 {
		return nil, false
	}
	return remarkRow, true
}

func (o *ProtoSheetProxy) getFieldItems(rowId int, itemRow, remarkRow *excel.ExcelRow,
	fieldSize int, remarkExist bool) (fieldItems []ProtoFieldItem, err error) {
	ps := Setting.Excel.Proto
	fieldItems = make([]ProtoFieldItem, fieldSize)
	colIdx := ps.StartFieldColIndex()
	idx := 0
	for idx < fieldSize {
		fieldStr := itemRow.Cell[colIdx]
		cellLoc := excel.GetCellName(colIdx+1, rowId)
		dataRemark := ""
		if remarkExist {
			dataRemark = strings.TrimSpace(remarkRow.Cell[colIdx])
		}
		fieldItem, err1 := o.getFieldItem(cellLoc, fieldStr, dataRemark)
		if nil != err1 {
			return nil, err1
		}
		fieldItems[idx] = *fieldItem
		colIdx += 1
		idx += 1
	}
	return fieldItems, nil
}

func (o *ProtoSheetProxy) getFieldItem(loc string, fieldStr string, dataRemark string) (fieldItem *ProtoFieldItem, err error) {
	info := strings.Split(fieldStr, ":")
	if len(info) != 2 {
		return nil,
			errors.New(fmt.Sprintf("ProtoItemField[Loc=%s, Value=\"%s\"] Format Error!, ",
				loc, fieldStr))
	}
	pointerCode := Setting.System.PointerCode
	dataName := strings.TrimSpace(info[0])
	originalType := strings.TrimSpace(info[1])
	if len(dataName) == 0 || len(originalType) == 0 {
		return nil,
			errors.New(fmt.Sprintf("ProtoItemField[Loc=%s, Value=\"%s\"] Field Type Empty Error!, ",
				loc, fieldStr))
	}
	formattedType, isArr, arrSize, isPointer, err1 := o.parseFieldType(originalType, pointerCode)
	if nil != err1 {
		return nil,
			errors.New(fmt.Sprintf("ProtoItemField[Loc=%s, Value=\"%s\"] Field Type Parse Error!, ",
				loc, fieldStr))
	}
	lang, _ := Setting.System.FindProgramLanguage(o.ProtoCtx.Language)
	landDataTypeDefine, ok := lang.Setting.GetDataTypeDefine(formattedType)
	fieldItem = &ProtoFieldItem{
		Remark:        dataRemark,
		Name:          dataName,
		Lang:          o.ProtoCtx.Language,
		OriginalType:  originalType,
		FormattedType: formattedType,
		IsPointer:     isPointer,
		IsArray:       isArr,
		ArraySize:     arrSize,
	}
	if ok {
		fieldItem.LangTypeDefine, fieldItem.LangType = landDataTypeDefine, landDataTypeDefine.LangTypeName
	} else {
		if isCustom := o.containsCustomProtoName(formattedType); isCustom {
			fieldItem.LangType, fieldItem.IsCustomType = formattedType, true
		} else {
			err = errors.New(fmt.Sprintf("ProtoItemField[Loc=%s, Value=\"%s\"] Format LangType Error!, ", loc, fieldStr))
		}
	}
	return
}

func (o *ProtoSheetProxy) parseFieldType(fieldType string, pointerCode string) (name string, isArr bool, arrSize int, isPointer bool, err error) {
	fieldType = setting.Format2FieldType(fieldType)
	arrStr := setting.RegArray.FindString(fieldType)
	if len(arrStr) < 2 {
		isArr, arrSize = false, -1
	} else if len(arrStr) == 2 {
		isArr, arrSize = true, -1
	} else {
		size, err1 := strconv.ParseInt(arrStr[1:len(arrStr)-1], 10, 32)
		if nil != err1 {
			err = err1
			isArr, arrSize = true, -1
		} else {
			isArr, arrSize = true, int(size)
		}
	}
	isPointer = strings.Contains(fieldType, pointerCode)
	name = setting.RegArray.ReplaceAllString(fieldType, "")
	name = strings.ReplaceAll(name, pointerCode, "")
	return
}
