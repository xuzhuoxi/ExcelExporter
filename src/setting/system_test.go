//go:build !skiptest

package setting

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestSystemSetting(t *testing.T) {
	str, err := os.ReadFile(SystemPath)
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
