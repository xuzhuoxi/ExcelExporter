package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/core/tools"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
)

// 表头导出上下文
type TitleContext struct {
	EnablePrefix  string         // 开启前缀
	RangeName     string         // 使用的字段索引名称
	RangeType     FieldRangeType // 使用的字段索引
	Language      string         // 使用的编程语言
	StartColIndex int            // 开始列索引
}

func (o TitleContext) String() string {
	return fmt.Sprintf("TitleContent(Prefix=%s, RangeName=%s, RangeType=%v, Language=%s, StartColIndex=%d)",
		o.EnablePrefix, o.RangeName, o.RangeType, o.Language, o.StartColIndex)
}

type TitleFieldItem struct {
	Index          int                  // 字段索引
	TitleName      string               // 字段名称：Excel原始值
	TitleRemark    string               // 字段描述：Excel原始值
	FieldLangName  string               // 字段名称：编程语言值
	OriginalType   string               // 字段数据类型：原始值
	FormattedType  string               // 字段数据类型：格式化值
	LangType       string               // 字段数据类型：编程语言值
	LangTypeDefine setting.LangDataType // 字段数据类型：编程语言定义

	titleProxy *TempTitleProxy
}

func (o *TitleFieldItem) GetFileKey(fileType string) string {
	return o.titleProxy.GetFieldFileKey(o.Index, fileType)
}

// 数据表代理
type TempTitleProxy struct {
	Sheet      *excel.ExcelSheet // 当前执行的Sheet数据对象
	Excel      *excel.ExcelProxy // 当前Excel代理，可能包含多个Excel
	TitleCtx   *TitleContext     // 当前执行的表头上下文数据
	FileName   string            // 表头导出类文件名
	ClassName  string            // 表头导出类名
	Namespace  string            // 表头导出类命名空间
	FieldIndex []int             // 当前选中的字段索引
}

func (o *TempTitleProxy) Language() string {
	return o.TitleCtx.Language
}

func (o *TempTitleProxy) LanguageDefine() *setting.ProgramLanguage {
	ld, _ := Setting.System.FindProgramLanguage(o.TitleCtx.Language)
	return ld
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

func (o *TempTitleProxy) GetFields() []TitleFieldItem {
	fieldItems := make([]TitleFieldItem, 0, len(o.FieldIndex))
	for _, idx := range o.FieldIndex {
		fieldItems = append(fieldItems, o.GetField(idx))
	}
	return fieldItems
}

func (o *TempTitleProxy) GetField(index int) TitleFieldItem {
	origin, formatted, lang, define, err := o.getFieldDataType(index)
	if nil != err {
		Logger.Errorln(err)
	}
	return TitleFieldItem{
		Index:          index,
		TitleName:      o.GetTitleName(index),
		TitleRemark:    o.GetTitleRemark(index),
		FieldLangName:  o.GetFieldName(index),
		OriginalType:   origin,
		FormattedType:  formatted,
		LangType:       lang,
		LangTypeDefine: define,
		titleProxy:     o,
	}
}

func (o *TempTitleProxy) GetTitleName(index int) string {
	value, err := o.getTitleName(index)
	if err != nil {
		Logger.Error(err)
	}
	return value
}

func (o *TempTitleProxy) GetTitleRemark(index int) string {
	value, err := o.getTitleRemark(index)
	if err != nil {
		Logger.Error(err)
	}
	return value
}

func (o *TempTitleProxy) GetTitleLangDefine(index int) setting.LangDataType {
	_, _, _, define, err := o.getFieldDataType(index)
	if err != nil {
		Logger.Error(err)
	}
	return define
}

func (o *TempTitleProxy) GetFieldName(index int) string {
	return o.GetFieldLangName(index, o.TitleCtx.Language)
}

func (o *TempTitleProxy) GetFieldLangName(index int, langName string) string {
	name, err := o.getFieldLangName(index, langName)
	if nil != err {
		Logger.Error(err)
	}
	return name
}

func (o *TempTitleProxy) GetFieldFileKey(index int, fileType string) string {
	key, err := o.getFieldFileKey(index, fileType)
	if nil != err {
		Logger.Error(err)
	}
	return key
}

func (o *TempTitleProxy) getTitleName(index int) (titleName string, err error) {
	nameRowIndex := Setting.Excel.TitleData.NameRow - 1
	value, err := o.Sheet.GetRowAt(nameRowIndex).ValueAtIndex(index)
	if err != nil {
		return "", errors.New(fmt.Sprintf("GetFieldLangName Error At %d", index))
	}
	return value, nil
}

func (o *TempTitleProxy) getTitleRemark(index int) (titleRemark string, err error) {
	remarkRowIndex := Setting.Excel.TitleData.RemarkRow - 1
	value, err := o.Sheet.GetRowAt(remarkRowIndex).ValueAtIndex(index)
	if err != nil {
		return "", errors.New(fmt.Sprintf("GetFieldRemark Error At %d", index))
	}
	return tools.Format2HtmlNewline(value), nil
}

func (o *TempTitleProxy) getFieldDataType(index int) (origin, formatted, lang string,
	define setting.LangDataType, err error) {
	formatRowIndex := Setting.Excel.TitleData.FieldFormatRow - 1
	value, err1 := o.Sheet.GetRowAt(formatRowIndex).ValueAtIndex(index)
	if err1 != nil {
		origin, err = value, errors.New(fmt.Sprintf("GetTitleLangDefine Error At %d: %v", index, err1))
		return
	}
	ld := o.LanguageDefine()
	formatted = setting.Format2FieldType(value)
	dtDefine, ok := ld.Setting.GetDataTypeDefine(formatted)
	if !ok {
		err = errors.New(fmt.Sprintf("GetTitleLangDefine Error At[%d]: Get Lang Define Fail with %s ", index, formatted))
		return
	}
	lang, define = dtDefine.LangTypeName, dtDefine
	return
}

func (o *TempTitleProxy) getFieldLangName(index int, langName string) (fieldLangName string, err error) {
	//fmt.Println("TempTitleProxy.GetFieldLangName:", index)
	langFormatRowIndex := Setting.Excel.TitleData.GetFieldLangNameRow(langName) - 1
	value, err := o.Sheet.GetRowAt(langFormatRowIndex).ValueAtIndex(index)
	if err != nil {
		return "", errors.New(fmt.Sprintf("getFieldLangName Error At [%s,%d]", langName, index))
	}
	return value, nil
}

func (o *TempTitleProxy) getFieldFileKey(index int, fileType string) (fieldFileKey string, err error) {
	//fmt.Println("TempTitleProxy.getFieldFileKey:", index)
	langKeyRowIndex := Setting.Excel.TitleData.GetFieldFileKeyRow(fileType) - 1
	value, err := o.Sheet.GetRowAt(langKeyRowIndex).ValueAtIndex(index)
	if err != nil {
		return "", errors.New(fmt.Sprintf("getFieldFileKey Error At [%s,%d]", fileType, index))
	}
	return value, nil
}
