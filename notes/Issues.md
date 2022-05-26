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
9. 尾部空字符串会数组越界。

[Info] 2022/05/26 11:52:28 [core.executeDataContext] Sheet[Data_角色好感度表]
panic: runtime error: index out of range [4] with length 4

goroutine 1 [running]:
github.com/xuzhuoxi/ExcelExporter/src/core.getRowData(0xc000272600, 0xc0002725d0, 0xc0002726c0, 0xc000010640, 0x5, 0x8, 0xc000254ff0, 0x5, 0x5)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:455 +0x22c
github.com/xuzhuoxi/ExcelExporter/src/core.executeDataContext(0xc0000fc8c0, 0xc00020a150, 0x0, 0x0)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:307 +0x7bb
github.com/xuzhuoxi/ExcelExporter/src/core.executeExcelFile(0xc00001a420, 0x55, 0x1, 0x1)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:149 +0x2d6
github.com/xuzhuoxi/ExcelExporter/src/core.loadExcelFile(0xc00001a420, 0x55, 0x138fed8, 0xc00017e1c0)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:117 +0x38f
github.com/xuzhuoxi/ExcelExporter/src/core.loadExcelFilesFromFolder.func1(0xc00001a420, 0x55, 0x138fed8, 0xc00017e1c0, 0x0, 0x0, 0x6, 0xc00001a420)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:102 +0x50
github.com/xuzhuoxi/infra-go/filex.WaldAllFiles.func1(0xc00001a420, 0x55, 0x138fed8, 0xc00017e1c0, 0x0, 0x0, 0x53, 0xc00001a1e0)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/infra-go/filex/path.go:144 +0xa9
github.com/xuzhuoxi/infra-go/filex.WalkAll.func1(0xc00001a360, 0x55, 0x138fed8, 0xc00017e1c0, 0x0, 0x0, 0x55, 0xc00022e000)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/infra-go/filex/path.go:131 +0x91
path/filepath.walk(0xc00001a360, 0x55, 0x138fed8, 0xc00017e1c0, 0xc00060dcc0, 0x0, 0x0)
        C:/Program Files/Go/src/path/filepath/path.go:414 +0x457
path/filepath.walk(0xc0000c72c0, 0x3b, 0x138fed8, 0xc000206540, 0xc000609cc0, 0x0, 0x1)
        C:/Program Files/Go/src/path/filepath/path.go:438 +0x31b
path/filepath.Walk(0xc0000c72c0, 0x3b, 0xc0000d7cc0, 0x3b, 0x5)
        C:/Program Files/Go/src/path/filepath/path.go:501 +0x125
github.com/xuzhuoxi/infra-go/filex.WalkAll(0xc0000c72c0, 0x3b, 0xc0000d7d08, 0x3b, 0xc0000d7d50)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/infra-go/filex/path.go:130 +0x85
github.com/xuzhuoxi/infra-go/filex.WaldAllFiles(0xc0000b20a8, 0x3b, 0x1309bd8, 0x0, 0x0)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/infra-go/filex/path.go:140 +0x85
github.com/xuzhuoxi/ExcelExporter/src/core.loadExcelFilesFromFolder(0xc0000b20a8, 0x3b)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:101 +0x48
github.com/xuzhuoxi/ExcelExporter/src/core.execExcelFiles()
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:93 +0x12f
github.com/xuzhuoxi/ExcelExporter/src/core.Execute(0xc0000982e8, 0xc0000d8340, 0x1, 0x1, 0xc0000d8348, 0x1, 0x1, 0xc0000d8350, 0x1, 0x1)
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/core/core.go:72 +0x35d
main.main()
        D:/workspaces/GoPath/src/github.com/xuzhuoxi/ExcelExporter/src/main.go:54 +0x8e8

10. 优化表结尾处理。