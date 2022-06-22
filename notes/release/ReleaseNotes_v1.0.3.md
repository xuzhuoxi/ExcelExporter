## Release notes

### Known Issues in v1.0.3

- yaml, toml, hcl, env, properties数据导出时，key会转为**小写**，本意要求**大小写相关**。
- project.yaml中encoding与buff相关的配置未实现。
- C++表头模板未实现， C++常量模板未实现。
- sql导出未支持主键设置

### Improvements

- 实现system.yaml中datafile_formats开放数据文件格式验证，若命令行中使用未设置的数据文件格式，则记录警告并跳过。
- 完善"client,server,db"中"db"的表头导出功能与数据导出功能，实现数据sql脚本(包括表结构与数据)的导出。
- 实现mysql表结构模板与数据模板，并支持自定义数据库模板设置。
- 增加sql导出时，合并为单一文件的功能。 可在命令行参数-merge=true进行开启。
- 增加模板自定义函数：Add(加法)，Sub(减法)
- 常量导出逻辑代码修改：对TempConstProxy.GetItems中空行进行提前处理。
- 对命令行参数进行trim处理，忽略头尾空格。
- 针对sql导出，增加了定制数据库数据类型的支持，可通过在excel.yaml中的sql_field_format_row设置为对应Excel行号开启。
- 针对字段string(n)动态映射到char(*)或varchar(*)的长度的功能。
- 针对字段string映射到char(*)或varchar(*)时，长度内数据最大字节数或最大字符类决定的功能。
- 数据库配置格式增加array字段，用于标记是否为数组类字段类型。

### Fixes

- ValueAtAxis上一版本修改后可能会返回空值，补全了这部分的验证。
- 更新mysql的脚本导出模板，修复部分换行带来的脚本错误。

### Changes

- 上下文对象中EndRow属性统一改为不包含。
- 增加ExcelSheet中对当前Excel路径的引用

### API Changes

- func (o *TempConstProxy) GetItem(row int) (item ConstItem, err error)
提前进行空行处理，避免调用FindProgramLanguage(o.Language)产生错误
- 新增func (es *ExcelSheet) ValueAtIndex(colIndex int, rowIndex int) (value string, err error)
用于根据行列号取数据的功能

## Library changes in v1.0.13

### library updated

- infra-go 更新为v1.0.1
