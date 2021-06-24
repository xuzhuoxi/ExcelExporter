package data

import (
	"strings"
)

var (
	charNewLine = "\n"
	charIndent  = "\t"
)

func NewTextBuilderDefault() *TextBuilder {
	return NewTextBuilder(charNewLine, charIndent)
}

func NewTextBuilder(line string, indent string) *TextBuilder {
	return &TextBuilder{line: line, indent: indent}
}

type TextBuilder struct {
	line   string
	indent string

	builder     strings.Builder
	indentLevel int
	indentStr   string
}

func (b *TextBuilder) IndentLevel() int {
	return b.indentLevel
}

func (b *TextBuilder) WriteNewLine() {
	b.builder.WriteString(b.line)
}

func (b *TextBuilder) WriteIndent() {
	b.builder.WriteString(b.indentStr)
}

func (b *TextBuilder) WriteText(text string) {
	b.builder.WriteString(text)
}

func (b *TextBuilder) ResetIndentLevel() {
	b.indentLevel = 0
	b.updateIndentStr()
}

func (b *TextBuilder) IncreaseIndentLevel() {
	b.indentLevel += 1
	b.updateIndentStr()
}

func (b *TextBuilder) DecreaseIndentLevel() {
	b.indentLevel -= 1
	if b.indentLevel < 0 {
		b.indentLevel = 0
	}
	b.updateIndentStr()
}

func (b *TextBuilder) updateIndentStr() {
	b.indentStr = ""
	for i := 0; i < b.indentLevel; i++ {
		b.indentStr += b.indent
	}
}
