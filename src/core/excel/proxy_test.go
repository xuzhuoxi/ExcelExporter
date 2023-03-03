package excel

import (
	"fmt"
	"testing"
)

var (
	path = `D:\workspaces\GoPath\src\github.com\xuzhuoxi\ExcelExporter\res\source\const.xlsx`
)

func TestLoadExcel(t *testing.T) {
	excelProxy := &ExcelProxy{}
	err := excelProxy.LoadExcel(path, true)
	if nil != err {
		fmt.Println(err)
		return
	}
	err = excelProxy.LoadSheetsByPrefix("Const_", 0, true)
	if nil != err {
		fmt.Println(err)
		return
	}
	err = excelProxy.MergedRowsByFilter(2, func(row *ExcelRow) bool {
		return !row.Empty()
	})
	//err = excelProxy.MergedRows(2)
	if nil != err {
		fmt.Println(err)
		return
	}

	fmt.Println(excelProxy.Sheets)
	fmt.Println(excelProxy.DataRows, len(excelProxy.DataRows))
}

//func TestLoadSheets(t *testing.T) {
//	path := osxu.RunningBaseDir() + "Source/const.xlsx"
//	excel, err := LoadExcel(path)
//	if nil != err {
//		fmt.Println(err)
//		return
//	}
//	excel.LoadSheetsByPrefixes("Const_", 2)
//	for _, s := range excel.Sheets {
//		fmt.Println(*s)
//	}
//}
