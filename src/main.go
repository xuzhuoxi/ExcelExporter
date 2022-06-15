package main

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core"
	"github.com/xuzhuoxi/ExcelExporter/src/core/cmd"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/osxu"
)

const (
	ToolName = "ExcelExporter"
)

func main() {
	logger := logx.NewLogger()
	logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	logger.SetConfig(logx.LogConfig{Type: logx.TypeRollingFile, Level: logx.LevelAll,
		FileDir: osxu.GetRunningDir(), FileName: ToolName, FileExtName: ".log", MaxSize: 10 * mathx.MB})
	core.SetLogger(logger)

	flags, err := cmd.ParseFlag()
	if nil != err {
		logger.Error(err)
		return
	}
	logger.Infoln(fmt.Sprintf("[main] Flag=%s", flags))
	s := &setting.Settings{}
	s.Init(flags.EnvPath)
	s.System.UpgradeEnvPath(setting.EnvPath)
	s.Project.UpdateSource(flags.Source)
	s.Project.UpdateTarget(flags.Target)
	s.Project.UpgradeEnvPath(setting.EnvPath)
	logger.Infoln(fmt.Sprintf("[main] %v", s.System))
	logger.Infoln(fmt.Sprintf("[main] %v", s.Excel))
	logger.Infoln(fmt.Sprintf("[main] %v", s.Project))

	cmdParams := flags.GetCommandParams()
	logger.Infoln(fmt.Sprintf("[main] Command=%s", cmdParams))

	for _, lang := range cmdParams.LangRefs {
		err = s.InitLangSetting(lang)
		if nil != err {
			logger.Infoln(fmt.Sprintf("[main] Error: %s", err))
			return
		}
	}

	titleCtxArr := cmdParams.GenTitleContexts()
	dataCtxArr := cmdParams.GenDataContexts()
	constCtxArr := cmdParams.GenConstContexts()
	sqlCtx := cmdParams.GenSqlContext()
	core.Execute(s, titleCtxArr, dataCtxArr, constCtxArr, sqlCtx)
}
