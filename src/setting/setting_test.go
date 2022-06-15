package setting

import (
	"github.com/xuzhuoxi/infra-go/filex"
	"testing"
)

var (
	LangGoPath  = filex.Combine(RootPath, "./lang/go.yaml")
	DbMysqlPath = filex.Combine(RootPath, "./db/mysql.yaml")
)

func init() {
	RootPath = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res`
	SystemPath = filex.Combine(RootPath, "system.yaml")
	ProjectPath = filex.Combine(RootPath, "project.yaml")
	ExcelPath = filex.Combine(RootPath, "excel.yaml")
	LangGoPath = filex.Combine(RootPath, "./lang/go.yaml")
	DbMysqlPath = filex.Combine(RootPath, "./db/mysql.yaml")
}

func TestFormatStringField(t *testing.T) {

}
