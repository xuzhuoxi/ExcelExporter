package data

import (
	"fmt"
	"github.com/tidwall/sjson"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"os"
	"io/fs"
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

func (b *jsonDataBuilder) FinishWriteData() {
	c, _ := sjson.Set(b.content, countName, b.rowIndex+1)
	b.content = c
}

func (b *jsonDataBuilder) WriteRow(ktvArr []*KTValue) error {
	for _, ktv := range ktvArr {
		err := b.writeCell(ktv)
		if nil != err {
			return err
		}
	}
	b.startNewRow()
	return nil
}

func (b *jsonDataBuilder) WriteDataToFile(filePath string) error {
	return os.WriteFile(filePath, []byte(b.content), fs.ModePerm)
}

func (b *jsonDataBuilder) startNewRow() {
	b.rowIndex += 1
}

func (b *jsonDataBuilder) writeCell(ktv *KTValue) error {
	v, err := ktv.GetValue()
	if nil != err {
		return err
	}
	if ktv.Type == setting.FieldUint8Arr {
		v = b.unit8ArrToUint16Arr(v)
	}
	path := fmt.Sprintf("%s.%d.%s", dataName, b.rowIndex, ktv.Key)
	c, err := sjson.Set(b.content, path, v)
	if nil != err {
		return err
	}
	b.content = c
	return nil
}

func (b *jsonDataBuilder) unit8ArrToUint16Arr(uint8Arr interface{}) interface{} {
	v := uint8Arr.([]uint8)
	if nil == v {
		return nil
	}
	if len(v) == 0 {
		return []uint16{}
	}
	rs := make([]uint16, len(v))
	for index := range v {
		rs[index] = uint16(v[index])
	}
	return rs
}
