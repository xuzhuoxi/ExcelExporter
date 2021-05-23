package core

type ModeType uint

const (
	ModeTitle ModeType = iota + 1
	ModeData
	ModeConst
)

type FieldType uint

const (
	FieldClient = iota + 1
	FieldServer
	FieldDatabase
)
