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
func LoadExcels(path string) (excels []*excelize.File, err error) {
	paths := strings.Split(strings.TrimSpace(path), ",")
	if len(paths) == 0 {
		return nil, errors.New("Path Empty:" + path)
	}
	var filePaths []string

	for _, path := range paths {
		fp := filex.FormatPath(path)
		filex.WalkAll(fp, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			name := info.Name()
			if filex.CheckExt(name, "xls") || filex.CheckExt(name, "xlsx") {
				filePaths = append(filePaths, path)
			}
			return nil
		})
	}
	if len(filePaths) == 0 {
		return nil, errors.New("Path Empty:" + path)
	}
	for _, filePath := range filePaths {
		excel, err := LoadExcel(filePath)
		if err != nil {
			return nil, err
		}
		excels = append(excels, excel)
	}
	return excels, nil
}

// 加载Excel文件，过滤无Sheet情况
func LoadExcel(FileName string) (excel *excelize.File, err error) {
	excelFile, err := excelize.OpenFile(FileName)
	if nil != err {
		return nil, err
	}
	if excelFile.SheetCount <= 0 {
		return nil, errors.New("No Sheets! ")
	}
	return excelFile, nil
}

// 通过SheetPrefix作为限制加载Sheet
// 指定NickRow所在行为别名,NickRow=0时，使用列号作为别名
func LoadSheets(excelFile *excelize.File, SheetPrefix string, NickRow int) (sheets []*ExcelSheet, err error) {
	var names []string
	var indexs []int
	for index, name := range excelFile.GetSheetMap() {
		if strings.Index(name, SheetPrefix) == 0 {
			names = append(names, name)
			indexs = append(indexs, index)
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
		es := &ExcelSheet{SheetIndex: indexs[i], SheetName: n}
		es.RowLength = len(rows)
		if es.RowLength > 0 {
			es.ColLength = len(rows[0])
			es.Axis = GenAxis(es.ColLength)
			for rowIndex, row := range rows {
				er := &ExcelRow{Index: rowIndex, Length: es.ColLength, Axis: es.Axis, Row: row}
				es.Rows = append(es.Rows, er)
			}
		}
		if NickRow > 0 {
			es.SetNick(rows[NickRow-1])
		} else {
			es.SetNick(es.Axis)
		}
		sheets = append(sheets, es)
	}
	return sheets, nil
}
