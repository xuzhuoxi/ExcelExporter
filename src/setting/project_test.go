//go:build !skiptest

package setting

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestProjectSetting(t *testing.T) {
	str, err := os.ReadFile(ProjectPath)
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
