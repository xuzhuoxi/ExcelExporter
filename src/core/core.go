package core

import (
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/logx"
)

type HandleMark uint

const (
	TitleMark HandleMark = 1 << iota
	DataMark
	ConstMark
)

var (
	Setting  *setting.Settings
	TitleCtx []*TitleContext
	DataCtx  []*DataContext
	Logger   logx.ILogger
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
	// 遍历Source
}
