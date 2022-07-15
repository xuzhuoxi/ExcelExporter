# ExcelExporter

A tool for exporting Excel data.

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

-