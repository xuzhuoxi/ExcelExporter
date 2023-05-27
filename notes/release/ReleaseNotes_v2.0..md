## Release notes with v2.0

### Known Issues 
- yaml, toml, hcl, env, properties数据导出时，key会转为**小写**，本意要求**大小写相关**。
- project.yaml中encoding与buff相关的配置未实现。
- C++表头模板未实现， C++常量模板未实现。

### New Features
+ 增加协议表(Proto)的解释结构导出功能。 
  + 新增 TempProtoProxy 结构，用于注入协议导出模板。
  + 新增 ProtoItem 结构，用于记录单条协议的信息。
  + 新增 ProtoFieldItem 结构，用于记录协议单个属性的信息。

### Fixes  

### Optimization  
+ 优化了Title导出时TempTitleProxy结构功能。
+ 优化了Title导出时各编译语言的模板实现，使之更清晰。
+ 优化了system.yaml、project.yaml、excel.yaml的配置。
+ 优化命令行参数的验证。

##### Changes  
+ 调整了 Title 导出时使用到的模板。
+ 调整了 Const 导出时使用到的模板。

##### API Changes  
+ TempTitleProxy
  + 重命名部分公开的函数。
  + 增加 TitleFieldItem 结构，用于记录字段信息. 
  + 增加 GetFields 函数。用于获取全部导出字段信息。返回 TitleFieldItem 的数组。
+ TitleContext 
  + ProgramLanguage  =》Language
+ TempSqlProxy
  + 重命名部分公开的函数
  + 重命名结构 FieldItem =》SqlFieldItem。
+ SqlContext
  + ProgramLanguage  =》Language
+ ConstContext
  + ProgramLanguage  =》Language

## Library changes

### library updated  