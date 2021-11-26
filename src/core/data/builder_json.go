package data

import (
	"fmt"
	"github.com/tidwall/sjson"
	"io/fs"
	"os"
)

func newJsonDataBuilder() IDataBuilder {
	return &jsonDataBuilder{}
}

type jsonDataBuilder struct {
	rowIndex int
	content  string
}

func (b *jsonDataBuilder) StartWriteData() {
	b.rowIndex = 0
}

func (b *jsonDataBuilder) StartNewRow() {
	b.rowIndex += 1
}

func (b *jsonDataBuilder) FinishWriteData() {
	c, err := sjson.Set(b.content, countName, b.rowIndex+1)
	if nil != err {

	}
	b.content = c
}

func (b *jsonDataBuilder) WriteCell(ktv *KTValue) error {
	v, err := ktv.GetValue()
	if nil != err {
		return err
	}
	path := fmt.Sprintf("%s.%d.%s", dataName, b.rowIndex, ktv.Key)
	c, err := sjson.Set(b.content, path, v)
	if nil != err {
		return err
	}
	b.content = c
	return nil
}

func (b *jsonDataBuilder) WriteRow(ktvArr []*KTValue) error {
	for _, ktv := range ktvArr {
		err := b.WriteCell(ktv)
		if nil != err {
			return err
		}
	}
	b.StartNewRow()
	return nil
}

func (b *jsonDataBuilder) WriteDataToFile(filePath string) error {
	return os.WriteFile(filePath, []byte(b.content), fs.ModePerm)
}
