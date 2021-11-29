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
	FileNameYml        = "yml"
	FileNameToml       = "toml"
	FileNameHcl        = "hcl"
	FileNameEnv        = "env"
	FileNameProperties = "properties"
)

const (
	FieldBool       = "bool"
	FieldBoolArr    = "[]bool"
	FieldInt8       = "int8"
	FieldInt8Arr    = "[]int8"
	FieldInt16      = "int16"
	FieldInt16Arr   = "[]int16"
	FieldInt32      = "int32"
	FieldInt32Arr   = "[]int32"
	FieldInt64      = "int64"
	FieldInt64Arr   = "[]int64"
	FieldUint8      = "uint8"
	FieldUint8Arr   = "[]uint8"
	FieldUint16     = "uint16"
	FieldUint16Arr  = "[]uint16"
	FieldUint32     = "uint32"
	FieldUint32Arr  = "[]uint32"
	FieldUint64     = "uint64"
	FieldUint64Arr  = "[]uint64"
	FieldFloat32    = "float32"
	FieldFloat32Arr = "[]float32"
	FieldFloat64    = "float64"
	FieldFloat64Arr = "[]float64"
	FieldString     = "string"
	FieldStringArr  = "[]string"
	FieldJson       = "json"
	FieldJsonArr    = "[]json"
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
