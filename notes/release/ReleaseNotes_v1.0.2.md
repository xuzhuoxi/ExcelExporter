## Release notes

### Known Issues in v1.0.1

- yaml, toml, hcl, env, properties数据导出时，key为**大小写无关**，本意要求**大小写相关**。

### Improvements

- 增加exce.yaml中数据控制列control_row，用于指定数据范围。

### API Changes

- ~~core.parseRangeRow(sheet *excel.ExcelSheet, rangeRow *excel.ExcelRow, rangeIndex uint) (selects []int, err error)~~
  core.parseRangeRow(sheet *excel.ExcelSheet, rangeRow *excel.ExcelRow, rangeIndex uint, maxSize int) (selects []int, err error)
  增加范围限制

### API Add

- core.getControlSize(sheet *excel.ExcelSheet) (size int)
  计算处理范围：扫描数据，遇到空字符或空格字符串结束

### Fixes

- 修复尾部空字符串会忽略带来的数组越界。