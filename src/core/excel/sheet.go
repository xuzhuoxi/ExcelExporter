package excel

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/mathx"
)

type ExcelSheet struct {
	FilePath   string //文件路径
	SheetIndex int    // Sheet的索引
	SheetName  string // Sheet的名称

	Axis []string // 列名
	Nick []string // 列别名

	RowLength int         // 行数
	Rows      []*ExcelRow // 全部行内容
}

func (es *ExcelSheet) AxisLength() int {
	return len(es.Axis)
}

func (es *ExcelSheet) NickLength() int {
	return len(es.Axis)
}

func (es *ExcelSheet) String() string {
	strRows := ""
	for _, r := range es.Rows {
		strRows = strRows + "\t" + fmt.Sprint(r) + "\n"
	}
	return fmt.Sprintf("ExcelSheet{SheetIndex=%d, Name=%s, Axis=%s, Nick=%s, RowLen=%d, \nRow=\n%s}",
		es.SheetIndex, es.SheetName, fmt.Sprint(es.Axis), fmt.Sprint(es.Nick), es.RowLength, strRows)
}

func (es *ExcelSheet) SetNick(nick []string) {
	es.Nick = nick
}

func (es *ExcelSheet) GetRowAt(rowIndex int) (row *ExcelRow) {
	if rowIndex < 0 || rowIndex >= es.RowLength {
		return nil
	}
	return es.Rows[rowIndex]
}

func (es *ExcelSheet) GetRowsFrom(startIndex int) (rows []*ExcelRow) {
	if startIndex < 0 || startIndex >= es.RowLength {
		return nil
	}
	return es.Rows[startIndex:]
}

func (es *ExcelSheet) GetRowsBetween(startIndex int, endIndex int) (rows []*ExcelRow) {
	if startIndex == endIndex {
		return nil
	}
	min := mathx.MinInt(startIndex, endIndex)
	if min < 0 {
		return nil
	}
	max := mathx.MaxInt(startIndex, endIndex)
	if max >= es.RowLength {
		return nil
	}
	return es.Rows[min:max]
}

func (es *ExcelSheet) GetRowsByFilter(startIndex int, filter func(row *ExcelRow) bool) (rows []*ExcelRow) {
	for _, row := range es.Rows[startIndex:] {
		if filter(row) {
			rows = append(rows, row)
		}
	}
	return
}

// Open to templates
// 通过坐标取值，坐标格式：B4
func (es *ExcelSheet) ValueAtAxis(axis string) (value string, err error) {
	cellIndex, rowIndex, err := ParseAxisIndex(axis)
	if nil != err {
		return "", err
	}
	row := es.Rows[rowIndex]
	if cellIndex >= row.CellLength() {
		return "", nil
	}
	return row.ValueAtIndex(cellIndex)
}

// Open to templates
// 通过坐标取值，
func (es *ExcelSheet) ValueAtIndex(colIndex int, rowIndex int) (value string, err error) {
	row := es.Rows[rowIndex]
	if colIndex >= row.CellLength() {
		return "", nil
	}
	return row.ValueAtIndex(colIndex)
}
