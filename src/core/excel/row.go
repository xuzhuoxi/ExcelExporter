package excel

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strings"
)

type ExcelRow struct {
	Index int
	Cell  []string

	Sheet *ExcelSheet
}

func (er *ExcelRow) CellLength() int {
	return len(er.Cell)
}

func (er *ExcelRow) NickLength() int {
	if nil == er.Sheet {
		return 0
	}
	return er.Sheet.NickLength()
}

func (er *ExcelRow) AxisLength() int {
	if nil == er.Sheet {
		return 0
	}
	return er.Sheet.AxisLength()
}

func (er *ExcelRow) Nick() []string {
	if nil == er.Sheet {
		return nil
	}
	return er.Sheet.Nick
}

func (er *ExcelRow) Axis() []string {
	if nil == er.Sheet {
		return nil
	}
	return er.Sheet.Axis
}

func (er *ExcelRow) String() string {
	return fmt.Sprintf("ExcelRow{Index=%d,Cell(%d)=%s}", er.Index, len(er.Cell), fmt.Sprint(er.Cell))
}

func (er *ExcelRow) Empty() bool {
	if er.CellLength() == 0 {
		return true
	}
	for _, cell := range er.Cell {
		if cell != "" && strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}

// Open to templates
// 通过索引号取值，索引号从0开始
func (er *ExcelRow) ValueAtIndex(index int) (value string, err error) {
	if index < 0 {
		return "", errors.New(fmt.Sprintf("Index(%d) out of range! ", index))
	}
	if index >= er.CellLength() {
		return "", nil
	}
	return er.Cell[index], nil
}

// Open to templates
// 通过别名取值
func (er *ExcelRow) ValueAtNick(nick string) (value string, err error) {
	index, ok := slicex.IndexString(er.Nick(), nick)
	if !ok {
		return "", errors.New(fmt.Sprintf("Nick(%s) is not exist! ", nick))
	}
	if index >= er.CellLength() {
		return "", nil
	}
	return er.Cell[index], nil
}

// Open to templates
// 通过列名取值
func (er *ExcelRow) ValueAtAxis(axis string) (value string, err error) {
	axis = strings.ToUpper(axis)
	index := GetColNum(axis) - 1
	if index < 0 {
		return "", errors.New(fmt.Sprintf("Axis(%s) is not exist! ", axis))
	}
	if index >= er.CellLength() {
		return "", nil
	}
	return er.Cell[index], nil
}
