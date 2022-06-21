package excel

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/mathx"
	"os"
	"strconv"
	"strings"
)

const (
	ExtXlsx = "xlsx"
)

// 1=>A, 2=>B
func GetAxisName(col int) string {
	return mathx.System10To26(col)
}

// A=>1，B=>2
func GetColNum(axis string) int {
	return mathx.System26To10(axis)
}

// 3 => [A, B, C]
func GenAxis(length int) []string {
	rs := make([]string, length, length)
	for index := 0; index < length; index += 1 {
		rs[index] = mathx.System10To26(index + 1)
	}
	return rs
}

// "A1" => 0, 0, nil
func ParseAxis(axis string) (colIndex int, rowIndex int, err error) {
	Axis := strings.ToUpper(strings.TrimSpace(axis))
	bs := []byte(Axis)
	var c, r []byte
	for index, b := range bs {
		if !(b >= 'A' && b <= 'Z') {
			c, r = bs[:index], bs[index:]
		}
	}
	if nil == c && nil == r {
		return 0, 0, errors.New("Axis Error:" + axis)
	}
	colNum := mathx.System26To10(string(c))
	rowNum, err := strconv.Atoi(string(r))
	if nil != err {
		return 0, 0, err
	}
	return colNum - 1, rowNum - 1, nil
}

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
			name := info.Name()
			if filex.CheckExt(name, ExtXlsx) {
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

// 通过SheetPrefix作为限制加载Sheet, 使用""代表不限制
// 指定NickRow所在行为别名,NickRow=0时，使用列号作为别名
func LoadSheets(excelPath string, excelFile *excelize.File, sheetPrefix string, nickRow int) (sheets []*ExcelSheet, err error) {
	var names []string
	var indexes []int
	for index, name := range excelFile.GetSheetMap() {
		if strings.Index(name, sheetPrefix) == 0 {
			names = append(names, name)
			indexes = append(indexes, index)
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
