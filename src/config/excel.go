package config

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
	ClassName string `yaml:"class_name"`
	// 数据文件名
	DataName string `yaml:"data_name"`
}

func (o ExcelOutputElement) String() string {
	return fmt.Sprintf("{Class=%s, Data=%s}", o.ClassName, o.DataName)
}

// 导出定义
type ExcelOutput struct {
	// 客户端定义
	Client ExcelOutputElement `yaml:"client"`
	// 客户端定义
	Server ExcelOutputElement `yaml:"server"`
	// 客户端定义
	Database ExcelOutputElement `yaml:"database"`
}

func (o ExcelOutput) String() string {
	return fmt.Sprintf("Output{Client=%s, Server=%s, Database=%s}", o.Client, o.Server, o.Database)
}

// 不同编程语言对应的字段名称，
type ExcelLangFieldNameRow struct {
	As3Row        int `yaml:"as3"`
	CPlusRow      int `yaml:"c++"`
	CSharpRow     int `yaml:"c#"`
	GoRow         int `yaml:"go"`
	JavaRow       int `yaml:"java"`
	TypeScriptRow int `yaml:"ts"`
	JsonRow       int `yaml:"json"`
	DbRow         int `yaml:"db"`
}

func (o ExcelLangFieldNameRow) String() string {
	return fmt.Sprintf("FieldRow{as3=%d, c++=%d, c#=%d, go=%d, java=%d, ts=%d, json=%d, db=%d}",
		o.As3Row, o.CPlusRow, o.CSharpRow, o.GoRow, o.JsonRow, o.TypeScriptRow, o.JsonRow, o.DbRow)
}

// 表头定义
type ExcelTitle struct {
	// 数据名称所在行号，与Excel行号一致
	NameRow int `yaml:"name"`
	// 数据注释所在行号，与Excel行号一致
	RemarkRow int `yaml:"remark"`
	// 输出选择，格式: 'c,s,d'，c、s、d的格式只能是0或1，c指前端，s指后端，d指数据库，顺序不能颠倒。从1开始
	ValidMarkRow int `yaml:"valid_mark"`
	// 数据格式,单元格格式目前支持{uint8,uint16,uint32,int8,int16,int32,float32,boolean,string,string(*)}
	DataTypeRow int `yaml:"data_type"`
	// 语言或数据格式使用的字段名称
	FieldNames ExcelLangFieldNameRow `yaml:"field_names"`
}

func (o ExcelTitle) String() string {
	return fmt.Sprintf("TitleRow{name=%d, remark=%d, valid_mark=%d, data_type=%d, fields=%v}",
		o.NameRow, o.RemarkRow, o.ValidMarkRow, o.DataTypeRow, o.FieldNames)
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
	return fmt.Sprintf("Excel{\n%s, \n%s, \n%v, \nData=%v\n}",
		o.Prefix, o.Output, o.Title, o.Data)
}
