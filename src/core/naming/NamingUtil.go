package naming

import (
	"github.com/xuzhuoxi/infra-go/stringx"
	"regexp"
	"strings"
)

func clearUnderScore(naming string) string {
	name := []rune(naming)
	index := stringx.LastIndexOfString(naming, "_")
	if index >= 0 && index < len(name)-1 {
		return strings.ToUpper(string(name[index+1])) + string(name[index+2:])
	}
	return naming
}

func ClearUnderScore(naming string) string {
	naming = strings.TrimSpace(naming)
	ln := len(naming)
	if 0 == ln {
		return naming
	}
	reg := regexp.MustCompile(`[_]+[a-zA-Z](\w.)`)
	return reg.ReplaceAllStringFunc(naming, clearUnderScore)
}

func ToLowerCamelCase(naming string) string {
	naming = ClearUnderScore(naming)
	ln := len(naming)
	if 0 == ln {
		return naming
	}
	n1, n2 := stringx.CutString(naming, 1, true)
	return strings.ToLower(n1) + n2
}

func ToUpperCamelCase(naming string) string {
	naming = ClearUnderScore(naming)
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
