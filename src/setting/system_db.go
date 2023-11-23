package setting

import (
	"fmt"
	"strings"
)

// DbFieldType 数据库数据类型描述
type DbFieldType struct {
	CfgType   string `yaml:"name"`   // 字段名称(标准化后，如string(5)=>string(*))
	FieldType string `yaml:"type"`   // 对应数据的字段数据类型
	IsNumber  bool   `yaml:"number"` // 是否为数值类型
	IsArray   bool   `yaml:"array"`  // 是否为数组类型
}

func (t DbFieldType) IsDynamicLen() bool {
	return strings.Index(t.FieldType, "*") != -1
}

func (t DbFieldType) IsDynamicVarchar() bool {
	return "VARCHAR(*)" == strings.ToUpper(t.FieldType)
}

func (t DbFieldType) IsDynamicChar() bool {
	return "CHAR(*)" == strings.ToUpper(t.FieldType)
}

func (t DbFieldType) String() string {
	return fmt.Sprintf("DbField{%s,%s,%v}", t.CfgType, t.FieldType, t.IsNumber)
}

// DatabaseExtend 数据库配置
type DatabaseExtend struct {
	DatabaseName string        `yaml:"db_name"`       // 数据库名称
	ScaleChar    float64       `yaml:"scale_char"`    // Char字符比例
	ScaleVarchar float64       `yaml:"scale_varchar"` // Varchar字符比例
	FieldTypes   []DbFieldType `yaml:"types"`         // 数据库数据类型描述列表
}

func (d DatabaseExtend) GetFieldType(formattedFieldType string) (t DbFieldType, ok bool) {
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
