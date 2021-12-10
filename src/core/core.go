package core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/data"
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
	ConstCtx []*ConstContext
	Excel    *excel.ExcelProxy
)

var (
	TitleLanguageTemps = make(map[string]*temps.TemplateProxy)
	ConstLanguageTemps = make(map[string]*temps.TemplateProxy)
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

func Execute(setting *setting.Settings, titleCtx []*TitleContext, dataCtx []*DataContext, constCtx []*ConstContext) {
	Setting = setting
	TitleCtx = titleCtx
	DataCtx = dataCtx
	ConstCtx = constCtx
	Logger.Infof("[core.Execute] Setting=%s", setting)
	Logger.Infof("[core.Execute] TitleContext=%v", titleCtx)
	Logger.Infof("[core.Execute] DataContext=%v", dataCtx)
	Logger.Infof("[core.Execute] ConstCtx=%v", constCtx)
	execExcelFiles()
}

func execExcelFiles() {
	if nil == Setting {
		Logger.Infoln("[core.execExcelFiles] Execution stop with error settings. ")
		return
	}
	if len(TitleCtx) == 0 && len(DataCtx) == 0 {
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
	filex.WaldAllFiles(folderPath, func(filePath string, fileInfo os.FileInfo, err error) error {
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
				Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] %s", et))
			}
		}
	}

	if len(DataCtx) > 0 {
		for _, dataCtx := range DataCtx {
			ed := executeDataContext(Excel, dataCtx)
			if nil != ed {
				Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] %s", ed))
			}
		}
	}

	if len(ConstCtx) > 0 {
		for _, constCtx := range ConstCtx {
			ed := executeConstContext(Excel, constCtx)
			if nil != ed {
				Logger.Warnln(fmt.Sprintf("[core.executeExcelFile] %s", ed))
			}
		}
	}
	return
}

func executeTitleContext(excel *excel.ExcelProxy, titleCtx *TitleContext) error {
	lang := titleCtx.ProgramLanguage
	temp, err := getTitleLanguageTemp(lang)
	if nil != err {
		return err
	}

	langDefine, ok := Setting.System.FindProgramLanguage(titleCtx.ProgramLanguage)
	if !ok {
		err = errors.New(fmt.Sprintf("-lang error at %s", titleCtx.ProgramLanguage))
		Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] %s ", err))
		return err
	}

	//prefix := Setting.Excel.Prefix.Data
	prefix := Setting.Excel.TitleData.Prefix
	Logger.Infoln(fmt.Sprintf("[core.executeTitleContext] Start Execute Content: %s", titleCtx))
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		//outEle, ok := Setting.Excel.Output.GetElement(titleCtx.RangeName)
		outEle, ok := Setting.Excel.TitleData.GetOutputInfo(titleCtx.RangeName)
		if !ok {
			err = errors.New(fmt.Sprintf("-field error at %s", titleCtx.RangeName))
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Error A %s ", err))
			return err
		}

		//fieldRangeRow := sheet.GetRowAt(Setting.Excel.Title.FieldRangeRow - 1)
		fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
		if nil == fieldRangeRow || fieldRangeRow.Empty() {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Sheet execute pass at '%s' with filed type empty! ", sheet.SheetName))
			continue
		}
		selects, err := parseFileTypeRow(sheet, fieldRangeRow, uint(titleCtx.RangeType)-1)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Parse file type error: %s ", err))
			return err
		}
		if len(selects) == 0 {
			continue
		}

		titleName, err := sheet.ValueAtAxis(outEle.Title)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] GetTitleFileName error: %s ", err))
			return err
		}
		targetDir := Setting.Project.Target.GetTitleDir(titleCtx.RangeName)
		if !filex.IsExist(targetDir) {
			os.MkdirAll(targetDir, os.ModePerm)
		}
		extendName := langDefine.ExtendName
		filePath := filex.Combine(targetDir, titleName+"."+extendName)

		// 创建模板数据代理
		tempDataProxy := &TempDataProxy{Sheet: sheet, Excel: excel, Index: selects,
			TitleName: titleName, Language: titleCtx.ProgramLanguage}

		//fileName, err := sheet.ValueAtAxis(outEle.TitleName)
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
	//prefix := Setting.Excel.Prefix.Data
	prefix := Setting.Excel.TitleData.Prefix
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		Logger.Infoln(fmt.Sprintf("[core.executeDataContext] Sheet[%s]", sheet.SheetName))
		//outEle, ok := Setting.Excel.Output.GetElement(dataCtx.RangeName)
		outEle, ok := Setting.Excel.TitleData.GetOutputInfo(dataCtx.RangeName)
		if !ok {
			err := errors.New(fmt.Sprintf("-field error at \"%s\"", dataCtx.RangeName))
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Error A %s ", err))
			return err
		}
		//fieldRangeRow := sheet.GetRowAt(Setting.Excel.Title.FieldRangeRow - 1)
		fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
		if nil == fieldRangeRow || fieldRangeRow.Empty() {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Sheet execute pass at '%s' with filed type empty! ", sheet.SheetName))
			continue
		}
		selects, err := parseFileTypeRow(sheet, fieldRangeRow, uint(dataCtx.RangeType)-1)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Parse file type error: %s ", err))
			return err
		}
		if len(selects) == 0 {
			continue
		}

		//fileName, err := sheet.ValueAtAxis(outEle.DataName)
		fileName, err := sheet.ValueAtAxis(outEle.Data)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] GetDataFileName error: %s ", err))
			return err
		}
		//keyRowNum := Setting.Excel.Title.FileKeyRows.GetRowNum(dataCtx.DataFileFormat)
		keyRowNum := Setting.Excel.TitleData.GetFileKeyRow(dataCtx.DataFileFormat)
		if -1 == keyRowNum {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Parse file format: %s ", dataCtx.DataFileFormat))
			continue
		}
		keyRow := sheet.GetRowAt(keyRowNum - 1)
		//typeRow := sheet.GetRowAt(Setting.Excel.Title.FieldFormatRow - 1)
		typeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldFormatRow - 1)
		//startRow := Setting.Excel.Data.StartRow
		startRow := Setting.Excel.TitleData.DataStartRow
		builder := data.GenBuilder(dataCtx.DataFileFormat)
		builder.StartWriteData()
		for startRow > 0 {
			dataRow := sheet.GetRowAt(startRow - 1)
			if nil == dataRow || len(dataRow.Cell) == 0 || dataRow.Cell[0] == "" { // 到达 表尾、空白头
				break
			}
			ktvRow := getRowData(keyRow, typeRow, dataRow, selects)
			err := builder.WriteRow(ktvRow)
			if nil != err {
				Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Error:%s", err))
				return err
			}
			startRow += 1
		}
		builder.FinishWriteData()

		targetDir := Setting.Project.Target.GetDataDir(dataCtx.RangeName)
		if !filex.IsExist(targetDir) {
			os.MkdirAll(targetDir, os.ModePerm)
		}
		extendName := dataCtx.DataFileFormat
		filePath := filex.Combine(targetDir, fileName+"."+extendName)
		Logger.Infoln(fmt.Sprintf("[core.executeDataContext] Sheet[%s]", filePath))
		err = builder.WriteDataToFile(filePath)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] WriteDataFile error: %s ", err))
		}
	}
	return nil
}

func executeConstContext(excel *excel.ExcelProxy, constCtx *ConstContext) error {
	lang := constCtx.ProgramLanguage
	temp, err := getConstLanguageTemp(lang)
	if nil != err {
		return err
	}
	langDefine, ok := Setting.System.FindProgramLanguage(lang)
	if !ok {
		err := errors.New(fmt.Sprintf("-lang error at %s", lang))
		Logger.Warnln(fmt.Sprintf("[core.executeConstContext] %s ", err))
		return err
	}
	prefix := Setting.Excel.Const.Prefix
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		Logger.Infoln(fmt.Sprintf("[core.executeConstContext] Sheet[%s]", sheet.SheetName))
		outEle, ok := Setting.Excel.Const.GetOutputInfo(constCtx.RangeName)
		if !ok {
			err := errors.New(fmt.Sprintf("-field error at \"%s\"", constCtx.RangeName))
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] Error A %s ", err))
			return err
		}
		if outEle.Value == "" {
			continue
		}
		fileName, err := sheet.ValueAtAxis(outEle.Value)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] GetTitleFileName error: %s ", err))
			return err
		}
		targetDir := Setting.Project.Target.GetConstDir(constCtx.RangeName)
		if !filex.IsExist(targetDir) {
			os.MkdirAll(targetDir, os.ModePerm)
		}
		extendName := langDefine.ExtendName
		filePath := filex.Combine(targetDir, fileName+"."+extendName)
		fmt.Println("6666:", filePath)

		// 创建模板数据代理
		tempDataProxy := &TempConstProxy{Sheet: sheet, Excel: excel,
			TitleName: fileName, Language: constCtx.ProgramLanguage,
			StartRow: Setting.Excel.Const.DataStartRow, EndRow: len(sheet.Rows)}

		buff := bytes.NewBuffer(nil)
		err = temp.Execute(buff, tempDataProxy, false)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] Execute Template error: %s ", err))
			return err
		}
		ioutil.WriteFile(filePath, buff.Bytes(), os.ModePerm)
		Logger.Infoln(fmt.Sprintf("[core.executeConstContext] Generate file : %s", filePath))
	}
	return nil
}

func getTitleLanguageTemp(lang string) (t *temps.TemplateProxy, err error) {
	if _, ok := TitleLanguageTemps[lang]; ok {
		return TitleLanguageTemps[lang], nil
	}
	if l, ok := Setting.System.FindProgramLanguage(lang); ok {
		temp, err := temps.LoadTemplates(l.GetTempsTitlePath())
		if nil != err {
			return nil, err
		}
		TitleLanguageTemps[lang] = temp

		return temp, nil
	}
	return nil, errors.New(fmt.Sprintf("Undefined Program Lanaguage for Title: %s", lang))
}

func getConstLanguageTemp(lang string) (t *temps.TemplateProxy, err error) {
	if _, ok := ConstLanguageTemps[lang]; ok {
		return ConstLanguageTemps[lang], nil
	}
	if l, ok := Setting.System.FindProgramLanguage(lang); ok {
		temp, err := temps.LoadTemplates(l.GetTempsConstPath())
		if nil != err {
			return nil, err
		}
		ConstLanguageTemps[lang] = temp

		return temp, nil
	}
	return nil, errors.New(fmt.Sprintf("Undefined Program Lanaguage for Const: %s", lang))
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

func getRowData(keyRow *excel.ExcelRow, typeRow *excel.ExcelRow, valueRow *excel.ExcelRow,
	selects []int) (dataRow []*data.KTValue) {
	dataRow = make([]*data.KTValue, len(selects), len(selects))
	for index, rowIndex := range selects {
		k := keyRow.Cell[rowIndex]
		t := typeRow.Cell[rowIndex]
		v := valueRow.Cell[rowIndex]
		dataRow[index] = &data.KTValue{Key: k, Type: t, Value: v}
	}
	return
}
