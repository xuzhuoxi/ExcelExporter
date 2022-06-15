package core

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
)

func executeSqlContext(excel *excel.ExcelProxy, sqlCtx *SqlContext) error {
	fmt.Println("executeSqlContext:", sqlCtx)
	return nil
}
