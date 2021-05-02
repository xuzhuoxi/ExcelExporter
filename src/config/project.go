package config

type ProjectIO struct {
	// 目录路径或文件路径
	Value string `yaml:"value"`
	// 编码格式
	Encoding string `yaml:"encoding"`
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

// 项目配置
type ProjectSetting struct {
	// 默认处理的文件或目录,以'':''开关，或路径中包含'':''的，视为绝对路径
	Source ProjectIO `yaml:"source"`
	// 输出目录,以'':''开关，或路径中包含'':''的，视为绝对路径,encoding属性作用于字符文件的输出,和字节文件中字符串的编码
	Target ProjectIO   `yaml:"target"`
	Buff   ProjectBuff `yaml:"buff"`
}
