package config

// 数据结构定义所支持的编程语言
type SystemLangRef struct {
	As3Ref        string `yaml:"as3"`
	CPlusRef      string `yaml:"c++"`
	CSharpRef     string `yaml:"c#"`
	GoRef         string `yaml:"go"`
	JavaRef       string `yaml:"java"`
	TypeScriptRef string `yaml:"ts"`
	JsonRef       string `yaml:"json"`
	DbRef         string `yaml:"db"`
}

type SystemSetting struct {
	// 数据结构定义所支持的编程语言
	RefLangs SystemLangRef `yaml:"languages"`
	// 支持的数据文件格式
	SupportFiles []string `yaml:"files"`
	// 支持的数据字段格式
	// 其中string中的*代表字符数，范围[1,1024]。
	// 浮点数最多支持6位小数，而且当数值越大，精度就越低，反之亦然
	// 使用浮点数时，如果是负数，序列化为二进制后再读取，部分编程语言会出现抖动现象，如AS3: -2.3 => [-64,19,51,51] => -2.299999952316284
	SupportFields []string `yaml:"fields"`
}
