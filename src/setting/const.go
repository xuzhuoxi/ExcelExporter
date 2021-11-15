package setting

import "regexp"

type FieldType uint

const (
	FieldTypeClient FieldType = iota + 1
	FieldTypeServer
	FieldTypeDatabase
)

const (
	FieldBool       = "boolean"
	FieldInt8       = "int8"
	FieldInt16      = "int16"
	FieldInt32      = "int32"
	FieldInt64      = "int64"
	FieldUint8      = "uint8"
	FieldUint16     = "uint16"
	FieldUint32     = "uint32"
	FieldUint64     = "uint64"
	FieldFloat32    = "float32"
	FieldFloat64    = "float64"
	FieldString     = "string"
	FieldJson       = "json"
	FieldBoolArr    = "boolean[]"
	FieldInt8Arr    = "int8[]"
	FieldInt16Arr   = "int16[]"
	FieldInt32Arr   = "int32[]"
	FieldInt64Arr   = "int64[]"
	FieldUint8Arr   = "uint8[]"
	FieldUint16Arr  = "uint16[]"
	FieldUint32Arr  = "uint32[]"
	FieldUint64Arr  = "uint64[]"
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
)

const (
	FileJson   = "json"
	FileSql    = "sql"
	FileBinary = "binary"
	FileYaml   = "yaml"
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
