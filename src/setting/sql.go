package setting

import "fmt"

// 数据库数据类型描述
type DbFieldType struct {
	CfgType   string `yaml:"name"`   // 字段名称(标准化后，如string(5)=>string(*))
	FieldType string `yaml:"type"`   // 对应数据的字段数据类型
	IsNumber  bool   `yaml:"number"` // 是否为数值类型
}

func (t DbFieldType) String() string {
	return fmt.Sprintf("DbField{%s,%s,%v}", t.CfgType, t.FieldType, t.IsNumber)
}

// 数据库配置
type DatabaseCfg struct {
	DatabaseName string        `yaml:"name"`  // 数据库名称
	FieldTypes   []DbFieldType `yaml:"types"` // 数据库数据类型描述列表
}

func (d DatabaseCfg) GetFieldType(formattedFieldType string) (t DbFieldType, ok bool) {
	ok = false
	if len(formattedFieldType) == 0 {
		return
	}
	for index := range d.FieldTypes {
		if d.FieldTypes[index].CfgType == formattedFieldType {
			return d.FieldTypes[index], true
		}
	}
	return
}
