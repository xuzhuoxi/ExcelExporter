package data

import (
	"bytes"
	"encoding/binary"
	"github.com/xuzhuoxi/infra-go/filex"
)

func newIBinDataBuilderDefault() IDataBuilder {
	return newBinDataBuilderDefault()
}

func newBinDataBuilderDefault() *binaryDataBuilder {
	return newBinDataBuilder(binary.BigEndian)
}

func newBinDataBuilder(order binary.ByteOrder) *binaryDataBuilder {
	return &binaryDataBuilder{order: order}
}

const (
	rowCountSize = 4
)

var (
	rowIndexValue = []byte{0, 0, 0, 0}
)

type binaryDataBuilder struct {
	order    binary.ByteOrder
	title    []byte
	content  *bytes.Buffer
	rowIndex int
}

func (b *binaryDataBuilder) StartWriteData() {
	b.title = make([]byte, rowCountSize, 1024)
	b.content = bytes.NewBuffer(nil)
	b.rowIndex = 0
	b.logRowIndex()
}

func (b *binaryDataBuilder) FinishWriteData() {
	return
}

func (b *binaryDataBuilder) WriteRow(ktvArr []*KTValue) error {
	for _, ktv := range ktvArr {
		err := b.writeCell(ktv)
		if nil != err {
			return err
		}
	}
	b.startNewRow()
	return nil
}

func (b *binaryDataBuilder) WriteDataToFile(filePath string) error {
	data := append(b.title, b.content.Bytes()...)
	return filex.WriteFile(filePath, data, filex.ModePerm)
}

func (b *binaryDataBuilder) startNewRow() {
	b.rowIndex += 1
	b.logRowIndex()
}

func (b *binaryDataBuilder) writeCell(ktv *KTValue) error {
	v, err := ktv.GetValue()
	if nil != err {
		return err
	}
	switch value := v.(type) {
	case string:
		b.writeString(value)
		return nil
	case *string:
		b.writeString(*value)
		return nil
	case []string:
		b.writeStringArr(value)
		return nil
	default:
		return binary.Write(b.content, b.order, v)
	}
}

func (b *binaryDataBuilder) logRowIndex() {
	b.order.PutUint32(b.title[:rowCountSize], uint32(b.rowIndex))
	if b.rowIndex == 0 {
		return
	}
	b.order.PutUint32(rowIndexValue, uint32(b.content.Len()))
	b.title = append(b.title, rowIndexValue...)
}

func (b *binaryDataBuilder) writeString(str string) {
	bs := []byte(str)
	binary.Write(b.content, b.order, uint16(len(bs)))
	binary.Write(b.content, b.order, bs)
}

func (b *binaryDataBuilder) writeStringArr(strArr []string) {
	binary.Write(b.content, b.order, uint16(len(strArr)))
	for index := range strArr {
		b.writeString(strArr[index])
	}
}
