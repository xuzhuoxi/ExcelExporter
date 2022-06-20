## Release notes

### Known Issues in v1.0.3

- yaml, toml, hcl, env, properties数据导出时，key为**大小写无关**，本意要求**大小写相关**。
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

### Fixes

- ValueAtAxis上一版本修改后可能会返回空值，补全了这部分的验证。

### Changes

- 上下文对象中EndRow属性统一改为不包含。
- 增加ExcelSheet中对当前Excel路径的引用

### API Changes

## Library changes in v1.0.13

### library updated

#### Updated

#### No longer available

#### Added