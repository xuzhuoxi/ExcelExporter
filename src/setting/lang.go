package setting

import (
	"fmt"
	"reflect"
)

// 指定数据文件的字段读写方法
type FileAttrOperation struct {
	// 文件类型
	FileName string `yaml:"file_name"`
	// 读取方法
	Get string `yaml:"get,omitempty"`
	// 写入方法
	Set string `yaml:"set,omitempty"`
}

func (o *FileAttrOperation) String() string {
	return fmt.Sprintf("FileAttrOperation{Name=%s, Get=%s, Set=%s}", o.FileName, o.Get, o.Set)
}

// 字段
type LangDataType struct {
	FieldTypeName string              `yaml:"name"`
	LangTypeName  string              `yaml:"lang"`
	Operates      []FileAttrOperation `yaml:"operates"`
}

func (o LangDataType) GoString() string {
	return fmt.Sprintf("{%v,%v}", o.FieldTypeName, o.LangTypeName)
}

func (o LangDataType) String() string {
	return fmt.Sprintf("LangDataType{FieldTypeName=%s, Operates=%v}", o.FieldTypeName, o.Operates)
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

type LangTemp struct {
	Path string `yaml:"path"`
}

func (o LangTemp) String() string {
	return fmt.Sprintf("Temp{Path=%s}", o.Path)
}

type LangSetting struct {
	LangName  string         `yaml:"lang_name"`
	DataTypes []LangDataType `yaml:"data_types"`
}

func (o *LangSetting) GetLangDefine(name string) (format LangDataType, ok bool) {
	for index := range o.DataTypes {
		if o.DataTypes[index].FieldTypeName == name {
			return o.DataTypes[index], true
		}
	}
	return
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

func (o *LangSetting) String() string {
	return fmt.Sprintf("Lang(%s)[%v]", o.LangName, o.DataTypes)

}
