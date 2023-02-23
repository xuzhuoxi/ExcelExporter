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
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	Logger logx.ILogger

	Setting  *setting.Settings
	TitleCtx []*TitleContext
	DataCtx  []*DataContext
	ConstCtx []*ConstContext
	SqlCtx   *SqlContext
	Excel    *excel.ExcelProxy
)

var (
	TitleLanguageTemps = make(map[string]*temps.TemplateProxy)
	ConstLanguageTemps = make(map[string]*temps.TemplateProxy)
	SqlTableTemps      *temps.TemplateProxy
	SqlDataTemps       *temps.TemplateProxy
)

func SetLogger(logger logx.ILogger) {
	if nil == logger {
		logger = logx.NewLogger()
		logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	}
	Logger = logger
}

func Execute(setting *setting.Settings, titleCtx []*TitleContext, dataCtx []*DataContext, constCtx []*ConstContext, sqlCtx *SqlContext) {
	Setting = setting
	TitleCtx = titleCtx
	DataCtx = dataCtx
	ConstCtx = constCtx
	SqlCtx = sqlCtx
	Logger.Infof("[core.Execute] Setting=%s", setting)
	Logger.Infof("[core.Execute] TitleContext=%v", titleCtx)
	Logger.Infof("[core.Execute] DataContext=%v", dataCtx)
	Logger.Infof("[core.Execute] ConstCtx=%v", constCtx)
	Logger.Infof("[core.Execute] SqlCtx=%v", sqlCtx)
	execExcelFiles()
	handleSqlMerge()
}

func execExcelFiles() {
	if nil == Setting {
		Logger.Infoln("[core.execExcelFiles] Execution stop with error settings. ")
		return
	}
	if len(TitleCtx) == 0 && len(DataCtx) == 0 && len(ConstCtx) == 0 {
		Logger.Infoln("[core.execExcelFiles] Execution finish with doing nothing. ")
		return
	}
	Excel = &excel.ExcelProxy{}
	// 遍历Source
	sourcePath := Setting.Project.Source.Value
	for _, path := range sourcePath {
		if !filex.IsExist(path) {
			Logger.Warnln(fmt.Sprintf("[core.execExcelFiles] Source(%s) is not exist. ", path))
			continue
		}
		if filex.IsFolder(path) {
			loadExcelFilesFromFolder(path)
		} else {
			loadExcelFile(path, nil)
		}
	}
}

func loadExcelFilesFromFolder(folderPath string) {
	filex.WalkAllFiles(folderPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		loadExcelFile(filePath, fileInfo)
		return nil
	})
}

func loadExcelFile(filePath string, fileInfo os.FileInfo) {
	isFileMatching := Setting.Project.Source.CheckFileFormat(filePath)
	isFileEmpty := nil != fileInfo && fileInfo.Size() == 0
	if !isFileMatching || isFileEmpty {
		Logger.Infoln(fmt.Sprintf("[core.loadExcelFile] Ignore file: %s", filePath))
		return
	}

	Logger.Println()
	Logger.Infoln(fmt.Sprintf("[core.loadExcelFile] Start At %s", filePath))
	err := executeExcelFile(filePath)
	if nil != err {
		Logger.Warnln(fmt.Sprintf("[core.loadExcelFile] Error At %s", err))
	} else {
		Logger.Infoln(fmt.Sprintf("[core.loadExcelFile] Finish At %s", filePath))
	}
}

func executeExcelFile(dataFilePath string) (err error) {
	err = Excel.LoadExcel(dataFilePath, true)
	if nil != err {
		return
	}

	//colNameRow := Setting.Excel.Title.ColNickRow
	colNameRow := Setting.Excel.TitleData.NickRow
	err = Excel.LoadSheets("", colNameRow, true) //加载全部
	if nil != err {
		return
	}

	if len(TitleCtx) > 0 {
		for _, titleCtx := range TitleCtx {
			et := executeTitleContext(Excel, titleCtx)
			if nil != et {
				Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] Error:%s", et))
			}
		}
	}

	if len(DataCtx) > 0 {
		for _, dataCtx := range DataCtx {
			ed := executeDataContext(Excel, dataCtx)
			if nil != ed {
				Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] Error:%s", ed))
			}
		}
	}

	if len(ConstCtx) > 0 {
		for _, constCtx := range ConstCtx {
			ec := executeConstContext(Excel, constCtx)
			if nil != ec {
				Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] Error:%s", ec))
			}
		}
	}

	if nil != SqlCtx {
		es := executeSqlContext(Excel, SqlCtx)
		if nil != es {
			Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] Error:%s", es))
		}
	}

	return
}

func handleSqlMerge() {
	if sqlBuffMergeExist {
		writeSqlMergeContext(SqlCtx)
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

func parseRangeRow(sheet *excel.ExcelSheet, rangeRow *excel.ExcelRow, rangeIndex uint, startIndex int, maxSize int) (selects []int, err error) {
	cellLen := rangeRow.CellLength()
	if maxSize > cellLen {
		return nil, errors.New(fmt.Sprintf("Range Row Lack At (%s)[%s]", sheet.SheetName, rangeRow.Axis()[cellLen]))
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
			return nil, errors.New(fmt.Sprintf("Cells Value Error At Sheet(%s)[%s]", sheet.SheetName, rangeRow.Axis()[index]))
		}
		ss := strings.Split(cell, ",")
		value, _ := strconv.Atoi(ss[rangeIndex])
		if value == 0 {
			continue
		}
		selects = append(selects, index)
	}
	return
}

func getRowData(keyRow *excel.ExcelRow, typeRow *excel.ExcelRow, valueRow *excel.ExcelRow, selects []int) (dataRow []*data.KTValue) {
	dataRow = make([]*data.KTValue, len(selects), len(selects))
	for index, rowIndex := range selects {
		k, _ := keyRow.ValueAtIndex(rowIndex)
		t, _ := typeRow.ValueAtIndex(rowIndex)
		v, _ := valueRow.ValueAtIndex(rowIndex)
		dataRow[index] = &data.KTValue{Key: k, Type: t, Value: v}
	}
	return
}
