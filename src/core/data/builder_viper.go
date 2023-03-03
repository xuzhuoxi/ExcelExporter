package data

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func newIViperDataBuilder(format string) IDataBuilder {
	return NewViperDataBuilder(format)
}

func NewViperDataBuilder(format string) *viperDataBuilder {
	dataViper := viper.New()
	dataViper.SetConfigType(format)
	return &viperDataBuilder{dataViper: dataViper}
}

type viperDataBuilder struct {
	dataViper *viper.Viper
	rowIndex  int
}

func (b *viperDataBuilder) StartWriteData() {
	b.rowIndex = 0
}

func (b *viperDataBuilder) FinishWriteData() {
	b.dataViper.Set(countName, b.rowIndex+1)
	return
}

func (b *viperDataBuilder) WriteRow(ktvArr []*KTValue) error {
	for _, ktv := range ktvArr {
		err := b.writeCell(ktv)
		if nil != err {
			return errors.New(fmt.Sprintf("[%s]{%s}", ktv.Loc, err))
		}
	}
	b.startNewRow()
	return nil
}

func (b *viperDataBuilder) WriteRow2(ktvArr []*KTValue) error {
	row := make(map[string]interface{})
	for _, ktv := range ktvArr {
		value, err := ktv.GetValue()
		if nil != err {
			return err
		}
		row[ktv.Key] = value
	}
	path := fmt.Sprintf("%s.%d", dataName, b.rowIndex)
	b.dataViper.Set(path, row)
	b.startNewRow()
	return nil
}

func (b *viperDataBuilder) WriteDataToFile(filePath string) error {
	return b.dataViper.WriteConfigAs(filePath)
}

func (b *viperDataBuilder) startNewRow() {
	b.rowIndex += 1
}

func (b *viperDataBuilder) writeCell(ktv *KTValue) error {
	v, err := ktv.GetValue()
	if nil != err {
		return err
	}
	path := fmt.Sprintf("%s.%d.%s", dataName, b.rowIndex, ktv.Key)
	b.dataViper.Set(path, v)
	//fmt.Println("writeCell:", path, v)
	return nil
}
