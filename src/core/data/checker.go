package data

import (
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"regexp"
)

var (
	valueRegexps  = make(map[string]*regexp.Regexp)
	valueCheckers = make(map[string]func(data string) bool)
)

var (
	regexpTypeFixedString    *regexp.Regexp
	regexpTypeFixedStringArr *regexp.Regexp
)

func init() {
	initTypeRegexps()
	initValueRegexps()
	initValueChecker()
}

func initTypeRegexps() {
	r, err := regexp.Compile(`^string\([1-9]\d*\)$`)
	if nil != err {
		panic(err)
	}
	regexpTypeFixedString = r
	rArr, errArr := regexp.Compile(`^\[\]string\([1-9]\d*\)$`)
	if nil != errArr {
		panic(errArr)
	}
	regexpTypeFixedStringArr = rArr
}

func initValueRegexps() {
	//regexpBool, err := regexp.Compile(`^[01][(true)(false)]$`)
	//if nil != err {
	//	panic(err)
	//}
	//valueRegexps[setting.FieldBool] = regexpBool
	//regexpBoolArr, err := regexp.Compile(`^(\[\])|\[([01][(true)(false)])(\,([01][(true)(false)]))+\]$`)
	//if nil != err {
	//	panic(err)
	//}
	//valueRegexps[setting.FieldBoolArr] = regexpBoolArr
	//
	//regexpInt, err := regexp.Compile(regexpx.Int)
	//if nil != err {
	//	panic(err)
	//}
	//valueRegexps[setting.FieldInt8] = regexpInt
	//valueRegexps[setting.FieldInt16] = regexpInt
	//valueRegexps[setting.FieldInt32] = regexpInt
	//valueRegexps[setting.FieldInt64] = regexpInt
	//valueRegexps[setting.FieldUint8] = regexpInt
	//valueRegexps[setting.FieldUint16] = regexpInt
	//valueRegexps[setting.FieldUint32] = regexpInt
	//valueRegexps[setting.FieldUint64] = regexpInt
	//regexpIntArr, err := regexp.Compile(`^(\[\])|(\[(-?[1-9]\d*|0)(\,(-?[1-9]\d*|0))*\])$`)
	//if nil != err {
	//	panic(err)
	//}
	//valueRegexps[setting.FieldInt8Arr] = regexpIntArr
	//valueRegexps[setting.FieldInt16Arr] = regexpIntArr
	//valueRegexps[setting.FieldInt32Arr] = regexpIntArr
	//valueRegexps[setting.FieldInt64Arr] = regexpIntArr
	//valueRegexps[setting.FieldUint8Arr] = regexpIntArr
	//valueRegexps[setting.FieldUint16Arr] = regexpIntArr
	//valueRegexps[setting.FieldUint32Arr] = regexpIntArr
	//valueRegexps[setting.FieldUint64Arr] = regexpIntArr
	//
	//regexpFloat, err := regexp.Compile(regexpx.Float)
	//if nil != err {
	//	panic(err)
	//}
	//valueRegexps[setting.FieldFloat32] = regexpFloat
	//valueRegexps[setting.FieldFloat64] = regexpFloat
	//regexpFloatArr, err := regexp.Compile(`^(\[\])|(\[(-?([1-9]\d*.d*|0.\d*[1-9]\d*|0?.0+|0))(\,(-?([1-9]\d*.d*|0.\d*[1-9]\d*|0?.0+|0))*\]))$`)
	//if nil != err {
	//	panic(err)
	//}
	//valueRegexps[setting.FieldFloat32Arr] = regexpFloatArr
	//valueRegexps[setting.FieldFloat64Arr] = regexpFloatArr

	regexpStringArr, err := regexp.Compile(`^(\[\])|(\[(.*)(\,.*)*\])$`)
	if nil != err {
		panic(err)
	}
	valueRegexps[setting.FieldStringArr] = regexpStringArr
	//valueRegexps[setting.FieldJsonArr] = regexpStringArr
}

func initValueChecker() {
	//valueCheckers[setting.FieldBool] = CheckBool
	//valueCheckers[setting.FieldBoolArr] = CheckBoolArr
	//valueCheckers[setting.FieldInt8] = CheckInt8
	//valueCheckers[setting.FieldInt8Arr] = CheckInt8Arr
	//valueCheckers[setting.FieldInt16] = CheckInt16
	//valueCheckers[setting.FieldInt16Arr] = CheckInt16Arr
	//valueCheckers[setting.FieldInt32] = CheckInt32
	//valueCheckers[setting.FieldInt32Arr] = CheckInt32Arr
	//valueCheckers[setting.FieldInt64] = CheckInt64
	//valueCheckers[setting.FieldInt64Arr] = CheckInt64Arr
	//valueCheckers[setting.FieldUint8] = CheckUInt8
	//valueCheckers[setting.FieldUint8Arr] = CheckUInt8Arr
	//valueCheckers[setting.FieldUint16] = CheckUInt16
	//valueCheckers[setting.FieldUint16Arr] = CheckUInt16Arr
	//valueCheckers[setting.FieldUint32] = CheckUInt32
	//valueCheckers[setting.FieldUint32Arr] = CheckUInt32Arr
	//valueCheckers[setting.FieldUint64] = CheckUInt64
	//valueCheckers[setting.FieldUint64Arr] = CheckUInt64Arr
	//valueCheckers[setting.FieldFloat32] = CheckFloat32
	//valueCheckers[setting.FieldFloat32Arr] = CheckFloat32Arr
	//valueCheckers[setting.FieldFloat64] = CheckFloat64
	//valueCheckers[setting.FieldFloat64Arr] = CheckFloat64Arr
	//valueCheckers[setting.FieldString] = CheckString
	valueCheckers[setting.FieldStringArr] = CheckStringArr
	//valueCheckers[setting.FieldJson] = CheckJson
	//valueCheckers[setting.FieldJsonArr] = CheckJsonArr
}

func CheckData(dataType string, data string) bool {
	if checkFunc, ok := valueCheckers[dataType]; ok {
		return checkFunc(data)
	}
	return false
}

func CheckFormat(dataType string, data string) bool {
	if checkRegexp, ok := valueRegexps[dataType]; ok {
		if nil == checkRegexp {
			return false
		}
		return checkRegexp.MatchString(data)
	}
	return true
}

//func CheckBool(data string) bool {
//	data = strings.ToLower(data)
//	return data == "0" || data == "1" || data == "true" || data == "false"
//}
//
//func CheckBoolArr(data string) bool {
//	return CheckFormat(setting.FieldBoolArr, data)
//}
//
//func CheckInt8(data string) bool {
//	if CheckFormat(setting.FieldInt8, data) {
//		_, err := strconv.ParseInt(data, 10, 8)
//		return err != nil
//	}
//	return false
//}
//
//func CheckInt8Arr(data string) bool {
//	return CheckFormat(setting.FieldInt8Arr, data)
//}
//
//func CheckInt16(data string) bool {
//	if CheckFormat(setting.FieldInt16, data) {
//		_, err := strconv.ParseInt(data, 10, 16)
//		return err != nil
//	}
//	return false
//}
//
//func CheckInt16Arr(data string) bool {
//	return CheckFormat(setting.FieldInt16Arr, data)
//}
//
//func CheckInt32(data string) bool {
//	if CheckFormat(setting.FieldInt32, data) {
//		_, err := strconv.ParseInt(data, 10, 32)
//		return err != nil
//	}
//	return false
//}
//
//func CheckInt32Arr(data string) bool {
//	return CheckFormat(setting.FieldInt32Arr, data)
//}
//
//func CheckInt64(data string) bool {
//	if CheckFormat(setting.FieldInt64, data) {
//		_, err := strconv.ParseInt(data, 10, 64)
//		return err != nil
//	}
//	return false
//}
//
//func CheckInt64Arr(data string) bool {
//	return CheckFormat(setting.FieldInt64Arr, data)
//}
//
//func CheckUInt8(data string) bool {
//	if CheckFormat(setting.FieldUint8, data) {
//		_, err := strconv.ParseUint(data, 10, 8)
//		return err != nil
//	}
//	return false
//}
//
//func CheckUInt8Arr(data string) bool {
//	return CheckFormat(setting.FieldUint8Arr, data)
//}
//
//func CheckUInt16(data string) bool {
//	if CheckFormat(setting.FieldUint16, data) {
//		_, err := strconv.ParseUint(data, 10, 16)
//		return err != nil
//	}
//	return false
//}
//
//func CheckUInt16Arr(data string) bool {
//	return CheckFormat(setting.FieldUint16Arr, data)
//}
//
//func CheckUInt32(data string) bool {
//	if CheckFormat(setting.FieldUint32, data) {
//		_, err := strconv.ParseUint(data, 10, 32)
//		return err != nil
//	}
//	return false
//}
//
//func CheckUInt32Arr(data string) bool {
//	return CheckFormat(setting.FieldUint32Arr, data)
//}
//
//func CheckUInt64(data string) bool {
//	if CheckFormat(setting.FieldUint64, data) {
//		_, err := strconv.ParseUint(data, 10, 64)
//		return err != nil
//	}
//	return false
//}
//
//func CheckUInt64Arr(data string) bool {
//	return CheckFormat(setting.FieldUint64Arr, data)
//}
//
//func CheckFloat32(data string) bool {
//	if CheckFormat(setting.FieldFloat32, data) {
//		_, err := strconv.ParseFloat(data, 32)
//		return err != nil
//	}
//	return false
//}
//
//func CheckFloat32Arr(data string) bool {
//	return CheckFormat(setting.FieldFloat32Arr, data)
//}
//
//func CheckFloat64(data string) bool {
//	if CheckFormat(setting.FieldFloat64, data) {
//		_, err := strconv.ParseFloat(data, 64)
//		return err != nil
//	}
//	return false
//}
//
//func CheckFloat64Arr(data string) bool {
//	return CheckFormat(setting.FieldFloat64Arr, data)
//}
//
//func CheckString(data string) bool {
//	return true
//}

func CheckStringArr(data string) bool {
	return CheckFormat(setting.FieldStringArr, data)
}

//
//func CheckJson(data string) bool {
//	return jsoniter.Valid([]byte(data))
//}
//
//func CheckJsonArr(data string) bool {
//	return CheckFormat(setting.FieldJsonArr, data)
//}
