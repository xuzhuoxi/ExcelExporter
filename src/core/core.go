package core

type HandleMark uint

const (
	DataMark HandleMark = 1 << iota
	DefinitionMark
	ConstMark
)
