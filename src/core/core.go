package core

import (
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/filex"
	"fmt"
	"os"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
)

type HandleMark uint

const (
	TitleMark HandleMark = 1 << iota
	DataMark
	ConstMark
)

var (
	Setting    *setting.Settings
	TitleCtx   []*TitleContext
	DataCtx    []*DataContext
	Logger     logx.ILogger
	ExcelProxy *excel.ExcelProxy
)

func Execute(setting *setting.Settings, titleCtx []*TitleContext, dataCtx []*DataContext, logger logx.ILogger) {
	Setting = setting
	TitleCtx = titleCtx
	DataCtx = dataCtx
	if nil == logger {
		logger = logx.NewLogger()
		logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	}
	Logger = logger
	execute()
}

func execute() {
	if nil == Setting {
		Logger.Infoln("Execution stop with error settings. ")
		return
	}
	if len(TitleCtx) == 0 && len(DataCtx) == 0 {
		Logger.Infoln("Execution finish with doing nothing. ")
		return
	}
	ExcelProxy = &excel.ExcelProxy{}
	// 遍历Source
	sourcePath := Setting.Project.Source.Value
	for _, path := range sourcePath {
		if !filex.IsExist(path) {
			Logger.Warnln(fmt.Sprintf("Source(%s) is not exist. ", path))
			continue
		}
		if filex.IsFolder(path) {
			executeFolder(path)
		} else {
			executeFile(path, nil)
		}
	}
}

func executeFolder(folderPath string) {
	filex.WaldAllFiles(folderPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		executeFile(filePath, fileInfo)
		return nil
	})
}

func executeFile(filePath string, fileInfo os.FileInfo) {
	isFileMatching := Setting.Project.Source.CheckFileMatching(filePath)
	isFileEmpty := nil != fileInfo && fileInfo.Size() == 0;
	if !isFileMatching || isFileEmpty {
		Logger.Infoln(fmt.Sprintf("Ignore file: %s", filePath))
		return
	}
	executeDataFile(filePath)
}

func executeDataFile(dataFilePath string) (err error) {
	err = ExcelProxy.LoadExcel(dataFilePath, true)

	if nil != err {
		Logger.Warnln(fmt.Sprintf("Load excel error(%s) at :%s", err.Error(), dataFilePath))
		return
	}
	Logger.Infoln(fmt.Sprintf("Load excel success at :%s", dataFilePath))

	prefix := Setting.Excel.Prefix.Data
	colNameRow := Setting.Excel.Title.ColNameRow
	ExcelProxy.LoadSheets(prefix, colNameRow, true)

	Logger.Infoln(prefix, colNameRow, ExcelProxy.Excels, ExcelProxy.Sheets)

	for _, sheet := range ExcelProxy.Sheets {
		Logger.Infoln(sheet.SheetName)
	}
	return
}
