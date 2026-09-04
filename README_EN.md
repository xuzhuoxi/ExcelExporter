# ExcelExporter

Export Excel sheets into multi-language source files, data files, and SQL scripts through templates. Supports **multiple data formats**, **extensible programming languages**, and **multiple operating systems**.

[中文](README.md) | English

## 1. Compatibility

- Minimum Go version: **1.24.0** (see `go.mod`)
- Default Excel extension: `.xlsx`
- Release binaries: Linux / Windows / macOS, amd64 and arm64
- Local cross-build scripts also cover freebsd and openbsd

## 2. How to get it

### 2.1. Download from Releases

[Releases](https://github.com/xuzhuoxi/ExcelExporter/releases)

+ Download the executable package for your platform (includes `LICENSE`, `README.md`, and `README_EN.md`)
+ Download the environment package `env_<tag>.zip` (extracts to `evn_<tag>/`; `target` is an empty directory)

### 2.2. Build from source

```shell
git clone https://github.com/xuzhuoxi/ExcelExporter.git
cd ExcelExporter
```

The project uses Go Modules. The main package is `./src`.

**Build locally:**

```shell
go build -o ExcelExporter ./src
```

**Cross-compile (recommended):**

+ Windows: [build/build.bat](/build/build.bat)
+ Linux / macOS: [build/build.sh](/build/build.sh)

The scripts run unit tests first, then write binaries to `build/release/`.

> [goxc_build.bat](/build/goxc_build.bat) and [goxc_build.sh](/build/goxc_build.sh) are legacy and are no longer the recommended path.

## 3. Getting started

### 3.1. Environment

1. Unzip `env_<tag>.zip` from the Release, or use [`res/`](/res) in this repository as the environment directory.
2. Remove sample Excel files under `source` if you do not need them.
3. Place the executable next to `system.yaml`, `project.yaml`, and `excel.yaml`.
4. Run commands from that directory. Logs are written to `ExcelExporter.log` in the working directory.

If the environment files sit beside the executable, `-env` is optional. For development you can point `-env` at `res/`.

### 3.2. Command line

Full example:

```shell
ExcelExporter -env=<env-dir> -mode=title,data,const,proto -range=client,server,db -lang=as3,c#,go,java,ts -file=json,bin,sql -merge=false -source=<excel-path> -target=<output-root>
```

`-mode` and `-range` are required. Multiple values are comma-separated. Flag values are case-insensitive.

**Examples:**

1. Export client titles and JSON data in C#:

```shell
ExcelExporter -mode=title,data -range=client -lang=c# -file=json
```

2. Export DB table and data SQL, merged into one file:

```shell
ExcelExporter -mode=title,data -range=db -file=sql -merge=true
```

`-lang` can be omitted for this SQL-only case.

3. Export client constants and protocols in Go:

```shell
ExcelExporter -mode=const,proto -range=client -lang=go
```

### 3.3. Flags (`*` = required)

- `-env`
  + Override the environment directory (relative paths are resolved from the executable directory)
  + Optional: defaults to the executable directory

- **`-mode` \***
  + Export modes, comma-separated
  + Values: `title`, `data`, `const`, `proto`
  + Extra flags:
    + `title`: `-range`, `-lang`
    + `data`: `-range`, `-file`
    + `const`: `-range`, `-lang` (`client` and `server` only)
    + `proto`: `-range`, `-lang` (`client` and `server` only)

- **`-range` \***
  + Field range / output subdirectory / protocol filter
  + Values: `client`, `server`, `db`
  + `title` / `data`: filter fields from the range row and choose the output subdirectory
  + `const`: output subdirectory; `db` is ignored
  + `proto`: matched against the protocol sheet range; `db` is ignored

- `-lang`
  + Languages, comma-separated
  + Complete templates: `as3`, `c#`, `go`, `java`, `ts`
  + `c++` is registered in `system.yaml` but C++ templates are not complete
  + Used by `title`, `const`, `proto`

- `-file`
  + Data formats, comma-separated
  + Enabled by default: `json`, `bin`, `sql`
  + Implemented but disabled: `yaml`, `yml`, `toml`, `hcl`, `env`, `properties` (keys become lowercase)
  + Used by `data`; `sql` plus `-range=db` enables SQL export with `title` / `data`

- `-merge`
  + When `-file` includes `sql`, merge all SQL into one file. Default `false`
  + Values: `true`, `false`

- `-source`
  + Override `source.value` in `project.yaml` (comma-separated paths)

- `-target`
  + Override `target.root` in `project.yaml`

## 4. Configuration

The tool loads three files from the environment directory:

| File | Role |
| --- | --- |
| [system.yaml](./res/system.yaml) | Languages, databases, field types, export formats |
| [project.yaml](./res/project.yaml) | Source and output paths |
| [excel.yaml](./res/excel.yaml) | Excel layout conventions |

### 4.1. Environment layout

```text
env root
├── db/                      database configs and SQL templates
│   ├── mysql.yaml
│   ├── mysql_table.temp
│   └── mysql_data.temp
├── lang/                    per-language types and get/set names
│   ├── as3.yaml / c#.yaml / c++.yaml / go.yaml / java.yaml / ts.yaml
├── proxy/                   optional runtime reader helpers
│   ├── as / go / java / ts
├── template/                Go text/template files
│   ├── <lang>_title.temp
│   ├── <lang>_const.temp
│   └── <lang>_proto.temp
├── source/                  default Excel input
├── target/                  default output root
├── excel.yaml
├── project.yaml
└── system.yaml
```

### 4.2. [System configuration](./res/system.yaml)

#### 4.2.1 Scope

1. Languages: extension, type-mapping file, title / const / proto templates
2. Databases: type mapping, table template, data template
3. Field types: `field_datatypes`
4. Export formats: `export_files`
5. Pointer token: `pointer_code` (used by proto custom types, default `*`)

#### 4.2.2 When to edit

+ Add or change a language
+ Add or change a database
+ Extend field types
+ Enable extra data formats via `export_files`

Default field types:

`bool`, `int8`/`16`/`32`/`64`, `uint8`/`16`/`32`/`64`, `float32`/`float64`, `string`, `string(*)`, `json`, and their `[]` array forms.

In `string(*)`, `*` is the character count, recommended range `[1, 1024]`. Floats should stay within about 6 decimal places. Negative floats written as binary may jitter when read back in some languages (for example AS3).

### 4.3. [Project configuration](./res/project.yaml)

#### 4.3.1 Scope

1. `source`: Excel paths and extensions (default `xlsx`)
2. `target`:
   + `root`: output root
   + `title` / `data` / `const` / `proto`: client / server / db subdirectories relative to `root`
   + `sql.dir`: SQL output subdirectory
3. `buff` and `encoding` exist in the file but are **not used** (binary export is always big-endian)

Relative paths are resolved against the environment directory. Existing absolute paths are kept.

#### 4.3.2 When to edit

+ Multiple Excel directories in one run: extend `source.value`
+ Custom source without `-source`: change `source.value`
+ Custom output layout: change `target`

### 4.4. [Excel configuration](./res/excel.yaml)

#### 4.4.1 Scope

1. `ignore`: skip files by **name prefix**; defaults `_` and `~$`
2. `title&data` (shared by title, data, and SQL)
   + `prefix`: Sheet name prefix, default `Data_`
   + `outputs`: cell axes for file name, class name, and namespace per range
   + `sql`: table name, script file prefix, primary-key axes (composite keys use column letters, e.g. `A,B`)
   + `control_row`: non-empty cells define field width
   + `nick_row`: alias row; `0` means use column letters
   + `name_row` / `remark_row`
   + `range_row`: must be `c,s,d` with each value `0` or `1` (client / server / db)
   + `data_type_row`
   + `sql_data_type_row`: custom SQL type row; `0` disables it
   + `ext_name_rows`: per-language field names
   + `file_key_rows`: per-format field keys
   + `data_start_axis`: data start cell, default `A8`
3. `const`
   + `prefix`: default `Const_`
   + `outputs` for client / server
   + `name_col` / `value_col` / `type_col` / `remark_col`
   + `data_start_row`; empty names are skipped
4. `proto`
   + `prefix`: default `Proto_`
   + `id_datatype` / `range_name` / `namespace` / `export`
   + `id_col` / `file_col` / `name_col` / `field_start_col`
   + `data_start_row`: protocol rows; fields are `name:type`
   + `remark_offset`
   + `blank_break`: stop at blank rows

A Sheet is processed only if its name starts with the matching prefix. An empty export file name skips that Sheet.

#### 4.4.2 When to edit

+ Ignore extra Excel files (`ignore`)
+ Move header cells
+ Change sheet prefixes or data start position

### 4.5. Language configuration

Go is used as the example.

#### 4.5.1 Language file

+ Default directory: [res/lang](./res/lang), linked by `languages[].ref` in `system.yaml`
+ Go file: [res/lang/go.yaml](./res/lang/go.yaml)
+ Fields:
  + `lang_name`
  + `custom`: how custom / pointer / array types are written (`T`, `TArray`, `TPointer`, `TPointerArray`)
  + `data_types`:
    + `name`: Excel field type (must cover `field_datatypes`)
    + `lang`: language type name
    + `operates`: `get` / `set` names per data file

#### 4.5.2 Templates

+ Default directory: [res/template](./res/template)
+ Linked by `temps_title`, `temps_const`, `temps_proto`; multiple files are allowed, **the first is the main template**
+ Syntax: Go `text/template` ([docs](https://pkg.go.dev/text/template))

### 4.6. Database configuration

MySQL is used as the example.

#### 4.6.1 Config

+ Default directory: [res/db](./res/db), linked by `databases.list[].ref`
+ [mysql.yaml](./res/db/mysql.yaml):
  + `db_name`
  + `scale_char` / `scale_varchar`
  + `types`: Excel type to SQL type

`databases.default` is `mysql`; SQL export uses that default.

#### 4.6.2 Templates

+ [mysql_table.temp](./res/db/mysql_table.temp): `CREATE TABLE`
+ [mysql_data.temp](./res/db/mysql_data.temp): `INSERT`

## 5. Features

Four export modes:

+ [Title export](#51-title-export): `-mode=title`
+ [Data export](#52-data-export): `-mode=data`
+ [Const export](#53-const-export): `-mode=const`
+ [Proto export](#54-proto-export): `-mode=proto`

Extra:

+ [SQL export](#55-sql-export): `-range` includes `db`, `-file` includes `sql`, and `-mode` includes `title` and/or `data`

Each Excel file is loaded and processed immediately. Merged SQL is written at the end.

### 5.1 Title export

Export field definitions as language structs or classes.

#### 5.1.1 Flow

1. Walk source paths; skip ignored prefixes, unmatched extensions, and empty files.
2. Keep Sheets whose names start with `title&data.prefix`.
3. Use the control row for width and the range row for `-range`.
4. Load the language config and title template from `-lang`.
5. Read file / class / namespace cells and write under `target.title.<range>`.

#### 5.1.2 Template data

Root object: [`*TempTitleProxy`](/src/core/context_title.go). Use `{{.}}` or `{{$proxy := .}}`.

**TempTitleProxy**

| Name | Type | Description |
| --- | --- | --- |
| `Excel` | `*excel.ExcelProxy` | Current workbook proxy |
| `Sheet` | `*excel.ExcelSheet` | Current sheet |
| `TitleCtx` | `*TitleContext` | Context (`RangeName`, `Language`, ...) |
| `FileName` | `string` | Output file name |
| `ClassName` | `string` | Class name |
| `Namespace` | `string` | Namespace / package |
| `FieldIndex` | `[]int` | Selected column indexes |
| `Language()` | `string` | Current language |
| `LanguageDefine()` | `*setting.ProgramLanguage` | Language definition |
| `ValueAtAxis(axis)` | `string` | Cell text |
| `FieldLen()` | `int` | Field count |
| `GetFields()` | `[]TitleFieldItem` | All fields |
| `GetField(index)` | `TitleFieldItem` | One field |
| `GetFieldName(index)` | `string` | Language field name |
| `GetFieldFileKey(index, fileType)` | `string` | Data-file key |

**TitleFieldItem**

| Name | Type | Description |
| --- | --- | --- |
| `Index` | `int` | Column index |
| `TitleName` | `string` | Excel field name |
| `TitleRemark` | `string` | Remark (newlines converted to `<br/>`) |
| `FieldLangName` | `string` | Language field name |
| `OriginalType` | `string` | Raw Excel type |
| `FormattedType` | `string` | Normalized type (`string(8)` → `string`) |
| `LangType` | `string` | Language type name |
| `LangTypeDefine` | `setting.LangDataType` | Type definition (`GetGetOperate` / `GetSetOperate`) |
| `GetFileKey(fileType)` | `string` | Key used in a data file |

### 5.2 Data export

- Default formats: `json`, `bin`. `sql` uses the SQL pipeline, not the data builder.
- Output: `target.data.<range>`, named from the sheet's data-file cell plus the format extension.
- Rows start at `data_start_axis` and stop on an empty row or empty first cell.
- Keys come from `file_key_rows` for the selected format.
- Array cells: `[a,b,c]`; empty arrays may be `[]`.

**JSON shape:**

```json
{
  "count": 2,
  "data": [
    { "<fileKey>": "<value>" }
  ]
}
```

`[]uint8` is exported as `[]uint16` so JSON does not encode it as a Base64 string.

**Binary:** always big-endian. A 4-byte `uint32` row count is followed by row-major field values. `buff.big_endian` in `project.yaml` does **not** change this.

To enable yaml-like formats, add them to `export_files`. Those writers go through Viper, so **keys become lowercase**.

### 5.3 Const export

#### 5.3.1 Flow

1. Sheets whose names start with `const.prefix`.
2. Read from `data_start_row` to the last row; skip empty names.
3. Render the const template and write to `target.const.client` or `target.const.server`.

#### 5.3.2 Template data

Root object: [`*TempConstProxy`](/src/core/context_const.go).

**TempConstProxy**

| Name | Type | Description |
| --- | --- | --- |
| `Excel` / `Sheet` | proxy / sheet | Current table |
| `ConstCtx` | `*ConstContext` | Context |
| `FileName` / `ClassName` / `Namespace` | `string` | Output info |
| `StartRow` / `EndRow` | `int` | Row range (end exclusive) |
| `Language()` | `string` | Current language |
| `ValueAtAxis(axis)` | `string` | Cell text |
| `GetItems()` | `[]ConstItem` | All constants |
| `GetItem(row)` | `ConstItem` | One Excel row |

**ConstItem**

| Name | Type | Description |
| --- | --- | --- |
| `Name` | `string` | Constant name |
| `Value` | `string` | Value; strings are already quoted |
| `Type` | `string` | Language type name |
| `Remark` | `string` | Comment |

### 5.4 Proto export

Field cells use `name:type`. `type` may be a primitive, a fixed-length array, a custom protocol name, or a pointer (`pointer_code`, default `*`). `json` and `[]json` are not supported. Custom types must be protocol names **already defined in the same sheet**.

#### 5.4.1 Flow

1. Sheets whose names start with `proto.prefix`.
2. Read Id type, range, namespace, and extra subdirectory.
3. Parse protocol rows from `data_start_row`; a row without `file` or `name` is not a protocol.
4. If the sheet range matches `-range`, render the proto template into `target.proto.<range>`, plus the optional `export` subdirectory.

#### 5.4.2 Template data

Root object: [`*TempProtoProxy`](/src/core/context_proto.go).

**TempProtoProxy**

| Name | Type | Description |
| --- | --- | --- |
| `ProtoItem` | `ProtoItem` | Current protocol |
| `SheetProxy` | `*ProtoSheetProxy` | Sheet proxy (`Excel`, `Sheet`, `ProtoCtx`, `Title`) |
| `ValueAtAxis(axis)` | `string` | Cell text |
| `Namespace()` | `string` | Namespace |
| `ProtoId()` | `string` | Protocol Id; strings are quoted |
| `ProtoIdDataType()` | `string` | Language type of the Id |
| `ClassName()` | `string` | Class name |
| `ClassRemark()` | `string` | Class comment |
| `GetFields()` | `[]ProtoFieldItem` | Fields |

**ProtoFieldItem**

| Name | Type | Description |
| --- | --- | --- |
| `Remark` | `string` | Field comment |
| `Name` | `string` | Field name |
| `Lang` | `string` | Current language |
| `OriginalType` / `FormattedType` | `string` | Raw / normalized type |
| `LangType` | `string` | Language type name |
| `LangTypeDefine` | `setting.LangDataType` | Primitive type definition |
| `IsCustomType` | `bool` | Custom protocol type |
| `IsPointer` / `IsArray` | `bool` | Pointer / array |
| `ArraySize` | `int` | Fixed length; `-1` if not a fixed array |
| `TempLangType()` | `string` | Final type string for templates |

### 5.5 SQL export

#### 5.5.1 Flow

SQL export runs only when all of the following are true:

1. `-range` includes `db`
2. `-file` includes `sql`
3. `-mode` includes `title` and/or `data`

Notes:

+ Sheet selection matches title / data export (`Data_` prefix)
+ `title` writes table scripts; `data` writes insert scripts
+ `-merge=true`: one file `all_merge.sql`
+ `-merge=false`: `<fileName>.table.sql` and `<fileName>.data.sql`
+ Output directory: `target.sql.dir` under `target.root`

When merging, `NeedTruncateData()` lets the data template skip extra `TRUNCATE` statements.

#### 5.5.2 Template data

Root object: [`*TempSqlProxy`](/src/core/context_sql.go).

**TempSqlProxy**

| Name | Type | Description |
| --- | --- | --- |
| `Excel` / `Sheet` | proxy / sheet | Current table |
| `SqlCtx` | `*SqlContext` | Context |
| `TableName` | `string` | Table name |
| `FieldIndex` | `[]int` | Selected columns |
| `StartRow` / `EndRow` | `int` | Data row range |
| `StartColIndex` | `int` | Start column index |
| `MergeOn()` | `bool` | Merge output |
| `NeedTruncateData()` | `bool` | Whether data SQL should TRUNCATE |
| `FieldLen()` / `ItemLen()` | `int` | Field count / row count |
| `ValueAtAxis(axis)` | `string` | Cell text |
| `GetFieldItems()` | `[]SqlFieldItem` | Field definitions |
| `GetPrimaryKeys()` | `[]SqlFieldItem` | Primary key fields |
| `PrimaryKeyLen()` | `int` | Primary key count |
| `GetItems()` | `[]SqlItem` | Data rows |
| `GetItem(row)` | `SqlItem` | One row |

**SqlFieldItem**

| Name | Type | Description |
| --- | --- | --- |
| `FieldName` | `string` | SQL column name |
| `FieldType` | `string` | Raw Excel type |
| `CustomFieldType` | `string` | Custom SQL type if enabled |
| `SqlFieldType()` | `string` | Final SQL type for the template |
| `MaxByteSize` / `MaxRuneSize` | `int` | Dynamic length stats |

**SqlItem**

| Name | Type | Description |
| --- | --- | --- |
| `FieldLen()` | `int` | Field count |
| `GetValues()` | `[]string` | SQL literals |
| `GetValue(i)` / `GetSqlValue(i)` | `string` | Raw / SQL literal |

Strings escape `'` as `''` and are quoted. Numeric values are written as-is.

### 5.6 Template helpers

Templates use Go `text/template`: [https://pkg.go.dev/text/template](https://pkg.go.dev/text/template)

Custom functions are available in every template. A function must return one value, or two values with `error` as the second.

| Function | Source | Role |
| --- | --- | --- |
| `ToLowerCamelCase` | [naming.go](/src/core/tools/naming.go) | lowerCamelCase |
| `ToUpperCamelCase` | [naming.go](/src/core/tools/naming.go) | UpperCamelCase |
| `Add` / `Sub` | [math.go](/src/core/tools/math.go) | integer add / sub |
| `NowTime` | [time.go](/src/core/tools/time.go) | `time.Time` |
| `NowTimeStr` | same | `2006-01-02 15:04:05` |
| `NowTimeFormat` | same | Go time layout |
| `NowYear` / `NowMonth` / `NowDay` / `NowWeekday` | same | year, month (1–12), day, weekday (Sunday = 0) |
| `NowHour` / `NowMinute` / `NowSecond` | same | clock fields |
| `NowUnix` / `NowUnixNano` | same | unix timestamps |

## 6. Known limitations

+ `yaml` / `toml` / `hcl` / `env` / `properties` force lowercase keys and are disabled by default.
+ `encoding` and `buff` in `project.yaml` are unused; binary export is always big-endian.
+ C++ language mapping exists, but title / const / proto templates are not provided.
+ Proto sheets do not support `json` or `[]json` fields.

## 7. Dependencies

- [infra-go](https://github.com/xuzhuoxi/infra-go) `v1.4.1`
- Spreadsheet I/O: vendored [excelize v2](/src/excelize.v2) (from [excelize](https://github.com/qax-os/excelize))
- [yaml.v2](https://gopkg.in/yaml.v2), [sjson](https://github.com/tidwall/sjson), [viper](https://github.com/spf13/viper)

## 8. Contact

xuzhuoxi  
<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## 9. License

ExcelExporter is released under the [MIT License](/LICENSE). The vendored excelize copy keeps its own BSD license.
