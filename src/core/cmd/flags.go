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
	envPath := flag.String("env", "", "Running Environment Path! ")
	modes := flag.String("mode", "", "Running Mode! ")
	ranges := flag.String("range", "", "Use Fields! ")
	langRefs := flag.String("lang", "", "Use Languages! ")
	dataFiles := flag.String("file", "", "Output Files! ")
	sqlMerge := flag.Bool("merge", false, "Merge Sql! ")

	source := flag.String("source", "", "Source Redefine! ")
	target := flag.String("target", "", "Target Redefine! ")
	flag.Parse()

	modesVal := strings.ToLower(*modes)
	rangesVal := strings.ToLower(*ranges)
	return &SysFlags{EnvPath: *envPath, Modes: modesVal, Ranges: rangesVal, LangRefs: *langRefs, DataFiles: *dataFiles,
		SqlMerge: *sqlMerge, Source: *source, Target: *target}, nil
}

type SysFlags struct {
	EnvPath   string // 运行时指定的运行环境路径(支持绝对路径与相对路径，相对路径以执行文件为基准)(空字符串时默认使用执行文件目录)
	Modes     string // 运行时使用的模式，支持多个，用英文逗号","分隔，支持参数请看src/setting/const.go文件中
	Ranges    string // 运行时指定的字段范围，支持多个，用英文逗号","分隔，支持参数请看src/setting/const.go文件中
	LangRefs  string // 运行时指定的编程语言，支持多个，用英文逗号","分隔，支持参数请看src/setting/const.go文件中
	DataFiles string // 运行时指定的数据文件类型，支持多个，用英文逗号","分隔，支持参数请看src/setting/const.go文件中
	SqlMerge  bool   // 是否使用sql文件合并，当DataFiles中包含sql时有效，为true时只产出一个sql文件
	Source    string // 运行时指定的源目录(空字符串代表使用EvnPath文件的默认配置)
	Target    string // 运行时指定的输出目录(空字符串代表使用EvnPath文件的默认配置)
}

func (f *SysFlags) String() string {
	return fmt.Sprintf("SysFlags(Modes=%s, Rangs=%s, LangRefs=%s, DataFiles=%s, SqlMerge=%v,  Source=%s, Target=%s)",
		f.Modes, f.Ranges, f.LangRefs, f.DataFiles, f.SqlMerge, f.Source, f.Target)
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

	langRefs := f.split(f.LangRefs, ParamsSep)
	dataFiles := f.split(f.DataFiles, ParamsSep)
	return &AppFlags{ModeNames: modeNames, ModeTypes: modeValues, RangeNames: rangeNames, RangeTypes: rangeValues,
		LangRefs: langRefs, DataFiles: dataFiles, SqlMerge: f.SqlMerge}
}

func (f *SysFlags) parseModes() (names []string, types []core.ModeType) {
	modes := f.split(f.Modes, ParamsSep)
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
	ranges := f.split(f.Ranges, ParamsSep)
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

func (f *SysFlags) split(value string, sep string) []string {
	rs := strings.Split(value, sep)
	for index := len(rs) - 1; index >= 0; index -= 1 {
		v := strings.TrimSpace(rs[index])
		if v == "" {
			rs = append(rs[:index], rs[index:+1]...)
		}
		rs[index] = v
	}
	return rs
}

type AppFlags struct {
	ModeNames  []string              // 运行时使用的模式列表(字符表达)
	ModeTypes  []core.ModeType       // 运行时使用的模式列表(具体类型表达)
	RangeNames []string              // 运行时指定的字段范围列表(字符表达)
	RangeTypes []core.FieldRangeType // 运行时指定的字段范围列表(枚举表达)
	LangRefs   []string              // 运行时指定的编程语言列表
	DataFiles  []string              // 运行时指定的数据文件类型列表
	SqlMerge   bool                  // 是否使用sql文件合并，当DataFiles中包含sql时有效，为true时只产出一个sql文件
}

func (o *AppFlags) CheckMode(mode core.ModeType) bool {
	for _, m := range o.ModeTypes {
		if m == mode {
			return true
		}
	}
	return false
}

func (o *AppFlags) CheckRange(rangeType core.FieldRangeType) bool {
	for _, m := range o.RangeTypes {
		if m == rangeType {
			return true
		}
	}
	return false
}

func (o *AppFlags) CheckDataFile(dataFile string) bool {
	for _, m := range o.DataFiles {
		if m == dataFile {
			return true
		}
	}
	return false
}

func (o *AppFlags) String() string {
	return fmt.Sprintf("AppFlags(ModeNames=%v, ModeTypes=%v, RangeNames=%v, RangeTypes=%v, LangRefs=%v, DataFiles=%v, SqlMerge=%v",
		o.ModeNames, o.ModeTypes, o.RangeNames, o.RangeTypes, o.LangRefs, o.DataFiles, o.SqlMerge)
}

func (o *AppFlags) GenTitleContexts(prefix string, startRowNum int, startColIndex int) (contexts []*core.TitleContext) {
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
			context := &core.TitleContext{EnablePrefix: prefix,
				RangeName: o.RangeNames[fieldIdx], RangeType: o.RangeTypes[fieldIdx],
				ProgramLanguage: o.LangRefs[langIdx],
				StartColIndex:   startColIndex}
			contexts = append(contexts, context)
		}
	}
	return
}

func (o *AppFlags) GenDataContexts(prefix string, startRowNum int, startColIndex int) (contexts []*core.DataContext) {
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
			if o.DataFiles[fileIdx] == setting.FileNameSql {
				continue
			}
			context := &core.DataContext{EnablePrefix: prefix,
				RangeName: o.RangeNames[fieldIdx], RangeType: o.RangeTypes[fieldIdx],
				DataFileFormat: o.DataFiles[fileIdx],
				StartRowNum:    startRowNum, StartColIndex: startColIndex}
			contexts = append(contexts, context)
		}
	}
	return
}

func (o *AppFlags) GenConstContexts(prefix string) (contexts []*core.ConstContext) {
	if !o.CheckMode(core.ModeConst) {
		return nil
	}
	rangeLen := len(o.RangeTypes)
	langLen := len(o.LangRefs)
	if rangeLen == 0 || langLen == 0 {
		return nil
	}
	ln := rangeLen * langLen
	contexts = make([]*core.ConstContext, 0, ln)
	for fieldIdx := 0; fieldIdx < rangeLen; fieldIdx += 1 {
		for langIdx := 0; langIdx < langLen; langIdx += 1 {
			rangeName := o.RangeNames[fieldIdx]
			if rangeName != setting.FieldRangeNameClient && rangeName != setting.FieldRangeNameServer {
				continue
			}
			context := &core.ConstContext{EnablePrefix: prefix,
				RangeName: o.RangeNames[fieldIdx], RangeType: o.RangeTypes[fieldIdx],
				ProgramLanguage: o.LangRefs[langIdx]}
			contexts = append(contexts, context)
		}
	}
	return
}

// 生成Sql导出相关
func (o *AppFlags) GenSqlContext(prefix string, startRowNum int, startColIndex int) (context *core.SqlContext) {
	if !o.CheckRange(core.FieldRangeDatabase) || !o.CheckDataFile(setting.FileNameSql) {
		return nil
	}
	titleOn := o.CheckMode(core.ModeTitle)
	dataOn := o.CheckMode(core.ModeData)
	if !titleOn && !dataOn {
		return nil
	}
	return &core.SqlContext{EnablePrefix: prefix,
		RangeName: setting.FieldRangeNameDb, RangeType: core.FieldRangeDatabase,
		TitleOn: titleOn, DataOn: dataOn, SqlMerge: o.SqlMerge,
		StartRowNum: startRowNum, StartColIndex: startColIndex}
}
