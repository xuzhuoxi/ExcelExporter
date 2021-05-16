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
	Mode            uint
	LangRefs        string
	FieldTypes      string
	DataFileFormats string
	Source          string
	Target          string
}

func (fp *FlagParams) GetCommandParams() *CommandParams {
	titleOn := (fp.Mode & uint(core.TitleMark)) > 0
	dataOn := (fp.Mode & uint(core.DataMark)) > 0
	constOn := (fp.Mode & uint(core.ConstMark)) > 0
	langRefs := strings.Split(fp.LangRefs, ParamsSep)
	fieldTypes := strings.Split(fp.FieldTypes, ParamsSep)
	dataFileFormats := strings.Split(fp.DataFileFormats, ParamsSep)
	return &CommandParams{HandleData: dataOn, HandleTitle: titleOn, HandleConst: constOn,
		FieldTypes: fieldTypes, LangRefs: langRefs, DataFileFormats: dataFileFormats}
}

type CommandParams struct {
	HandleTitle     bool
	HandleData      bool
	HandleConst     bool
	LangRefs        []string
	FieldTypes      []string
	DataFileFormats []string
}

func (o *CommandParams) GenDataContexts() (contexts []*core.DataContext) {
	fieldLen := len(o.FieldTypes)
	fileLen := len(o.DataFileFormats)
	if !o.HandleData || fieldLen == 0 || fileLen == 0 {
		return nil
	}
	len := fieldLen * fileLen
	contexts = make([]*core.DataContext, len)
	for i := 0; i < fieldLen; i += 1 {
		for j := 0; j < fileLen; j += 1 {
			context := &core.DataContext{DataField: o.FieldTypes[i], DataFile: o.DataFileFormats[j]}
			contexts[i*j] = context
		}
	}
	return
}

func (o *CommandParams) GenTitleContexts() (contexts []*core.TitleContext) {
	fieldLen := len(o.FieldTypes)
	langLen := len(o.LangRefs)
	if !o.HandleTitle || fieldLen == 0 || langLen == 0 {
		return nil
	}
	len := fieldLen * langLen
	contexts = make([]*core.TitleContext, len)
	for i := 0; i < fieldLen; i += 1 {
		for j := 0; j < langLen; j += 1 {
			context := &core.TitleContext{TitleField: o.FieldTypes[i], TitleLang: o.LangRefs[j]}
			contexts[i*j] = context
		}
	}
	return
}

func ParseFlag() (cfg *FlagParams, err error) {
	mode := flag.Uint("mode", 0, "Running Mode! ")
	langRefs := flag.String("lang", "", "Use Languages! ")
	fieldTypes := flag.String("field", "", "Use Fields! ")
	dataFileFormats := flag.String("file", "", "Output Files! ")
	source := flag.String("source", "", "Source Redefine! ")
	target := flag.String("target", "", "Target Redefine! ")
	flag.Parse()

	if *mode == 0 {
		return nil, errors.New("Mode Error! ")
	}
	if len(*fieldTypes) == 0 {
		return nil, errors.New("Field Type Error! ")
	}
	return &FlagParams{Mode: *mode, FieldTypes: *fieldTypes, LangRefs: *langRefs, DataFileFormats: *dataFileFormats, Source: *source, Target: *target}, nil
}
