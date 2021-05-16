package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/logx"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type HandleMark uint

const (
	TitleMark HandleMark = 1 << iota
	DataMark
	ConstMark
)

var (
	Logger logx.ILogger

	Setting  *setting.Settings
	TitleCtx []*TitleContext
	DataCtx  []*DataContext
	Excel    *excel.ExcelProxy
)

var (
	ProgramLanguageTemps map[string]*Template = make(map[string]*Template)
)

func SetLogger(logger logx.ILogger) {
	if nil == logger {
		logger = logx.NewLogger()
		logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	}
	Logger = logger
}

func Execute(setting *setting.Settings, titleCtx []*TitleContext, dataCtx []*DataContext) {
	Setting = setting
	TitleCtx = titleCtx
	DataCtx = dataCtx
	Logger.Infof("[core.Execute] Setting=%s", setting)
	Logger.Infof("[core.Execute] TitleContext=%s", titleCtx)
	Logger.Infof("[core.Execute] DataContext=%s", dataCtx)
	execute()
}

func execute() {
	if nil == Setting {
		Logger.Infoln("[core.execute] Execution stop with error settings. ")
		return
	}
	if len(TitleCtx) == 0 && len(DataCtx) == 0 {
		Logger.Infoln("[core.execute] Execution finish with doing nothing. ")
		return
	}
	Excel = &excel.ExcelProxy{}
	// 遍历Source
	sourcePath := Setting.Project.Source.Value
	for _, path := range sourcePath {
		if !filex.IsExist(path) {
			Logger.Warnln(fmt.Sprintf("[core.execute] Source(%s) is not exist. ", path))
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
	isFileMatching := Setting.Project.Source.CheckFileFormat(filePath)
	isFileEmpty := nil != fileInfo && fileInfo.Size() == 0
	if !isFileMatching || isFileEmpty {
		Logger.Infoln(fmt.Sprintf("[core.executeFile] Ignore file: %s", filePath))
		return
	}

	Logger.Println()
	Logger.Infoln(fmt.Sprintf("[core.executeFile] Start At %s", filePath))
	err := executeExcelFile(filePath)
	if nil != err {
		Logger.Warnln(fmt.Sprintf("[core.executeFile] Error At %s", err))
	} else {
		Logger.Infoln(fmt.Sprintf("[core.executeFile] Finish At %s", filePath))
	}
}

func executeExcelFile(dataFilePath string) (err error) {
	err = Excel.LoadExcel(dataFilePath, true)
	if nil != err {
		return
	}

	colNameRow := Setting.Excel.Title.ColNickRow
	err = Excel.LoadSheets("", colNameRow, true) //加载全部
	if nil != err {
		return
	}

	Logger.Infoln(fmt.Sprintf("[core.executeExcelFile] Load excel success at :%s", dataFilePath))

	for _, titleCtx := range TitleCtx {
		et := executeTitleContext(Excel, titleCtx)
		if nil != et {
			Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] %s", et))
		}
	}

	for _, dataCtx := range DataCtx {
		ed := executeDataContext(Excel, dataCtx)
		if nil != ed {
			Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] %s", ed))
		}
	}
	return
}

func executeTitleContext(excel *excel.ExcelProxy, titleCtx *TitleContext) error {
	lang := titleCtx.ProgramLanguage
	temp, err := getProgramLanguageTemp(lang)
	if nil != err {
		return err
	}

	Logger.Infoln(temp)

	prefix := Setting.Excel.Prefix.Data
	Logger.Infoln(fmt.Sprintf("[core.executeExcelFile] Start execute: %s", titleCtx))
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		fieldTypeRow := sheet.GetRowAt(Setting.Excel.Title.FieldTypeRow - 1)
		if nil == fieldTypeRow || fieldTypeRow.Empty() {
			Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] Sheet execute pass at '%s' with filed type empty! ", sheet.SheetName))
			continue
		}
		selects, err := parseFileTypeRow(sheet, fieldTypeRow, titleCtx.FieldTypeIndex)
		if nil != err {
			return err
		}
		Logger.Infoln(fmt.Sprintf("[core.executeExcelFile] %v", selects))
	}
	return nil
}

func executeDataContext(excel *excel.ExcelProxy, dataCtx *DataContext) error {
	return nil
}

func getProgramLanguageTemp(lang string) (temp *Template, err error) {
	if _, ok := ProgramLanguageTemps[lang]; ok {
		return ProgramLanguageTemps[lang], nil
	}
	if ok, l := Setting.System.FindProgramLanguage(lang); ok {
		temp, err := LoadTemplates(l.TempPaths())
		if nil == err {
			return nil, err
		}
		ProgramLanguageTemps[lang] = temp
		return temp, nil
	}
	return nil, errors.New(fmt.Sprintf("Undefined Program Lanaguage: %s", lang))
}

func parseFileTypeRow(sheet *excel.ExcelSheet, row *excel.ExcelRow, selectIndex int) (selects []int, err error) {
	for index, cell := range row.Cell {
		m, _ := regexp.MatchString(`[01],[01],[01]`, cell)
		if !m {
			return nil, errors.New(fmt.Sprintf("Cell Value Error At Sheet(%s)[%s]", sheet.SheetName, row.Axis()[index]))
		}
		ss := strings.Split(cell, ",")
		value, _ := strconv.Atoi(ss[selectIndex])
		if value == 0 {
			continue
		}
		selects = append(selects, index)
	}
	return
}
