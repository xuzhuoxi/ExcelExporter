package setting

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
)

// 名称与号记录项
type NameRow struct {
	Name string `yaml:"name"` // 名称(键)
	Row  int    `yaml:"row"`  // Excel行呈
}

func (o NameRow) String() string {
	return fmt.Sprintf("NameRow{Name=%s, Row=%d}", o.Name, o.Row)
}

// 名称与字符值记录项
type NameValue struct {
	Name  string `yaml:"name"`  // 名称(键)
	Value string `yaml:"value"` // 内容
}

func (o NameValue) String() string {
	return fmt.Sprintf("NameValue{Name=%s, Value=%s}", o.Name, o.Value)
}

// 导出标记
type TitleDataOutputInfo struct {
	RangeName     string `yaml:"range_name"` // 导出范围名称[client, server, db]
	TitleFileAxis string `yaml:"title_file"` // 导出类文件名坐标(Excel坐标)
	DataFileAxis  string `yaml:"data_file"`  // 导出数据文件名坐标(Excel坐标)
	ClassAxis     string `yaml:"class"`      // 导出类名坐标(Excel坐标)
	NamespaceAxis string `yaml:"namespace"`  // 导出类命名空间坐标(Excel坐标)
}

func (o TitleDataOutputInfo) String() string {
	return fmt.Sprintf("TitleDataOut{Name=%s, Title=%s, Data=%s, Class=%s, Namespace=%s}",
		o.RangeName, o.TitleFileAxis, o.DataFileAxis, o.ClassAxis, o.NamespaceAxis)
}

// Sql坐标信息
type TitleDataSqlInfo struct {
	TableNameAxis  string `yaml:"table"` // 表名坐标(Excel坐标)
	FileNameAxis   string `yaml:"file"`  // 数据文件名坐标(Excel坐标)
	PrimaryKeyAxis string `yaml:"key"`   // 主键信息坐标,复合主键使用英文逗号","分隔
}

type TitleData struct {
	// 启用前缀
	Prefix string `yaml:"prefix"`
	// 导出命名
	Outputs []TitleDataOutputInfo `yaml:"outputs"`
	// 表头导出类信息
	Sql TitleDataSqlInfo `yaml:"sql"`
	// 数据控制列，决定数据范围
	ControlRow int `yaml:"control_row"`
	// 字段别名行号，用于查找指定列，值为0时使用列号作为别名
	NickRow int `yaml:"nick_row"`
	// 数据名称所在行号，与Excel行号一致
	NameRow int `yaml:"name_row"`
	// 数据注释所在行号，与Excel行号一致
	RemarkRow int `yaml:"remark_row"`
	// 输出选择行号
	// 内容格式: 'c,s,d'，c、s、d的格式只能是0或1，c指前端，s指后端，d指数据库
	// 顺序不能颠倒
	// sql脚本导出只针对d值
	FieldRangeRow int `yaml:"field_range_row"`
	//  数据格式行号，内容格式支持:
	//  uint8,uint16,uint32,int8,int16,int32,float32,boolean,string,string(*),
	//  uint8[],uint16[],uint32[],int8[],int16[],int32[],float32[],boolean[],string[],string(*)[]
	FieldFormatRow int `yaml:"field_format_row"`
	// 数据库字段类型定制行号，0为不定制
	SqlFieldFormatRow int `yaml:"sql_field_format_row"`
	// 各语言使用的字段名称对应行号
	FieldNames []NameRow `yaml:"field_names"`
	// 数据文件使用的字段名称行号
	FileKeys []NameRow `yaml:"file_keys"`
	//// 数据的开始行号
	//DataStartRow int `yaml:"data_start_row"`
	// 数据的开始行号
	DataStart string `yaml:"data_start"`
}

// 数据开始行号
func (td TitleData) DataStartRow() int {
	_, row, _ := excel.SplitAxis(td.DataStart)
	return row
}

// 数据开始列号索引
func (td TitleData) DataStartColIndex() int {
	colIndex, _, _ := excel.ParseAxisIndex(td.DataStart)
	return colIndex
}

func (td TitleData) GetSqlInfo() TitleDataSqlInfo {
	return td.Sql
}

func (td TitleData) GetOutputInfo(rangeName string) (info TitleDataOutputInfo, ok bool) {
	for index := range td.Outputs {
		if td.Outputs[index].RangeName == rangeName {
			return td.Outputs[index], true
		}
	}
	return TitleDataOutputInfo{}, false
}

func (td TitleData) GetFieldNameInfo(name string) (row NameRow, ok bool) {
	for index := range td.FieldNames {
		if td.FieldNames[index].Name == name {
			return td.FieldNames[index], true
		}
	}
	return NameRow{}, false
}

func (td TitleData) GetFieldNameRow(name string) int {
	if row, ok := td.GetFieldNameInfo(name); ok {
		return row.Row
	}
	return 0
}

func (td TitleData) GetFileKeyInfo(fileName string) (row NameRow, ok bool) {
	for index := range td.FileKeys {
		if td.FileKeys[index].Name == fileName {
			return td.FileKeys[index], true
		}
	}
	return NameRow{}, false
}

func (td TitleData) GetFileKeyRow(name string) int {
	if row, ok := td.GetFileKeyInfo(name); ok {
		return row.Row
	}
	return 0
}

func (td TitleData) IsCustomSqlFieldType() bool {
	return td.SqlFieldFormatRow > 0
}

// 导出标记
type ConstOutputInfo struct {
	RangeName     string `yaml:"range_name"` // 导出范围名称[client, server, db]
	FileAxis      string `yaml:"file"`       // 导出类文件名坐标(Excel坐标)
	ClassAxis     string `yaml:"class"`      // 导出类名坐标(Excel坐标)
	NamespaceAxis string `yaml:"namespace"`  // 导出类命名空间坐标(Excel坐标)
}

type Const struct {
	Prefix       string            `yaml:"prefix"`         // 启用前缀
	Outputs      []ConstOutputInfo `yaml:"outputs"`        // 导出文件信息列表
	NameCol      string            `yaml:"name_col"`       // 常量名 对应列号(Excel列号)
	ValueCol     string            `yaml:"value_col"`      // 常量值 对应列号(Excel列号)
	TypeCol      string            `yaml:"type_col"`       // 常量值类型 对应列号(Excel列号)
	RemarkCol    string            `yaml:"remark_col"`     // 注释 对应列号(Excel列号)
	DataStartRow int               `yaml:"data_start_row"` // 数据的开始行号(Excel列号)
}

func (c Const) GetOutputInfo(rangeName string) (v ConstOutputInfo, ok bool) {
	for index := range c.Outputs {
		if c.Outputs[index].RangeName == rangeName {
			return c.Outputs[index], true
		}
	}
	return ConstOutputInfo{}, false
}

// Excel相关配置环境
type ExcelSetting struct {
	TitleData TitleData `yaml:"title&data"` //数据表配置
	Const     Const     `yaml:"const"`      // 常量表配置
}
