1. 导出json数据时，错误地把[]uint8类型的数据识别为[]byte类型，因此导出的类型成为了string类型。
(已经临时解决)

2. yaml, toml, hcl, env, properties数据导出时，key为**大小写无关**，本意要求**大小写相关**。

3. system.yaml中datafile_formats数据未使用，考虑是不删除掉。

4. project.yaml中encoding与buff相关的配置未实现。

5. excel.yaml中data.pass配置未实现。

6. C++表头模板未实现， C++常量模板未实现。

7. Sql表定义与数据导出未实现。

	- 模板分两部分，一是表结构定义(create)，二是数据更新

	- 创建表：
	
	```
	CREATE TABLE table_name
	(column1 datatype, column2 datatype,
	column3 datatype, column4 datatype,
	...	, primary key( column1, column2, ...))
	```

	- 删除全部数据：
	
	`truncate table table_name`
	
	- 批量插入数据：

	```
	Insert into table_name (column1, column2 column3, ...) 
	values (value1, value2 value3, ...), 
	(value1, value2 value3, ...), 
	(value1, value2 value3, ...) ...
	```

	- 更新数据：
	
	```
	update table_name
	set column=value, column1=value, ...
	where someColumn=someValue
	```
8. 增加主键定义
9. 由于数据库数据类型的限制

10. EndRow改为不包含
11. ValueAtAxis 可能会返回空值，应该补充检查
12. Sql数据转义判断错误