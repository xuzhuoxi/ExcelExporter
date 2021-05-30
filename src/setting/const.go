package setting

import "regexp"

type FieldType uint

const (
	FieldTypeClient FieldType = iota + 1
	FieldTypeServer
	FieldTypeDatabase
)

const (
	FieldBoolean    = "boolean"
	FieldInt8       = "int8"
	FieldInt16      = "int16"
	FieldInt32      = "int32"
	FieldInt64      = "int64"
	FieldUInt8      = "uint8"
	FieldUInt16     = "uint16"
	FieldUInt32     = "uint32"
	FieldUInt64     = "uint64"
	FieldFloat32    = "float32"
	FieldFloat64    = "float64"
	FieldString     = "string"
	FieldJson       = "json"
	FieldBooleanArr = "boolean[]"
	FieldInt8Arr    = "int8[]"
	FieldInt16Arr   = "int16[]"
	FieldInt32Arr   = "int32[]"
	FieldInt64Arr   = "int64[]"
	FieldUInt8Arr   = "uint8[]"
	FieldUInt16Arr  = "uint16[]"
	FieldUInt32Arr  = "uint32[]"
	FieldUInt64Arr  = "uint64[]"
	FieldFloat32Arr = "float32[]"
	FieldFloat64Arr = "float64[]"
	FieldStringArr  = "string[]"
	FieldJsonArr    = "json[]"
)

const (
	LangAs3        = "as3"
	LangCPlus      = "c++"
	LangCSharp     = "c#"
	LangGo         = "go"
	LangJava       = "java"
	LangTypeScript = "ts"

	FileJson = "json"
	FileSql  = "sql"
)

var (
	reg = regexp.MustCompile(`string\(\d+\)`)
)

func FormatStringField(fieldValue string) string {
	if !reg.MatchString(fieldValue) {
		return fieldValue
	}
	return reg.ReplaceAllString(fieldValue, "string")
}
