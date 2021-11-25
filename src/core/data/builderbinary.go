package data

func newIBinDataBuilder() IDataBuilder {
	return newBinDataBuilder()
}

func newBinDataBuilder() *binaryDataBuilder {
	return &binaryDataBuilder{}
}

type binaryDataBuilder struct {
}

func (b *binaryDataBuilder) StartWriteData() {
	panic("implement me")
}

func (b *binaryDataBuilder) StartNewRow() {
	panic("implement me")
}

func (b *binaryDataBuilder) FinishWriteData() {
	panic("implement me")
}

func (b *binaryDataBuilder) WriteCell(ktv *KTValue) error {
	panic("implement me")
}

func (b *binaryDataBuilder) WriteRow(ktvArr []*KTValue) error {
	panic("implement me")
}

func (b *binaryDataBuilder) WriteDataToFile(filePath string) error {
	panic("implement me")
}
