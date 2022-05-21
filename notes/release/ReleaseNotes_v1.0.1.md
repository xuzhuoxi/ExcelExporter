## Release notes

### Known Issues in v1.0.1

### Improvements

- 放宽ExcelRow取值方法的限制，减少无必要的error返回。
- Const类型的Excel表数据支持空行间隔，输出时忽略空行。

### API Changes

- (r*ExcelRow)ValueAtIndex(index int) (value string, err error)
  只有index<0时返回error，如果没有内容则返回空字符串”“.
- (r*ExcelRow)ValueAtNick(nick string) (value string, err error)
  只有nick不存在时返回error，如果没有内容则返回空字符串”“.
- (r*ExcelRow)ValueAtAxis(axis string) (value string, err error)
  只有axis非法时返回error，如果没有内容则返回空字符串”“.

### API Add

- excel.GetAxisName(col int) string
  十进制列号 转 Excel列号
- excel.GetColNum(axis string) int
  Excel列号 转 十进制死号
  
### Fixes

- 修复当导出模式只有常量(-mode=const)时无法导出常量表的问题。此问题由TempConstProxy.GetItems造成。
- 修复Const表数据空行无法导出的问题。