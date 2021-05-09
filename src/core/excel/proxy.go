package excel

import "errors"

type ExcelProxy struct {
	Sheets   []*ExcelSheet
	DataRows []*ExcelRow
}

func (ep *ExcelProxy) GetSheet(sheet string) (es *ExcelSheet, err error) {
	for _, s := range ep.Sheets {
		if s.SheetName == sheet {
			return s, nil
		}
	}
	return nil, errors.New("No Sheet is " + sheet)
}

// Open to templates
// 通过Sheet的名称和坐标取值，坐标格式：B4
func (ep *ExcelProxy) ValueAtAxis(sheet string, axis string) (value string, err error) {
	s, err := ep.GetSheet(sheet)
	if nil != err {
		return "", err
	}
	return s.ValueAtAxis(axis)
}

// 合并全部sheet的行数据。
// 从StartRow开始。
// 清除空数据
func (ep *ExcelProxy) MergedRows(StartRow int) (err error) {
	var rows []*ExcelRow
	for _, sheet := range ep.Sheets {
		rows = append(rows, sheet.GetDataRows(StartRow-1)...)
	}
	if len(rows) == 0 {
		return errors.New("Rows is empty! ")
	} else {
		ep.DataRows = rows
		return nil
	}
}

// 合并全部sheet的行数据。
// 从StartRow开始。
// 清除空数据
func (ep *ExcelProxy) MergedRowsByFilter(StartRow int, filter func(row *ExcelRow) bool) (err error) {
	var rows []*ExcelRow
	for _, sheet := range ep.Sheets {
		rows = append(rows, sheet.GetDataRowsByFilter(StartRow-1, filter)...)
	}
	if len(rows) == 0 {
		return errors.New("Rows is empty! ")
	} else {
		ep.DataRows = rows
		return nil
	}
}

// 加载SourcePath指定的一个或多个Excel文件。
// SourcePath支持多路径模式，用","分隔。
// SourcePath支持文件夹，不支持递归。
func (ep *ExcelProxy) LoadSheets(SourcePath string, SheetPrefix string, NickRow int) error {
	excels, err := LoadExcels(SourcePath)
	if nil != err {
		return err
	}
	for _, excel := range excels {
		sheets, err := LoadSheets(excel, SheetPrefix, NickRow)
		if nil != err {
			return err
		}
		ep.Sheets = append(ep.Sheets, sheets...)
	}
	return nil
}
