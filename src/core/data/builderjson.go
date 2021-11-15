package data

import (
	"fmt"
	"github.com/tidwall/sjson"
)

type BuilderJson struct {
	jsonText string
}

func (b *BuilderJson) WriteString(key string, value string) {
	panic("implement me")
}

func (b *BuilderJson) WriteData(key string, ktv *KTValue) error {
	text, err := sjson.Set(b.jsonText, key, ktv.Value)
	if nil != err {
		return err
	}
	b.jsonText = text
	return nil
}

func (b *BuilderJson) WriteDataArray(key string, ktvArr []*KTValue) error {
	var text string
	var err error
	for index, value := range ktvArr {
		text, err = sjson.Set(b.jsonText, fmt.Sprintf("%s.%d", key, index), value.Value)
		if nil != err {
			return err
		}
		b.jsonText = text
	}
	return nil
}
