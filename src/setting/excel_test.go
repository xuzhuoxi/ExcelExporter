//go:build !skiptest

package setting

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestExcelSetting(t *testing.T) {
	str, err := os.ReadFile(ExcelPath)
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
