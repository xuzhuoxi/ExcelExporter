package setting

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"gopkg.in/yaml.v2"
	"os"
)

// ProgramLanguage
// 数据结构定义所支持的编程语言定义
// 编程语言描述
type ProgramLanguage struct {
	Name       string   `yaml:"name"`        // 编程语言名称
	ExtendName string   `yaml:"ext"`         // 源代码文件扩展名
	RefPath    string   `yaml:"ref"`         // 基础数据读写配置文件路径(相对于配置根目录相对路径)
	TempsTitle []string `yaml:"temps_title"` // title导出类定义导出模板路径(相对于配置根目录相对路径)
	TempsConst []string `yaml:"temps_const"` // 常量定义导出模板路径(相对于配置根目录相对路径)
	TempsProto []string `yaml:"temps_proto"` // 协议定义导出模板路径(相对于配置根目录相对路径)

	Setting *LangSetting // 由RefPath加载进来的配置信息
}

func (o *ProgramLanguage) String() string {
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
	for index := range o.TempsProto {
		o.TempsProto[index] = filex.Combine(envPath, o.TempsProto[index])
	}
}

func (o *ProgramLanguage) GetTempsTitlePath() string {
	return getMergedTempsPath(o.TempsTitle)
}

func (o *ProgramLanguage) GetTempsConstPath() string {
	return getMergedTempsPath(o.TempsConst)
}

func (o *ProgramLanguage) GetTempsProtoPath() string {
	return getMergedTempsPath(o.TempsProto)
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

// 数据库相关配置
type Database struct {
	Name       string   `yaml:"name"`        // 数据库名称
	RefPath    string   `yaml:"ref"`         // 数据库具体配置文件所在路径
	TempsTable []string `yaml:"temps_table"` // 表结构sql生成模板列表
	TempsData  []string `yaml:"temps_data"`  // 表数据sql生成模板列表

	Extend *DatabaseExtend // 由RefPath加载进来的配置信息
}

func (o *Database) UpgradePaths(envPath string) {
	o.RefPath = filex.Combine(envPath, o.RefPath)
	for index := range o.TempsTable {
		o.TempsTable[index] = filex.Combine(envPath, o.TempsTable[index])
	}
	for index := range o.TempsData {
		o.TempsData[index] = filex.Combine(envPath, o.TempsData[index])
	}
}

func (o *Database) GetTempsTablePath() string {
	return getMergedTempsPath(o.TempsTable)
}

func (o *Database) GetTempsDataPath() string {
	return getMergedTempsPath(o.TempsData)
}

func (o *Database) GetDatabaseExtend() (dbTypes *DatabaseExtend, err error) {
	if nil != o.Extend {
		return o.Extend, nil
	}

	str, err := os.ReadFile(o.RefPath)

	if nil != err {
		return nil, err
	}
	types := &DatabaseExtend{}
	err = yaml.Unmarshal(str, types)
	if nil != err {
		return nil, err
	}
	o.Extend = types
	return types, nil
}

// 数据库相关配置
type Databases struct {
	Default      string      `yaml:"default"` // 默认使用的数据库配置，必须为DatabaseList中的一个
	DatabaseList []*Database `yaml:"list"`    // 数据库的配置静静列表
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

type SystemSettings struct {
	// 数据结构定义所支持的编程语言
	Languages []*ProgramLanguage `yaml:"languages"`
	// 数据库相关配置
	Databases *Databases `yaml:"databases"`
	// 支持的数据字段格式
	// 其中string中的*代表字符数，范围[1,1024]。
	// 浮点数最多支持6位小数，而且当数值越大，精度就越低，反之亦然
	// 使用浮点数时，如果是负数，序列化为二进制后再读取，部分编程语言会出现抖动现象，如AS3: -2.3 => [-64,19,51,51] => -2.299999952316284
	FieldDataTypes []string `yaml:"field_datatypes"`
	// 指针代码
	PointerCode string `yaml:"pointer_code"`
	// 支持的导出数据文件格式
	ExportDataFiles []string `yaml:"export_files"`
}

func (s *SystemSettings) String() string {
	return fmt.Sprintf("System{Languages=%v, Fields=%v, Files=%v}",
		s.Languages, s.FieldDataTypes, s.ExportDataFiles)
}

func (s *SystemSettings) UpgradeEnvPath(envPath string) {
	for index := range s.Languages {
		s.Languages[index].UpgradePaths(envPath)
	}
	s.Databases.UpgradePaths(envPath)
}

func (s *SystemSettings) FindProgramLanguage(lang string) (define *ProgramLanguage, ok bool) {
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

func (s *SystemSettings) CheckFieldDataType(dataType string) bool {
	if len(dataType) == 0 {
		return false
	}
	for _, ld := range s.FieldDataTypes {
		if ld == dataType {
			return true
		}
	}
	return false
}

func (s *SystemSettings) CheckExportDataFile(dataFileType string) bool {
	if len(dataFileType) == 0 {
		return false
	}
	for _, ld := range s.ExportDataFiles {
		if ld == dataFileType {
			return true
		}
	}
	return false
}

func (s *SystemSettings) GetDatabase() (db Database, ok bool) {
	return s.Databases.GetDefaultDatabase()
}
