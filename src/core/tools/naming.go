package tools

import (
	"github.com/xuzhuoxi/infra-go/stringx"
	"regexp"
	"strings"
)

func clearUnderscore(naming string) string {
	name := []rune(naming)
	index := stringx.LastIndexOfString(naming, "_")
	if index >= 0 && index < len(name)-1 {
		return strings.ToUpper(string(name[index+1])) + string(name[index+2:])
	}
	return naming
}

// 去除下划线，并把下划线后一个字符改为大写
func ClearUnderscore(naming string) string {
	naming = strings.TrimSpace(naming)
	ln := len(naming)
	if 0 == ln {
		return naming
	}
	reg := regexp.MustCompile(`[_]+[a-zA-Z](\w.)`)
	return reg.ReplaceAllStringFunc(naming, clearUnderscore)
}

// 转为小驼峰
func ToLowerCamelCase(naming string) string {
	naming = ClearUnderscore(naming)
	ln := len(naming)
	if 0 == ln {
		return naming
	}
	n1, n2 := stringx.CutString(naming, 1, true)
	return strings.ToLower(n1) + n2
}

// 转为大驼峰
func ToUpperCamelCase(naming string) string {
	naming = ClearUnderscore(naming)
	ln := len(naming)
	if 0 == ln {
		return naming
	}
	n1, n2 := stringx.CutString(naming, 1, true)
	return strings.ToUpper(n1) + n2
}

func GenLowerCamelCase(prefix string, naming string, suffix string) string {
	return prefix + ToLowerCamelCase(naming) + suffix
}

func GenUpperCamelCase(prefix string, naming string, suffix string) string {
	return prefix + ToUpperCamelCase(naming) + suffix

}
