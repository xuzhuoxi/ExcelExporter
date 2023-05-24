package setting

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"strings"
)

// 名称与号记录项
type NameRow struct {
	Name string `yaml:"name"` // 名称(键)
	Row  int    `yaml:"row"`  // Excel行号
}

func (o NameRow) String() string {
	return fmt.Sprintf("NameRow{Name=%s, Row=%d}", o.Name, o.Row)
}

// 名称与字符值记录项
type NameValue struct {
	Name  string `yaml:"name"`  // 名称(键)
	Value string `yaml:"value"` // 内容
}

func (o NameValue) String() string {
	return fmt.Sprintf("NameValue{Name=%s, Value=%s}", o.Name, o.Value)
}

// Excel相关配置环境
type ExcelSetting struct {
	Ignore    []string  `yaml:"ignore"`     // 忽略文件前缀
	TitleData TitleData `yaml:"title&data"` // 数据表配置
	Const     Const     `yaml:"const"`      // 常量表配置
	Proto     Proto     `yaml:"proto"`      // 协议表配置
}

func (s ExcelSetting) CheckIgnorePath(filePath string) bool {
	if len(s.Ignore) == 0 {
		return false
	}
	_, fileName := filex.Split(filePath)
	return s.checkName(fileName)
}

func (s ExcelSetting) CheckIgnoreName(fileName string) bool {
	if len(s.Ignore) == 0 {
		return false
	}
	return s.checkName(fileName)
}

func (s ExcelSetting) checkName(fileName string) bool {
	for _, ignore := range s.Ignore {
		if strings.HasPrefix(fileName, ignore) {
			return true
		}
	}
	return false
}
