package data

type IDataBuilder interface {
	// 开始写入数据
	StartWriteData()
	// 开始新一行
	StartNewRow()
	// 开始写入数据
	FinishWriteData()
	// 写入一个单元
	WriteCell(ktv *KTValue) error
	// 写入一行数据
	WriteRow(ktvArr []*KTValue) error
}

func GenBuilder(format string) IDataBuilder {
	if c, ok := builderMap[format]; ok {
		return c()
	}
	return nil
}

func init() {
	RegisterBuilder("json", newCharDataBuilder)
	RegisterBuilder("yaml", newCharDataBuilder)
}

var (
	builderMap = make(map[string]BuilderConstructor)
)

type BuilderConstructor func() IDataBuilder

func RegisterBuilder(format string, constructor BuilderConstructor) {
	builderMap[format] = constructor
}
