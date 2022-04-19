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

## 配置结构说明

<pre><code>.配置根目录
├── langs: 编程语言相关配置
│   ├── templates: 模板文件目录，只支持golang模板语法
│  		├── as3_const.temp: ActionScript3语言下，常量定义模板
│  		├── as3_title.temp: ActionScript3语言下，Title定义模板
│  		├── c#_const.temp: C#语言下，常量定义模板
│  		├── c#_title.temp: C#语言下，Title定义模板
│  		├── go_const.temp: golang语言下，常量定义模板
│  		├── go_title.temp: golang语言下，Title定义模板
│  		├── java_const.temp: java语言下，常量定义模板
│  		├── java_title.temp: java语言下，Title定义模板
│  		├── ts_const.temp: TypeScript语言下，常量定义模板
│  		├── ts_title.temp: TypeScript语言下，Title定义模板
│  		├── ...: 其它语言下，常量定义模板与Title定义模板
│   ├── as3.yaml: 针对ActionScript3，不同数据文件下各基础数据类型的读写语法配置
│   ├── c#.yaml: 针对c#，不同数据文件下各基础数据类型的读写语法配置
│   ├── c++.yaml: 针对c++，不同数据文件下各基础数据类型的读写语法配置
│   ├── go.yaml: 针对golang，不同数据文件下各基础数据类型的读写语法配置
│   ├── java.yaml: 针对java，不同数据文件下各基础数据类型的读写语法配置
│   ├── ...: 其它编程语言下，不同数据文件下各基础数据类型的读写语法配置
├── proxy: 代理代码集(非必要)
│   ├── as: ActionScript3相关的代理代码集
│   ├── go: golang相关的代理代码集
│   ├── java: java相关的代理代码集 
│   ├── ts: TypeScript相关的代理代码集
│   ├── ...: 其它编程语言相关的代理代码集
├── excel.yaml: Excel的表头配置，包括数据表头配置、常量表头配置
├── project.yaml: 项目配置，包括数据源配置、数据输出配置、缓冲配置、大小端配置等
├── system.yaml: 应用配置，包括支持的编程语言配置(扩展名、读写配置、模板关联等)、数据字段类型配置、数据文件配置等
</code></pre>

### 字段数据类型关联配置

- system.yaml

	<pre><code>
	.
	├── languages: 支持的编程语言
	│   ├── name: 语言名称
	│   ├── ext: 语言源文件扩展名
	│   ├── ref: 基础数据读写配置文件路径(相对于配置根目录相对路径)
	│   ├── temps_title: title定义导出模板路径(相对于配置根目录相对路径)
	│   ├── temps_const: 常量定义导出模板路径(相对于配置根目录相对路径)
	├── datafield_formats: 支持的基础数据类型
	├── datafile_formats: 支持的数据文件格式
	</code></pre>

- project.yaml

	<pre><code>
	.
	├── source: Excel源目录
	│   ├── value: Excel源目录路径，支持多个，相对于配置根目录相对路径
	│   ├── encoding: 编码格式(如果需要)
	│   ├── ext_name: 文件扩展名，支持多个
	├── target: 输出设置
	│   ├── root: 输出目录,以'':''开关，或路径中包含'':''的，视为绝对路径 
	│   ├── title: 
	│  		├── client: client表头定义输出目录
	│  		├── server: server表头定义输出目录
	│  		├── db: db表头定义输出目录
	│   ├── data: 
	│  		├── client: client数据文件输出目录
	│  		├── server: server数据文件输出目录
	│  		├── db: db数据文件输出目录
	│   ├── const: 
	│  		├── client: client常量表输出目录
	│  		├── server: server常量表输出目录
	│   ├── encoding: 输出文件的编码格式
	├── buff: 缓冲区定义
	│   ├── big_endian: 二进制数据文件的大小端设置(true|false)
	│   ├── token: 每一个字段的缓冲大小(未使用)
	│   ├── item: 每一条数据的缓冲大小(未使用)
	│   ├── sheet: 每个sheet表的缓冲大小(未使用)
	</code></pre>

- excel.yaml

	<pre><code>
	.
	├── title&data: 
	│   ├── prefix: 启用前缀
	│   ├── outputs: 导出命名设置
	│  		├── range_name：字段域名称(client|server|db)
	│  		├── title: Title定义文件的名称所在坐标
	│  		├── data： 数据文件的名称
	│   ├── classes: 表头导出类信息
	│  		├── name：字段域名称(client|server|db)
	│  		├── value: 导出类所在坐标
	│   ├── nick_row: 字段别名行号，用于查找指定列，值为0时使用列号作为别名
	│   ├── name_row: 数据名称所在行号
	│   ├── remark_row: 数据注释所在行号
	│   ├── field_range_row: 输出选择行号，内容格式: 'c,s,d'，c指前端，s指后端，d指数据库，(01)
	│   ├── field_format_row: 字段数据格式行号
	│   ├── field_names: Title定义文件字段名称配置 
	│  		├── name: 语言名称
	│  		    row: 语言属性所在的行号
	│   ├── file_keys: 数据文件字段名称配置
	│  		├── name: 数据文件格式(bin,json,yaml等)
	│  		    row: 数据字段名称所在的行号
	│   ├── data_start_row: 数据开始行号
	├── const: 
	│   ├── prefix: 启用前缀
	│   ├── outputs: 导出类名配置
	│  		├── name: 字段域(client|server)
	│  		    value: 坐标(如： "A1")
	│   ├── classes: 常量导出类信息
	│  		├── name：字段域名称(client|server|db)
	│  		├── value: 导出类所在坐标
	│   ├── name_col: 常量名列号
	│   ├── value_col: 常量值列号
	│   ├── type_col: 常量值类型列号
	│   ├── remark_col: 注释列号
	│   ├── data_start_row: 数据开始行号
	</code></pre>

- 具体语言.yaml

  配置文件位于[res/langs](/res/langs)目录中。

	<pre><code>
	.
	├── name: 当前语言名称
	├── bool: bool数据类型
	│   ├── name: 当前语言下数据类型表达
	│   ├── operates: 操作
	│  		├── file_name: 数据文件(json，bin等)
	│  		    get: 获得数据的函数方法
	│  		    set: 设置数据的函数方法
	├── ...: 其它数据类型
	</code></pre>

- 具体语言.temp

  配置文件位于[res/langs/templates](/res/langs/templates)目录中。

  golang语法支持下的模板文件，帮助可查看**[https://golang.google.cn/pkg/text/template/](https://golang.google.cn/pkg/text/template/)**

## 运行

程序只允许通过命令行运行

支持的命令行参数包括：-env, -mode, -range, -lang, -file, -source, -target

- -env

  重新指定运行环境，运行环境指的是配置根目录。

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

三种导出功能：表头导出，常量表导出，数据导出

### 表头导出

把Excel文件中的表头信息导出为对应语言的数据结构或类

#### 表头导出流程

1. 遍历源目录中每一个符合的Excel文件。

	- 源目录由project.yaml中的soruce.value列表给出。

	- 可以通过-source参数重新指定源目录。

	- 根据project.yaml中soruce.ext_name列表进行匹配。

2. 遍历Excel文件中匹配的的Sheet。

	- 根据excel.yaml中的title&data.prefix属性进行匹配。

3. 根据-range参数选择对应字段列表。

	- -range参数支持client,server,db三种类型，详细请[查看]()。

4. 根据-lang参数，选择对应语言的配置及导出模板。

	- -lang参数支持go, as3, ts, java, c#，详细请[查看]()。

5. 字段列表 => 数据结构或类的字段或属性。

6. 相应文件全生成到目标目录中。

	- 目标根目录由project.yaml中的target.root列表给出。
	
	- 表头输出目录中project.yaml中的target.title给出，为target.root的相对路径。
	
	- 可以通过-target参数重新指定源目录。

	- 根据-range参数的内容，文件分别生成到project.yaml中target.title.client、target.title.server、target.title.database对应的目录中去。

#### 表头模板说明

1. 注入的数据对象为 [*TempDataProxy](/src/core/context.go)

   可通过`{{.}}`、`{{$proxy := .}}`这类模板语法取得，结构定义为：

	```
	type TempDataProxy struct {
		Sheet      *excel.ExcelSheet
		Excel      *excel.ExcelProxy
		TitleCtx   *TitleContext
		FileName  string
		ClassName string
		Index     []int
		Language  string
	}```    

  - Excel:[*excel.ExcelProxy](/src/core/excel/proxy.go)
	
  	当前执行的Excel数据对象

  - Sheet:[*excel.ExcelSheet](/src/core/excel/sheet.go)
	
  	当前执行的Sheet数据对象
	
  - TitleCtx:*TitleContext
	
  	当前执行的上下文数据

  - FileName:string
	
  	导出文件名称

  - ClassName:string
	
  	导出类名称

  - Index:[]int
	
  	当前选中的字段索引

  - Language:string
	
  	选择的编程语言		

2. [自定义函数](#自定义函数)

### 常量表导出

1. 遍历源目录中每一个符合的Excel文件。

	- 源目录由project.yaml中的soruce.value列表给出。

	- 可以通过-source参数重新指定源目录。

	- 根据project.yaml中soruce.ext_name列表进行匹配。

2. 遍历Excel文件中匹配的的Sheet。

	- 根据excel.yaml中的const.prefix属性进行匹配。
	
	- 根据excel.yaml中的name_col、value_col、type_col、remark_col, 定位常量的名称、值、类型、注释。
	
	- 根据excel.yaml中的data_start_row开始向正描述数据，直到遇到空行结束。 

4. 根据-lang参数，选择对应语言的配置及导出模板。

	- -lang参数支持go, as3, ts, java, c#，详细请[查看]()。

6. 相应文件全生成到目标目录中。

	- 目标根目录由project.yaml中的target.root列表给出。
	
	- 常量表输出目录中project.yaml中的target.const给出，为target.root的相对路径。
	
	- 可以通过-target参数重新指定源目录。

	- 根据-range参数的内容，文件分别生成到project.yaml中target.const.client、target.const.server对应的目录中去。

#### 注入到常量模板中的数据及函数

1. 注入的数据对象为 [*TempConstProxy](/src/core/context.go)

  可通过`{{.}}`、`{{$proxy := .}}`这类模板语法取得，结构定义为：

  ```
	type TempConstProxy struct {
    	Sheet     *excel.ExcelSheet
		Excel     *excel.ExcelProxy
		ConstCtx  *ConstContext
		FileName  string
		ClassName string
		Language  string
		StartRow  int
		EndRow    int
	}``` 

  - Excel:[*excel.ExcelProxy](/src/core/excel/proxy.go)
	
  	当前执行的Excel数据对象

  - Sheet:[*excel.ExcelSheet](/src/core/excel/sheet.go)
	
  	当前执行的Sheet数据对象
	
  - ConstCtx:*ConstContext
	
  	当前执行的上下文数据

  - FileName:string
	
  	导出文件名称

  - ClassName:string
	
  	常量类名称

  - Index:[]int
	
  	当前选中的字段索引

  - Language:string
	
  	选择的编程语言	

  - StartRow:int
	
  	常量数据开始行号

  - EndRow:int
	
  	常量数据结束行号	

2. [自定义函数](#自定义函数)

### 数据导出

支持的数据导出格式：bin(二进制), json。

sql导出**未实现**。

yaml, toml, hcl, env, properties数据导出时，字段名称会**强制处理**为小写，本意要求**大小写相关**，固暂时不开放

### 模板定制

模板文件格式为go语言模板，文档说明地址如下:

[https://golang.google.cn/pkg/text/template/](https://golang.google.cn/pkg/text/template/)

#### 自定义函数

自定义函数对全部模板有效

**注意**：自定义函数的返回值必须是1个或2个，这是官方要求。

当返回值为2个时，第2个返回值类型必须是error。

- [ToLowerCamelCase](/src/core/naming/NamingUtil.go)

  把字符串内容转化为**小驼峰**格式

- [ToUpperCamelCase](/src/core/naming/NamingUtil.go)

  把字符串内容转化为**大驼峰**格式

- [NowTime](/src/core/tools/time.go)

  取当前时间

- [NowTimeStr](/src/core/tools/time.go)

  取当前时间默认格式字符串

- [NowTimeFormat](/src/core/tools/time.go)

  取当前时间	
  2006-01-02 15**:**04**:**05 PM Mon Jan
  2006-01-_2 15**:**04**:**05 PM Mon Jan

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

## 依赖性

- infra-go(库依赖) [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- excelize(库依赖) [https://github.com/360EntSecGroup-Skylar/excelize](https://github.com/360EntSecGroup-Skylar/excelize)

- goxc(编译依赖) [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## 联系作者

xuzhuoxi 

<xuzhuoxi@gmail.com> 或 <mailxuzhuoxi@163.com>

## 开源许可证

~~ExcelExporter 源代码基于[MIT许可证](/LICENSE)进行开源。~~