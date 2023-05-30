// Create on 2023/5/21
// @author xuzhuoxi
package setting

// 导出标记
type ConstOutputInfo struct {
	RangeName     string `yaml:"range_name"` // 导出范围名称[client, server, db]
	FileAxis      string `yaml:"file"`       // 导出类文件名坐标(Excel坐标)
	ClassAxis     string `yaml:"class"`      // 导出类名坐标(Excel坐标)
	NamespaceAxis string `yaml:"namespace"`  // 导出类命名空间坐标(Excel坐标)
	ExportAxis    string `yaml:"export"`     // 额外的导出子目录(Excel坐标)
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
