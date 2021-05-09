package excel

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/slicex"
	"strings"
)

type ExcelRow struct {
	Index  int
	Length int

	Axis []string
	Nick []string
	Row  []string
}

func (er *ExcelRow) String() string {
	return fmt.Sprintf("ExcelRow{Index=%d, Length=%d, Axis=%s, Nick=%s, Row=%s}", er.Index, er.Length,
		fmt.Sprint(er.Axis), fmt.Sprint(er.Nick), fmt.Sprint(er.Row))
}

func (er *ExcelRow) Empty() bool {
	for _, value := range er.Row {
		if value != "" && strings.TrimSpace(value) != "" {
			return false
		}
	}
	return true
}

// Open to templates
// 通过索引号取值，索引号从0开始
func (er *ExcelRow) ValueAtIndex(index int) (value string, err error) {
	if index < 0 || index >= er.Length {
		return "", errors.New(fmt.Sprintf("Index(%d) out of range! ", index))
	}
	return er.Row[index], nil
}

// Open to templates
// 通过别名取值
func (er *ExcelRow) ValueAtNick(nick string) (value string, err error) {
	index, ok := slicex.IndexString(er.Nick, nick)
	if !ok {
		return "", errors.New(fmt.Sprintf("Nick(%s) is not exist! ", nick))
	}
	return er.Row[index], nil
}

// Open to templates
// 通过列名取值
func (er *ExcelRow) ValueAtAxis(axis string) (value string, err error) {
	index, ok := slicex.IndexString(er.Axis, axis)
	if !ok {
		return "", errors.New(fmt.Sprintf("Axis(%s) is not exist! ", axis))
	}
	return er.Row[index], nil
}
