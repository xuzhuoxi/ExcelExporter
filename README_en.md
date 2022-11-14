# ExcelExporter  
A tool for exporting Excel data.  
Export Excel data according to templates. Supports **multiple data formats** and **any programming language**. **Multi-OS** support.  

[中文](README.md) | English 

## Compatibility  
go 1.16.15  

## How to get  
The executable file can be obtained by downloading the executable file or downloading the source code and compiling  

- Execution file [download link](https://github.com/xuzhuoxi/ExcelExporter/releases).  
- Download source code to compile through github  

1. Execute the code.  

````
go get -u github.com/xuzhuoxi/ExcelExporter
````

2. Compile the project.  
Execute [goxc_build.bat](/build/goxc_build.bat) under Windows  
Execute [goxc_build.sh](/build/goxc_build.sh) under Linux  

## Configuration environment description

<pre><code>.Configuration root directory
├── db: database related configuration and sql template
├── lang: programming language related configuration
│ ├── as3.yaml: For ActionScript3, read and write syntax configuration of each basic data type under different data files
│ ├── c#.yaml: For c#, read and write syntax configuration of each basic data type under different data files
│ ├── c++.yaml: For C++, read and write syntax configuration of each basic data type under different data files
│ ├── go.yaml: For golang, read and write syntax configuration of each basic data type under different data files
│ ├── java.yaml: For java, read and write syntax configuration of each basic data type under different data files
│ ├── ...: In other programming languages, read and write syntax configuration of each basic data type in different data files
├── proxy: proxy code set (optional)
│ ├── as: ActionScript3 related proxy code set
│ ├── go: golang-related proxy code set
│ ├── java: Java-related proxy code set
│ ├── ts: TypeScript-related proxy code set
│ ├── ...: other programming language-related proxy code sets
├── excel.yaml: Excel header configuration, including data header configuration, constant header configuration
├── project.yaml: project configuration, including data source configuration, data output configuration, buffer configuration, big and small end configuration, etc.
├── system.yaml: Application configuration, including supported programming language configuration (extension, read-write configuration, template association, etc.), data field type configuration, data file configuration, etc.
│ templates: Template file directory, only supports golang template syntax
│ ├── as3_const.temp: In ActionScript3 language, constant definition template
│ ├── as3_title.temp: In ActionScript3 language, Title defines the template
│ ├── c#_const.temp: In C# language, constant definition template
│ ├── c#_title.temp: In C# language, Title defines the template
│ ├── go_const.temp: In golang language, constant definition template
│ ├── go_title.temp: In golang language, Title defines the template
│ ├── java_const.temp: constant definition template in java language
│ ├── java_title.temp: In the java language, Title defines the template
│ ├── ts_const.temp: Constant definition template in TypeScript language
│ ├── ts_title.temp: In TypeScript language, Title defines the template
│ ├── ...: In other languages, constant definition templates and Title definition templates
</code></pre>

### Application environment configuration instructions

- system.yaml

<pre><code>. Apply system level configuration
├── languages: Supported programming language configuration
│ ├── name: programming language name
│ ├── ext: source code file extension
│ ├── ref: basic data read and write configuration file path (relative to the configuration root directory)
│ ├── temps_title: title export class defines the export template path (relative to the configuration root directory)
│ ├── temps_const: constant defines the export template path (relative to the configuration root directory)
├── database: supported database configurations
│ ├── default: The default database configuration, which must be one of the DatabaseList
│ ├── list: database configuration list
│ ├── name: database name
│ ref: The path where the database configuration file is located (relative to the configuration root directory)
│ temps_table: table structure sql generates template list
│ temps_data: table data sql generates template list
├── datafield_formats: list of supported basic data types
├── datafile_formats: list of supported data file formats
</code></pre>

- project.yaml

<pre><code>. Apply project-level configuration
├── source: Excel source directory
│ ├── value: Excel source directory path, support multiple, relative to the configuration root directory
│ ├── encoding: encoding format (if needed)
│ ├── ext_name: file extension, supports multiple
├── target: output settings
│ ├── root: output directory, with '':'' switch, or if the path contains '':'', it is regarded as an absolute path
│ ├── title:
│ ├── client: The client header defines the output directory
│ ├── server: The server header defines the output directory
│ ├── db: db header defines the output directory
│ ├── data:
│ ├── client: client data file output directory
│ ├── server: server data file output directory
│ ├── db: db data file output directory
│ ├── const:
│ ├── client: client constant table output directory
│ ├── server: server constant table output directory
│ ├── encoding: The encoding format of the output file
├── buff: buffer definition
│ ├── big_endian: The endian of the binary data file (true|false)
│ ├── token: buffer size of each field (unused)
│ ├── item: the buffer size of each piece of data (unused)
│ ├── sheet: buffer size of each sheet (unused)
</code></pre>

- excel.yaml

<pre><code>. Apply Excel Configuration
├── title&data:
│ ├── prefix: enable prefix
│ ├── outputs: export naming settings
│ ├── range_name: field domain name (client|server|db)
│ ├── title: The coordinates of the name of the Title definition file
│ ├── data: The name of the data file
│ ├── classes: header export class information
│ ├── name: field domain name (client|server|db)
│ ├── value: the coordinates of the exported class
│ ├── nick_row: field alias row number, used to find the specified column, when the value is 0, the column number is used as an alias
│ ├── name_row: The row number where the data name is located
│ ├── remark_row: the row number where the data comment is located
│ ├── field_range_row: output selection row number, content format: 'c,s,d', c means front end, s means back end, d means database, (01)
│ ├── field_format_row: field data format row number
│ ├── sql_field_format_row: database field type custom row number, 0 is not customizable
│ ├── field_names: Title definition file field name configuration
│ ├── name: language name
│ row: the row number where the language attribute is located
│ ├── file_keys: data file field name configuration
│ ├── name: data file format (bin, json, yaml, etc.)
│ row: the row number where the data field name is located
│ ├── data_start_row: data start row number
├── const:
│ ├── prefix: enable prefix
│ ├── outputs: export class name configuration
│ ├── name: field domain (client|server)
│ value: Coordinate (eg: "A1")
│ ├── classes: constant export class information
│ ├── name: field domain name (client|server|db)
│ ├── value: the coordinates of the exported class
│ ├── name_col: constant name column number
│ ├── value_col: constant value column number
│ ├── type_col: constant value type column number
│ ├── remark_col: remark column number
│ ├── data_start_row: data start row number
</code></pre>

#### Programming language configuration instructions  

- Specific language.yaml (eg [go.yaml](/res/lang/go.yaml))  
  The configuration files are located in the [res/lang] (/res/lang) directory.  

<pre><code>. Specific programming language configuration
├── lang_name: current language name
├── data_types: database configuration list
│ ├── name: field data type (filled in Excel sheet)
│ lang: The data type corresponding to the programming language
│ operates: operation methods for different data files
│ ├── file_name: data file type (json, bin, etc.)
│ get: read method character expression
│ set: write method character expression
</code></pre>

- Template file .temp (eg [go_titel.temp](/res/template/go_titel.temp), [go_const.temp](/res/template/go_const.temp))  
  Template files are located in the [res/template] (/res/template) directory.  
  Template files supported by golang syntax, help can be viewed [**https://golang.google.cn/pkg/text/template/**](https://golang.google.cn/pkg/text/template/ )  

#### Database configuration instructions  

- Database.yaml (eg [mysql.yaml](/res/db/mysql.yaml))  
  The configuration files are located in the [res/db] (/res/db) directory.  

<pre><code>. Specific data configuration
├── db_name: database name
├── scale_char: Char character scale
├── scale_varchar: archar character scale
├── types: database data type description list
│ ├── name: field name (normalized, such as string(5)=>string(*))
│ type: The field data type of the corresponding data
│ number: is it a numeric type
│ array: whether it is an array type
</code></pre>

- Template file.temp (eg [mysql_table.temp](/res/db/mysql_table.temp), [mysql_data.temp](/res/db/mysql_data.temp))  
   Template files are located in the [/res/db](/res/db) directory.  
   Template files supported by golang syntax, help can be viewed [**https://golang.google.cn/pkg/text/template/**](https://golang.google.cn/pkg/text/template/ )  
   
## Run

The program is only allowed to run via the command line  

Supported command line parameters include: -env, -mode, -range, -lang, -file, -merge, -source, -target  

- -env  
  Re-specify the running environment, which refers to the configuration root directory.  

- -mode  
  The running mode supports **header export**, **data export**, **constant table export**  
  - Supported values: title, data, const  
  title is the header export, data is the data export, const is the constant table export  
  - Supports multiple values, which can be separated by English commas ","  

- -range  
  The field range selected at runtime is valid for **header export** and **data export**  
  - Supported values: client, server, db  
  - Supports multiple values, which can be separated by English commas ","  
  
- -lang  
  The programming language selected at runtime is only valid for **header export**  
  - Supported values: go, as3, ts, java, c#  
  - Supports multiple values, which can be separated by English commas ","  
  
- -file  
  The data file type of the running output, valid for **data export**  
  - Supported values: bin, sql, json, yaml, toml, hcl, env, properties  
  - Supports multiple values, which can be separated by English commas ","  

- -merge  
  When exporting sql, it is not merged into one file, true means merged, false means not merged  
  - Supported values: true, false  

- -source  
  Specify the data source directory at runtime to override the value of source.value in the configuration file project.yaml  
  Optional, valid for **header export**, **data export**, **constant table export**  

- -target  
  Specify the data source directory at runtime to override the value of target.value in the configuration file project.yaml  
  Optional, valid for **header export**, **data export**, **constant table export**  

## Function Description

- Three basic export functions: [**Header Export**](#Header Export), [**Constant Table Export**](#Constant Table Export), [**Data Export**](#Data Export)  
- Special export function: [**Sql export**](#Sql export)  

### Header export  
Export the header information in the Excel file as a data structure or class of the corresponding language  

#### Header export process  

1. Traverse every matching Excel file in the source directory.  
  - The source directory is given by the soruce.value list in project.yaml.  
  - The source directory can be re-specified by the -source parameter.  
  - Match according to the soruce.ext_name list in project.yaml.  

2. Traverse the matched Sheets in the Excel file.  
  - Match according to the title&data.prefix attribute in excel.yaml.  

3. Select the corresponding field list according to the -range parameter.  
  - The -range parameter supports three types: client, server, and db. For details, please [view]().  

4. According to the -lang parameter, select the configuration and export template corresponding to the language.  
  - The -lang parameter supports go, as3, ts, java, c#, please [view]() for details.  

5. Field list => fields or properties of a data structure or class.  

6. The corresponding files are all generated into the target directory.  
  - The target root directory is given by the target.root list in project.yaml.  
  - The target.title in project.yaml in the header output directory is given as the relative path of target.root.  
  - The source directory can be re-specified by the -target parameter.  
  - According to the content of the -range parameter, the files are generated into the directories corresponding to target.title.client, target.title.server, and target.title.database in project.yaml.  

#### Header Template Description  

1. The injected data object is [\*TempTitelProxy](/src/core/context_title.go)
It can be obtained through template syntax such as `{{.}}`, `{{$proxy := .}}`, and the structure is defined as:  

```golang
  type TempTitleProxy struct {
    Sheet *excel.ExcelSheet // currently executed Sheet data object
    Excel *excel.ExcelProxy // current Excel proxy, may contain multiple Excel
    TitleCtx *TitleContext // currently executed header context data
    FileName string // header export class file name
    ClassName string // header export class name
    FieldIndex []int // currently selected field index
    Language string // currently selected programming language
  }
```

  - Excel:[\*excel.ExcelProxy](/src/core/excel/proxy.go)  
    The currently executing Excel data proxy object  
  - Sheet:[\*excel.ExcelSheet](/src/core/excel/sheet.go)  
    Currently executing Sheet data object  
  - TitleCtx:[\*TitleContext](/src/core/context_title.go)  
    Header context data of the current execution  
  - FileName: string  
    header export class file name  
  - ClassName: string  
    header export class name  
  - FieldIndex: []int  
    The currently selected field index  
  - Language: string  
    The current programming language of choice  

2. [Custom function](#Custom function)  

### Constant table export  

1. Traverse every matching Excel file in the source directory.  
  - The source directory is given by the soruce.value list in project.yaml.  
  - The source directory can be re-specified by the -source parameter.  
  - Match according to the soruce.ext_name list in project.yaml.  

2. Traverse the matched Sheets in the Excel file.  
  - Match according to the const.prefix property in excel.yaml.  
  - According to the name_col, value_col, type_col, and remark_col in excel.yaml, locate the name, value, type, and comment of the constant.  
  - According to the data_start_row in excel.yaml, start to describe the data positively, until it ends with a blank row.  

4. According to the -lang parameter, select the configuration and export template corresponding to the language.  
  - The -lang parameter supports go, as3, ts, java, c#, please [view]() for details.  

6. The corresponding files are all generated into the target directory.  
  - The target root directory is given by the target.root list in project.yaml.  
  - given by target.const in project.yaml in the constant table output directory, it is the relative path of target.root.  
  - The source directory can be re-specified by the -target parameter.  
  - According to the content of the -range parameter, the files are generated into the directories corresponding to target.const.client and target.const.server in project.yaml respectively.  

#### Data and functions injected into constant templates  

1. The injected data object is [\*TempConstProxy](/src/core/context_const.go)
It can be obtained through template syntax such as `{{.}}`, `{{$proxy := .}}`, and the structure is defined as:  

```golang
  type TempConstProxy struct {
    Sheet *excel.ExcelSheet // currently executed Sheet data object
    Excel *excel.ExcelProxy // current Excel proxy, may contain multiple Excel
    ConstCtx *ConstContext // currently executed context data
    FileName string // export file name
    ClassName string // export constant class name
    Language string // export the corresponding programming language
    StartRow int // Data start row number
    EndRow int // data end row number
  }
```

  - Excel:[\*excel.ExcelProxy](/src/core/excel/proxy.go)  
    The currently executing Excel data proxy object  
  - Sheet:[\*excel.ExcelSheet](/src/core/excel/sheet.go)  
    Currently executing Sheet data object  
  - ConstCtx:[\*ConstContext](/src/core/context_const.go)  
    context data for the current execution  
  - FileName: string  
    export file name  
  - ClassName: string  
    export constant class name  
  - Language: string  
    Export the corresponding programming language  
  - StartRow:int  
    Data start line number  
  - EndRow:int  
    Data end line number  

2. [Custom function](#Custom function)  

### Data output  
- Supported data export formats: bin (binary), json, sql.  
- When yaml, toml, hcl, env, properties data is exported, the field name will be **forced** to be lowercase, the original intention requires **case dependent**, not open by default  
- To enable data export such as yaml, please modify the system.yaml file and add it to the "datafiel_formats" list.  

### Sql export  

- **Sql export depends on the settings of header export and data export. **  

- When the following three conditions **consist at the same time**, perform sql export.  
  1. -ragne contains the db item  
  2. -file contains sql items  
  3. -mode contains at least one of title or data.  

- Export process:  
  1. Traversing Excel files and Sheets is consistent with [**Table Header Export**](#Table Header Export) and [**Data Export**](#Data Export).  
  2. When the -merge parameter is set to true, only one sql file (all_merge.sql) will be produced  
  3. When the -merge parameter is turned off or set to false, "filename.talbe.sql" and "filename.data.sql" will be generated. The table.sql file is the table structure update script, and data.sql is the data update script.  

#### Data and functions injected into constant templates  

1. The injected data object is [\*TempSqlProxy](/src/core/context_sql.go)
It can be obtained through template syntax such as `{{.}}`, `{{$proxy := .}}`, and the structure is defined as:  

```golang
  type TempSqlProxy struct {
    Sheet *excel.ExcelSheet // currently executed Sheet data object
    Excel *excel.ExcelProxy // currently executing Excel data proxy object
    SqlCtx *SqlContext // currently executed Sql context
    TableName string // database table name
    FieldIndex []int // field selection index
    StartRow int // start row number
    EndRow int // end row number
  }
```

  - Excel:[\*excel.ExcelProxy](/src/core/excel/proxy.go)  
    The currently executing Excel data proxy object  
  - Sheet:[\*excel.ExcelSheet](/src/core/excel/sheet.go)  
    Currently executing Sheet data object  
  - SqlCtx:[\*SqlContext](/src/core/context_sql.go)  
    The currently executing Sql context  
  - TableName: string  
    database table name  
  - FieldIndex: string  
    field selection index  
  - StartRow: string  
    start line number  
  - EndRow:int  
    end line number  

2. [Custom function](#Custom function)  

### Template customization  
The template file format is a go language template, and the document description address is as follows:  
[https://golang.google.cn/pkg/text/template/](https://golang.google.cn/pkg/text/template/)  

#### Custom Functions  
Custom functions are valid for all templates  
**Note**: The return value of the custom function must be 1 or 2, which is an official requirement.  
When there are 2 return values, the type of the second return value must be error.  

- [ToLowerCamelCase](/src/core/tools/naming.go)  
  Convert the string content to **small camelcase** format  

- [ToUpperCamelCase](/src/core/tools/naming.go)  
  Convert string content to **big camelcase** format  

- [Add](/src/core/tools/math.go)  
  addition  

- [Sub](/src/core/tools/math.go)  
  subtraction  

- [NowTime](/src/core/tools/time.go)  
  get current time  

- [NowTimeStr](/src/core/tools/time.go)  
  Get the current time default format string  

- [NowTimeFormat](/src/core/tools/time.go)  
  get current time  
  2006-01-02 15**:**04**:**05 PM Mon Jan  
  2006-01-\_2 15**:**04**:**05 PM Mon Jan  

- [NowYear](/src/core/tools/time.go)  
  current time year  

- [NowMonth](/src/core/tools/time.go)  
  current time month  
  January: 1  

- [NowDay](/src/core/tools/time.go)  
  current date  

- [NowWeekday](/src/core/tools/time.go)  
  current time day of the week  
  Sunday: 0  

- [NowHour](/src/core/tools/time.go)  
  current time hour  

- [NowMinute](/src/core/tools/time.go)  
  current time in minutes  

- [NowSecond](/src/core/tools/time.go)  
  current time in seconds  

- [NowUnix](/src/core/tools/time.go)  
  current timestamp (s)  

- [NowUnixNano](/src/core/tools/time.go)  
  Current timestamp (ns)  

## Dependencies  
- infra-go (library dependency) [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)  
- excelize (library dependency) [https://github.com/360EntSecGroup-Skylar/excelize](https://github.com/360EntSecGroup-Skylar/excelize)  
- goxc (compilation dependencies) [https://github.com/laher/goxc](https://github.com/laher/goxc)  

## Contact Me 
xuzhuoxi  
<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>  

## Open Source License  
ExcelExporter source code is open source based on [MIT license](/LICENSE).  