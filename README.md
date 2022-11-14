# ExcelExporter  
一个用于导出Excel数据的工具。  
根据模板导出Excel数据。支持**多种数据格式**和**任何一种编程语言**。支持**多操作系统**。  

中文 | [English](/README_EN.md)  

## 兼容性
go 1.16.15  

## 如何获取
可通过下载执行文件或下载源码编译获得执行文件  

- 执行文件[下载地址](https://github.com/xuzhuoxi/ExcelExporter/releases)。
- 通过github下载源码编译

1. 执行代码。  
```
go get -u github.com/xuzhuoxi/ExcelExporter
```

2. 编译工程。 
Windows下执行[goxc_build.bat](/build/goxc_build.bat)  
Linux下执行[goxc_build.sh](/build/goxc_build.sh)  

## 配置环境说明

<pre><code>.配置根目录
├── db: 数据库相关配置与sql模板
├── lang: 编程语言相关配置
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
│   templates: 模板文件目录，只支持golang模板语法
│    ├── as3_const.temp: ActionScript3语言下，常量定义模板
│    ├── as3_title.temp: ActionScript3语言下，Title定义模板
│    ├── c#_const.temp: C#语言下，常量定义模板
│    ├── c#_title.temp: C#语言下，Title定义模板
│    ├── go_const.temp: golang语言下，常量定义模板
│    ├── go_title.temp: golang语言下，Title定义模板
│    ├── java_const.temp: java语言下，常量定义模板
│    ├── java_title.temp: java语言下，Title定义模板
│    ├── ts_const.temp: TypeScript语言下，常量定义模板
│    ├── ts_title.temp: TypeScript语言下，Title定义模板
│    ├── ...: 其它语言下，常量定义模板与Title定义模板
</code></pre>

### 应用环境配置说明

- system.yaml

  <pre><code>.应用系统级配置
  ├── languages: 支持的编程语言配置
  │   ├── name: 编程语言名称
  │   ├── ext:  源代码文件扩展名
  │   ├── ref:  基础数据读写配置文件路径(相对于配置根目录相对路径)
  │   ├── temps_title: title导出类定义导出模板路径(相对于配置根目录相对路径)
  │   ├── temps_const: 常量定义导出模板路径(相对于配置根目录相对路径)
  ├── database: 支持的数据库配置
  │   ├── default: 默认使用的数据库配置，必须为DatabaseList中的一个
  │   ├── list:    数据库的配置静静列表
  │      ├── name: 数据库名称
  │          ref:  数据库具体配置文件所在路径(相对于配置根目录相对路径)
  │          temps_table: 表结构sql生成模板列表
  │          temps_data:  表数据sql生成模板列表
  ├── datafield_formats: 支持的基础数据类型列表
  ├── datafile_formats:  支持的数据文件格式列表
  </code></pre>

- project.yaml

  <pre><code>.应用项目级配置
  ├── source: Excel源目录
  │   ├── value: Excel源目录路径，支持多个，相对于配置根目录相对路径
  │   ├── encoding: 编码格式(如果需要)
  │   ├── ext_name: 文件扩展名，支持多个
  ├── target: 输出设置
  │   ├── root: 输出目录,以'':''开关，或路径中包含'':''的，视为绝对路径 
  │   ├── title: 
  │      ├── client: client表头定义输出目录
  │      ├── server: server表头定义输出目录
  │      ├── db: db表头定义输出目录
  │   ├── data: 
  │      ├── client: client数据文件输出目录
  │      ├── server: server数据文件输出目录
  │      ├── db: db数据文件输出目录
  │   ├── const: 
  │      ├── client: client常量表输出目录
  │      ├── server: server常量表输出目录
  │   ├── encoding: 输出文件的编码格式
  ├── buff: 缓冲区定义
  │   ├── big_endian: 二进制数据文件的大小端设置(true|false)
  │   ├── token: 每一个字段的缓冲大小(未使用)
  │   ├── item: 每一条数据的缓冲大小(未使用)
  │   ├── sheet: 每个sheet表的缓冲大小(未使用)
  </code></pre>

- excel.yaml

  <pre><code>.应用Excel配置
  ├── title&data: 
  │   ├── prefix: 启用前缀
  │   ├── outputs: 导出命名设置
  │      ├── range_name：字段域名称(client|server|db)
  │      ├── title: Title定义文件的名称所在坐标
  │      ├── data： 数据文件的名称
  │   ├── classes: 表头导出类信息
  │      ├── name：字段域名称(client|server|db)
  │      ├── value: 导出类所在坐标
  │   ├── nick_row: 字段别名行号，用于查找指定列，值为0时使用列号作为别名
  │   ├── name_row: 数据名称所在行号
  │   ├── remark_row: 数据注释所在行号
  │   ├── field_range_row: 输出选择行号，内容格式: 'c,s,d'，c指前端，s指后端，d指数据库，(01)
  │   ├── field_format_row: 字段数据格式行号
  │   ├── sql_field_format_row: 数据库字段类型定制行号，0为不定制
  │   ├── field_names: Title定义文件字段名称配置 
  │      ├── name: 语言名称
  │          row: 语言属性所在的行号
  │   ├── file_keys: 数据文件字段名称配置
  │      ├── name: 数据文件格式(bin,json,yaml等)
  │          row: 数据字段名称所在的行号
  │   ├── data_start_row: 数据开始行号
  ├── const: 
  │   ├── prefix: 启用前缀
  │   ├── outputs: 导出类名配置
  │      ├── name: 字段域(client|server)
  │          value: 坐标(如： "A1")
  │   ├── classes: 常量导出类信息
  │      ├── name：字段域名称(client|server|db)
  │      ├── value: 导出类所在坐标
  │   ├── name_col: 常量名列号
  │   ├── value_col: 常量值列号
  │   ├── type_col: 常量值类型列号
  │   ├── remark_col: 注释列号
  │   ├── data_start_row: 数据开始行号
  </code></pre>

#### 编程语言配置说明  

- 具体语言.yaml(如[go.yaml](/res/lang/go.yaml))  
  配置文件位于[res/lang](/res/lang)目录中。  

  <pre><code>.具体编程语言配置
  ├── lang_name: 当前语言名称
  ├── data_types: 数据库配置列表
  │  ├── name: 字段数据类型(Excel表上填的)
  │       lang: 编程语言对应的数据类型
  │       operates: 针对不同数据文件的操作方法
  │      ├── file_name: 数据文件类型(json，bin等)
  │          get: 读取方法字符表达
  │          set: 写入方法字符表达
  </code></pre>

- 模板文件.temp(如[go_titel.temp](/res/template/go_titel.temp)、[go_const.temp](/res/template/go_const.temp))
  模板文件位于[res/template](/res/template)目录中。  
  golang语法支持下的模板文件，帮助可查看[**https://golang.google.cn/pkg/text/template/**](https://golang.google.cn/pkg/text/template/)  

#### 数据库配置说明  

- 数据库.yaml(如[mysql.yaml](/res/db/mysql.yaml))
  配置文件位于[res/db](/res/db)目录中。  

  <pre><code>.具体数据配置
  ├── db_name: 数据库名称
  ├── scale_char: Char字符比例
  ├── scale_varchar: archar字符比例
  ├── types: 数据库数据类型描述列表
  │  ├── name: 字段名称(标准化后，如string(5)=>string(*))
  │       type: 对应数据的字段数据类型
  │       number: 是否为数值类型
  │       array: 是否为数组类型
  </code></pre>

- 模板文件.temp(如[mysql_table.temp](/res/db/mysql_table.temp)、[mysql_data.temp](/res/db/mysql_data.temp))
  模板文件位于[/res/db](/res/db)目录中。  
  golang语法支持下的模板文件，帮助可查看[**https://golang.google.cn/pkg/text/template/**](https://golang.google.cn/pkg/text/template/)  

## 运行

程序只允许通过命令行运行  

支持的命令行参数包括：-env, -mode, -range, -lang, -file, -merge, -source, -target  

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

- -merge  
  导出sql时是不合并为一个文件，true代表合并，false代表不合并  
  - 支持值：true, false  

- -source  
  运行时指定数据来源目录，用于覆盖配置文件project.yaml中source.value的值  
  可选项，对**表头导出**、**数据导出**、**常量表导出**有效  

- -target  
  运行时指定数据来源目录，用于覆盖配置文件project.yaml中target.value的值  
  可选项，对**表头导出**、**数据导出**、**常量表导出**有效  

## 功能说明

- 三种基础导出功能：[**表头导出**](#表头导出)，[**常量表导出**](#常量表导出)，[**数据导出**](#数据导出)  
- 特殊导出功能：[**Sql导出**](#Sql导出)  

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

1. 注入的数据对象为 [\*TempTitelProxy](/src/core/context_title.go)  
可通过`{{.}}`、`{{$proxy := .}}`这类模板语法取得，结构定义为：  

```golang
  type TempTitleProxy struct {
    Sheet      *excel.ExcelSheet // 当前执行的Sheet数据对象
    Excel      *excel.ExcelProxy // 当前Excel代理，可能包含多个Excel
    TitleCtx   *TitleContext     // 当前执行的表头上下文数据
    FileName   string            // 表头导出类文件名
    ClassName  string            // 表头导出类名
    FieldIndex []int             // 当前选中的字段索引
    Language   string            // 当前的选择的编程语言
  }
```

  - Excel:[\*excel.ExcelProxy](/src/core/excel/proxy.go)  
    当前执行的Excel数据代理对象  
  - Sheet:[\*excel.ExcelSheet](/src/core/excel/sheet.go)  
    当前执行的Sheet数据对象  
  - TitleCtx:[\*TitleContext](/src/core/context_title.go)  
    当前执行的表头上下文数据  
  - FileName:string  
    表头导出类文件名  
  - ClassName:string  
    表头导出类名  
  - FieldIndex:[]int  
    当前选中的字段索引  
  - Language:string  
    当前的选择的编程语言  

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

1. 注入的数据对象为[\*TempConstProxy](/src/core/context_const.go)  
可通过`{{.}}`、`{{$proxy := .}}`这类模板语法取得，结构定义为：  

```golang
  type TempConstProxy struct {
    Sheet     *excel.ExcelSheet // 当前执行的Sheet数据对象
    Excel     *excel.ExcelProxy // 当前Excel代理，可能包含多个Excel
    ConstCtx  *ConstContext     // 当前执行的上下文数据
    FileName  string            // 导出文件名
    ClassName string            // 导出常量类名
    Language  string            // 导出对应的编程语言
    StartRow  int               // 数据开始行号
    EndRow    int               // 数据结束行号
  }
```

  - Excel:[\*excel.ExcelProxy](/src/core/excel/proxy.go)  
    当前执行的Excel数据代理对象  
  - Sheet:[\*excel.ExcelSheet](/src/core/excel/sheet.go)  
    当前执行的Sheet数据对象  
  - ConstCtx:[\*ConstContext](/src/core/context_const.go)  
    当前执行的上下文数据  
  - FileName:string  
    导出文件名称  
  - ClassName:string  
    导出常量类名  
  - Language:string  
    导出对应的编程语言  
  - StartRow:int  
    数据开始行号  
  - EndRow:int  
    数据结束行号  

2. [自定义函数](#自定义函数)  

### 数据导出
- 支持的数据导出格式：bin(二进制), json, sql。  
- yaml, toml, hcl, env, properties数据导出时，字段名称会**强制处理**为小写，本意要求**大小写相关**，默认不开放  
- 要开放yaml等数据导出，请修改system.yaml文件，在"datafiel_formats"列表中补充。  

### Sql导出

- **Sql导出依赖于表头导出与数据导出的设置。**   

- 当以下三个条件**同时具备**时，进行sql导出。  
  1. -ragne中包含db项  
  2. -file中包含sql项  
  3. -mode中至少包含title或data其中之一。  

- 导出流程：  
  1. 遍历Excel文件及Sheet与[**表头导出**](#表头导出)和[**数据导出**](#数据导出)一致。  
  2. 设置-merge参数为true时，只产出一个sql文件(all_merge.sql)  
  3. 关闭-merge参数或设置为false时，产出"文件名.talbe.sql"和"文件名.data.sql", table.sql文件为表结构更新脚本，data.sql为数据更新脚本。  

#### 注入到常量模板中的数据及函数

1. 注入的数据对象为[\*TempSqlProxy](/src/core/context_sql.go)  
可通过`{{.}}`、`{{$proxy := .}}`这类模板语法取得，结构定义为：  

```golang
  type TempSqlProxy struct {
    Sheet      *excel.ExcelSheet // 当前执行的Sheet数据对象
    Excel      *excel.ExcelProxy // 当前执行的Excel数据代理对象
    SqlCtx     *SqlContext       // 当前执行的Sql上下文
    TableName  string            // 数据库表名
    FieldIndex []int             // 字段选择索引
    StartRow   int               // 开始行号
    EndRow     int               // 结束行号
  }
```

  - Excel:[\*excel.ExcelProxy](/src/core/excel/proxy.go)  
    当前执行的Excel数据代理对象  
  - Sheet:[\*excel.ExcelSheet](/src/core/excel/sheet.go)  
    当前执行的Sheet数据对象  
  - SqlCtx:[\*SqlContext](/src/core/context_sql.go)  
    当前执行的Sql上下文  
  - TableName:string  
    数据库表名  
  - FieldIndex:string  
    字段选择索引  
  - StartRow:string  
    开始行号  
  - EndRow:int  
    结束行号  

2. [自定义函数](#自定义函数)  

### 模板定制  

模板文件格式为go语言模板，文档说明地址如下:  
[https://golang.google.cn/pkg/text/template/](https://golang.google.cn/pkg/text/template/)  

#### 自定义函数  
自定义函数对全部模板有效  
**注意**：自定义函数的返回值必须是1个或2个，这是官方要求。  
当返回值为2个时，第2个返回值类型必须是error。  

- [ToLowerCamelCase](/src/core/tools/naming.go)  
  把字符串内容转化为**小驼峰**格式  

- [ToUpperCamelCase](/src/core/tools/naming.go)  
  把字符串内容转化为**大驼峰**格式  

- [Add](/src/core/tools/math.go)  
  加法  

- [Sub](/src/core/tools/math.go)  
  减法  

- [NowTime](/src/core/tools/time.go)  
  取当前时间  

- [NowTimeStr](/src/core/tools/time.go)  
  取当前时间默认格式字符串  

- [NowTimeFormat](/src/core/tools/time.go)  
  取当前时间  
  2006-01-02 15**:**04**:**05 PM Mon Jan  
  2006-01-\_2 15**:**04**:**05 PM Mon Jan  

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
ExcelExporter 源代码基于[MIT许可证](/LICENSE)进行开源。