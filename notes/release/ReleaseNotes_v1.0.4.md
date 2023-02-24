## Release notes

### Known Issues in v1.0.4

- yaml, toml, hcl, env, properties数据导出时，key会转为**小写**，本意要求**大小写相关**。
- project.yaml中encoding与buff相关的配置未实现。
- C++表头模板未实现， C++常量模板未实现。

### Improvements

- excel配置：数据表头配置中的"data_start_row"修改为"data_start"，使用行列号进行配置。。

### Fixes

### Changes

- excel.yaml配置中，'title&data'下‘data_start_row’删除。
- excel.yaml配置中，'title&data'下新增‘data_start’配置，格式为行列号。

### API Changes

- 修改：core/core.go中函数parseRangeRow增加了开始索引参数startIndex。
- 修改：core/excel.utils.go中函数ParseAxis改名为ParseAxisIndex。
- 新增：TitleData结构体增加DataStartRow函数与DataStartColIndex函数。

## Library changes in v1.0.4

### library updated