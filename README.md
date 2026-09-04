# ExcelExporter

根据模板把 Excel 表导出为多语言代码、数据文件和 SQL 脚本。支持 **多种数据格式**、**可扩展的编程语言** 和 **多操作系统**。

中文 | [English](/README_EN.md)

## 1. 兼容性

- 最低 Go 版本：**1.24.0**（见 `go.mod`）
- 默认 Excel 扩展名：`.xlsx`
- Release 预编译平台：Linux / Windows / macOS，架构含 amd64、arm64
- 本地交叉编译脚本额外覆盖 freebsd、openbsd

## 2. 如何获取

可通过以下两种方式获取 ExcelExporter：

### 2.1. 从 Release 下载

[下载页面](https://github.com/xuzhuoxi/ExcelExporter/releases)

+ 下载对应平台的可执行文件包（内含 `LICENSE`、`README.md`、`README_EN.md`）
+ 下载环境配置包 `env_<tag>.zip`（解压后根目录为 `evn_<tag>/`，其中 `target` 为空目录）

### 2.2. 下载源码编译

```shell
git clone https://github.com/xuzhuoxi/ExcelExporter.git
cd ExcelExporter
```

当前工程使用 Go Modules，入口在 `./src`。

**本机直接编译：**

```shell
go build -o ExcelExporter ./src
```

**交叉编译（推荐）：**

+ Windows 执行 [build/build.bat](/build/build.bat)
+ Linux / macOS 执行 [build/build.sh](/build/build.sh)

脚本会先跑单元测试，再产出多平台可执行文件到 `build/release/`。

> 仓库中仍保留 [goxc_build.bat](/build/goxc_build.bat) / [goxc_build.sh](/build/goxc_build.sh)，属于旧构建方式，不再作为推荐路径。

## 3. 开始

### 3.1. 环境准备

1. 解压 Release 中的 `env_<tag>.zip`，或直接使用仓库中的 [`res/`](/res) 作为环境目录。
2. 清理 `source` 中的测试 Excel（如不需要）。
3. 把可执行文件放到环境目录中（与 `system.yaml`、`project.yaml`、`excel.yaml` 同级）。
4. 在该目录执行命令。日志默认写到运行目录下的 `ExcelExporter.log`。

环境目录与可执行文件同级时，不必传 `-env`。开发调试也可把 `-env` 指到 `res/`。

### 3.2. 执行命令

完整参数示意：

```shell
ExcelExporter -env=<环境目录> -mode=title,data,const,proto -range=client,server,db -lang=as3,c#,go,java,ts -file=json,bin,sql -merge=false -source=<Excel源路径> -target=<导出根目录>
```

`-mode`、`-range` 为必填。多个值用英文逗号 `,` 分隔，参数大小写不敏感。

**示例：**

1. 导出 client 范围的表头与 JSON 数据，语言为 C#：

```shell
ExcelExporter -mode=title,data -range=client -lang=c# -file=json
```

2. 导出 db 范围的建表脚本与数据脚本，并合并为一个 SQL 文件：

```shell
ExcelExporter -mode=title,data -range=db -file=sql -merge=true
```

该场景只出 SQL 时可以不传 `-lang`。

3. 导出 client 范围的常量表与协议表，语言为 Go：

```shell
ExcelExporter -mode=const,proto -range=client -lang=go
```

### 3.3. 参数说明（带 \* 为必要参数）

- `-env`
  + 作用：指定环境配置目录（相对路径以可执行文件所在目录为基准）
  + 可选：不指定则使用可执行文件所在目录

- **`-mode` \***
  + 作用：运行模式，多个用 `,` 分隔
  + 支持值：`title`、`data`、`const`、`proto`
  + 各模式后续参数：
    + `title`：需要 `-range`、`-lang`
    + `data`：需要 `-range`、`-file`
    + `const`：需要 `-range`、`-lang`（仅 `client`、`server` 生效）
    + `proto`：需要 `-range`、`-lang`（仅 `client`、`server` 生效）

- **`-range` \***
  + 作用：字段范围 / 导出目录 / 协议过滤
  + 支持值：`client`、`server`、`db`
  + 适用：
    + `title`、`data`：按 Sheet 中范围行筛选字段，并决定输出子目录
    + `const`：决定输出子目录；`db` 会被忽略
    + `proto`：与协议表上的范围配置匹配；`db` 会被忽略

- `-lang`
  + 作用：编程语言，多个用 `,` 分隔
  + 已提供完整模板：`as3`、`c#`、`go`、`java`、`ts`
  + `c++` 已在 `system.yaml` 中登记语言配置，但 C++ 模板尚未齐备，不能作为可用导出语言
  + 适用模式：`title`、`const`、`proto`

- `-file`
  + 作用：数据文件格式，多个用 `,` 分隔
  + 默认启用：`json`、`bin`、`sql`
  + 代码已实现但默认未开放：`yaml`、`yml`、`toml`、`hcl`、`env`、`properties`（字段名会被转为小写）
  + 适用模式：`data`；当包含 `sql` 且 `-range` 含 `db` 时，配合 `title` / `data` 导出 SQL

- `-merge`
  + 作用：`-file` 包含 `sql` 时，是否把全部 SQL 合并到一个文件。默认 `false`
  + 支持值：`true`、`false`

- `-source`
  + 作用：覆盖 `project.yaml` 中 `source.value`（多个路径用 `,` 分隔）
  + 适用全部模式

- `-target`
  + 作用：覆盖 `project.yaml` 中 `target.root`
  + 适用全部模式

## 4. 配置文件

工具启动后从环境目录加载三份主配置：

| 文件 | 职责 |
| --- | --- |
| [system.yaml](./res/system.yaml) | 语言、数据库、字段类型、导出格式 |
| [project.yaml](./res/project.yaml) | 源目录、输出目录 |
| [excel.yaml](./res/excel.yaml) | Excel 表结构约定 |

### 4.1. 环境目录及文件说明

```text
环境根目录
├── db/                      数据库配置与 SQL 模板
│   ├── mysql.yaml
│   ├── mysql_table.temp
│   └── mysql_data.temp
├── lang/                    各语言基础类型与读写方法名
│   ├── as3.yaml / c#.yaml / c++.yaml / go.yaml / java.yaml / ts.yaml
├── proxy/                   运行时读取代理（非导出必需）
│   ├── as / go / java / ts
├── template/                Go text/template 模板
│   ├── <lang>_title.temp
│   ├── <lang>_const.temp
│   └── <lang>_proto.temp
├── source/                  默认 Excel 源目录
├── target/                  默认导出根目录
├── excel.yaml
├── project.yaml
└── system.yaml
```

### 4.2. [System 配置](./res/system.yaml)

#### 4.2.1 职能范围

1. 编程语言：扩展名、类型映射文件、title / const / proto 模板路径
2. 数据库：类型映射、建表模板、数据模板
3. 基础字段类型 `field_datatypes`
4. 导出数据格式 `export_files`
5. 指针记号 `pointer_code`（协议表自定义类型使用，默认 `*`）

#### 4.2.2 定制场合

出现以下需求时修改 `system.yaml`：

+ 增加或调整编程语言
+ 增加或调整数据库
+ 扩展基础字段类型
+ 开放额外数据格式（需同步改 `export_files`）

当前默认字段类型：

`bool`、`int8`/`16`/`32`/`64`、`uint8`/`16`/`32`/`64`、`float32`/`float64`、`string`、`string(*)`、`json` 及其 `[]` 数组形式。

`string(*)` 中 `*` 表示字符数，范围建议 `[1, 1024]`。浮点建议最多 6 位小数；负数写入二进制后再读回时，部分语言（如 AS3）可能出现精度抖动。

### 4.3. [Project 配置](./res/project.yaml)

#### 4.3.1 职能范围

1. `source`：Excel 源路径列表、扩展名（默认 `xlsx`）
2. `target`：
   + `root`：导出根目录
   + `title` / `data` / `const` / `proto`：相对 `root` 的 client / server / db 子目录
   + `sql.dir`：SQL 输出子目录
3. `buff`、`encoding`：配置项存在，**当前实现未使用**（二进制字节序固定为大端）

相对路径会相对环境目录解析；已存在的绝对路径保持不变。

#### 4.3.2 定制场合

+ 多个 Excel 目录希望一次处理：扩展 `source.value`
+ 不传 `-source` 时自定义源目录：改 `source.value`
+ 自定义输出布局：改 `target` 下各子目录

### 4.4. [Excel 配置](./res/excel.yaml)

#### 4.4.1 职能范围

1. `ignore`：按**文件名前缀**忽略，默认 `_`、`~$`
2. `title&data`（表头 / 数据 / SQL 共用）
   + `prefix`：参与处理的 Sheet 名前缀，默认 `Data_`
   + `outputs`：client / server / db 在 Sheet 中的导出文件名、类名、命名空间坐标
   + `sql`：表名、脚本文件名前缀、主键坐标（复合主键用 `,` 分隔列号，如 `A,B`）
   + `control_row`：控制行，本行非空单元格数量决定字段宽度
   + `nick_row`：列别名行；`0` 表示用列号作别名
   + `name_row` / `remark_row`：字段名称、备注行号
   + `range_row`：范围行，单元格格式必须为 `c,s,d`（三个 `0` 或 `1`，分别对应 client / server / db）
   + `data_type_row`：字段类型行
   + `sql_data_type_row`：SQL 定制类型行；`0` 表示不定制
   + `ext_name_rows`：各语言专用字段名行
   + `file_key_rows`：各数据文件专用字段名行
   + `data_start_axis`：数据起始单元格，默认 `A8`
3. `const`
   + `prefix`：默认 `Const_`
   + `outputs`：client / server 的文件名、类名、命名空间、额外子目录坐标
   + `name_col` / `value_col` / `type_col` / `remark_col`
   + `data_start_row`：常量数据起始行；空名称行跳过
4. `proto`
   + `prefix`：默认 `Proto_`
   + `id_datatype` / `range_name` / `namespace` / `export`：表头单元格
   + `id_col` / `file_col` / `name_col` / `field_start_col`
   + `data_start_row`：协议行起始；字段格式为 `name:type`
   + `remark_offset`：备注相对协议行的行偏移
   + `blank_break`：是否在空行处停止解析

Sheet 名必须以对应 `prefix` 开头才会被该模式处理。单元格中导出文件名为空时，该 Sheet 被跳过。

#### 4.4.2 定制场合

+ 过滤部分 Excel 文件（改 `ignore`）
+ 调整表头信息在 Sheet 中的位置
+ 修改 Sheet 启用前缀或数据起始位置

### 4.5. 编程语言配置

以 Go 为例。

#### 4.5.1 语言配置

+ 默认位于 [res/lang](./res/lang)，由 `system.yaml` 中 `languages[].ref` 关联
+ Go 配置：[res/lang/go.yaml](./res/lang/go.yaml)
+ 主要字段：
  + `lang_name`：语言名
  + `custom`：自定义类型 / 指针 / 数组在该语言中的写法（`T`、`TArray`、`TPointer`、`TPointerArray`）
  + `data_types`：
    + `name`：Excel 字段类型名（需覆盖 `system.yaml` 的 `field_datatypes`）
    + `lang`：该语言中的类型名
    + `operates`：按数据文件给出 `get` / `set` 方法名

#### 4.5.2 模板配置

+ 默认位于 [res/template](./res/template)
+ 分别由 `temps_title`、`temps_const`、`temps_proto` 关联；可配置多个模板，**第一个为主模板**
+ 语法为 Go `text/template`：[官方文档](https://pkg.go.dev/text/template)

### 4.6. 数据库配置

以 MySQL 为例。

#### 4.6.1 配置

+ 默认位于 [res/db](./res/db)，由 `system.yaml` 的 `databases.list[].ref` 关联
+ [mysql.yaml](./res/db/mysql.yaml) 字段：
  + `db_name`
  + `scale_char` / `scale_varchar`：动态长度估算比例
  + `types`：Excel 类型到数据库类型的映射

当前 `databases.default` 为 `mysql`，导出 SQL 时使用该默认库。

#### 4.6.2 模板

+ [mysql_table.temp](./res/db/mysql_table.temp)：建表脚本，对应 `temps_table`
+ [mysql_data.temp](./res/db/mysql_data.temp)：插入数据脚本，对应 `temps_data`

## 5. 功能

四种基础导出：

+ [表头导出](#51-表头导出)：`-mode=title`
+ [数据导出](#52-数据导出)：`-mode=data`
+ [常量表导出](#53-常量表导出)：`-mode=const`
+ [协议导出](#54-协议导出)：`-mode=proto`

附加：

+ [SQL 导出](#55-sql-导出)：`-range` 含 `db`，`-file` 含 `sql`，且 `-mode` 含 `title` 或 `data`

处理顺序：按 `project.yaml` 的源路径加载每个 Excel，立刻处理该文件中全部匹配 Sheet，最后在需要时写出合并 SQL。

### 5.1 表头导出

把数据表字段导出为对应语言的结构体 / 类。

#### 5.1.1 处理流程

1. 遍历源路径中扩展名匹配、且未被 `ignore` 前缀或空文件过滤的 Excel。
2. 仅处理 Sheet 名以 `title&data.prefix` 开头的表。
3. 用控制行确定字段宽度，用范围行按 `-range` 选出字段。
4. 按 `-lang` 加载语言配置与 title 模板。
5. 从 Sheet 读取文件名、类名、命名空间，渲染后写入 `target.title.<range>`。

#### 5.1.2 模板数据

根对象为 [`*TempTitleProxy`](/src/core/context_title.go)，可用 `{{.}}` 或 `{{$proxy := .}}`。

**TempTitleProxy**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `Excel` | `*excel.ExcelProxy` | 当前 Excel 代理 |
| `Sheet` | `*excel.ExcelSheet` | 当前 Sheet |
| `TitleCtx` | `*TitleContext` | 表头上下文（含 `RangeName`、`Language` 等） |
| `FileName` | `string` | 导出文件名 |
| `ClassName` | `string` | 导出类名 |
| `Namespace` | `string` | 命名空间 / 包名 |
| `FieldIndex` | `[]int` | 选中字段的列索引 |
| `Language()` | `string` | 当前语言 |
| `LanguageDefine()` | `*setting.ProgramLanguage` | 语言定义 |
| `ValueAtAxis(axis)` | `string` | 取单元格文本 |
| `FieldLen()` | `int` | 字段数量 |
| `GetFields()` | `[]TitleFieldItem` | 全部字段 |
| `GetField(index)` | `TitleFieldItem` | 指定列字段 |
| `GetFieldName(index)` | `string` | 当前语言字段名 |
| `GetFieldFileKey(index, fileType)` | `string` | 数据文件字段键 |

**TitleFieldItem**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `Index` | `int` | 列索引 |
| `TitleName` | `string` | Excel 原始字段名 |
| `TitleRemark` | `string` | 备注（换行已转为 `<br/>`） |
| `FieldLangName` | `string` | 当前语言字段名 |
| `OriginalType` | `string` | Excel 原始类型 |
| `FormattedType` | `string` | 标准化类型（如 `string(8)` → `string`） |
| `LangType` | `string` | 语言类型名 |
| `LangTypeDefine` | `setting.LangDataType` | 语言类型定义，可用 `GetGetOperate` / `GetSetOperate` |
| `GetFileKey(fileType)` | `string` | 指定数据文件中的字段键 |

### 5.2 数据导出

- 默认格式：`json`、`bin`。`sql` 走独立 SQL 流程，不走数据 Builder。
- 输出目录：`target.data.<range>`，文件名为 Sheet 中配置的数据文件名 + 格式扩展名。
- 数据从 `data_start_axis` 对应行开始，遇到空行或首单元格为空即停止。
- 字段键取自 `file_key_rows` 中对应格式的行。
- 数组单元格格式：`[a,b,c]`；空数组可用 `[]`。

**JSON 结构：**

```json
{
  "count": 2,
  "data": [
    { "<fileKey>": "<value>" }
  ]
}
```

导出 `[]uint8` 时会转为 `[]uint16`，避免 JSON 把它当成 Base64 字符串。

**二进制：** 固定大端。文件头 4 字节为行数（`uint32`），随后按字段顺序逐行写入。`project.yaml` 的 `buff.big_endian` **不会改变**该行为。

要开放 yaml 等格式：在 `system.yaml` 的 `export_files` 中追加对应项。这些格式经 Viper 写出，**键名会被转成小写**。

### 5.3 常量表导出

#### 5.3.1 处理流程

1. 处理 Sheet 名以 `const.prefix` 开头的表。
2. 从 `data_start_row` 读到表尾，名称空的行跳过。
3. 按 `-lang` 渲染 const 模板，输出到 `target.const.client` 或 `target.const.server`。

#### 5.3.2 模板数据

根对象为 [`*TempConstProxy`](/src/core/context_const.go)。

**TempConstProxy**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `Excel` / `Sheet` | 代理 / Sheet | 当前表 |
| `ConstCtx` | `*ConstContext` | 上下文 |
| `FileName` / `ClassName` / `Namespace` | `string` | 导出信息 |
| `StartRow` / `EndRow` | `int` | 数据行范围（含起不含终） |
| `Language()` | `string` | 当前语言 |
| `ValueAtAxis(axis)` | `string` | 取单元格 |
| `GetItems()` | `[]ConstItem` | 全部常量（已跳过空行） |
| `GetItem(row)` | `ConstItem` | 指定 Excel 行号 |

**ConstItem**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `Name` | `string` | 常量名 |
| `Value` | `string` | 常量值；`string` 类型已加双引号 |
| `Type` | `string` | 语言类型名 |
| `Remark` | `string` | 备注 |

### 5.4 协议表导出

协议属性单元格格式：`name:type`。`type` 可以是基础类型、固定长度数组、自定义协议名，以及指针（由 `pointer_code` 标记，默认 `*`）。不支持 `json` / `[]json`。自定义类型必须是**同一张表中已经出现过的协议名**。

#### 5.4.1 处理流程

1. 处理 Sheet 名以 `proto.prefix` 开头的表。
2. 读取 Id 类型、范围、命名空间、额外子目录。
3. 从 `data_start_row` 解析协议行；`file` 或 `name` 为空则不是协议行。
4. 范围与 `-range` 匹配时，按 `-lang` 渲染 proto 模板，输出到 `target.proto.<range>`，并可叠加 Sheet 中的 `export` 子目录。

#### 5.4.2 模板数据

根对象为 [`*TempProtoProxy`](/src/core/context_proto.go)。

**TempProtoProxy**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `ProtoItem` | `ProtoItem` | 当前协议 |
| `SheetProxy` | `*ProtoSheetProxy` | Sheet 代理（含 `Excel`、`Sheet`、`ProtoCtx`、`Title`） |
| `ValueAtAxis(axis)` | `string` | 取单元格 |
| `Namespace()` | `string` | 命名空间 |
| `ProtoId()` | `string` | 协议 Id；string 类型已加双引号 |
| `ProtoIdDataType()` | `string` | Id 的语言类型 |
| `ClassName()` | `string` | 导出类名 |
| `ClassRemark()` | `string` | 类备注 |
| `GetFields()` | `[]ProtoFieldItem` | 属性列表 |

**ProtoFieldItem**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `Remark` | `string` | 属性备注 |
| `Name` | `string` | 属性名 |
| `Lang` | `string` | 当前语言 |
| `OriginalType` / `FormattedType` | `string` | 原始 / 标准化类型 |
| `LangType` | `string` | 语言类型名 |
| `LangTypeDefine` | `setting.LangDataType` | 基础类型定义 |
| `IsCustomType` | `bool` | 是否为自定义协议类型 |
| `IsPointer` / `IsArray` | `bool` | 指针 / 数组 |
| `ArraySize` | `int` | 固定长度；非数组或不定长为 `-1` |
| `TempLangType()` | `string` | 模板中应使用的最终类型写法（含指针与数组） |

### 5.5 SQL 导出

#### 5.5.1 处理流程

同时满足以下条件才会导出 SQL：

1. `-range` 包含 `db`
2. `-file` 包含 `sql`
3. `-mode` 至少包含 `title` 或 `data` 之一

说明：

+ 遍历规则与表头 / 数据导出相同（`Data_` 前缀）
+ `-mode=title` 生成建表脚本，`-mode=data` 生成插入脚本，二者可同时开启
+ `-merge=true`：全部写入 `all_merge.sql`
+ `-merge=false`：分别写出 `<fileName>.table.sql` 与 `<fileName>.data.sql`
+ 输出目录为 `target.sql.dir`（相对 `target.root`）

合并写出时，数据模板可通过 `NeedTruncateData()` 避免在同一文件里重复 `TRUNCATE`。

#### 5.5.2 模板数据

根对象为 [`*TempSqlProxy`](/src/core/context_sql.go)。

**TempSqlProxy**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `Excel` / `Sheet` | 代理 / Sheet | 当前表 |
| `SqlCtx` | `*SqlContext` | 上下文 |
| `TableName` | `string` | 表名 |
| `FieldIndex` | `[]int` | 选中列索引 |
| `StartRow` / `EndRow` | `int` | 数据行范围 |
| `StartColIndex` | `int` | 起始列索引 |
| `MergeOn()` | `bool` | 是否合并输出 |
| `NeedTruncateData()` | `bool` | 数据脚本是否应 TRUNCATE |
| `FieldLen()` / `ItemLen()` | `int` | 字段数 / 数据行数 |
| `ValueAtAxis(axis)` | `string` | 取单元格 |
| `GetFieldItems()` | `[]SqlFieldItem` | 字段定义 |
| `GetPrimaryKeys()` | `[]SqlFieldItem` | 主键字段 |
| `PrimaryKeyLen()` | `int` | 主键数量 |
| `GetItems()` | `[]SqlItem` | 数据行 |
| `GetItem(row)` | `SqlItem` | 指定行 |

**SqlFieldItem**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `FieldName` | `string` | SQL 字段名 |
| `FieldType` | `string` | Excel 原始类型 |
| `CustomFieldType` | `string` | 定制 SQL 类型（若启用定制行） |
| `SqlFieldType()` | `string` | 最终写入模板的数据库类型 |
| `MaxByteSize` / `MaxRuneSize` | `int` | 动态长度统计 |

**SqlItem**

| 名称 | 类型 | 说明 |
| --- | --- | --- |
| `FieldLen()` | `int` | 字段数 |
| `GetValues()` | `[]string` | 已按 SQL 字面量处理的值列表 |
| `GetValue(i)` / `GetSqlValue(i)` | `string` | 原始值 / SQL 字面量 |

字符串值会把 `'` 转义为 `''` 并加单引号；数值类型原样输出。

### 5.6 模板定制

模板为 Go `text/template`：[https://pkg.go.dev/text/template](https://pkg.go.dev/text/template)

自定义函数对全部模板有效。返回值必须是 1 个，或第 2 个为 `error` 的 2 个返回值。

| 函数 | 源码 | 作用 |
| --- | --- | --- |
| `ToLowerCamelCase` | [naming.go](/src/core/tools/naming.go) | 小驼峰 |
| `ToUpperCamelCase` | [naming.go](/src/core/tools/naming.go) | 大驼峰 |
| `Add` / `Sub` | [math.go](/src/core/tools/math.go) | 整数加减 |
| `NowTime` | [time.go](/src/core/tools/time.go) | `time.Time` |
| `NowTimeStr` | 同上 | `2006-01-02 15:04:05` |
| `NowTimeFormat` | 同上 | 按 Go 时间格式化 |
| `NowYear` / `NowMonth` / `NowDay` / `NowWeekday` | 同上 | 年、月（1–12）、日、星期（周日为 0） |
| `NowHour` / `NowMinute` / `NowSecond` | 同上 | 时分秒 |
| `NowUnix` / `NowUnixNano` | 同上 | 秒 / 纳秒时间戳 |

## 6. 已知限制

+ `yaml` / `toml` / `hcl` / `env` / `properties` 导出时键名会变为小写，默认未开放。
+ `project.yaml` 中 `encoding`、`buff` 未生效；二进制始终大端。
+ C++ 已有语言映射，但 title / const / proto 模板文件尚未提供。
+ 协议表不支持 `json`、`[]json` 字段类型。

## 7. 依赖性

- [infra-go](https://github.com/xuzhuoxi/infra-go) `v1.4.1`
- Excel 读写：内嵌 [excelize v2](/src/excelize.v2)（源自 [excelize](https://github.com/qax-os/excelize)）
- [yaml.v2](https://gopkg.in/yaml.v2)、[sjson](https://github.com/tidwall/sjson)、[viper](https://github.com/spf13/viper)

## 8. 联系作者

xuzhuoxi  
<xuzhuoxi@gmail.com> 或 <mailxuzhuoxi@163.com>

## 9. 许可证

ExcelExporter 源码基于 [MIT 许可证](/LICENSE) 开源。内嵌 excelize 使用其自身的 BSD 许可证。
