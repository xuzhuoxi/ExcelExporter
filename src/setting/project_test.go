package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestProjectSetting(t *testing.T) {
	str, err := ioutil.ReadFile(ProjectPath)
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
