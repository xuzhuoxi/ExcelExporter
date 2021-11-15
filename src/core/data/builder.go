package data

type IDataBuilder interface {
	WriteString(key string, value string)
	WriteData(ktv *KTValue)
	WriteDataArray(key string, ktvArr []*KTValue)
}
