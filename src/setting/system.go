package setting

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"gopkg.in/yaml.v2"
	"os"
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

func (o *ProgramLanguage) UpgradePaths(envPath string) {
	o.RefPath = filex.Combine(envPath, o.RefPath)
	for index := range o.TempsTitle {
		o.TempsTitle[index] = filex.Combine(envPath, o.TempsTitle[index])
	}
	for index := range o.TempsConst {
		o.TempsConst[index] = filex.Combine(envPath, o.TempsConst[index])
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

type Database struct {
	Name       string   `yaml:"name"`
	RefPath    string   `yaml:"ref"`
	TempsTitle []string `yaml:"temps_title"`
	TempsData  []string `yaml:"temps_data"`

	DataTypes *SqlDataTypes
}

func (o *Database) UpgradePaths(envPath string) {
	o.RefPath = filex.Combine(envPath, o.RefPath)
	for index := range o.TempsTitle {
		o.TempsTitle[index] = filex.Combine(envPath, o.TempsTitle[index])
	}
	for index := range o.TempsData {
		o.TempsData[index] = filex.Combine(envPath, o.TempsData[index])
	}
}

func (o *Database) GetDataTypes() (dbTypes *SqlDataTypes, err error) {
	if nil != o.DataTypes {
		return o.DataTypes, nil
	}

	str, err := os.ReadFile(o.RefPath)

	if nil != err {
		return nil, err
	}
	types := &SqlDataTypes{}
	err = yaml.Unmarshal(str, types)
	if nil != err {
		return nil, err
	}
	o.DataTypes = types
	return types, nil
}

type Databases struct {
	Default      string      `yaml:"default"`
	DatabaseList []*Database `yaml:"list"`
}

func (d *Databases) UpgradePaths(envPath string) {
	for index := range d.DatabaseList {
		d.DatabaseList[index].UpgradePaths(envPath)
	}
}

func (d *Databases) GetDefaultDatabase() (db Database, exist bool) {
	return d.FindDatabase(d.Default)
}

func (d *Databases) FindDatabase(name string) (db Database, exist bool) {
	exist = false
	if len(name) == 0 || len(d.DatabaseList) == 0 {
		return
	}

	for index := range d.DatabaseList {
		if d.DatabaseList[index].Name == name {
			return *d.DatabaseList[index], true
		}
	}
	return
}

type SystemSetting struct {
	// 数据结构定义所支持的编程语言
	Languages []*ProgramLanguage `yaml:"languages"`
	// 数据库相关配置
	Databases *Databases `yaml:"databases"`
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

func (s *SystemSetting) UpgradeEnvPath(envPath string) {
	for index := range s.Languages {
		s.Languages[index].UpgradePaths(envPath)
	}
	s.Databases.UpgradePaths(envPath)
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

func (s *SystemSetting) GetDatabase() (db Database, ok bool) {
	return s.Databases.GetDefaultDatabase()
}
