package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/data"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/core/temps"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/slicex"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	Logger         logx.ILogger
	Setting        *setting.Settings
	Excel          *excel.ExcelProxy
	EnablePrefixes []string
)

var (
	TitleCtx []*TitleContext
	DataCtx  []*DataContext
	SqlCtx   *SqlContext
	ConstCtx []*ConstContext
	ProtoCtx []*ProtoContext
)

var (
	TitleLanguageTemps = make(map[string]*temps.TemplateProxy)
	ConstLanguageTemps = make(map[string]*temps.TemplateProxy)
	SqlTableTemps      *temps.TemplateProxy
	SqlDataTemps       *temps.TemplateProxy
	ProtoLanguageTemps = make(map[string]*temps.TemplateProxy)
)

type funcExec = func(showFileInfo bool) error

func SetLogger(logger logx.ILogger) {
	if nil == logger {
		logger = logx.NewLogger()
		logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	}
	Logger = logger
}

func Execute(setting *setting.Settings, titleCtx []*TitleContext, dataCtx []*DataContext, sqlCtx *SqlContext,
	constCtx []*ConstContext, protoCtx []*ProtoContext) {
	Setting = setting

	TitleCtx = titleCtx
	DataCtx = dataCtx
	SqlCtx = sqlCtx
	ConstCtx = constCtx
	ProtoCtx = protoCtx
	EnablePrefixes = getMergedPrefixes()

	Logger.Infoln(fmt.Sprintf("[core.Execute][Settings]: %v", setting))
	Logger.Infoln(fmt.Sprintf("[core.Execute][Title][%d]: %v", len(titleCtx), titleCtx))
	Logger.Infoln(fmt.Sprintf("[core.Execute][Data][%d]: %v", len(dataCtx), dataCtx))
	Logger.Infoln(fmt.Sprintf("[core.Execute][Sql]: %v", sqlCtx))
	Logger.Infoln(fmt.Sprintf("[core.Execute][Const][%d]: %v", len(constCtx), constCtx))
	Logger.Infoln(fmt.Sprintf("[core.Execute][Proto][%d]: %v", len(protoCtx), protoCtx))

	// 每加载一个文件，马上处理当前文件的全部Sheet
	loadExcelFiles(execExcelSheets)

	// 加载全部文件，统一处理全部sheet
	//Logger.Println()
	//loadExcelFiles(nil)
	//execExcelSheets(true)

	// 合并Sql输出
	handleSqlMerge()
}

func loadExcelFiles(execFunc funcExec) {
	logPrefix := "core.loadExcelFiles"
	if nil == Setting {
		Logger.Infoln(fmt.Sprintf("[%s] Execution stop with error settings. ", logPrefix))
		return
	}
	if len(TitleCtx) == 0 && len(DataCtx) == 0 && len(ConstCtx) == 0 && nil == SqlCtx && len(ProtoCtx) == 0 {
		Logger.Infoln(fmt.Sprintf("[%s] Execution finish with doing nothing. ", logPrefix))
		return
	}
	Excel = &excel.ExcelProxy{}
	// 遍历Source
	sourcePath := Setting.Project.Source.Value
	for _, path := range sourcePath {
		if !filex.IsExist(path) {
			Logger.Errorln(fmt.Sprintf("[%s] Source(%s) is not exist. ", logPrefix, path))
			continue
		}
		if Setting.Excel.CheckIgnorePath(path) {
			Logger.Warnln(fmt.Sprintf("[%s] Source(%s) ignored . ", logPrefix, path))
			continue
		}
		if filex.IsFolder(path) {
			loadExcelFilesFromFolder(path, execFunc)
		} else {
			loadExcelFile(path, nil, execFunc)
		}
	}
}

func loadExcelFilesFromFolder(folderPath string, execFunc funcExec) {
	filex.WalkAllFiles(folderPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		loadExcelFile(filePath, fileInfo, execFunc)
		return nil
	})
}

func loadExcelFile(filePath string, fileInfo os.FileInfo, execFunc funcExec) {
	logPrefix := "core.loadExcelFile"
	isFileIgnore := Setting.Excel.CheckIgnorePath(filePath)
	isFileMatching := Setting.Project.Source.CheckFileFormat(filePath)
	isFileEmpty := nil != fileInfo && fileInfo.Size() == 0
	if isFileIgnore || !isFileMatching || isFileEmpty {
		if nil != execFunc {
			Logger.Println()
		}
		if isFileIgnore {
			Logger.Traceln(fmt.Sprintf("[%s] Ignore file by ignore prefix: %s", logPrefix, filePath))
		} else if !isFileMatching {
			Logger.Traceln(fmt.Sprintf("[%s] Ignore file by format unmatching: %s", logPrefix, filePath))
		} else {
			Logger.Traceln(fmt.Sprintf("[%s] Ignore file by empty size: %s", logPrefix, filePath))
		}
		return
	}
	if nil != execFunc {
		Logger.Println()
		Logger.Infoln(fmt.Sprintf("[%s] Start At %s", logPrefix, filePath))
		err := Excel.LoadExcel(filePath, true) // 注意：这里使用了覆盖
		if nil != err {
			Logger.Errorln(fmt.Sprintf("[%s] Error At %s", logPrefix, err))
			return
		}

		err = execFunc(false)
		if nil != err {
			Logger.Errorln(fmt.Sprintf("[%s] Error At %s", logPrefix, err))
			return
		}
		Logger.Infoln(fmt.Sprintf("[%s] Finish At %s", logPrefix, filePath))
	} else {
		err := Excel.LoadExcel(filePath, false) // 注意：这里不覆盖
		if nil != err {
			Logger.Errorln(fmt.Sprintf("[%s] Error At %s", logPrefix, err))
		}
	}
}

func execExcelSheets(showFileInfo bool) (err error) {
	logPrefix := "core.execExcelSheets"
	colNameRow := Setting.Excel.TitleData.NickRow
	err = Excel.LoadSheetsByPrefixes(EnablePrefixes, colNameRow, true) //加载全部
	if nil != err {
		return
	}
	filePath := ""
	for _, sheet := range Excel.Sheets {
		if showFileInfo && filePath != sheet.FilePath {
			if len(filePath) != 0 {
				Logger.Infoln(fmt.Sprintf("[%s] Finish At %s", logPrefix, filePath))
			}
			filePath = sheet.FilePath
			Logger.Println()
			Logger.Infoln(fmt.Sprintf("[%s] Start At %s", logPrefix, filePath))
		}
		execSheet(sheet)
	}
	if showFileInfo && len(filePath) != 0 {
		Logger.Infoln(fmt.Sprintf("[%s] Finish At %s", logPrefix, filePath))
	}
	return
}

func execSheet(sheet *excel.ExcelSheet) {
	logPrefix := "core.execSheet"
	if len(TitleCtx) > 0 || len(DataCtx) > 0 || len(ConstCtx) > 0 || nil != SqlCtx {
		Logger.Infoln(fmt.Sprintf("[%s] [SheetName=%s, FileName=%s]", logPrefix, sheet.SheetName, sheet.FileName()))
	}
	if len(TitleCtx) > 0 {
		for _, titleCtx := range TitleCtx {
			et := execSheetTitleContext(Excel, sheet, titleCtx)
			if nil != et {
				Logger.Errorln(fmt.Sprintf("[%s] %s", logPrefix, et))
			}
		}
	}

	if len(DataCtx) > 0 {
		for _, dataCtx := range DataCtx {
			ed := execSheetDataContext(Excel, sheet, dataCtx)
			if nil != ed {
				Logger.Errorln(fmt.Sprintf("[%s] %s", logPrefix, ed))
			}
		}
	}

	if nil != SqlCtx {
		es := execSheetSqlContext(Excel, sheet, SqlCtx)
		if nil != es {
			Logger.Errorln(fmt.Sprintf("[%s] %s", logPrefix, es))
		}
	}

	if len(ConstCtx) > 0 {
		for _, constCtx := range ConstCtx {
			ec := execSheetConstContext(Excel, sheet, constCtx)
			if nil != ec {
				Logger.Errorln(fmt.Sprintf("[%s] %s", logPrefix, ec))
			}
		}
	}

	if len(ProtoCtx) > 0 {
		for _, protoCtx := range ProtoCtx {
			ec := execSheetProtoContext(Excel, sheet, protoCtx)
			if nil != ec {
				Logger.Errorln(fmt.Sprintf("[%s] %s", logPrefix, ec))
			}
		}
	}
}

func handleSqlMerge() {
	if sqlMergeBuffExist {
		writeMergedSql(SqlCtx)
	}
}

func getControlSize(sheet *excel.ExcelSheet) (size int) {
	controlRow := sheet.GetRowAt(Setting.Excel.TitleData.ControlRow - 1)
	for index, value := range controlRow.Cell {
		str := strings.TrimSpace(value)
		if len(str) == 0 {
			return index
		}
	}
	return len(controlRow.Cell)
}

func parseRangeRow(sheet *excel.ExcelSheet, rangeRow *excel.ExcelRow, rangeIndex uint, startIndex int, maxSize int) (selects []int, colNames []string, err error) {
	cellLen := rangeRow.CellLength()
	if maxSize > cellLen {
		return nil, nil, errors.New(fmt.Sprintf("Range Row Lack At (%s)[%s]", sheet.SheetName, rangeRow.Axis()[cellLen]))
	}
	for index, cell := range rangeRow.Cell {
		if index == maxSize {
			return
		}
		if index < startIndex {
			continue
		}
		m, _ := regexp.MatchString(RegexPatternRange, cell)
		if !m {
			return nil, nil, errors.New(fmt.Sprintf("Cells Value Error At Sheet(%s)[%s]", sheet.SheetName, rangeRow.Axis()[index]))
		}
		ss := strings.Split(cell, ",")
		value, _ := strconv.Atoi(ss[rangeIndex])
		if value == 0 {
			continue
		}
		selects = append(selects, index)
		colNames = append(colNames, excel.GetAxisName(index+1))
	}
	return
}

func getRowData(keyRow *excel.ExcelRow, typeRow *excel.ExcelRow, valueRow *excel.ExcelRow,
	selects []int, selectNames []string, rowNum int) (dataRow []*data.KTValue) {
	dataRow = make([]*data.KTValue, len(selects), len(selects))
	for index, rowIndex := range selects {
		k, _ := keyRow.ValueAtIndex(rowIndex)
		t, _ := typeRow.ValueAtIndex(rowIndex)
		v, _ := valueRow.ValueAtIndex(rowIndex)
		loc := fmt.Sprintf("%s%d", selectNames[index], rowNum)
		dataRow[index] = &data.KTValue{Key: k, Type: t, Value: v, Loc: loc}
	}
	return
}

func getMergedPrefixes() []string {
	var prefixes []string
	if len(TitleCtx) > 0 {
		prefixes = appendPrefix(prefixes, TitleCtx[0].EnablePrefix)
	}
	if len(DataCtx) > 0 {
		prefixes = appendPrefix(prefixes, DataCtx[0].EnablePrefix)
	}
	if nil != SqlCtx {
		prefixes = appendPrefix(prefixes, SqlCtx.EnablePrefix)
	}
	if len(ConstCtx) > 0 {
		prefixes = appendPrefix(prefixes, ConstCtx[0].EnablePrefix)
	}
	if len(ProtoCtx) > 0 {
		prefixes = appendPrefix(prefixes, ProtoCtx[0].EnablePrefix)
	}
	return prefixes
}

func appendPrefix(prefixes []string, prefix string) []string {
	if _, ok := slicex.IndexString(prefixes, prefix); ok {
		return prefixes
	}
	return append(prefixes, prefix)
}
