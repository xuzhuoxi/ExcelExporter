package setting

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"strings"
)

type OutputCfg struct {
	// 前端定义输出目录
	Client string `yaml:"client"`
	// 后端定义输出目录
	Server string `yaml:"server"`
	// 数据库定义输出目录
	Database string `yaml:"database"`
}

func (o OutputCfg) GetValue(fieldRangeName string) string {
	switch fieldRangeName {
	case FieldRangeNameClient:
		return o.Client
	case FieldRangeNameServer:
		return o.Server
	case FieldRangeNameDb:
		return o.Database
	default:
		return ""
	}
}

type SourceCfg struct {
	// 目录路径或文件路径
	Value []string `yaml:"value"`
	// 编码格式(如果需要)
	Encoding string `yaml:"encoding"`
	// 文件扩展名
	ExtName []string `yaml:"ext_name"`
}

func (o *SourceCfg) UpgradePath(basePath string) {
	if len(o.Value) == 0 {
		return
	}
	for index := range o.Value {
		if filex.IsExist(o.Value[index]) {
			continue
		}
		o.Value[index] = filex.Combine(basePath, o.Value[index])
	}
}

func (o SourceCfg) CheckFileFormat(filePath string) bool {
	if len(o.ExtName) == 0 {
		return false
	}
	_, _, ext := filex.SplitFileName(filePath)
	for _, e := range o.ExtName {
		if e == strings.ToLower(strings.TrimSpace(ext)) {
			return true
		}
	}
	return false
}

func (o SourceCfg) String() string {
	return fmt.Sprintf("IO{Path=%s, Encoding=%s}", o.Value, o.Encoding)
}

type TargetCfg struct {
	// 根目录
	RootDir string `yaml:"root"`
	// 定义目录
	Title OutputCfg `yaml:"title"`
	// 数据目录
	Data OutputCfg `yaml:"data"`
	// 常量目录
	Const string `yaml:"const"`
}

func (o *TargetCfg) UpgradePath(basePath string) {
	if !filex.IsFolder(o.RootDir) {
		o.RootDir = filex.Combine(basePath, o.RootDir)
	}
}

func (o *TargetCfg) GetTitleDir(fieldRangeName string) string {
	return filex.Combine(o.RootDir, o.Title.GetValue(fieldRangeName))
}

func (o *TargetCfg) GetDataDir(fieldRangeName string) string {
	return filex.Combine(o.RootDir, o.Data.GetValue(fieldRangeName))
}

func (o *TargetCfg) ConstDir() string {
	return filex.Combine(o.RootDir, o.Const)
}

func (o TargetCfg) String() string {
	return fmt.Sprintf("IO{Path=%s, Encoding=%s}", o.RootDir, o.Title, o.Data, o.Const)
}

// 缓冲区定义
type ProjectBuff struct {
	// 数据导出是否使用高位在前
	IsBigEndian bool `yaml:"big_endian"`
	// 每个token的最大缓冲区
	TokenSize int `yaml:"token"`
	// 每个item的最大缓冲区
	ItemSize int `yaml:"item"`
	// 每个sheet的最大缓冲区
	SheetSize int `yaml:"sheet"`
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

func (ps *ProjectSetting) UpgradePath(basePath string) {
	ps.Source.UpgradePath(basePath)
	ps.Target.UpgradePath(basePath)
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
