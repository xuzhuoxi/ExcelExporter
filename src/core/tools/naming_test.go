package tools

import (
	"fmt"
	"testing"
)

var (
	namings = []string{
		"dbc",
		"_dbc",
		"__dbc",
		"dbc_dbc",
		"_dbc_dbc",
		"__dbc_dbc",
	}
)

func TestClearUnderScore(t *testing.T) {
	for _, naming := range namings {
		//ClearUnderScore(naming)
		fmt.Println(naming, ClearUnderscore(naming))
	}
}
