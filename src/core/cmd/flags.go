package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"strings"
)

const (
	ParamsSep = ","
)

func ModeNameToType(modeName string) core.ModeType {
	if modeName == setting.ModeNameTitle {
		return core.ModeTitle
	}
	if modeName == setting.ModeNameData {
		return core.ModeData
	}
	if modeName == setting.ModeNameConst {
		return core.ModeConst
	}
	return core.ModeNone
}

func FieldRangeNameToType(rangeName string) core.FieldRangeType {
	if rangeName == setting.FieldRangeNameClient {
		return core.FieldRangeClient
	}
	if rangeName == setting.FieldRangeNameServer {
		return core.FieldRangeServer
	}
	if rangeName == setting.FieldRangeNameDb {
		return core.FieldRangeDatabase
	}
	return core.FieldRangeNone
}

func ParseFlag() (cfg *SysFlags, err error) {
	modes := flag.String("mode", "", "Running Mode! ")
	ranges := flag.String("range", "", "Use Fields! ")
	langRefs := flag.String("lang", "", "Use Languages! ")
	dataFiles := flag.String("file", "", "Output Files! ")
	source := flag.String("source", "", "Source Redefine! ")
	target := flag.String("target", "", "Target Redefine! ")
	flag.Parse()

	modesVal := strings.ToLower(*modes)
	rangesVal := strings.ToLower(*ranges)
	return &SysFlags{Modes: modesVal, Ranges: rangesVal, LangRefs: *langRefs, DataFiles: *dataFiles,
		Source: *source, Target: *target}, nil
}

type SysFlags struct {
	Modes     string
	Ranges    string
	LangRefs  string
	DataFiles string
	Source    string
	Target    string
}

func (f *SysFlags) String() string {
	return fmt.Sprintf("SysFlags(Modes=%s, Rangs=%s, LangRefs=%s, DataFiles=%s, Source=%s, Target=%s)",
		f.Modes, f.Ranges, f.LangRefs, f.DataFiles, f.Source, f.Target)
}

func (f *SysFlags) GetCommandParams() *AppFlags {
	modeNames, modeValues := f.parseModes()
	if len(modeValues) == 0 {
		panic(errors.New("Command -mode error! "))
	}
	rangeNames, rangeValues := f.parseFieldRanges()
	if len(rangeValues) == 0 {
		panic(errors.New("Command -range error! "))
	}

	langRefs := strings.Split(f.LangRefs, ParamsSep)
	dataFiles := strings.Split(f.DataFiles, ParamsSep)
	return &AppFlags{ModeNames: modeNames, ModeTypes: modeValues, RangeNames: rangeNames, RangeTypes: rangeValues,
		LangRefs: langRefs, DataFiles: dataFiles}
}

func (f *SysFlags) parseModes() (names []string, types []core.ModeType) {
	modes := strings.Split(f.Modes, ParamsSep)
	if len(modes) == 0 {
		return nil, nil
	}
	names = make([]string, 0, len(modes))
	types = make([]core.ModeType, 0, len(modes))
	for _, o := range modes {
		t := ModeNameToType(o)
		if t == core.ModeNone {
			continue
		}
		names = append(names, o)
		types = append(types, t)
	}
	return
}

func (f *SysFlags) parseFieldRanges() (names []string, types []core.FieldRangeType) {
	ranges := strings.Split(f.Ranges, ParamsSep)
	if len(ranges) == 0 {
		return nil, nil
	}
	names = make([]string, 0, len(ranges))
	types = make([]core.FieldRangeType, 0, len(ranges))
	for _, o := range ranges {
		t := FieldRangeNameToType(o)
		if t == core.FieldRangeNone {
			continue
		}
		names = append(names, o)
		types = append(types, t)
	}
	return
}

type AppFlags struct {
	ModeNames  []string
	ModeTypes  []core.ModeType
	RangeNames []string
	RangeTypes []core.FieldRangeType
	LangRefs   []string
	DataFiles  []string
}

func (o *AppFlags) CheckMode(mode core.ModeType) bool {
	for _, m := range o.ModeTypes {
		if m == mode {
			return true
		}
	}
	return false
}

func (o *AppFlags) String() string {
	return fmt.Sprintf("AppFlags(ModeNames=%v, ModeTypes=%v, RangeNames=%v, RangeTypes=%v, LangRefs=%v, DataFiles=%v)",
		o.ModeNames, o.ModeTypes, o.RangeNames, o.RangeTypes, o.LangRefs, o.DataFiles)
}

func (o *AppFlags) GenTitleContexts() (contexts []*core.TitleContext) {
	if !o.CheckMode(core.ModeTitle) {
		return nil
	}
	rangeLen := len(o.RangeTypes)
	langLen := len(o.LangRefs)
	if rangeLen == 0 || langLen == 0 {
		return nil
	}
	ln := rangeLen * langLen
	contexts = make([]*core.TitleContext, 0, ln)
	for fieldIdx := 0; fieldIdx < rangeLen; fieldIdx += 1 {
		for langIdx := 0; langIdx < langLen; langIdx += 1 {
			context := &core.TitleContext{RangeName: o.RangeNames[fieldIdx], RangeType: o.RangeTypes[fieldIdx],
				ProgramLanguage: o.LangRefs[langIdx]}
			contexts = append(contexts, context)
		}
	}
	return
}

func (o *AppFlags) GenDataContexts() (contexts []*core.DataContext) {
	if !o.CheckMode(core.ModeData) {
		return nil
	}
	rangeLen := len(o.RangeTypes)
	fileLen := len(o.DataFiles)
	if rangeLen == 0 || fileLen == 0 {
		return nil
	}
	ln := rangeLen * fileLen
	contexts = make([]*core.DataContext, 0, ln)
	for fieldIdx := 0; fieldIdx < rangeLen; fieldIdx += 1 {
		for fileIdx := 0; fileIdx < fileLen; fileIdx += 1 {
			context := &core.DataContext{RangeName: o.RangeNames[fieldIdx], RangeType: o.RangeTypes[fieldIdx],
				DataFileFormat: o.DataFiles[fileIdx]}
			contexts = append(contexts, context)
		}
	}
	return
}

func (o *AppFlags) GenConstContexts() (contexts []*core.ConstContext) {
	if !o.CheckMode(core.ModeConst) {
		return nil
	}
	return nil
}
