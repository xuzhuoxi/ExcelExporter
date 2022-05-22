package core

type ModeType uint

const (
	ModeNone ModeType = iota
	ModeTitle
	ModeData
	ModeConst
)

type FieldRangeType uint

const (
	FieldRangeNone FieldRangeType = iota
	FieldRangeClient
	FieldRangeServer
	FieldRangeDatabase
)

const RegexPatternRange = `[01],[01],[01]`
