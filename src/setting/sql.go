package setting

import "fmt"

type SqlDataType struct {
	Name          string `yaml:"name"`
	DataTypeValue string `yaml:"type"`
}

func (t SqlDataType) String() string {
	return fmt.Sprintf("SqlDataType{%s,%s}", t.Name, t.DataTypeValue)
}

type SqlDataTypes struct {
	DatabaseName string        `yaml:"name"`
	DataTypes    []SqlDataType `yaml:"types"`
}

func (d SqlDataTypes) GetType(name string) (t SqlDataType, ok bool) {
	ok = false
	if len(name) == 0 {
		return
	}
	for index := range d.DataTypes {
		if d.DataTypes[index].Name == name {
			return d.DataTypes[index], true
		}
	}
	return
}
