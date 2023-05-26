// Create on 2023/5/21
// @author xuzhuoxi
package setting

import "github.com/xuzhuoxi/ExcelExporter/src/core/excel"

type Proto struct {
	Prefix string `yaml:"prefix"` // 启用前缀

	IdDataType string `yaml:"id_datatype"` // Id数据类型 单元格
	RangeName  string `yaml:"range_name"`  // 导出范围名称[client, server]
	Namespace  string `yaml:"namespace"`   // 命名空间(包名) 单元格
	Export     string `yaml:"export"`      // 额外的导出子目录 单元格

	IdCol         string `yaml:"id_col"`          // 协议Id  配置列号
	FileCol       string `yaml:"file_col"`        // 导出文件名称（此处为空，代表不导出文件）  配置列号
	NameCol       string `yaml:"name_col"`        // 协议名称(导出类名)  配置列号
	FieldStartCol string `yaml:"field_start_col"` // 协议属性开始列号  配置列号

	DataStartRow   int  `yaml:"data_start_row"` // 协议数据开始行号: 格式(key:type)
	RemarkOffset   int  `yaml:"remark_offset"`  // 属性备注行号偏移
	BlankLineBreak bool `yaml:"blank_break"`    // 空行中断
}

func (p Proto) IdColIndex() int {
	return excel.GetColNum(p.IdCol) - 1
}

func (p Proto) FileColIndex() int {
	return excel.GetColNum(p.FileCol) - 1
}

func (p Proto) NameColIndex() int {
	return excel.GetColNum(p.NameCol) - 1
}

func (p Proto) StartFieldColIndex() int {
	return excel.GetColNum(p.FieldStartCol) - 1
}
