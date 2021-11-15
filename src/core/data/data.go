package data

import "errors"

var (
	ErrTypeMismatch     = errors.New("Type Mismatch(类型不匹配)! ")
	ErrTypeUndefined    = errors.New("Type Undefined(类型未定义)! ")
	ErrValueFormatWrong = errors.New("Value Format Wrong(内容格式错误)! ")
)

var (
	ArrEmptyString = make([]string, 0)
)
