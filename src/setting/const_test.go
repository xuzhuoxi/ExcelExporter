// Package setting
// Create on 2024/1/14
// @author xuzhuoxi
package setting

import (
	"fmt"
	"testing"
)

func TestRegArray(t *testing.T) {
	str := "66[100]77"
	rs := RegArray.FindString(str)
	fmt.Println(rs)
}
