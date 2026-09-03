//go:build !skiptest

package setting

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestLangSetting(t *testing.T) {
	str, err := os.ReadFile(LangGoPath)
	if nil != err {
		t.Fatal(err)
		return
	}
	system := &LangSetting{}
	err = yaml.Unmarshal(str, system)
	if nil != err {
		t.Fatal(err)
		return
	}
	fmt.Println(system)
}
