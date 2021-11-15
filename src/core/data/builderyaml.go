package data

import "fmt"

type BuilderYaml struct {
	textBuilder *TextBuilder
}

func (b *BuilderYaml) WriteKey(key string) {
	b.textBuilder.WriteText(fmt.Sprintf("'%s':", key))
}

func (b *BuilderYaml) WriteStringValue(key string, value string) {
	b.textBuilder.WriteNewLine()
	b.textBuilder.WriteIndent()
	b.textBuilder.WriteText(fmt.Sprintf("'%s' : '%s'", key, value))
}
