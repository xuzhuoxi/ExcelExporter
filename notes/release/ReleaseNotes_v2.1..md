## Release notes with v2.1

### Known Issues 
- yaml, toml, hcl, env, properties数据导出时，key会转为**小写**，本意要求**大小写相关**。
- project.yaml中encoding与buff相关的配置未实现。
- C++表头模板未实现， C++常量模板未实现。

### New Features

### Fixes  

### Optimization  
+ 协议表(Proto)导出属性支持对自定义结构的引用支持。
+ 协议表(Proto)导出属性支持对自定义结构数组的引用支持。

##### Changes  
+ 调整了 Proto 导出时使用到的模板。

##### API Changes

## Library changes

### library updated 
- gjson 更新为 v1.17.0
- sjson 更新为 v1.2.5
- infra-go 更新为 v1.0.4