package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestLangSetting(t *testing.T) {
	str, err := ioutil.ReadFile(LangGoPath)
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
