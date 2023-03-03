## Release notes  

### Known Issues in v1.0.5  
- yaml, toml, hcl, env, properties数据导出时，key会转为**小写**，本意要求**大小写相关**。
- project.yaml中encoding与buff相关的配置未实现。
- C++表头模板未实现， C++常量模板未实现。

### Improvements  
- 优化excel.yaml配置结构，使之更容易理解。  
- 优化Excel处理顺序逻辑，使之符合人类一般逻辑思维。  
- 优化Log信息显示， 进行分组打印，减小打印量，提高清晰度。  
- 优化Log信息显示，对数据处理异常时增加坐标的打印，以便于问题的定位。  

### Fixes  

### Changes  
- excel配置："title&data" 配置中合并classes配置项到outputs中，并增加namespace的坐标配置。  
- 模板：优化表头模板与常量模板，支持namespace的内容填充。  
- 执行逻辑： 优化表处理顺序。  
  + 旧顺序： 每个Excel -> 命令mode、lang、ragne、 file、merge的组合处理 -> Sheet.  
  + 新顺序： 每个Excel =》 Sheets =》 命令mode、lang、ragne、 file、merge的组合处理  
- Log: 优化Log信息，减小打印信息，并提高清晰度。  
- Log: 增加数据处理异常时的坐标信息打印。   

### API Changes  
- data.KTValue：增加Loc属性，用于记录坐标值。  
- excel.ExcelProxy:  增加公共函数 GetExcelByPath(filePath string) \*excelize.File  
- excel.ExcelProxy:  函数重命名 LoadSheets =》 LoadSheetsByPrefix  
- excel.ExcelProxy:  增加公共函数 LoadSheetsByPrefixes(sheetPrefix []string, colNickRow int, overwrite bool) error   
- excel.ExcelSheet: 增加公共函数 FileName() string   
- 包excel:  修改函数 LoadSheets(excelPath string, excelFile \*excelize.File, sheetPrefix string, nickRow int) (sheets []\*ExcelSheet, err error) 为 LoadSheetsByPrefixes(excelPath string, excelFile *excelize.File, sheetPrefixes []string, nickRow int) (sheets []*ExcelSheet, err error)  
- cmd.AppFlags: GenTitleContexts 函数增加三个参数： (prefix string, startRowNum int, startColIndex int)  
- cmd.AppFlags: GenDataContexts 函数增加三个参数：( prefix string, startRowNum int, startColIndex int)  
- cmd.AppFlags: GenConstContexts 函数增加一个参数： (prefix string)  
- cmd.AppFlags: GenSqlContext 函数增加三个参数： (prefix string, startRowNum int, startColIndex int)  

## Library changes in v1.0.5  

### library updated  