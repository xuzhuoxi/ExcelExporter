package data

import "github.com/xuzhuoxi/ExcelExporter/src/setting"

const (
	countName = "count"
	dataName  = "data"
)

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
	// 把数据写到文件
	WriteDataToFile(filePath string) error
}

func GenBuilder(format string) IDataBuilder {
	if c, ok := builderMap[format]; ok {
		return c()
	}
	return nil
}

func init() {
	RegisterBuilder(setting.FileNameBin, newIBinDataBuilder)
	RegisterBuilder(setting.FileNameSql, nil)
	RegisterBuilder(setting.FileNameJson, newJsonDataBuilder)
	RegisterBuilder(setting.FileNameYaml, func() IDataBuilder {
		return newIViperDataBuilder(setting.FileNameYaml)
	})
	RegisterBuilder(setting.FileNameYml, func() IDataBuilder {
		return newIViperDataBuilder(setting.FileNameYml)
	})
	RegisterBuilder(setting.FileNameToml, func() IDataBuilder {
		return newIViperDataBuilder(setting.FileNameToml)
	})
	RegisterBuilder(setting.FileNameHcl, func() IDataBuilder {
		return newIViperDataBuilder(setting.FileNameHcl)
	})
	RegisterBuilder(setting.FileNameEnv, func() IDataBuilder {
		return newIViperDataBuilder(setting.FileNameEnv)
	})
	RegisterBuilder(setting.FileNameProperties, func() IDataBuilder {
		return newIViperDataBuilder(setting.FileNameProperties)
	})
}

var (
	builderMap = make(map[string]BuilderConstructor)
)

type BuilderConstructor func() IDataBuilder

func RegisterBuilder(format string, constructor BuilderConstructor) {
	builderMap[format] = constructor
}
