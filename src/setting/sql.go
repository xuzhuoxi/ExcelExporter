package setting

import "fmt"

type DbFieldType struct {
	CfgType   string `yaml:"name"`
	FieldType string `yaml:"type"`
	IsNumber  bool   `yaml:"number"`
}

func (t DbFieldType) String() string {
	return fmt.Sprintf("DbField{%s,%s,%v}", t.CfgType, t.FieldType, t.IsNumber)
}

type DatabaseCfg struct {
	DatabaseName string        `yaml:"name"`
	FieldTypes   []DbFieldType `yaml:"types"`
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
