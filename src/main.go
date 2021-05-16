package main

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/cmd"
	"github.com/xuzhuoxi/ExcelExporter/src/core"
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
	s.Init()

	runningPath := osxu.GetRunningDir()

	s.System.UpgradePath(runningPath)

	s.Project.UpdateSource(flags.Source)
	s.Project.UpdateTarget(flags.Target)
	s.Project.UpgradePath(runningPath)

	cmdParams := flags.GetCommandParams()
	logger.Infoln(fmt.Sprintf("[main] Command=%s", cmdParams))

	titleContexts := cmdParams.GenTitleContexts()
	dataContexts := cmdParams.GenDataContexts()

	core.Execute(s, titleContexts, dataContexts)
}
