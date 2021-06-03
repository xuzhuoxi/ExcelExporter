package setting

import (
	"fmt"
)

// 要求输出的Sheet的名称前缀，只有带这个前缀的Sheet才参与处理
type ExcelPrefix struct {
	// 定义及数据处理
	Data string `yaml:"data"`
	// 常量表处理
	Const string `yaml:"const"`
}

func (o ExcelPrefix) String() string {
	return fmt.Sprintf("Prefix{Data=%s, Const=%s}", o.Data, o.Const)
}

// 值为Excel对应单元格号，用'_'分隔
type ExcelOutputElement struct {
	// 数据结构定义名(类名)
	TitleName string `yaml:"title_name"`
	// 数据文件名
	DataName string `yaml:"data_name"`
}

func (o ExcelOutputElement) String() string {
	return fmt.Sprintf("{Class=%s, Data=%s}", o.TitleName, o.DataName)
}

// 导出定义
type ExcelOutput struct {
	// 客户端定义
	Client ExcelOutputElement `yaml:"client"`
	// 服务端定义
	Server ExcelOutputElement `yaml:"server"`
	// 数据库定义
	Database ExcelOutputElement `yaml:"database"`
}

func (o ExcelOutput) String() string {
	return fmt.Sprintf("Output{Client=%s, Server=%s, Database=%s}", o.Client, o.Server, o.Database)
}

func (o ExcelOutput) GetElement(fieldType FieldType) (ele ExcelOutputElement, ok bool) {
	switch fieldType {
	case FieldTypeClient:
		return o.Client, true
	case FieldTypeServer:
		return o.Server, true
	case FieldTypeDatabase:
		return o.Database, true
	default:
		ok = false
		return
	}
}

// 不同编程语言对应的字段名称，
type LangKeyRows struct {
	As3Row        int `yaml:"as3"`
	CPlusRow      int `yaml:"c++"`
	CSharpRow     int `yaml:"c#"`
	GoRow         int `yaml:"go"`
	JavaRow       int `yaml:"java"`
	TypeScriptRow int `yaml:"ts"`
}

func (o LangKeyRows) GetRowNum(name string) int {
	switch name {
	case LangAs3:
		return o.As3Row
	case LangCPlus:
		return o.CPlusRow
	case LangCSharp:
		return o.CSharpRow
	case LangGo:
		return o.GoRow
	case LangJava:
		return o.JavaRow
	case LangTypeScript:
		return o.TypeScriptRow
	default:
		return -1
	}
}

func (o LangKeyRows) String() string {
	return fmt.Sprintf("LangKeyRows{as3=%d, c++=%d, c#=%d, go=%d, java=%d, ts=%d}",
		o.As3Row, o.CPlusRow, o.CSharpRow, o.GoRow, o.JavaRow, o.TypeScriptRow)
}

// 不同编程语言对应的字段名称，
type FileKeyRows struct {
	JsonRow int `yaml:"json"`
	SqlRow  int `yaml:"sql"`
}

func (o FileKeyRows) GetRowNum(name string) int {
	switch name {
	case FileJson:
		return o.JsonRow
	case FileSql:
		return o.SqlRow
	default:
		return -1
	}
}

func (o FileKeyRows) String() string {
	return fmt.Sprintf("FileKeyRows{json=%d, db=%d}", o.JsonRow, o.SqlRow)
}

// 表头定义
type ExcelTitle struct {
	// 字段别名，用于查找指定列，值为0时使用列号作为别名
	ColNickRow int `yaml:"col_nick_row"`
	// 数据名称所在行号，与Excel行号一致
	NameRow int `yaml:"name_row"`
	// 数据注释所在行号，与Excel行号一致
	RemarkRow int `yaml:"remark_row"`
	// 输出开关选择，格式: 'c,s,d'，c、s、d的格式只能是0或1，c指前端，s指后端，d指数据库，顺序不能颠倒。从1开始
	FieldSwitchRow int `yaml:"field_switch_row"`
	// 数据格式,单元格格式目前支持{uint8,uint16,uint32,int8,int16,int32,float32,boolean,string,string(*)}
	FieldFormatRow int `yaml:"field_format_row"`
	// 语言使用的字段名称
	LangKeyRows LangKeyRows `yaml:"lang_key_rows"`
	// 数据文件使用的字段名称
	FileKeyRows FileKeyRows `yaml:"file_key_rows"`
}

func (o ExcelTitle) String() string {
	return fmt.Sprintf("TitleRow{name=%d, remark=%d, valid_mark=%d, data_type=%d, fields=%v}",
		o.NameRow, o.RemarkRow, o.FieldSwitchRow, o.FieldFormatRow, o.LangKeyRows)
}

// 数据定义
type ExcelData struct {
	// 数据的开始行号
	StartRow int
	// 数据忽略
	Pass string
}

func (o ExcelData) String() string {
	return fmt.Sprintf("Data{start=%d, pass=%s}", o.StartRow, o.Pass)
}

type ExcelSetting struct {
	Prefix ExcelPrefix `yaml:"prefix"`
	Output ExcelOutput `yaml:"output"`
	Title  ExcelTitle  `yaml:"title"`
	Data   ExcelData   `yaml:"data"`
}

func (o *ExcelSetting) String() string {
	return fmt.Sprintf("Excel{%s, %s, %v, Data=%v}", o.Prefix, o.Output, o.Title, o.Data)
}
