package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestSystemSetting(t *testing.T) {
	str, err := ioutil.ReadFile(SystemPath)
	if nil != err {
		t.Fatal(err)
		return
	}
	system := &SystemSettings{}
	err = yaml.Unmarshal(str, system)
	if nil != err {
		t.Fatal(err)
		return
	}
	fmt.Println(system)
}
