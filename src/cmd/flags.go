package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core"
	"regexp"
	"strconv"
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

func (fp *FlagParams) String() string {
	return fmt.Sprintf("FlagParams(Mode=%d, Lang=%s, Field=%s, File=%s, Source=%s, Target=%s)",
		fp.Mode, fp.LangRefs, fp.FieldTypes, fp.DataFileFormats, fp.Source, fp.Target)
}

func (fp *FlagParams) GetCommandParams() *CommandParams {
	titleOn := (fp.Mode & uint(core.TitleMark)) > 0
	dataOn := (fp.Mode & uint(core.DataMark)) > 0
	constOn := (fp.Mode & uint(core.ConstMark)) > 0
	langRefs := strings.Split(fp.LangRefs, ParamsSep)
	fieldTypes := strings.Split(fp.FieldTypes, ParamsSep)
	filedTypeValues := make([]int, len(fieldTypes))
	for index, o := range fieldTypes {
		filedTypeValues[index], _ = strconv.Atoi(o)
	}
	dataFileFormats := strings.Split(fp.DataFileFormats, ParamsSep)
	return &CommandParams{DataOn: dataOn, TitleOn: titleOn, ConstOn: constOn,
		FieldTypes: filedTypeValues, LangRefs: langRefs, DataFileFormats: dataFileFormats}
}

type CommandParams struct {
	TitleOn         bool
	DataOn          bool
	ConstOn         bool
	LangRefs        []string
	FieldTypes      []int
	DataFileFormats []string
}

func (o *CommandParams) String() string {
	return fmt.Sprintf("CommandParams(TitleOn=%t, DataOn=%t, ConstOn=%t, Langs=%s, Fields=%v, Files=%s)",
		o.TitleOn, o.DataOn, o.ConstOn, o.LangRefs, o.FieldTypes, o.DataFileFormats)
}

func (o *CommandParams) GenTitleContexts() (contexts []*core.TitleContext) {
	fieldLen := len(o.FieldTypes)
	langLen := len(o.LangRefs)
	if !o.TitleOn || fieldLen == 0 || langLen == 0 {
		return nil
	}
	len := fieldLen * langLen
	contexts = make([]*core.TitleContext, 0, len)
	for fieldIdx := 0; fieldIdx < fieldLen; fieldIdx += 1 {
		for langIdx := 0; langIdx < langLen; langIdx += 1 {
			context := &core.TitleContext{FieldTypeIndex: o.FieldTypes[fieldIdx], ProgramLanguage: o.LangRefs[langIdx]}
			contexts = append(contexts, context)
		}
	}
	return
}

func (o *CommandParams) GenDataContexts() (contexts []*core.DataContext) {
	fieldLen := len(o.FieldTypes)
	fileLen := len(o.DataFileFormats)
	if !o.DataOn || fieldLen == 0 || fileLen == 0 {
		return nil
	}
	len := fieldLen * fileLen
	contexts = make([]*core.DataContext, 0, len)
	for fieldIdx := 0; fieldIdx < fieldLen; fieldIdx += 1 {
		for fileIdx := 0; fileIdx < fileLen; fileIdx += 1 {
			context := &core.DataContext{FieldTypeIndex: o.FieldTypes[fieldIdx], DataFileFormat: o.DataFileFormats[fileIdx]}
			contexts = append(contexts, context)
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
	m, err := regexp.MatchString(`\d[,\d]+`, *fieldTypes)
	if nil != err {
		return nil, err
	}
	if !m {
		return nil, errors.New("-field config error! ")
	}
	return &FlagParams{Mode: *mode, FieldTypes: *fieldTypes, LangRefs: *langRefs, DataFileFormats: *dataFileFormats, Source: *source, Target: *target}, nil
}
