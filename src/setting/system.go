package setting

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
)

// 数据结构定义所支持的编程语言定义
type ProgramLanguage struct {
	Name       string   `yaml:"name"`
	ExtendName string   `yaml:"ext"`
	RefPath    string   `yaml:"ref"`
	TempsTitle []string `yaml:"temps_title"`
	TempsConst []string `yaml:"temps_const"`

	Setting *LangSetting
}

func (o ProgramLanguage) String() string {
	return fmt.Sprintf("Lang{Name=%s, RefPath=%s, TempsTitle=%s, TempsConst=%s}",
		o.Name, o.RefPath, o.TempsTitle, o.TempsConst)
}

func (o *ProgramLanguage) UpgradePaths(basePath string) {
	o.RefPath = filex.Combine(basePath, o.RefPath)
	for index := range o.TempsTitle {
		o.TempsTitle[index] = filex.Combine(basePath, o.TempsTitle[index])
	}
	for index := range o.TempsConst {
		o.TempsConst[index] = filex.Combine(basePath, o.TempsConst[index])
	}
}

func (o *ProgramLanguage) GetTempsTitlePath() string {
	return getMergedTempsPath(o.TempsTitle)
}

func (o *ProgramLanguage) GetTempsConstPath() string {
	return getMergedTempsPath(o.TempsConst)
}

func getMergedTempsPath(paths []string) string {
	ln := len(paths)
	if ln == 0 {
		return ""
	}
	if ln == 1 {
		return paths[0]
	}
	rs := paths[0]
	for i := 1; i < ln; i += 1 {
		rs = rs + "," + paths[i]
	}
	return rs
}

type SystemSetting struct {
	// 数据结构定义所支持的编程语言
	Languages []*ProgramLanguage `yaml:"languages"`
	// 支持的数据字段格式
	// 其中string中的*代表字符数，范围[1,1024]。
	// 浮点数最多支持6位小数，而且当数值越大，精度就越低，反之亦然
	// 使用浮点数时，如果是负数，序列化为二进制后再读取，部分编程语言会出现抖动现象，如AS3: -2.3 => [-64,19,51,51] => -2.299999952316284
	DataFieldFormats []string `yaml:"datafield_formats"`
	// 支持的数据文件格式
	DataFileFormats []string `yaml:"datafile_formats"`
}

func (s *SystemSetting) String() string {
	return fmt.Sprintf("System{Languages=%v, Fields=%v, Files=%v}",
		s.Languages, s.DataFieldFormats, s.DataFileFormats)
}

func (s *SystemSetting) UpgradePath(basePath string) {
	for index := range s.Languages {
		s.Languages[index].UpgradePaths(basePath)
	}
}

func (s *SystemSetting) FindProgramLanguage(lang string) (define *ProgramLanguage, ok bool) {
	if len(lang) == 0 {
		return nil, false
	}
	for _, ld := range s.Languages {
		if ld.Name == lang {
			return ld, true
		}
	}
	return nil, false
}

func (s *SystemSetting) CheckDataFieldFormat(dataFieldFormat string) bool {
	if len(dataFieldFormat) == 0 {
		return false
	}
	for _, ld := range s.DataFieldFormats {
		if ld == dataFieldFormat {
			return true
		}
	}
	return false
}

func (s *SystemSetting) CheckDataFileFormat(dataFileFormat string) bool {
	if len(dataFileFormat) == 0 {
		return false
	}
	for _, ld := range s.DataFileFormats {
		if ld == dataFileFormat {
			return true
		}
	}
	return false
}
