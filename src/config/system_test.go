package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var sysPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res\system.yaml`

func TestSystemSetting(t *testing.T) {
	str, err := ioutil.ReadFile(sysPath)
	if nil != err {
		t.Fatal(err)
		return
	}
	system := &SystemSetting{}
	err = yaml.Unmarshal(str, system)
	if nil != err {
		t.Fatal(err)
		return
	}
	fmt.Println(system)
}
