package setting

import (
	"fmt"
	"reflect"
)

type LangDefine struct {
	Name      string `yaml:"name"`
	JsonGet   string `yaml:"json_get,omitempty"`
	JsonSet   string `yaml:"json_set,omitempty"`
	YamlGet   string `yaml:"yaml_get,omitempty"`
	YamlSet   string `yaml:"yaml_set,omitempty"`
	BinaryGet string `yaml:"bin_get,omitempty"`
	BinarySet string `yaml:"bin_set,omitempty"`
}

func (o LangDefine) String() string {
	return fmt.Sprintf("{Name=%s, JsonGet=%s, JsonSet=%s, YamlGet=%s, YamlSet=%s, BinaryGet=%s, BinarySet=%s}",
		o.Name, o.JsonGet, o.JsonSet, o.YamlGet, o.YamlSet, o.BinaryGet, o.BinarySet)
}

type LangTemp struct {
	Path string `yaml:"path"`
}

func (o LangTemp) String() string {
	return fmt.Sprintf("Temp{Path=%s}", o.Path)
}

type LangSetting struct {
	Name string `yaml:"name"`

	Boolean    LangDefine `yaml:"boolean,omitempty"`
	Int8       LangDefine `yaml:"int8,omitempty"`
	Int16      LangDefine `yaml:"int16,omitempty"`
	Int32      LangDefine `yaml:"int32,omitempty"`
	Int64      LangDefine `yaml:"int64,omitempty"`
	UInt8      LangDefine `yaml:"uint8,omitempty"`
	UInt16     LangDefine `yaml:"uint16,omitempty"`
	UInt32     LangDefine `yaml:"uint32,omitempty"`
	UInt64     LangDefine `yaml:"uint64,omitempty"`
	Float32    LangDefine `yaml:"float32,omitempty"`
	Float64    LangDefine `yaml:"float64,omitempty"`
	Str        LangDefine `yaml:"string,omitempty"`
	Json       LangDefine `yaml:"json,omitempty"`
	BooleanArr LangDefine `yaml:"boolean[],omitempty"`
	Int8Arr    LangDefine `yaml:"int8[],omitempty"`
	Int16Arr   LangDefine `yaml:"int16[],omitempty"`
	Int32Arr   LangDefine `yaml:"int32[],omitempty"`
	Int64Arr   LangDefine `yaml:"int64[],omitempty"`
	UInt8Arr   LangDefine `yaml:"uint8[],omitempty"`
	UInt16Arr  LangDefine `yaml:"uint16[],omitempty"`
	UInt32Arr  LangDefine `yaml:"uint32[],omitempty"`
	UInt64Arr  LangDefine `yaml:"uint64[],omitempty"`
	Float32Arr LangDefine `yaml:"float32[],omitempty"`
	Float64Arr LangDefine `yaml:"float64[],omitempty"`
	StrArr     LangDefine `yaml:"string[],omitempty"`
	JsonArr    LangDefine `yaml:"json[],omitempty"`
}

func (o *LangSetting) GetLangDefine(name string) (format LangDefine, ok bool) {
	switch name {
	case FieldBool:
		return o.Boolean, true
	case FieldBoolArr:
		return o.BooleanArr, true
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
		return LangDefine{}, false
	}
}

func (o *LangSetting) getFormat(name string) (format LangDefine, ok bool) {
	t := reflect.TypeOf(o)
	elem := t.Elem()
	ln := elem.NumField()
	for index := 0; index < ln; index += 1 {
		f := elem.Field(index)
		tag := f.Tag.Get("yaml")
		fmt.Println(tag)
	}
	return o.Boolean, ok
}

func (o *LangSetting) String() string {
	format := "Lang(%s)[\n" +
		"Boolean=%v, \n" +
		"Int8=%v, \nInt16=%v, \nInt32=%v, \nInt64=%v, \n" +
		"UInt8=%v, \nUInt16=%v, \nUInt32=%v, \nUInt64=%v, \n" +
		"Float32=%v, \nFloat64=%v, \n" +
		"String=%v, \nJson=%v, \n" +
		"Boolean[]=%v, \n" +
		"Int8[]=%v, \nInt16[]=%v, \nInt32[]=%v, \nInt64[]=%v, \n" +
		"UInt8[]=%v, \nUInt16[]=%v, \nUInt32[]=%v, \nUInt64[]=%v, \n" +
		"Float32[]=%v, \nFloat64[]=%v, \n" +
		"String[]=%v, \nJson[]=%v"
	return fmt.Sprintf(format, o.Name,
		o.Boolean, o.Int8, o.Int16, o.Int32, o.Int64, o.UInt8, o.UInt16, o.UInt32, o.UInt64, o.Float32, o.Float64, o.Str, o.Json,
		o.BooleanArr, o.Int8Arr, o.Int16Arr, o.Int32Arr, o.Int64Arr, o.UInt8Arr, o.UInt16Arr, o.UInt32Arr, o.UInt64Arr, o.Float32Arr, o.Float64Arr, o.StrArr, o.JsonArr)
}
