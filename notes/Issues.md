1. 导出json数据时，错误地把[]uint8类型的数据识别为[]byte类型，因此导出的类型成为了string类型。
(已经临时解决)

2. yaml, toml, hcl, env, properties数据导出时，key为**大小写无关**，本意要求**大小写相关**。

4. project.yaml中encoding与buff相关的配置未实现。

5. excel.yaml中data.pass配置未实现。

6. C++表头模板未实现， C++常量模板未实现。

