package setting

import (
	"fmt"
)

// 数据结构定义所支持的编程语言定义
type ProgramLanguage struct {
	Name string `yaml:"name"`
	Ref  string `yaml:"ref"`
}

func (o ProgramLanguage) String() string {
	return fmt.Sprintf("Lang{Name=%s, Ref=%s}", o.Name, o.Ref)
}

type SystemSetting struct {
	// 数据结构定义所支持的编程语言
	Languages []ProgramLanguage `yaml:"program_languages"`
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

func (s *SystemSetting) FindProgramLanguage(lang string) (ok bool, define *ProgramLanguage) {
	if len(lang) == 0 {
		return false, nil
	}
	for _, ld := range s.Languages {
		if ld.Name == lang {
			return true, &ld
		}
	}
	return false, nil
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
