package setting

import (
	"fmt"
)

// 数据结构定义所支持的编程语言定义
type LangDefine struct {
	Name string `yaml:"name"`
	Ref  string `yaml:"ref"`
}

func (o LangDefine) String() string {
	return fmt.Sprintf("Lang{Name=%s, Ref=%s}", o.Name, o.Ref)
}

type SystemSetting struct {
	// 数据结构定义所支持的编程语言
	LangRef []LangDefine `yaml:"languages"`
	// 支持的数据文件格式
	SupportFiles []string `yaml:"out_files"`
	// 支持的数据字段格式
	// 其中string中的*代表字符数，范围[1,1024]。
	// 浮点数最多支持6位小数，而且当数值越大，精度就越低，反之亦然
	// 使用浮点数时，如果是负数，序列化为二进制后再读取，部分编程语言会出现抖动现象，如AS3: -2.3 => [-64,19,51,51] => -2.299999952316284
	SupportFields []string `yaml:"fields"`
}

func (s *SystemSetting) String() string {
	return fmt.Sprintf("System{Langs=%v, Files=%v, Fields=%v}",
		s.LangRef, s.SupportFiles, s.SupportFields)
}

func (s *SystemSetting) FindLangRef(lang string) (ok bool, define *LangDefine) {
	if len(lang) == 0 {
		return false, nil
	}
	for _, ld := range s.LangRef {
		if ld.Name == lang {
			return true, &ld
		}
	}
	return false, nil
}

func (s *SystemSetting) CheckDataOutputFileSupport(fileType string) bool {
	if len(fileType) == 0 {
		return false
	}
	for _, ld := range s.SupportFiles {
		if ld == fileType {
			return true
		}
	}
	return false
}

func (s *SystemSetting) CheckFieldSupport(field string) bool {
	if len(field) == 0 {
		return false
	}
	for _, ld := range s.SupportFields {
		if ld == field {
			return true
		}
	}
	return false
}
