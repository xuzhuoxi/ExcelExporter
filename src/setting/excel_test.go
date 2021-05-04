package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var excelPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res\excel.yaml`

func TestExcelSetting(t *testing.T) {
	str, err := ioutil.ReadFile(excelPath)
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
