# ExcelExporter

一个用于导出Excel数据的工具。

### 测试命名行参数

程序只允许通过命令行运行

支持的命令行参数包括：-mode, -range, -lang, -file, -source, -target

- -mode

    运行的模式，支持**Title导出**、**数据导出**、**常量表导出**
    
    - 支持值：title, data, const
    
        title为Title导出，data为数据导出，const为常量表导出
    
    - 支持多值，可用英文逗号","分隔

- -range

    运行时选择的字段范围，对**Title导出**和**数据导出**有效
    
    - 支持值：client, server, db
    
    - 支持多值，可用英文逗号","分隔
  
- -lang

    运行时选择的编程语言，只针对**Title导出**有效
    
    - 支持值：go, as3, ts, java, c#
    
    - 支持多值，可用英文逗号","分隔
    
- -file

    运行输出的数据文件类型，对**数据导出**有效

    - 支持值：bin, sql, json, yaml, toml, hcl, env, properties
    
    - 支持多值，可用英文逗号","分隔

- -source

    运行时指定数据来源目录，用于覆盖配置文件project.yaml中source.value的值
    
    可选项，对**Title导出**、**数据导出**、**常量表导出**有效

- -target

    运行时指定数据来源目录，用于覆盖配置文件project.yaml中target.value的值
    
    可选项，对**Title导出**、**数据导出**、**常量表导出**有效

-mode=1,2 -lang=go,as3,ts,java,c# -field=1,2,3 -file=json
    