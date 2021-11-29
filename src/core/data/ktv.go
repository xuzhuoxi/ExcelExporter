package data

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"strconv"
	"strings"
)

type KTValue struct {
	Key   string // 键
	Type  string // 类型
	Value string // 字符串值

	cacheValue interface{}
	cacheErr   error
}

func (ktv *KTValue) IsEmpty() bool {
	return ktv.Value == ""
}

func (ktv *KTValue) Set(k, t, v string) {
	ktv.Key, ktv.Type, ktv.Value = k, t, v
	ktv.cacheValue, ktv.cacheErr = nil, nil
}

func (ktv *KTValue) String() string {
	return fmt.Sprintf("{Key=%v,Type=%v,Value=%v}", ktv.Key, ktv.Type, ktv.Value)
}

func (ktv *KTValue) GetValue() (val interface{}, err error) {
	if ktv.cacheErr != nil || ktv.cacheValue != nil {
		return ktv.cacheValue, ktv.cacheErr
	}
	return ktv.updateCacheValue()
}

func (ktv *KTValue) isFixedString() bool {
	return regexpTypeFixedString.MatchString(ktv.Type)
}

func (ktv *KTValue) isFixedStringArr() bool {
	return regexpTypeFixedStringArr.MatchString(ktv.Type)
}

func (ktv *KTValue) updateCacheValue() (val interface{}, err error) {
	switch ktv.Type {
	case setting.FieldBool:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueBool()
	case setting.FieldBoolArr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueBoolArr()
	case setting.FieldInt8:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt8()
	case setting.FieldInt8Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt8Arr()
	case setting.FieldInt16:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt16()
	case setting.FieldInt16Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt16Arr()
	case setting.FieldInt32:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt32()
	case setting.FieldInt32Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt32Arr()
	case setting.FieldInt64:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt64()
	case setting.FieldInt64Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueInt64Arr()
	case setting.FieldUint8:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint8()
	case setting.FieldUint8Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint8Arr()
	case setting.FieldUint16:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint16()
	case setting.FieldUint16Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint16Arr()
	case setting.FieldUint32:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint32()
	case setting.FieldUint32Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint32Arr()
	case setting.FieldUint64:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint64()
	case setting.FieldUint64Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueUint64Arr()
	case setting.FieldFloat32:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueFloat32()
	case setting.FieldFloat32Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueFloat32Arr()
	case setting.FieldFloat64:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueFloat64()
	case setting.FieldFloat64Arr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueFloat64Arr()
	case setting.FieldString:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueString()
	case setting.FieldStringArr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueStringArr()
	case setting.FieldJson:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueJson()
	case setting.FieldJsonArr:
		ktv.cacheValue, ktv.cacheErr = ktv.toValueJsonArr()
	default:
		if ktv.isFixedString() {
			ktv.cacheValue, ktv.cacheErr = ktv.toValueFixedString()
			return ktv.cacheValue, ktv.cacheErr
		}
		if ktv.isFixedStringArr() {
			ktv.cacheValue, ktv.cacheErr = ktv.toValueFixedStringArr()
			return ktv.cacheValue, ktv.cacheErr
		}
		ktv.cacheValue, ktv.cacheErr = nil, ErrTypeUndefined
	}
	return ktv.cacheValue, ktv.cacheErr
}

func (ktv *KTValue) ValueBool() (value bool, err error) {
	if ktv.Type != setting.FieldBool {
		return false, ErrTypeMismatch
	}
	return ktv.toValueBool()
}

func (ktv *KTValue) ValueBoolArr() (value []bool, err error) {
	if ktv.Type != setting.FieldBoolArr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueBoolArr()
}

func (ktv *KTValue) ValueInt8() (value int8, err error) {
	if ktv.Type != setting.FieldInt8 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueInt8()
}

func (ktv *KTValue) ValueInt8Arr() (value []int8, err error) {
	if ktv.Type != setting.FieldInt8Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueInt8Arr()
}

func (ktv *KTValue) ValueInt16() (value int16, err error) {
	if ktv.Type != setting.FieldInt16 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueInt16()
}

func (ktv *KTValue) ValueInt16Arr() (value []int16, err error) {
	if ktv.Type != setting.FieldInt16Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueInt16Arr()
}

func (ktv *KTValue) ValueInt32() (value int32, err error) {
	if ktv.Type != setting.FieldInt32 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueInt32()
}

func (ktv *KTValue) ValueInt32Arr() (value []int32, err error) {
	if ktv.Type != setting.FieldInt32Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueInt32Arr()
}

func (ktv *KTValue) ValueInt64() (value int64, err error) {
	if ktv.Type != setting.FieldInt64 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueInt64()
}

func (ktv *KTValue) ValueInt64Arr() (value []int64, err error) {
	if ktv.Type != setting.FieldInt64Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueInt64Arr()
}

func (ktv *KTValue) ValueUint8() (value uint8, err error) {
	if ktv.Type != setting.FieldUint8 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueUint8()
}

func (ktv *KTValue) ValueUint8Arr() (value []uint8, err error) {
	if ktv.Type != setting.FieldUint8Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueUint8Arr()
}

func (ktv *KTValue) ValueUint16() (value uint16, err error) {
	if ktv.Type != setting.FieldUint16 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueUint16()
}

func (ktv *KTValue) ValueUint16Arr() (value []uint16, err error) {
	if ktv.Type != setting.FieldUint16Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueUint16Arr()
}

func (ktv *KTValue) ValueUint32() (value uint32, err error) {
	if ktv.Type != setting.FieldUint32 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueUint32()
}

func (ktv *KTValue) ValueUint32Arr() (value []uint32, err error) {
	if ktv.Type != setting.FieldUint32Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueUint32Arr()
}

func (ktv *KTValue) ValueUint64() (value uint64, err error) {
	if ktv.Type != setting.FieldUint64 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueUint64()
}

func (ktv *KTValue) ValueUint64Arr() (value []uint64, err error) {
	if ktv.Type != setting.FieldUint64Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueUint64Arr()
}

func (ktv *KTValue) ValueFloat32() (value float32, err error) {
	if ktv.Type != setting.FieldFloat32 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueFloat32()
}

func (ktv *KTValue) ValueFloat32Arr() (value []float32, err error) {
	if ktv.Type != setting.FieldFloat32Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueFloat32Arr()
}

func (ktv *KTValue) ValueFloat64() (value float64, err error) {
	if ktv.Type != setting.FieldFloat64 {
		return 0, ErrTypeMismatch
	}
	return ktv.toValueFloat64()
}

func (ktv *KTValue) ValueFloat64Arr() (value []float64, err error) {
	if ktv.Type != setting.FieldFloat64Arr {
		return nil, ErrTypeMismatch
	}
	return ktv.toValueFloat64Arr()
}

func (ktv *KTValue) toValueBool() (value bool, err error) {
	return strconv.ParseBool(ktv.Value)
}

func (ktv *KTValue) toValueBoolArr() (value []bool, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]bool, len(arr))
	for index := range arr {
		val, err := strconv.ParseBool(arr[index])
		if nil != err {
			return nil, err
		}
		value[index] = val
	}
	return
}

func (ktv *KTValue) toValueInt8() (value int8, err error) {
	val, err := strconv.ParseInt(ktv.Value, 10, 8)
	if nil != err {
		return 0, err
	}
	return int8(val), nil
}

func (ktv *KTValue) toValueInt8Arr() (value []int8, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]int8, len(arr))
	for index := range arr {
		val, err := strconv.ParseInt(arr[index], 10, 8)
		if nil != err {
			return nil, err
		}
		value[index] = int8(val)
	}
	return
}

func (ktv *KTValue) toValueInt16() (value int16, err error) {
	val, err := strconv.ParseInt(ktv.Value, 10, 16)
	if nil != err {
		return 0, err
	}
	return int16(val), nil
}

func (ktv *KTValue) toValueInt16Arr() (value []int16, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]int16, len(arr))
	for index := range arr {
		val, err := strconv.ParseInt(arr[index], 10, 16)
		if nil != err {
			return nil, err
		}
		value[index] = int16(val)
	}
	return
}

func (ktv *KTValue) toValueInt32() (value int32, err error) {
	val, err := strconv.ParseInt(ktv.Value, 10, 32)
	if nil != err {
		return 0, err
	}
	return int32(val), nil
}

func (ktv *KTValue) toValueInt32Arr() (value []int32, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]int32, len(arr))
	for index := range arr {
		val, err := strconv.ParseInt(arr[index], 10, 32)
		if nil != err {
			return nil, err
		}
		value[index] = int32(val)
	}
	return
}

func (ktv *KTValue) toValueInt64() (value int64, err error) {
	val, err := strconv.ParseInt(ktv.Value, 10, 64)
	if nil != err {
		return 0, err
	}
	return int64(val), nil
}

func (ktv *KTValue) toValueInt64Arr() (value []int64, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]int64, len(arr))
	for index := range arr {
		val, err := strconv.ParseInt(arr[index], 10, 64)
		if nil != err {
			return nil, err
		}
		value[index] = val
	}
	return
}

func (ktv *KTValue) toValueUint8() (value uint8, err error) {
	val, err := strconv.ParseUint(ktv.Value, 10, 8)
	if nil != err {
		return 0, err
	}
	return uint8(val), nil
}

func (ktv *KTValue) toValueUint8Arr() (value []uint8, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]uint8, len(arr))
	for index := range arr {
		val, err := strconv.ParseUint(arr[index], 10, 8)
		if nil != err {
			return nil, err
		}
		value[index] = uint8(val)
	}
	return
}

func (ktv *KTValue) toValueUint16() (value uint16, err error) {
	val, err := strconv.ParseUint(ktv.Value, 10, 16)
	if nil != err {
		return 0, err
	}
	return uint16(val), nil
}

func (ktv *KTValue) toValueUint16Arr() (value []uint16, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]uint16, len(arr))
	for index := range arr {
		val, err := strconv.ParseUint(arr[index], 10, 16)
		if nil != err {
			return nil, err
		}
		value[index] = uint16(val)
	}
	return
}

func (ktv *KTValue) toValueUint32() (value uint32, err error) {
	val, err := strconv.ParseUint(ktv.Value, 10, 32)
	if nil != err {
		return 0, err
	}
	return uint32(val), nil
}

func (ktv *KTValue) toValueUint32Arr() (value []uint32, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]uint32, len(arr))
	for index := range arr {
		val, err := strconv.ParseUint(arr[index], 10, 32)
		if nil != err {
			return nil, err
		}
		value[index] = uint32(val)
	}
	return
}

func (ktv *KTValue) toValueUint64() (value uint64, err error) {
	val, err := strconv.ParseUint(ktv.Value, 10, 64)
	if nil != err {
		return 0, err
	}
	return uint64(val), nil
}

func (ktv *KTValue) toValueUint64Arr() (value []uint64, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]uint64, len(arr))
	for index := range arr {
		val, err := strconv.ParseUint(arr[index], 10, 64)
		if nil != err {
			return nil, err
		}
		value[index] = val
	}
	return
}

func (ktv *KTValue) toValueFloat32() (value float32, err error) {
	val, err := strconv.ParseFloat(ktv.Value, 32)
	if nil != err {
		return 0, err
	}
	return float32(val), nil
}

func (ktv *KTValue) toValueFloat32Arr() (value []float32, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]float32, len(arr))
	for index := range arr {
		val, err := strconv.ParseFloat(arr[index], 32)
		if nil != err {
			return nil, err
		}
		value[index] = float32(val)
	}
	return
}

func (ktv *KTValue) toValueFloat64() (value float64, err error) {
	return strconv.ParseFloat(ktv.Value, 64)
}

func (ktv *KTValue) toValueFloat64Arr() (value []float64, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	value = make([]float64, len(arr))
	for index := range arr {
		val, err := strconv.ParseFloat(arr[index], 64)
		if nil != err {
			return nil, err
		}
		value[index] = val
	}
	return
}

func (ktv *KTValue) toValueString() (value string, err error) {
	return ktv.Value, nil
}

func (ktv *KTValue) toValueStringArr() (value []string, err error) {
	return ktv.valueToArray()
}

func (ktv *KTValue) toValueFixedString() (value string, err error) {
	n, err := ktv.getTypeN()
	if nil != err {
		return "", err
	}
	return ktv.toFixedRuneStr(ktv.Value, n), nil
}

func (ktv *KTValue) toValueFixedStringArr() (value []string, err error) {
	arr, err := ktv.valueToArray()
	if nil != err {
		return nil, err
	}
	n, err := ktv.getTypeArrN()
	if nil != err {
		return nil, err
	}
	value = make([]string, len(arr))
	for index := range arr {
		value[index] = ktv.toFixedRuneStr(arr[index], n)
	}
	return
}

func (ktv *KTValue) toValueJson() (value string, err error) {
	return ktv.Value, nil
}

func (ktv *KTValue) toValueJsonArr() (value []string, err error) {
	return ktv.valueToArray()
}

func (ktv *KTValue) valueToArray() (arr []string, err error) {
	if ktv.Value == "" || ktv.Value == "[]" {
		return ArrEmptyString, nil
	}
	if !CheckStringArr(ktv.Value) {
		return nil, ErrValueFormatWrong
	}
	str := ktv.Value[1 : len(ktv.Value)-1]
	return strings.Split(str, ","), nil
}

func (ktv *KTValue) getTypeN() (n int, err error) {
	nStr := ktv.Type[len("string(") : len(ktv.Type)-1]
	rs, err := strconv.ParseUint(nStr, 10, 16)
	if nil != err {
		return 0, err
	}
	return int(rs), nil
}

func (ktv *KTValue) getTypeArrN() (n int, err error) {
	nStr := ktv.Type[len("[]string(") : len(ktv.Type)-1]
	rs, err := strconv.ParseUint(nStr, 10, 16)
	if nil != err {
		return 0, err
	}
	return int(rs), nil
}

func (ktv *KTValue) toFixedRuneStr(str string, ln int) string {
	rn := []rune(str)
	strLen := len(rn)
	if strLen > ln {
		return string(rn[:ln])
	}
	if strLen < ln {
		rs := make([]rune, ln)
		copy(rs, rn)
		for index := strLen; index < ln; index += 1 {
			rs[index] = rune(' ')
		}
		return string(rs)
	}
	return ktv.Value
}
