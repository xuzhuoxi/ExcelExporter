package setting

import (
	"fmt"
	"reflect"
)

type PropertyFileOperate struct {
	FileName string `yaml:"file_name"`
	Get      string `yaml:"get,omitempty"`
	Set      string `yaml:"set,omitempty"`
}

func (o *PropertyFileOperate) String() string {
	return fmt.Sprintf("PropertyFileOperate{Name=%s, Get=%s, Set=%s}", o.FileName, o.Get, o.Set)
}

type FieldOperate struct {
	Name     string                `yaml:"name"`
	Operates []PropertyFileOperate `yaml:"operates"`
}

func (o FieldOperate) String() string {
	return fmt.Sprintf("FieldOperate{FieldName=%s, Operates=%v}", o.Name, o.Operates)
}

func (o FieldOperate) GetPropertyOperate(fileName string) (op PropertyFileOperate, ok bool) {
	for index := range o.Operates {
		if o.Operates[index].FileName == fileName {
			return o.Operates[index], true
		}
	}
	return PropertyFileOperate{}, false
}

func (o FieldOperate) GetGetOperate(fileName string) string {
	if op, ok := o.GetPropertyOperate(fileName); ok {
		return op.Get
	}
	return ""
}

func (o FieldOperate) GetSetOperate(fileName string) string {
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
	Name string `yaml:"name"`

	Bool       FieldOperate `yaml:"bool,omitempty"`
	Int8       FieldOperate `yaml:"int8,omitempty"`
	Int16      FieldOperate `yaml:"int16,omitempty"`
	Int32      FieldOperate `yaml:"int32,omitempty"`
	Int64      FieldOperate `yaml:"int64,omitempty"`
	UInt8      FieldOperate `yaml:"uint8,omitempty"`
	UInt16     FieldOperate `yaml:"uint16,omitempty"`
	UInt32     FieldOperate `yaml:"uint32,omitempty"`
	UInt64     FieldOperate `yaml:"uint64,omitempty"`
	Float32    FieldOperate `yaml:"float32,omitempty"`
	Float64    FieldOperate `yaml:"float64,omitempty"`
	Str        FieldOperate `yaml:"string,omitempty"`
	Json       FieldOperate `yaml:"json,omitempty"`
	BoolArr    FieldOperate `yaml:"[]bool,omitempty"`
	Int8Arr    FieldOperate `yaml:"[]int8,omitempty"`
	Int16Arr   FieldOperate `yaml:"[]int16,omitempty"`
	Int32Arr   FieldOperate `yaml:"[]int32,omitempty"`
	Int64Arr   FieldOperate `yaml:"[]int64,omitempty"`
	UInt8Arr   FieldOperate `yaml:"[]uint8,omitempty"`
	UInt16Arr  FieldOperate `yaml:"[]uint16,omitempty"`
	UInt32Arr  FieldOperate `yaml:"[]uint32,omitempty"`
	UInt64Arr  FieldOperate `yaml:"[]uint64,omitempty"`
	Float32Arr FieldOperate `yaml:"[]float32,omitempty"`
	Float64Arr FieldOperate `yaml:"[]float64,omitempty"`
	StrArr     FieldOperate `yaml:"[]string,omitempty"`
	JsonArr    FieldOperate `yaml:"[]json,omitempty"`
}

func (o *LangSetting) GetLangDefine(name string) (format FieldOperate, ok bool) {
	switch name {
	case FieldBool:
		return o.Bool, true
	case FieldBoolArr:
		return o.BoolArr, true
	case FieldInt8:
		return o.Int8, true
	case FieldInt8Arr:
		return o.Int8Arr, true
	case FieldInt16:
		return o.Int16, true
	case FieldInt16Arr:
		return o.Int16Arr, true
	case FieldInt32:
		return o.Int32, true
	case FieldInt32Arr:
		return o.Int32Arr, true
	case FieldInt64:
		return o.Int64, true
	case FieldInt64Arr:
		return o.Int64Arr, true
	case FieldUint8:
		return o.UInt8, true
	case FieldUint8Arr:
		return o.UInt8Arr, true
	case FieldUint16:
		return o.UInt16, true
	case FieldUint16Arr:
		return o.UInt16Arr, true
	case FieldUint32:
		return o.UInt32, true
	case FieldUint32Arr:
		return o.UInt32Arr, true
	case FieldUint64:
		return o.UInt64, true
	case FieldUint64Arr:
		return o.UInt64Arr, true
	case FieldFloat32:
		return o.Float32, true
	case FieldFloat32Arr:
		return o.Float32Arr, true
	case FieldFloat64:
		return o.Float64, true
	case FieldFloat64Arr:
		return o.Float64Arr, true
	case FieldString:
		return o.Str, true
	case FieldStringArr:
		return o.StrArr, true
	case FieldJson:
		return o.Json, true
	case FieldJsonArr:
		return o.JsonArr, true
	default:
		return FieldOperate{}, false
	}
}

func (o *LangSetting) getFormat(name string) (format FieldOperate, ok bool) {
	t := reflect.TypeOf(o)
	elem := t.Elem()
	ln := elem.NumField()
	for index := 0; index < ln; index += 1 {
		f := elem.Field(index)
		tag := f.Tag.Get("yaml")
		fmt.Println(tag)
	}
	return o.Bool, ok
}

func (o *LangSetting) String() string {
	format := "Lang(%s)[\n" +
		"Bool=%v, \n" +
		"Int8=%v, \nInt16=%v, \nInt32=%v, \nInt64=%v, \n" +
		"UInt8=%v, \nUInt16=%v, \nUInt32=%v, \nUInt64=%v, \n" +
		"Float32=%v, \nFloat64=%v, \n" +
		"String=%v, \nJson=%v, \n" +
		"Bool[]=%v, \n" +
		"Int8[]=%v, \nInt16[]=%v, \nInt32[]=%v, \nInt64[]=%v, \n" +
		"UInt8[]=%v, \nUInt16[]=%v, \nUInt32[]=%v, \nUInt64[]=%v, \n" +
		"Float32[]=%v, \nFloat64[]=%v, \n" +
		"String[]=%v, \nJson[]=%v"
	return fmt.Sprintf(format, o.Name,
		o.Bool, o.Int8, o.Int16, o.Int32, o.Int64, o.UInt8, o.UInt16, o.UInt32, o.UInt64, o.Float32, o.Float64, o.Str, o.Json,
		o.BoolArr, o.Int8Arr, o.Int16Arr, o.Int32Arr, o.Int64Arr, o.UInt8Arr, o.UInt16Arr, o.UInt32Arr, o.UInt64Arr, o.Float32Arr, o.Float64Arr, o.StrArr, o.JsonArr)
}
