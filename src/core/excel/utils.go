package excel

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/excelize.v2"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"os"
	"strconv"
	"strings"
)

var (
	ExtXlsx = "xlsx"
)

// GetAxisName
// 1=>A, 2=>B
func GetAxisName(col int) string {
	return mathx.System10To26(col)
}

func GetCellName(col int, row int) string {
	return fmt.Sprintf("%s%d", mathx.System10To26(col), row)
}

// GetColNum
// A=>1，B=>2
func GetColNum(axis string) int {
	return mathx.System26To10(axis)
}

// GenAxis
// 3 => [A, B, C]
func GenAxis(length int) []string {
	rs := make([]string, length, length)
	for index := 0; index < length; index += 1 {
		rs[index] = mathx.System10To26(index + 1)
	}
	return rs
}

// SplitAxis
// "A1" => "A", 1, nil
func SplitAxis(axis string) (colName string, rowNum int, err error) {
	Axis := strings.ToUpper(strings.TrimSpace(axis))
	bs := []byte(Axis)
	var c, r []byte
	for index, b := range bs {
		if !(b >= 'A' && b <= 'Z') {
			c, r = bs[:index], bs[index:]
		}
	}
	if nil == c && nil == r {
		return "", 0, errors.New("Axis Error:" + axis)
	}
	colName = string(c)
	rowNum, err = strconv.Atoi(string(r))
	return
}

// ParseAxisIndex
// "A1" => 0, 0, nil
func ParseAxisIndex(axis string) (colIndex int, rowIndex int, err error) {
	colName, rowNum, err := SplitAxis(axis)
	if nil != err {
		return 0, 0, err
	}
	colNum := mathx.System26To10(colName)
	return colNum - 1, rowNum - 1, nil
}

// LoadExcels
// 加载路径下的Excel文件，多个路径用","分割
// 支持文件夹路径
func LoadExcels(path string) (excels []*excelize.File, excelPaths []string, err error) {
	paths := strings.Split(strings.TrimSpace(path), ",")
	if len(paths) == 0 {
		return nil, nil, errors.New("Path Empty:" + path)
	}
	var filePaths []string

	for _, path := range paths {
		fp := filex.FormatPath(path)
		filex.WalkAll(fp, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if filex.CheckExt(info.Name(), ExtXlsx) {
				filePaths = append(filePaths, path)
			}
			return nil
		})
	}
	if len(filePaths) == 0 {
		return nil, nil, errors.New("Path Empty:" + path)
	}
	for _, filePath := range filePaths {
		excel, err := LoadExcel(filePath)
		if err != nil {
			return nil, nil, err
		}
		excels = append(excels, excel)
		excelPaths = append(excelPaths, filePath)
	}
	return excels, excelPaths, nil
}

// LoadExcel
// 加载Excel文件，过滤无Sheet情况
func LoadExcel(filePath string) (excel *excelize.File, err error) {
	excelFile, err := excelize.OpenFile(filePath)
	if nil != err {
		return nil, err
	}
	if excelFile.SheetCount <= 0 {
		return nil, errors.New("No Sheets! ")
	}
	return excelFile, nil
}

// LoadSheetsByPrefixes
// 通过SheetPrefix作为限制加载Sheet, 使用""代表不限制
// 指定NickRow所在行为别名,NickRow=0时，使用列号作为别名
func LoadSheetsByPrefixes(excelPath string, excelFile *excelize.File, sheetPrefixes []string, nickRow int) (sheets []*ExcelSheet, err error) {
	var names []string
	var indexes []int
	for index, name := range excelFile.GetSheetMap() {
		if len(sheetPrefixes) == 0 {
			names = append(names, name)
			indexes = append(indexes, index)
			break
		}
		for _, prefix := range sheetPrefixes {
			if strings.Index(name, prefix) == 0 {
				names = append(names, name)
				indexes = append(indexes, index)
				break
			}
		}
	}
	if len(names) == 0 {
		return nil, nil
	}
	for i, n := range names {
		rows, err := excelFile.GetRows(n)
		if nil != err {
			return nil, err
		}

		rowLen := len(rows)
		if rowLen == 0 { // 空Sheet过滤
			continue
		}

		sheet := &ExcelSheet{FilePath: excelPath, SheetIndex: indexes[i], SheetName: n, RowLength: rowLen}

		maxCellLen := 0
		for _, cells := range rows {
			cellLen := len(cells)
			maxCellLen = mathx.MaxInt(maxCellLen, cellLen)
		}
		sheet.Axis = GenAxis(maxCellLen)
		if nickRow > 0 {
			sheet.SetNick(rows[nickRow-1])
		} else {
			sheet.SetNick(sheet.Axis)
		}

		sheetRows := make([]*ExcelRow, rowLen)
		for rowIndex, cells := range rows {
			er := &ExcelRow{Index: rowIndex, Cell: cells, Sheet: sheet}
			sheetRows[rowIndex] = er
		}
		sheet.Rows = sheetRows

		sheets = append(sheets, sheet)
	}
	return sheets, nil
}
