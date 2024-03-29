package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestExcelSetting(t *testing.T) {
	str, err := ioutil.ReadFile(ExcelPath)
	if nil != err {
		t.Fatal(err)
		return
	}
	system := &ExcelSetting{}
	err = yaml.Unmarshal(str, system)
	if nil != err {
		t.Fatal(err)
		return
	}
	fmt.Println(system)
}
