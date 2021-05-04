package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var proPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res\project.yaml`

func TestProjectSetting(t *testing.T) {
	str, err := ioutil.ReadFile(proPath)
	if nil != err {
		t.Fatal(err)
		return
	}
	system := &ProjectSetting{}
	err = yaml.Unmarshal(str, system)
	if nil != err {
		t.Fatal(err)
		return
	}
	fmt.Println(system)
}
