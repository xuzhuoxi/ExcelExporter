package setting

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"strings"
)

// 输出目录
type OutDir struct {
	Client string `yaml:"client"` // 前端定义输出目录
	Server string `yaml:"server"` // 后端定义输出目录
	Db     string `yaml:"db"`     // 数据库定义输出目录
}

func (o OutDir) GetValue(fieldRangeName string) string {
	switch fieldRangeName {
	case FieldRangeNameClient:
		return o.Client
	case FieldRangeNameServer:
		return o.Server
	case FieldRangeNameDb:
		return o.Db
	default:
		return ""
	}
}

type OutSql struct {
	Dir string `yaml:"dir"` // 输出目录
}

// 数据源配置
type SourceCfg struct {
	Value    []string `yaml:"value"`    // 目录路径或文件路径
	Encoding string   `yaml:"encoding"` // 编码格式(如果需要)
	ExtName  []string `yaml:"ext_name"` // 支持文件扩展名
}

func (o *SourceCfg) UpgradeEnvPath(envPath string) {
	if len(o.Value) == 0 {
		return
	}
	for index := range o.Value {
		if filex.IsExist(o.Value[index]) {
			continue
		}
		o.Value[index] = filex.Combine(envPath, o.Value[index])
	}
}

func (o SourceCfg) CheckFileType(filePath string) bool {
	if len(o.ExtName) == 0 {
		return false
	}
	for _, e := range o.ExtName {
		if filex.CheckExt(filePath, e) {
			return true
		}
	}
	return false
}

func (o SourceCfg) String() string {
	return fmt.Sprintf("IO{Path=%s, Encoding=%s}", o.Value, o.Encoding)
}

// 输出配置
type TargetCfg struct {
	RootDir  string `yaml:"root"`     // 输出根目录
	Title    OutDir `yaml:"title"`    // (定义)导出类输出配置
	Data     OutDir `yaml:"data"`     // 数据文件输出配置
	Const    OutDir `yaml:"const"`    // 常量类输出配置
	Sql      OutSql `yaml:"sql"`      // Sql输出配置
	Proto    OutDir `yaml:"proto"`    // Proto输出配置
	Encoding string `yaml:"encoding"` // 字符文件编码，暂时未使用
}

func (o *TargetCfg) UpgradeEnvPath(envPath string) {
	if !filex.IsRelativeFormat(o.RootDir) || filex.IsDir(o.RootDir) {
		return
	}
	o.RootDir = filex.Combine(envPath, o.RootDir)
}

func (o *TargetCfg) GetTitleDir(fieldRangeName string) string {
	return filex.Combine(o.RootDir, o.Title.GetValue(fieldRangeName))
}

func (o *TargetCfg) GetDataDir(fieldRangeName string) string {
	return filex.Combine(o.RootDir, o.Data.GetValue(fieldRangeName))
}

func (o *TargetCfg) GetConstDir(fieldRangeName string) string {
	return filex.Combine(o.RootDir, o.Const.GetValue(fieldRangeName))
}

func (o *TargetCfg) GetSqlDir(fieldRangeName string) string {
	return filex.Combine(o.RootDir, o.Sql.Dir)
}

func (o *TargetCfg) GetProtoDir(fieldRangeName string) string {
	return filex.Combine(o.RootDir, o.Proto.GetValue(fieldRangeName))
}

func (o TargetCfg) String() string {
	return fmt.Sprintf("TargetCfg{RootDir=%s, Title=%v, Data=%v, Const=%v, Proto=%v}",
		o.RootDir, o.Title, o.Data, o.Const, o.Proto)
}

// 缓冲区定义(未使用)
type ProjectBuff struct {
	IsBigEndian bool `yaml:"big_endian"` // 数据导出是否使用高位在前
	TokenSize   int  `yaml:"token"`      // 每个token的最大缓冲区
	ItemSize    int  `yaml:"item"`       // 每个item的最大缓冲区
	SheetSize   int  `yaml:"sheet"`      // 每个sheet的最大缓冲区
}

func (b ProjectBuff) String() string {
	return fmt.Sprintf("Buff{BigEndian=%t, TokenSize=%d, ItemSize=%d, SheetSize=%d}",
		b.IsBigEndian, b.TokenSize, b.ItemSize, b.SheetSize)
}

// 项目配置
type ProjectSetting struct {
	// 默认处理的文件或目录,以'':''开关，或路径中包含'':''的，视为绝对路径
	Source SourceCfg `yaml:"source"`
	// 输出目录,以'':''开关，或路径中包含'':''的，视为绝对路径,encoding属性作用于字符文件的输出,和字节文件中字符串的编码
	Target TargetCfg `yaml:"target"`
	// 处理时缓存设置
	Buff ProjectBuff `yaml:"buff"`
}

func (ps *ProjectSetting) UpgradeEnvPath(envPath string) {
	ps.Source.UpgradeEnvPath(envPath)
	ps.Target.UpgradeEnvPath(envPath)
}
func (ps *ProjectSetting) String() string {
	return fmt.Sprintf("Project{Source=%s, Target=%s, Buff=%v}",
		ps.Source, ps.Target, ps.Buff)
}

func (ps *ProjectSetting) UpdateSource(source string) {
	if len(source) == 0 || len(strings.TrimSpace(source)) == 0 {
		return
	}
	source = strings.TrimSpace(source)
	ss := strings.Split(source, ",")
	ps.Source.Value = ss
}

func (ps *ProjectSetting) UpdateTarget(target string) {
	if len(target) == 0 || len(strings.TrimSpace(target)) == 0 {
		return
	}
	ps.Target.RootDir = strings.TrimSpace(target)

}
