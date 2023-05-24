package setting

import (
	"fmt"
	"reflect"
)

// 指定数据文件的字段读写方法
type FileAttrOperation struct {
	FileName string `yaml:"file_name"`     // 数据文件类型(json，bin等)
	Get      string `yaml:"get,omitempty"` // 读取方法字符表达
	Set      string `yaml:"set,omitempty"` // 写入方法字符表达
}

func (o *FileAttrOperation) String() string {
	return fmt.Sprintf("FileAttrOperation{Name=%s, Get=%s, Set=%s}", o.FileName, o.Get, o.Set)
}

// 字段
type LangDataType struct {
	FieldTypeName string              `yaml:"name"`     // 字段数据类型(Excel表上填的)
	LangTypeName  string              `yaml:"lang"`     // 编程语言对应的数据类型
	Operates      []FileAttrOperation `yaml:"operates"` // 针对不同数据文件的操作方法
}

func (o LangDataType) String() string {
	return fmt.Sprintf("{%v,%v}", o.FieldTypeName, o.LangTypeName)
}

func (o LangDataType) GetPropertyOperate(fileName string) (op FileAttrOperation, ok bool) {
	for index := range o.Operates {
		if o.Operates[index].FileName == fileName {
			return o.Operates[index], true
		}
	}
	return FileAttrOperation{}, false
}

func (o LangDataType) GetGetOperate(fileName string) string {
	if op, ok := o.GetPropertyOperate(fileName); ok {
		return op.Get
	}
	return ""
}

func (o LangDataType) GetSetOperate(fileName string) string {
	if op, ok := o.GetPropertyOperate(fileName); ok {
		return op.Set
	}
	return ""
}

// 编程语言配置
type LangSetting struct {
	LangName  string         `yaml:"lang_name"`  // 编程语言名称
	DataTypes []LangDataType `yaml:"data_types"` // 编程语言数据类型配置
}

func (o *LangSetting) GetDataTypeDefine(name string) (format LangDataType, ok bool) {
	for index := range o.DataTypes {
		if o.DataTypes[index].FieldTypeName == name {
			return o.DataTypes[index], true
		}
	}
	return
}

func (o *LangSetting) String() string {
	return fmt.Sprintf("Lang(%s)[%v]", o.LangName, o.DataTypes)

}

func (o *LangSetting) getFormat(name string) (format LangDataType, ok bool) {
	t := reflect.TypeOf(o)
	elem := t.Elem()
	ln := elem.NumField()
	for index := 0; index < ln; index += 1 {
		f := elem.Field(index)
		tag := f.Tag.Get("yaml")
		fmt.Println(tag)
	}
	return
}
