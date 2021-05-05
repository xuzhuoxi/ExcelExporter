package cmd

import (
	"errors"
	"flag"
	"github.com/xuzhuoxi/ExcelExporter/src/core"
	"strings"
)

const (
	ParamsSep = ","
)

type FlagParams struct {
	Mode   uint
	Field  string
	Lang   string
	File   string
	Source string
	Target string
}

func (fp *FlagParams) GetCommandParams() *CommandParams {
	dataOn := (fp.Mode & uint(core.DataMark)) > 0
	definitionOn := (fp.Mode & uint(core.DefinitionMark)) > 0
	constOn := (fp.Mode & uint(core.ConstMark)) > 0
	fields := strings.Split(fp.Field, ParamsSep)
	langs := strings.Split(fp.File, ParamsSep)
	files := strings.Split(fp.Field, ParamsSep)
	return &CommandParams{HandleData: dataOn, HandleDefinition: definitionOn, HandleConst: constOn,
		Fields: fields, Lands: langs, Files: files}
}

type CommandParams struct {
	HandleData       bool
	HandleDefinition bool
	HandleConst      bool
	Fields           []string
	Lands            []string
	Files            []string
}

func (o *CommandParams) GenDataContexts() (contexts []*core.DataContext) {
	fieldLen := len(o.Fields)
	fileLen := len(o.Files)
	if !o.HandleData || fieldLen == 0 || fileLen == 0 {
		return nil
	}
	len := fieldLen * fileLen
	contexts = make([]*core.DataContext, len)
	for i := 0; i < fieldLen; i += 1 {
		for j := 0; j < fileLen; j += 1 {
			context := &core.DataContext{DataField: o.Fields[i], DataFile: o.Files[j]}
			contexts[i*j] = context
		}
	}
	return
}

func (o *CommandParams) GenDefinitionContexts() (contexts []*core.DefinitionContext) {
	fieldLen := len(o.Fields)
	langLen := len(o.Lands)
	if !o.HandleDefinition || fieldLen == 0 || langLen == 0 {
		return nil
	}
	len := fieldLen * langLen
	contexts = make([]*core.DefinitionContext, len)
	for i := 0; i < fieldLen; i += 1 {
		for j := 0; j < langLen; j += 1 {
			context := &core.DefinitionContext{DefinedField: o.Fields[i], DefinedLang: o.Lands[j]}
			contexts[i*j] = context
		}
	}
	return
}

func ParseFlag() (cfg *FlagParams, err error) {
	mode := flag.Uint("mode", 0, "Running Mode! ")
	field := flag.String("field", "", "Use Fields! ")
	lang := flag.String("lang", "", "Use Languages! ")
	file := flag.String("file", "", "Output Files! ")
	source := flag.String("source", "", "Source Redefine! ")
	target := flag.String("target", "", "Target Redefine! ")
	flag.Parse()

	if *mode == 0 {
		return nil, errors.New("Mode Error! ")
	}
	if len(*field) == 0 {
		return nil, errors.New("Field Error! ")
	}
	return &FlagParams{Mode: *mode, Field: *field, Lang: *lang, File: *file, Source: *source, Target: *target}, nil
}
