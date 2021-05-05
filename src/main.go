package main

import (
	"github.com/xuzhuoxi/ExcelExporter/src/cmd"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/logx"
)

var (
	GlobalLogger logx.ILogger
)

func main() {
	GlobalLogger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})

	flags, err := cmd.ParseFlag()
	if nil != err {
		GlobalLogger.Error(err)
		return
	}

	s := &setting.Settings{}
	s.Init()
	s.Project.UpdateSource(flags.Source)
	s.Project.UpdateTarget(flags.Target)

	cmdParams := flags.GetCommandParams()
	cmdParams.GenDataContexts()
	cmdParams.GenDefinitionContexts()
}
