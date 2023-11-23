package core

// ModeType 执行模式
type ModeType uint

const (
	ModeNone ModeType = iota
	ModeTitle
	ModeData
	ModeConst
	ModeProto
)

// FieldRangeType 字段类型
type FieldRangeType uint

const (
	FieldRangeNone FieldRangeType = iota
	FieldRangeClient
	FieldRangeServer
	FieldRangeDatabase
)

// RegexPatternRange 字段类型格式对应 正则表达式
const RegexPatternRange = `[01],[01],[01]`
