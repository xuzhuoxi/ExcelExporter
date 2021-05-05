package core

type DataContext struct {
	// 数据来源
	Source string
	// 数据输出
	Target string
	// 使用的字段类型
	DataField string
	// 输出的文件类型
	DataFile string
}

type DefinitionContext struct {
	// 数据来源
	Source string
	// 数据输出
	Target string
	// 使用的字段类型
	DefinedField string
	// 使用的编程语言
	DefinedLang string
}

type ConstContext struct {
	// 数据来源
	Source string
	// 数据输出
	Target string
}
