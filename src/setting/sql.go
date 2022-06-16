package setting

import "fmt"

type SqlDataType struct {
	Name          string `yaml:"name"`
	DataTypeValue string `yaml:"type"`
	IsNumber      bool   `yaml:"number"`
}

func (t SqlDataType) String() string {
	return fmt.Sprintf("SqlDataType{%s,%s}", t.Name, t.DataTypeValue)
}

type SqlDataTypes struct {
	DatabaseName string        `yaml:"name"`
	DataTypes    []SqlDataType `yaml:"types"`
}

func (d SqlDataTypes) GetType(formattedFieldType string) (t SqlDataType, ok bool) {
	ok = false
	if len(formattedFieldType) == 0 {
		return
	}
	for index := range d.DataTypes {
		if d.DataTypes[index].Name == formattedFieldType {
			return d.DataTypes[index], true
		}
	}
	return
}
