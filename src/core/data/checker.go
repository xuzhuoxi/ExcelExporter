package data

import (
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
)

var (
	checker = make(map[string]func(data string) bool)
)

func init() {
	checker[setting.FieldBool] = CheckDataBoolean
	checker[setting.FieldBoolArr] = CheckDataBooleanArr
	checker[setting.FieldInt8] = CheckDataInt8
	checker[setting.FieldInt8Arr] = CheckDataInt8Arr
	checker[setting.FieldInt16] = CheckDataInt16
	checker[setting.FieldInt16Arr] = CheckDataInt16Arr
	checker[setting.FieldInt32] = CheckDataInt32
	checker[setting.FieldInt32Arr] = CheckDataInt32Arr
	checker[setting.FieldInt64] = CheckDataInt64
	checker[setting.FieldInt64Arr] = CheckDataInt64Arr
	checker[setting.FieldUInt8] = CheckDataUInt8
	checker[setting.FieldUInt8Arr] = CheckDataUInt8Arr
	checker[setting.FieldUInt16] = CheckDataUInt16
	checker[setting.FieldUInt16Arr] = CheckDataUInt16Arr
	checker[setting.FieldUInt32] = CheckDataUInt32
	checker[setting.FieldUInt32Arr] = CheckDataUInt32Arr
	checker[setting.FieldUInt64] = CheckDataUInt64
	checker[setting.FieldUInt64Arr] = CheckDataUInt64Arr
	checker[setting.FieldFloat32] = CheckDataFloat32
	checker[setting.FieldFloat32Arr] = CheckDataFloat32Arr
	checker[setting.FieldFloat64] = CheckDataFloat64
	checker[setting.FieldFloat64Arr] = CheckDataFloat64Arr
	checker[setting.FieldString] = CheckDataString
	checker[setting.FieldStringArr] = CheckDataStringArr
	checker[setting.FieldJson] = CheckDataJson
	checker[setting.FieldJsonArr] = CheckDataJsonArr
}

func CheckData(dataType string, data string) bool {
	if checkFunc, ok := checker[dataType]; ok {
		return checkFunc(data)
	}
	return false
}

func CheckDataBoolean(data string) bool {
	return data == "0" || data == "1" || data == "true" || data == "false"
}

func CheckDataBooleanArr(data string) bool {
	return data == "0" || data == "1" || data == "true" || data == "false"
}

func CheckDataInt8(data string) bool {
	return false
}

func CheckDataInt8Arr(data string) bool {
	return false
}

func CheckDataInt16(data string) bool {
	return false
}

func CheckDataInt16Arr(data string) bool {
	return false
}

func CheckDataInt32(data string) bool {
	return false
}

func CheckDataInt32Arr(data string) bool {
	return false
}

func CheckDataInt64(data string) bool {
	return false
}

func CheckDataInt64Arr(data string) bool {
	return false
}

func CheckDataUInt8(data string) bool {
	return false
}

func CheckDataUInt8Arr(data string) bool {
	return false
}

func CheckDataUInt16(data string) bool {
	return false
}

func CheckDataUInt16Arr(data string) bool {
	return false
}

func CheckDataUInt32(data string) bool {
	return false
}

func CheckDataUInt32Arr(data string) bool {
	return false
}

func CheckDataUInt64(data string) bool {
	return false
}

func CheckDataUInt64Arr(data string) bool {
	return false
}

func CheckDataFloat32(data string) bool {
	return false
}

func CheckDataFloat32Arr(data string) bool {
	return false
}

func CheckDataFloat64(data string) bool {
	return false
}

func CheckDataFloat64Arr(data string) bool {
	return false
}

func CheckDataString(data string) bool {
	return false
}

func CheckDataStringArr(data string) bool {
	return false
}

func CheckDataJson(data string) bool {
	return false
}

func CheckDataJsonArr(data string) bool {
	return false
}
