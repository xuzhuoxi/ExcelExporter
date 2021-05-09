package excel

import "fmt"

type ExcelSheet struct {
	SheetIndex int
	SheetName  string

	Axis []string
	Nick []string
	Rows []*ExcelRow

	RowLength int //行数
	ColLength int //列数
}

func (es *ExcelSheet) String() string {
	strRows := ""
	for _, r := range es.Rows {
		strRows = strRows + "\t" + fmt.Sprint(r) + "\n"
	}
	return fmt.Sprintf("ExcelSheet{Index=%d, Name=%s, Axis=%s, Nick=%s, RowLen=%d, ColLen=%d,\nRow=\n%s}",
		es.SheetIndex, es.SheetName, fmt.Sprint(es.Axis), fmt.Sprint(es.Nick), es.RowLength, es.ColLength, strRows)
}

func (es *ExcelSheet) SetNick(nick []string) {
	es.Nick = nick
	for _, r := range es.Rows {
		r.Nick = nick
	}
}

func (es *ExcelSheet) GetDataRows(startIndex int) (rows []*ExcelRow) {
	return es.Rows[startIndex:]
}

func (es *ExcelSheet) GetDataRowsByFilter(startIndex int, filter func(row *ExcelRow) bool) (rows []*ExcelRow) {
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
	colIndex, rowIndex, err := ParseAxis(axis)
	if nil != err {
		return "", err
	}
	return es.Rows[rowIndex].Row[colIndex], nil
}
