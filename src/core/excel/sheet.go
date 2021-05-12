package excel

import "fmt"

type ExcelSheet struct {
	SheetIndex int    // Sheet的索引
	SheetName  string // Sheet的名称

	Axis []string // 列名
	Nick []string // 列别名

	RowLength int         // 行数
	Rows      []*ExcelRow // 行内容
}

func (es *ExcelSheet) AxisLength() int{
	return len(es.Axis)
}

func (es *ExcelSheet) NickLength() int{
	return len(es.Axis)
}

func (es *ExcelSheet) String() string {
	strRows := ""
	for _, r := range es.Rows {
		strRows = strRows + "\t" + fmt.Sprint(r) + "\n"
	}
	return fmt.Sprintf("ExcelSheet{Index=%d, Name=%s, Axis=%s, Nick=%s, RowLen=%d, \nRow=\n%s}",
		es.SheetIndex, es.SheetName, fmt.Sprint(es.Axis), fmt.Sprint(es.Nick), es.RowLength, strRows)
}

func (es *ExcelSheet) SetNick(nick []string) {
	es.Nick = nick
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
	cellIndex, rowIndex, err := ParseAxis(axis)
	if nil != err {
		return "", err
	}
	return es.Rows[rowIndex].Cell[cellIndex], nil
}
