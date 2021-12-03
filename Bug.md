1. 导出json数据是，错误地把[]uint8类型的数据识别为[]byte类型，因此导出的类型成为了string类型。

2. yaml, toml, hcl, env, properties数据导出时，key为**大小写无关**，本意要求**大小写相关**。

3. const导出未完成。

4. system.yaml中datafile_formats数据未使用，考虑是不删除掉。

5. project.yaml中encoding与buff相关的配置未实现。

6. excel.yaml中data.pass配置未实现。

7. C++表头模板未实现， C++常量模板未实现。