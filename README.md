# ExcelExporter

一个用于导出Excel数据的工具。

## 兼容性

go 1.16.15

## 获取

可通过下载执行文件或下载源码编译获得执行文件

- 执行文件[下载地址](/)。

- 通过github下载源码编译

	1. 执行代码。

	```
	go get -u github.com/xuzhuoxi/ExcelExporter
	```
	
	2. 编译工程。 

## 运行

程序只允许通过命令行运行

支持的命令行参数包括：-mode, -range, -lang, -file, -source, -target

- -mode

  运行的模式，支持**表头导出**、**数据导出**、**常量表导出**
    
	- 支持值：title, data, const
    
  title为表头导出，data为数据导出，const为常量表导出
    
  - 支持多值，可用英文逗号","分隔

- -range

  运行时选择的字段范围，对**表头导出**和**数据导出**有效
    
  - 支持值：client, server, db
    
  - 支持多值，可用英文逗号","分隔
  
- -lang

  运行时选择的编程语言，只针对**表头导出**有效
    
  - 支持值：go, as3, ts, java, c#
    
  - 支持多值，可用英文逗号","分隔
    
- -file

  运行输出的数据文件类型，对**数据导出**有效

  - 支持值：bin, sql, json, yaml, toml, hcl, env, properties
    
  - 支持多值，可用英文逗号","分隔

- -source

  运行时指定数据来源目录，用于覆盖配置文件project.yaml中source.value的值
    
  可选项，对**表头导出**、**数据导出**、**常量表导出**有效

- -target

  运行时指定数据来源目录，用于覆盖配置文件project.yaml中target.value的值
    
  可选项，对**表头导出**、**数据导出**、**常量表导出**有效
    
## 功能说明

三种导出功能：表头导出，数据导出，常量表导出

### 表头导出

把Excel文件中的表头信息导出为对应语言的数据结构或类

#### 表头导出流程

1. 遍历源目录中每一个符合的Excel文件。

	- 源目录由project.yaml中的soruce.value列表给出。

	- 可以通过-source参数重新指定源目录。

	- 根据project.yaml中soruce.ext_name列表进行匹配。

2. 遍历Excel文件中匹配的的Sheet。

	- 根据excel.yaml中的prefix.data属性进行匹配。

3. 根据-range参数选择对应字段列表。

	- -range参数支持client,server,db三种类型，详细请[查看]()。

4. 根据-lang参数，选择对应语言的配置及导出模板。

	- -lang参数支持go, as3, ts, java, c#，详细请[查看]()。

5. 字段列表 => 数据结构或类的字段或属性。

6. 相应文件全生成到目标目录中。

	- 目标目录由project.yaml中的target.root列表给出。
	
	- 可以通过-target参数重新指定源目录。

	- 根据-range参数的内容，文件分别生成到project.yaml中target.title.client、target.title.server、target.title.database对应的目录中去。

### 模板定制

模板文件格式为go语言模板，文档说明地址如下:

[https://golang.google.cn/pkg/text/template/](https://golang.google.cn/pkg/text/template/)

#### 注入到模板中的数据及函数

- [*TempDataProxy](/src/core/context.go)

  注入的数据结构，可通过`{{.}}`、`{{$proxy := .}}`这类模板语法取得，结构定义为：

	```
	type TempDataProxy struct {
		Sheet     *excel.ExcelSheet
		Excel     *excel.ExcelProxy
		Index     []int
		TitleName string
		Language  string
	}```    

	- Excel:[*excel.ExcelProxy](/src/core/excel/proxy.go)
	
  当前执行的Sheet

	- Sheet:[*excel.ExcelSheet](/src/core/excel/sheet.go)
	
  当前执行的Sheet

	- Index:[]int
	
  当前选中的字段索引

	- TitleName:string
	
  导出结构名称

	- Language:string
	
  选择的编程语言		

- 开放的自定义函数

	- [ToLowerCamelCase](/src/core/naming/NamingUtil.go)
	
	把字符串内容转化为**小驼峰**格式

	- [ToUpperCamelCase](/src/core/naming/NamingUtil.go)
	
	把字符串内容转化为**大驼峰**格式

	- [NowTime](/src/core/tools/time.go)
	
	取当前时间
	2006-01-02 15**:**04**:**05 PM Mon Jan
	2006-01-_2 15**:**04**:**05 PM Mon Jan

	- [NowTimeFormat](/src/core/tools/time.go)
	
	取当前时间	
	
	- [NowYear](/src/core/tools/time.go)
	 
	当前时间年份

	- [NowMonth](/src/core/tools/time.go)

	当前时间月份
	一月: 1

	- [NowDay](/src/core/tools/time.go)

	当前时间日期
	
	- [NowWeekday](/src/core/tools/time.go)

	当前时间星期几
	星期日： 0	

	- [NowHour](/src/core/tools/time.go)

	当前时间小时

	- [NowMinute](/src/core/tools/time.go)

	当前时间分钟

	- [NowSecond](/src/core/tools/time.go)

	当前时间秒
	
	- [NowUnix](/src/core/tools/time.go)

	当前时间秒戳（s）
	
	- [NowUnixNano](/src/core/tools/time.go)
 
	当前时间秒戳（ns）

### 字段数据类型关联配置

- excel.yaml

  lang_key_rows下配置Excel表中定义字段类型的行号。

- system.yaml

  program_languages列表列出来当前支持的语言的配置，其中ref为字段数据类型

- 具体语言.yaml

  配置文件位于[res/langs/templates](/res/langs)目录中。

  name:当前配置

- 具体语言.temp

  配置文件位于[res/langs/templates](/res/langs/templates)目录中。

### 数据导出

#### 选择字段名称

### 常量表导出

## 依赖性

- infra-go(库依赖) [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- excelize(库依赖) [https://github.com/360EntSecGroup-Skylar/excelize](https://github.com/360EntSecGroup-Skylar/excelize)

- goxc(编译依赖) [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## 联系作者

xuzhuoxi 

<xuzhuoxi@gmail.com> 或 <mailxuzhuoxi@163.com>

## 开源许可证

~~ExcelExporter 源代码基于[MIT许可证](/LICENSE)进行开源。~~