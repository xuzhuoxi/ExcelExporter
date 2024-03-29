package excel

import (
	"errors"
	"github.com/xuzhuoxi/ExcelExporter/src/excelize.v2"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strings"
)

type ExcelProxy struct {
	ExcelPaths []string         // Excel文件路径列表
	Excels     []*excelize.File // Excel文件初始列表
	Sheets     []*ExcelSheet    // 加载进来的Sheet列表
	DataRows   []*ExcelRow      // 加载进来的行数据
}

func (ep *ExcelProxy) GetExcelByPath(filePath string) *excelize.File {
	index, ok := slicex.IndexString(ep.ExcelPaths, filePath)
	if !ok {
		return nil
	}
	return ep.Excels[index]
}

func (ep *ExcelProxy) GetSheets(prefix string) (sheets []*ExcelSheet) {
	for _, sheet := range ep.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		sheets = append(sheets, sheet)
	}
	return
}

func (ep *ExcelProxy) GetSheet(sheetName string) (es *ExcelSheet, err error) {
	for _, s := range ep.Sheets {
		if s.SheetName == sheetName {
			return s, nil
		}
	}
	return nil, errors.New("No Sheet is " + sheetName)
}

// ValueAtAxis
// Open to templates
// 通过Sheet的名称和坐标取值，坐标格式：B4
func (ep *ExcelProxy) ValueAtAxis(sheet string, axis string) (value string, err error) {
	s, err := ep.GetSheet(sheet)
	if nil != err {
		return "", err
	}
	return s.ValueAtAxis(axis)
}

// MergedRows
// 合并全部sheet的行数据。
// 从StartRow开始。
// 清除空数据
func (ep *ExcelProxy) MergedRows(startRow int) (err error) {
	var rows []*ExcelRow
	for _, sheet := range ep.Sheets {
		rows = append(rows, sheet.GetRowsFrom(startRow-1)...)
	}
	if len(rows) == 0 {
		return errors.New("Rows is empty! ")
	} else {
		ep.DataRows = rows
		return nil
	}
}

// MergedRowsByFilter
// 合并全部sheet的行数据。
// 从StartRow开始。
// 清除空数据
func (ep *ExcelProxy) MergedRowsByFilter(startRow int, filter func(row *ExcelRow) bool) (err error) {
	var rows []*ExcelRow
	for _, sheet := range ep.Sheets {
		rows = append(rows, sheet.GetRowsByFilter(startRow-1, filter)...)
	}
	if len(rows) == 0 {
		return errors.New("Rows is empty! ")
	} else {
		ep.DataRows = rows
		return nil
	}
}

// LoadExcels
// 加载excelPath指定的一个Excel文件。
func (ep *ExcelProxy) LoadExcels(excelPath string, overwrite bool) error {
	excels, paths, err := LoadExcels(excelPath)
	if nil != err {
		return err
	}
	if overwrite {
		ep.Excels = excels
		ep.ExcelPaths = paths
	} else {
		ep.Excels = append(ep.Excels, excels...)
		ep.ExcelPaths = append(ep.ExcelPaths, paths...)
	}
	return nil
}

// LoadExcel
// 加载excelPath指定的一个或多个Excel文件。
// excelPath支持多路径模式，用","分隔。
// excelPath支持文件夹，不支持递归。
func (ep *ExcelProxy) LoadExcel(excelPath string, overwrite bool) error {
	excel, err := LoadExcel(excelPath)
	if nil != err {
		return err
	}
	if overwrite {
		if nil == ep.Excels {
			ep.Excels = []*excelize.File{excel}
			ep.ExcelPaths = []string{excelPath}
		} else {
			ep.Excels = ep.Excels[:1]
			ep.Excels[0] = excel
			ep.ExcelPaths = ep.ExcelPaths[:1]
			ep.ExcelPaths[0] = excelPath
		}
	} else {
		ep.Excels = append(ep.Excels, excel)
		ep.ExcelPaths = append(ep.ExcelPaths, excelPath)
	}
	return nil
}

// LoadSheetsByPrefix
// 从已经加载好的Excel中加载Sheets
// sheetPrefix用于限制
func (ep *ExcelProxy) LoadSheetsByPrefix(sheetPrefix string, colNickRow int, overwrite bool) error {
	if len(sheetPrefix) == 0 {
		return ep.LoadSheetsByPrefixes(nil, colNickRow, overwrite)
	} else {
		return ep.LoadSheetsByPrefixes([]string{sheetPrefix}, colNickRow, overwrite)
	}
}

func (ep *ExcelProxy) LoadSheetsByPrefixes(sheetPrefix []string, colNickRow int, overwrite bool) error {
	if overwrite {
		if nil != ep.Sheets {
			ep.Sheets = ep.Sheets[:0]
		}
	}
	if 0 == len(ep.Excels) {
		return nil
	}
	for index := range ep.Excels {
		sheets, err := LoadSheetsByPrefixes(ep.ExcelPaths[index], ep.Excels[index], sheetPrefix, colNickRow)
		if nil != err {
			return err
		}
		ep.Sheets = append(ep.Sheets, sheets...)
	}
	return nil
}
