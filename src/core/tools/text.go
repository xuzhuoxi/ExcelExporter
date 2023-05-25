// Create on 2023/5/25
// @author xuzhuoxi
package tools

import "strings"

// 格式化为Html格式的换行
func Format2HtmlNewline(text string) string {
	return strings.ReplaceAll(text, "\n", "<br/>")
}
