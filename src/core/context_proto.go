// Create on 2023/5/21
// @author xuzhuoxi
package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strings"
)

// 协议表上下文
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

// 协议表关定义
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

// 协议属性
type ProtoFieldItem struct {
	Name          string // 名称
	OriginalType  string // 原始的数据类型
	FormattedType string // 格式化后的数据类型
	Remark        string // 备注

	lang string // 编程语言
	loc  string // 坐标
}

func (o *ProtoFieldItem) IsCustomType() bool {
	_, ok := o.getTypeDefine()
	return !ok
}

func (o *ProtoFieldItem) Type() string {
	define, ok := o.getTypeDefine()
	if ok {
		return define.LangTypeName
	} else {
		return o.OriginalType
	}
}

func (o *ProtoFieldItem) TypeDefine() setting.LangDataType {
	define, ok := o.getTypeDefine()
	if !ok {
		return setting.LangDataType{}
	}
	return define
}

func (o *ProtoFieldItem) getTypeDefine() (dataType setting.LangDataType, ok bool) {
	lang, _ := Setting.System.FindProgramLanguage(o.lang)
	return lang.Setting.GetDataTypeDefine(o.FormattedType)
}

// 协议项
type ProtoItem struct {
	Id     string           // 协议Id
	Name   string           // 协议名称
	Remark string           // 协议备注
	File   string           // 协议文件名
	Fields []ProtoFieldItem // 协议属性
}

// 协议表模板代理
type TempProtoProxy struct {
	ProtoItem  ProtoItem
	SheetProxy *ProtoSheetProxy
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

// 取当前Sheet中对应坐标的字符数据，若数据不存在，返回空字符串
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

// 取当前Sheet全部协议数据列表(已经过滤中间的空行)
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

// 取当前Sheet指定行号数据，转换为协议项，格式非法则返回对应错误
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
	suc := o.appendProtoName(name)
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

func (o *ProtoSheetProxy) appendProtoName(protoName string) bool {
	if o.checkProtoName(protoName) {
		return false
	}
	o.protoNames = append(o.protoNames, protoName)
	return true
}

func (o *ProtoSheetProxy) checkProtoName(protoName string) bool {
	return slicex.ContainsString(o.protoNames, protoName)
}

func (o *ProtoSheetProxy) getFieldSize(excelRow *excel.ExcelRow, fieldStartIdx int) int {
	// len = 7 startIdx = 3	rs = 4
	celLen := excelRow.CellLength() //7
	if celLen <= 0 || celLen <= fieldStartIdx {
		return 0
	}
	for idx := fieldStartIdx; idx < celLen; idx += 1 {
		cellStr := excelRow.Cell[idx]
		if len(cellStr) == 0 {
			return idx - fieldStartIdx + 1
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
	fieldItems = make([]ProtoFieldItem, fieldSize, fieldSize)
	colIdx := ps.StartFieldColIndex()
	idx := 0
	for idx < fieldSize {
		fieldStr := itemRow.Cell[colIdx]
		info := strings.Split(fieldStr, ":")

		cellName := excel.GetCellName(colIdx+1, rowId)
		if len(info) != 2 {
			return nil, errors.New(fmt.Sprintf("ProtoItemField[Loc=%s, Value=\"%s\"] Format Size Error!, ",
				cellName, fieldStr))
		}
		nameStr := strings.TrimSpace(info[0])
		originalType := strings.TrimSpace(info[1])
		if len(nameStr) == 0 || len(originalType) == 0 {
			return nil, errors.New(fmt.Sprintf("ProtoItemField[Loc=%s, Value=\"%s\"] Format Empty Error!, ",
				cellName, fieldStr))
		}
		formattedType := setting.Format2FieldType(originalType)
		fieldItem := &ProtoFieldItem{Name: nameStr, OriginalType: originalType, FormattedType: formattedType,
			loc: cellName, lang: o.ProtoCtx.Language}
		if remarkExist {
			fieldItem.Remark = strings.TrimSpace(remarkRow.Cell[colIdx])
		}

		// 属性类型错误
		if fieldItem.IsCustomType() && !o.checkProtoName(formattedType) {
			return nil, errors.New(fmt.Sprintf("ProtoItemField[Loc=%s, Value=\"%s\"] Format Type Error!, ",
				cellName, fieldStr))
		}
		fieldItems[idx] = *fieldItem
		colIdx += 1
		idx += 1
	}
	return fieldItems, nil
}
