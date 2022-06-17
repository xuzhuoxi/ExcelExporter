package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestMysqlSetting(t *testing.T) {
	str, err := ioutil.ReadFile(DbMysqlPath)
	if nil != err {
		t.Fatal(err)
		return
	}
	mysql := &DatabaseCfg{}
	err = yaml.Unmarshal(str, mysql)
	if nil != err {
		t.Fatal(err)
		return
	}
	fmt.Println(mysql)
}
