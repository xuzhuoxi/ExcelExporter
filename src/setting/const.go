package setting

import "regexp"

const (
	ModeNameTitle = "title"
	ModeNameData  = "data"
	ModeNameConst = "const"
)

const (
	FieldRangeNameClient = "client"
	FieldRangeNameServer = "server"
	FieldRangeNameDb     = "db"
)

const (
	LangNameAs3    = "as3"
	LangNameCPlus  = "c++"
	LangNameCSharp = "c#"
	LangNameGo     = "go"
	LangNameJava   = "java"
	LangNameTs     = "ts"
)

const (
	FileNameBin        = "bin"
	FileNameSql        = "sql"
	FileNameJson       = "json"
	FileNameYaml       = "yaml"
	FileNameToml       = "toml"
	FileNameHcl        = "hcl"
	FileNameEnv        = "env"
	FileNameProperties = "properties"
)

//
//type FieldType uint
//
//const (
//	FieldTypeClient FieldType = iota + 1
//	FieldTypeServer
//	FieldTypeDatabase
//)

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

var (
	regFixedString = regexp.MustCompile(`string\(\d+\)`)
)

func FormatStringField(fieldValue string) string {
	if !regFixedString.MatchString(fieldValue) {
		return fieldValue
	}
	return regFixedString.ReplaceAllString(fieldValue, "string")
}
