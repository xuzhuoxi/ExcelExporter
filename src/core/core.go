package core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/core/naming"
	"github.com/xuzhuoxi/ExcelExporter/src/core/temps"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/logx"
	"io/ioutil"
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
	Excel    *excel.ExcelProxy
)

var (
	ProgramLanguageTemps = make(map[string]*temps.TemplateProxy)
)

func init() {
	temps.RegisterFunc("ToLowerCamelCase", naming.ToLowerCamelCase)
	temps.RegisterFunc("ToUpperCamelCase", naming.ToUpperCamelCase)
}

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

	langDefine, ok := Setting.System.FindProgramLanguage(titleCtx.ProgramLanguage)
	if !ok {
		err = errors.New(fmt.Sprintf("-lang error at %d", titleCtx.ProgramLanguage))
		Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] %s ", err))
		return err
	}

	prefix := Setting.Excel.Prefix.Data
	fieldType := setting.FieldType(titleCtx.FieldType)
	Logger.Infoln(fmt.Sprintf("[core.executeTitleContext] Start Execute Content: %s", titleCtx))
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		outEle, ok := Setting.Excel.Output.GetElement(fieldType)
		if !ok {
			err = errors.New(fmt.Sprintf("-field error at %d", titleCtx.FieldType))
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Error A %s ", err))
			return err
		}

		fieldTypeRow := sheet.GetRowAt(Setting.Excel.Title.FieldSwitchRow - 1)
		if nil == fieldTypeRow || fieldTypeRow.Empty() {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Sheet execute pass at '%s' with filed type empty! ", sheet.SheetName))
			continue
		}
		selects, err := parseFileTypeRow(sheet, fieldTypeRow, uint(fieldType)-1)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Parse file type error: %s ", err))
			return err
		}
		//Logger.Infoln("Selects:", selects)
		titleName, _ := sheet.ValueAtAxis(outEle.TitleName)
		// 创建模板数据代理
		tempDataProxy := &TempDataProxy{Sheet: sheet, Excel: excel, Index: selects,
			TitleName: titleName, Language: titleCtx.ProgramLanguage}

		targetDir := Setting.Project.Target.GetTitleDir(fieldType)
		if !filex.IsExist(targetDir) {
			os.MkdirAll(targetDir, os.ModePerm)
		}

		fileName, err := sheet.ValueAtAxis(outEle.TitleName)
		extendName := langDefine.ExtendName
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] GetTitleFileName error: %s ", err))
			return err
		}
		filePath := filex.Combine(targetDir, fileName+"."+extendName)
		buff := bytes.NewBuffer(nil)
		err = temp.Execute(buff, tempDataProxy, false)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Execute Template error: %s ", err))
			return err
		}
		ioutil.WriteFile(filePath, buff.Bytes(), os.ModePerm)
		Logger.Infoln(fmt.Sprintf("[core.executeTitleContext] Generate file : %s", filePath))
	}
	Logger.Infoln(fmt.Sprintf("[core.executeTitleContext] Finish execute: %s", titleCtx))
	return nil
}

func executeDataContext(excel *excel.ExcelProxy, dataCtx *DataContext) error {
	return nil
}

func getProgramLanguageTemp(lang string) (t *temps.TemplateProxy, err error) {
	if _, ok := ProgramLanguageTemps[lang]; ok {
		return ProgramLanguageTemps[lang], nil
	}
	if l, ok := Setting.System.FindProgramLanguage(lang); ok {
		temp, err := temps.LoadTemplates(l.TempPaths())
		if nil != err {
			return nil, err
		}
		ProgramLanguageTemps[lang] = temp

		return temp, nil
	}
	return nil, errors.New(fmt.Sprintf("Undefined Program Lanaguage: %s", lang))
}

func parseFileTypeRow(sheet *excel.ExcelSheet, row *excel.ExcelRow, selectIndex uint) (selects []int, err error) {
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
