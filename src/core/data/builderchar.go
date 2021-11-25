package data

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
)

func newICharDataBuilder(format string) IDataBuilder {
	return newCharDataBuilder(format)
}

func newCharDataBuilder(format string) *charDataBuilder {
	dataViper := viper.New()
	dataViper.SetConfigType(format)
	return &charDataBuilder{dataViper: dataViper}
}

type charDataBuilder struct {
	dataViper *viper.Viper
	rowIndex  int
}

func (b *charDataBuilder) StartWriteData() {
	b.rowIndex = 0
	b.dataViper.Set("data2", make([]int64, 2))
}

func (b *charDataBuilder) StartNewRow() {
	b.rowIndex += 1
}

func (b *charDataBuilder) FinishWriteData() {
	return
}

func (b *charDataBuilder) WriteCell(ktv *KTValue) error {
	v, err := ktv.GetValue()
	if nil != err {
		return err
	}
	path := fmt.Sprintf("data.%d.%s", b.rowIndex, ktv.Key)
	b.dataViper.Set(path, v)
	//fmt.Println("WriteCell:", path, v)
	return nil
}

func (b *charDataBuilder) WriteRow(ktvArr []*KTValue) error {
	for _, ktv := range ktvArr {
		err := b.WriteCell(ktv)
		if nil != err {
			return err
		}
	}
	b.StartNewRow()

	b.dataViper.AllSettings()

	return nil
}

func (b *charDataBuilder) WriteRow2(ktvArr []*KTValue) error {
	row := make(map[string]interface{})
	for _, ktv := range ktvArr {
		value, err := ktv.GetValue()
		if nil != err {
			return err
		}
		row[ktv.Key] = value
	}
	path := fmt.Sprintf("data.%d", b.rowIndex)
	b.dataViper.Set(path, row)
	b.StartNewRow()
	return nil
}

func (b *charDataBuilder) WriteDataToFile(filePath string) error {
	return b.dataViper.WriteConfigAs(filePath)
}

func newJsonDataBuilder() IDataBuilder {
	return &jsonDataBuilder{charDataBuilder: *newCharDataBuilder(setting.FileNameJson)}
}

type jsonDataBuilder struct {
	charDataBuilder
}

func (b *jsonDataBuilder) WriteDataToFile(filePath string) error {
	return b.charDataBuilder.WriteDataToFile(filePath)
}
