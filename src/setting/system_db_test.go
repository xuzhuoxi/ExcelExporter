//go:build !skiptest

package setting

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestMysqlSetting(t *testing.T) {
	str, err := os.ReadFile(DbMysqlPath)
	if nil != err {
		t.Fatal(err)
		return
	}
	mysql := &DatabaseExtend{}
	err = yaml.Unmarshal(str, mysql)
	if nil != err {
		t.Fatal(err)
		return
	}
	fmt.Println(mysql)
}
