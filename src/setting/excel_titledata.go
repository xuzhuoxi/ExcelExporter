// Create on 2023/5/21
// @author xuzhuoxi
package setting

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
)

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
	// 启用Sheet前缀
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
	FieldRangeRow int `yaml:"range_row"`
	//  数据格式行号，内容格式支持:
	//  uint8,uint16,uint32,int8,int16,int32,float32,boolean,string,string(*),
	//  uint8[],uint16[],uint32[],int8[],int16[],int32[],float32[],boolean[],string[],string(*)[]
	DataTypeRow int `yaml:"data_type_row"`
	// 数据库字段类型定制行号，0为不定制
	SqlDataTypeRow int `yaml:"sql_data_type_row"`
	// 各语言使用的字段名称对应行号
	ExtNameRows []NameRow `yaml:"ext_name_rows"`
	// 数据文件使用的字段名称行号
	FileKeyRows []NameRow `yaml:"file_key_rows"`
	// 数据的开始行号
	DataStartAxis string `yaml:"data_start_axis"`
}

// 数据开始行号
func (td TitleData) DataStartRow() int {
	_, row, _ := excel.SplitAxis(td.DataStartAxis)
	return row
}

// 数据开始列号索引
func (td TitleData) DataStartColIndex() int {
	colIndex, _, _ := excel.ParseAxisIndex(td.DataStartAxis)
	return colIndex
}

func (td TitleData) DataStart() (row int, col int) {
	col, row, _ = excel.ParseAxisIndex(td.DataStartAxis)
	return col + 1, row + 1
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

func (td TitleData) GetFieldLangNameInfo(langName string) (row NameRow, ok bool) {
	for index := range td.ExtNameRows {
		if td.ExtNameRows[index].Name == langName {
			return td.ExtNameRows[index], true
		}
	}
	return NameRow{}, false
}

func (td TitleData) GetFieldLangNameRow(langName string) int {
	if row, ok := td.GetFieldLangNameInfo(langName); ok {
		return row.Row
	}
	return 0
}

func (td TitleData) GetFieldFileKeyInfo(fileTypeName string) (row NameRow, ok bool) {
	for index := range td.FileKeyRows {
		if td.FileKeyRows[index].Name == fileTypeName {
			return td.FileKeyRows[index], true
		}
	}
	return NameRow{}, false
}

func (td TitleData) GetFieldFileKeyRow(fileTypeName string) int {
	if row, ok := td.GetFieldFileKeyInfo(fileTypeName); ok {
		return row.Row
	}
	return 0
}

func (td TitleData) IsCustomSqlFieldType() bool {
	return td.SqlDataTypeRow > 0
}
