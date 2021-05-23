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
	Mode            string
	LangRefs        string
	FieldTypes      string
	DataFileFormats string
	Source          string
	Target          string
}

func (fp *FlagParams) String() string {
	return fmt.Sprintf("FlagParams(Mode=%s, Lang=%s, Field=%s, File=%s, Source=%s, Target=%s)",
		fp.Mode, fp.LangRefs, fp.FieldTypes, fp.DataFileFormats, fp.Source, fp.Target)
}

func (fp *FlagParams) GetCommandParams() *CommandParams {
	modes := strings.Split(fp.Mode, ParamsSep)
	modeValues := make([]core.ModeType, len(modes))
	for index, o := range modes {
		value, err := strconv.Atoi(o)
		if err != nil {
			panic(err)
		}
		if value < 0 {
			panic(errors.New("Command -mode error! "))
		}
		modeValues[index] = core.ModeType(value)
	}

	langRefs := strings.Split(fp.LangRefs, ParamsSep)
	fieldTypes := strings.Split(fp.FieldTypes, ParamsSep)
	filedTypeValues := make([]core.FieldType, len(fieldTypes))
	for index, o := range fieldTypes {
		value, err := strconv.Atoi(o)
		if err != nil {
			panic(err)
		}
		if value < 0 {
			panic(errors.New("Commnad -field error! "))
		}
		filedTypeValues[index] = core.FieldType(value)
	}
	dataFileFormats := strings.Split(fp.DataFileFormats, ParamsSep)
	return &CommandParams{ModeTypes: modeValues, LangRefs: langRefs, FieldTypes: filedTypeValues, DataFileFormats: dataFileFormats}
}

type CommandParams struct {
	ModeTypes       []core.ModeType
	LangRefs        []string
	FieldTypes      []core.FieldType
	DataFileFormats []string
}

func (o *CommandParams) CheckMode(mode core.ModeType) bool {
	for _, m := range o.ModeTypes {
		if m == mode {
			return true
		}
	}
	return false
}

func (o *CommandParams) String() string {
	return fmt.Sprintf("CommandParams(Mode=%v, Langs=%s, Fields=%v, Files=%s)", o.ModeTypes, o.LangRefs, o.FieldTypes, o.DataFileFormats)
}

func (o *CommandParams) GenTitleContexts() (contexts []*core.TitleContext) {
	fieldLen := len(o.FieldTypes)
	langLen := len(o.LangRefs)
	if !o.CheckMode(core.ModeTitle) || fieldLen == 0 || langLen == 0 {
		return nil
	}
	ln := fieldLen * langLen
	contexts = make([]*core.TitleContext, 0, ln)
	for fieldIdx := 0; fieldIdx < fieldLen; fieldIdx += 1 {
		for langIdx := 0; langIdx < langLen; langIdx += 1 {
			context := &core.TitleContext{FieldType: o.FieldTypes[fieldIdx], ProgramLanguage: o.LangRefs[langIdx]}
			contexts = append(contexts, context)
		}
	}
	return
}

func (o *CommandParams) GenDataContexts() (contexts []*core.DataContext) {
	fieldLen := len(o.FieldTypes)
	fileLen := len(o.DataFileFormats)
	if !o.CheckMode(core.ModeData) || fieldLen == 0 || fileLen == 0 {
		return nil
	}
	ln := fieldLen * fileLen
	contexts = make([]*core.DataContext, 0, ln)
	for fieldIdx := 0; fieldIdx < fieldLen; fieldIdx += 1 {
		for fileIdx := 0; fileIdx < fileLen; fileIdx += 1 {
			context := &core.DataContext{FieldType: o.FieldTypes[fieldIdx], DataFileFormat: o.DataFileFormats[fileIdx]}
			contexts = append(contexts, context)
		}
	}
	return
}

func ParseFlag() (cfg *FlagParams, err error) {
	mode := flag.String("mode", "", "Running Mode! ")
	langRefs := flag.String("lang", "", "Use Languages! ")
	fieldTypes := flag.String("field", "", "Use Fields! ")
	dataFileFormats := flag.String("file", "", "Output Files! ")
	source := flag.String("source", "", "Source Redefine! ")
	target := flag.String("target", "", "Target Redefine! ")
	flag.Parse()

	modeVal := *mode
	mm, err := regexp.MatchString(`\d[,\d]+`, modeVal)
	if nil != err {
		return nil, err
	}
	if !mm {
		return nil, errors.New("-mode config error! ")
	}

	langRefsVal := *langRefs
	if len(langRefsVal) == 0 {
		return nil, errors.New("Field Type Error! ")
	}

	fieldTypesVal := *fieldTypes
	fm, err := regexp.MatchString(`\d[,\d]+`, fieldTypesVal)
	if nil != err {
		return nil, err
	}
	if !fm {
		return nil, errors.New("-field config error! ")
	}

	return &FlagParams{Mode: modeVal, FieldTypes: fieldTypesVal, LangRefs: langRefsVal, DataFileFormats: *dataFileFormats,
		Source: *source, Target: *target}, nil
}
