package data

import "fmt"

type IDataBuilder interface {
	WriteString(key string, value string)
}

type IJsonBuilder struct {
}

type IYamlBuilder struct {
	textBuilder *TextBuilder
}

func (b *IYamlBuilder) WriteKey(key string) {
	b.textBuilder.WriteText(fmt.Sprintf("'%s':", key))
}

func (b *IYamlBuilder) WriteStringValue(key string, value string) {
	b.textBuilder.WriteNewLine()
	b.textBuilder.WriteIndent()
	b.textBuilder.WriteText(fmt.Sprintf("'%s' : '%s'", key, value))
}
