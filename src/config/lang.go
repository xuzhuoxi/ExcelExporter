package config

type LangDataType struct {
	Name      string `yaml:"name"`
	LangName  string `yaml:"lang_name"`
	JsonGet   string `yaml:"json_get,omitempty"`
	JsonSet   string `yaml:"json_set,omitempty"`
	YamlGet   string `yaml:"yaml_get,omitempty"`
	YamlSet   string `yaml:"yaml_set,omitempty"`
	BinaryGet string `yaml:"bin_get,omitempty"`
	BinarySet string `yaml:"bin_set,omitempty"`
}

type LangSetting struct {
	DataTypes []LangDataType
	TempPath  string
}
