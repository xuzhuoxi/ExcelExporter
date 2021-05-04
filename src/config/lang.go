package config

import "fmt"

type LangDataType struct {
	Name      string `yaml:"name"`
	JsonGet   string `yaml:"json_get,omitempty"`
	JsonSet   string `yaml:"json_set,omitempty"`
	YamlGet   string `yaml:"yaml_get,omitempty"`
	YamlSet   string `yaml:"yaml_set,omitempty"`
	BinaryGet string `yaml:"binary_get,omitempty"`
	BinarySet string `yaml:"binary_set,omitempty"`
}

func (o LangDataType) String() string {
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
	Name     string   `yaml:"name"`
	TempPath LangTemp `yaml:"temp"`

	Boolean    LangDataType `yaml:"boolean,omitempty"`
	Int8       LangDataType `yaml:"int8,omitempty"`
	Int16      LangDataType `yaml:"int16,omitempty"`
	Int32      LangDataType `yaml:"int32,omitempty"`
	Int64      LangDataType `yaml:"int64,omitempty"`
	UInt8      LangDataType `yaml:"uint8,omitempty"`
	UInt16     LangDataType `yaml:"uint16,omitempty"`
	UInt32     LangDataType `yaml:"uint32,omitempty"`
	UInt64     LangDataType `yaml:"uint64,omitempty"`
	Float32    LangDataType `yaml:"float32,omitempty"`
	Float64    LangDataType `yaml:"float64,omitempty"`
	Str        LangDataType `yaml:"string,omitempty"`
	Json       LangDataType `yaml:"json,omitempty"`
	BooleanArr LangDataType `yaml:"boolean[],omitempty"`
	Int8Arr    LangDataType `yaml:"int8[],omitempty"`
	Int16Arr   LangDataType `yaml:"int16[],omitempty"`
	Int32Arr   LangDataType `yaml:"int32[],omitempty"`
	Int64Arr   LangDataType `yaml:"int64[],omitempty"`
	UInt8Arr   LangDataType `yaml:"uint8[],omitempty"`
	UInt16Arr  LangDataType `yaml:"uint16[],omitempty"`
	UInt32Arr  LangDataType `yaml:"uint32[],omitempty"`
	UInt64Arr  LangDataType `yaml:"uint64[],omitempty"`
	Float32Arr LangDataType `yaml:"float32[],omitempty"`
	Float64Arr LangDataType `yaml:"float64[],omitempty"`
	StrArr     LangDataType `yaml:"string[],omitempty"`
	JsonArr    LangDataType `yaml:"json[],omitempty"`
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
