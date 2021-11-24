package data

import (
	"fmt"
	"github.com/spf13/viper"
)

func newCharDataBuilder(format string) IDataBuilder {
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
	for index := range ktvArr {
		err := b.WriteCell(ktvArr[index])
		if nil != err {
			return err
		}
	}
	b.StartNewRow()
	return nil
}

func (b *charDataBuilder) WriteDataToFile(filePath string) error {
	return b.dataViper.WriteConfigAs(filePath)
}
