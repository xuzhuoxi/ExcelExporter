package data

import (
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/regexpx"
	"regexp"
	"strconv"
)

var (
	regexps = make(map[string]*regexp.Regexp)
	checker = make(map[string]func(data string) bool)
)

func init() {
	initRegexps()
	initChecker()
}

func initRegexps() {
	regexpBool, err := regexp.Compile(`^[01][(true)(false)]$`)
	if nil != err {
		panic(err)
	}
	regexps[setting.FieldBool] = regexpBool
	regexpBoolArr, err := regexp.Compile(`^(\[\])|\[([01][(true)(false)])(\,([01][(true)(false)]))+\]$`)
	if nil != err {
		panic(err)
	}
	regexps[setting.FieldBoolArr] = regexpBoolArr

	regexpInt, err := regexp.Compile(regexpx.Int)
	if nil != err {
		panic(err)
	}
	regexps[setting.FieldInt8] = regexpInt
	regexps[setting.FieldInt16] = regexpInt
	regexps[setting.FieldInt32] = regexpInt
	regexps[setting.FieldInt64] = regexpInt
	regexps[setting.FieldUInt8] = regexpInt
	regexps[setting.FieldUInt16] = regexpInt
	regexps[setting.FieldUInt32] = regexpInt
	regexps[setting.FieldUInt64] = regexpInt
	regexpIntArr, err := regexp.Compile(`^(\[\])|(\[(-?[1-9]\d*|0)(\,(-?[1-9]\d*|0))*\])$`)
	if nil != err {
		panic(err)
	}
	regexps[setting.FieldInt8Arr] = regexpIntArr
	regexps[setting.FieldInt16Arr] = regexpIntArr
	regexps[setting.FieldInt32Arr] = regexpIntArr
	regexps[setting.FieldInt64Arr] = regexpIntArr
	regexps[setting.FieldUInt8Arr] = regexpIntArr
	regexps[setting.FieldUInt16Arr] = regexpIntArr
	regexps[setting.FieldUInt32Arr] = regexpIntArr
	regexps[setting.FieldUInt64Arr] = regexpIntArr

	regexpFloat, err := regexp.Compile(regexpx.Float)
	if nil != err {
		panic(err)
	}
	regexps[setting.FieldFloat32] = regexpFloat
	regexps[setting.FieldFloat64] = regexpFloat
	regexpFloatArr, err := regexp.Compile(`^(\[\])|(\[(-?([1-9]\d*.d*|0.\d*[1-9]\d*|0?.0+|0))(\,(-?([1-9]\d*.d*|0.\d*[1-9]\d*|0?.0+|0))*\]))$`)
	if nil != err {
		panic(err)
	}
	regexps[setting.FieldFloat32Arr] = regexpFloatArr
	regexps[setting.FieldFloat64Arr] = regexpFloatArr

	regexpStringArr, err := regexp.Compile(`^(\[\])|(\[(.*)(\,.*)*\])$`)
	if nil != err {
		panic(err)
	}
	regexps[setting.FieldStringArr] = regexpStringArr
	regexps[setting.FieldJsonArr] = regexpStringArr
}

func initChecker() {
	checker[setting.FieldBool] = CheckBool
	checker[setting.FieldBoolArr] = CheckBoolArr
	checker[setting.FieldInt8] = CheckInt8
	checker[setting.FieldInt8Arr] = CheckInt8Arr
	checker[setting.FieldInt16] = CheckInt16
	checker[setting.FieldInt16Arr] = CheckInt16Arr
	checker[setting.FieldInt32] = CheckInt32
	checker[setting.FieldInt32Arr] = CheckInt32Arr
	checker[setting.FieldInt64] = CheckInt64
	checker[setting.FieldInt64Arr] = CheckInt64Arr
	checker[setting.FieldUInt8] = CheckUInt8
	checker[setting.FieldUInt8Arr] = CheckUInt8Arr
	checker[setting.FieldUInt16] = CheckUInt16
	checker[setting.FieldUInt16Arr] = CheckUInt16Arr
	checker[setting.FieldUInt32] = CheckUInt32
	checker[setting.FieldUInt32Arr] = CheckUInt32Arr
	checker[setting.FieldUInt64] = CheckUInt64
	checker[setting.FieldUInt64Arr] = CheckUInt64Arr
	checker[setting.FieldFloat32] = CheckFloat32
	checker[setting.FieldFloat32Arr] = CheckFloat32Arr
	checker[setting.FieldFloat64] = CheckFloat64
	checker[setting.FieldFloat64Arr] = CheckFloat64Arr
	checker[setting.FieldString] = CheckString
	checker[setting.FieldStringArr] = CheckStringArr
	checker[setting.FieldJson] = CheckJson
	checker[setting.FieldJsonArr] = CheckJsonArr
}

func CheckData(dataType string, data string) bool {
	if checkFunc, ok := checker[dataType]; ok {
		return checkFunc(data)
	}
	return false
}

func CheckFormat(dataType string, data string) bool {
	if checkRegexp, ok := regexps[dataType]; ok {
		if nil == checkRegexp {
			return false
		}
		return checkRegexp.MatchString(data)
	}
	return true
}

func CheckBool(data string) bool {
	return data == "0" || data == "1" || data == "true" || data == "false"
}

func CheckBoolArr(data string) bool {
	return CheckFormat(setting.FieldBoolArr, data)
}

func CheckInt8(data string) bool {
	if CheckFormat(setting.FieldInt8, data) {
		_, err := strconv.ParseInt(data, 10, 8)
		return err != nil
	}
	return false
}

func CheckInt8Arr(data string) bool {
	return CheckFormat(setting.FieldInt8Arr, data)
}

func CheckInt16(data string) bool {
	if CheckFormat(setting.FieldInt16, data) {
		_, err := strconv.ParseInt(data, 10, 16)
		return err != nil
	}
	return false
}

func CheckInt16Arr(data string) bool {
	return CheckFormat(setting.FieldInt16Arr, data)
}

func CheckInt32(data string) bool {
	if CheckFormat(setting.FieldInt32, data) {
		_, err := strconv.ParseInt(data, 10, 32)
		return err != nil
	}
	return false
}

func CheckInt32Arr(data string) bool {
	return CheckFormat(setting.FieldInt32Arr, data)
}

func CheckInt64(data string) bool {
	if CheckFormat(setting.FieldInt64, data) {
		_, err := strconv.ParseInt(data, 10, 64)
		return err != nil
	}
	return false
}

func CheckInt64Arr(data string) bool {
	return CheckFormat(setting.FieldInt64Arr, data)
}

func CheckUInt8(data string) bool {
	if CheckFormat(setting.FieldUInt8, data) {
		_, err := strconv.ParseUint(data, 10, 8)
		return err != nil
	}
	return false
}

func CheckUInt8Arr(data string) bool {
	return CheckFormat(setting.FieldUInt8Arr, data)
}

func CheckUInt16(data string) bool {
	if CheckFormat(setting.FieldUInt16, data) {
		_, err := strconv.ParseUint(data, 10, 16)
		return err != nil
	}
	return false
}

func CheckUInt16Arr(data string) bool {
	return CheckFormat(setting.FieldUInt16Arr, data)
}

func CheckUInt32(data string) bool {
	if CheckFormat(setting.FieldUInt32, data) {
		_, err := strconv.ParseUint(data, 10, 32)
		return err != nil
	}
	return false
}

func CheckUInt32Arr(data string) bool {
	return CheckFormat(setting.FieldUInt32Arr, data)
}

func CheckUInt64(data string) bool {
	if CheckFormat(setting.FieldUInt64, data) {
		_, err := strconv.ParseUint(data, 10, 64)
		return err != nil
	}
	return false
}

func CheckUInt64Arr(data string) bool {
	return CheckFormat(setting.FieldUInt64Arr, data)
}

func CheckFloat32(data string) bool {
	if CheckFormat(setting.FieldFloat32, data) {
		_, err := strconv.ParseFloat(data, 32)
		return err != nil
	}
	return false
}

func CheckFloat32Arr(data string) bool {
	return CheckFormat(setting.FieldFloat32Arr, data)
}

func CheckFloat64(data string) bool {
	if CheckFormat(setting.FieldFloat64, data) {
		_, err := strconv.ParseFloat(data, 64)
		return err != nil
	}
	return false
}

func CheckFloat64Arr(data string) bool {
	return CheckFormat(setting.FieldFloat64Arr, data)
}

func CheckString(data string) bool {
	return true
}

func CheckStringArr(data string) bool {
	return CheckFormat(setting.FieldStringArr, data)
}

func CheckJson(data string) bool {
	return true
}

func CheckJsonArr(data string) bool {
	return CheckFormat(setting.FieldJsonArr, data)
}
