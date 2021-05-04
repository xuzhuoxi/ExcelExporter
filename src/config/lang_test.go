package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var langGoPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res\langs\go.yaml`

func TestLangSetting(t *testing.T) {
	str, err := ioutil.ReadFile(langGoPath)
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
