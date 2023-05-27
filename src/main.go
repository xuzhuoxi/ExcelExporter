package main

import (
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core"
	"github.com/xuzhuoxi/ExcelExporter/src/core/cmd"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"os"
)

const (
	ToolName = "ExcelExporter"
)

var (
	Logger   logx.ILogger
	Settings *setting.Settings
	AppFlags *cmd.AppFlags
)

func main() {
	logger := logx.NewLogger()
	logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	logger.SetConfig(logx.LogConfig{Type: logx.TypeRollingFile, Level: logx.LevelAll,
		FileDir: osxu.GetRunningDir(), FileName: ToolName, FileExtName: ".log", MaxSize: 10 * mathx.MB})
	Logger = logger
	core.SetLogger(logger)

	flags, err := cmd.ParseFlag()
	if nil != err {
		logger.Error(err)
		return
	}
	logger.Infoln(fmt.Sprintf("[main] Flag=%s", flags))
	s := &setting.Settings{}
	Settings = s
	s.Init(flags.EnvPath)
	s.System.UpgradeEnvPath(setting.EnvPath)
	s.Project.UpdateSource(flags.Source)
	s.Project.UpdateTarget(flags.Target)
	s.Project.UpgradeEnvPath(setting.EnvPath)
	logger.Infoln(fmt.Sprintf("[main] %v", s.System))
	logger.Infoln(fmt.Sprintf("[main] %v", s.Excel))
	logger.Infoln(fmt.Sprintf("[main] %v", s.Project))

	cmdParams := flags.GetCommandParams()
	AppFlags = cmdParams
	logger.Infoln(fmt.Sprintf("[main] Command=%s", cmdParams))

	initAndFixLangs()
	checkAndFixDataFiles()

	startRowNum := s.Excel.TitleData.DataStartRow()
	startColIndex := s.Excel.TitleData.DataStartColIndex()
	titlePrefix := s.Excel.TitleData.Prefix
	dataPrefix := titlePrefix
	sqlPrefix := titlePrefix
	constPrefix := s.Excel.Const.Prefix
	protoPrefix := s.Excel.Proto.Prefix

	titleCtxArr := cmdParams.GenTitleContexts(titlePrefix, startRowNum, startColIndex)
	dataCtxArr := cmdParams.GenDataContexts(dataPrefix, startRowNum, startColIndex)
	sqlCtx := cmdParams.GenSqlContext(sqlPrefix, startRowNum, startColIndex)
	constCtxArr := cmdParams.GenConstContexts(constPrefix)
	protoCtxArr := cmdParams.GenProtoContexts(protoPrefix)

	core.Execute(s, titleCtxArr, dataCtxArr, sqlCtx, constCtxArr, protoCtxArr)
}

func initAndFixLangs() {
	// 只导出sql脚本
	if len(AppFlags.LangRefs) == 0 &&
		len(AppFlags.DataFiles) == 1 && AppFlags.CheckDataFile("sql") &&
		len(AppFlags.RangeNames) == 1 && AppFlags.CheckRange(core.FieldRangeDatabase) {
		return
	}
	for index := len(AppFlags.LangRefs) - 1; index >= 0; index -= 1 {
		lang := AppFlags.LangRefs[index]
		err := Settings.InitLangSetting(lang)
		if nil != err {
			AppFlags.LangRefs = append(AppFlags.LangRefs[:index], AppFlags.LangRefs[index+1:]...)
			Logger.Errorln(fmt.Sprintf("[main] Lang Config Error: %s", err))
			os.Exit(1)
		}
	}
}

func checkAndFixDataFiles() {
	for index := len(AppFlags.DataFiles) - 1; index >= 0; index -= 1 {
		datafile := AppFlags.DataFiles[index]
		if !Settings.System.CheckExportDataFile(datafile) {
			AppFlags.LangRefs = append(AppFlags.DataFiles[:index], AppFlags.DataFiles[index+1:]...)
			Logger.Errorln(fmt.Sprintf("[main] DataFile Config Error: %s", datafile))
		}
	}
}
