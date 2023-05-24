package core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/core/temps"
	"github.com/xuzhuoxi/infra-go/filex"
	"os"
	"strings"
)

func execExcelConstContext(excel *excel.ExcelProxy, constCtx *ConstContext) error {
	sheets := excel.GetSheets(constCtx.EnablePrefix)
	if len(sheets) == 0 {
		return nil
	}
	logPrefix := "core.execExcelConstContext"
	Logger.Infoln(fmt.Sprintf("[%s][--Start ConstContext]: %s", logPrefix, constCtx))
	for _, sheet := range sheets {
		err := execSheetConstContext(excel, sheet, constCtx)
		if nil != err {
			return err
		}
	}
	Logger.Infoln(fmt.Sprintf("[%s][--Finish ConstContext]: %s", logPrefix, constCtx))
	return nil
}

func execSheetConstContext(excel *excel.ExcelProxy, sheet *excel.ExcelSheet, constCtx *ConstContext) error {
	// 过滤Sheet的命名
	if strings.Index(sheet.SheetName, constCtx.EnablePrefix) != 0 {
		return nil
	}
	logPrefix := "core.execSheetConstContext"
	lang := constCtx.Language
	temp, err := getConstLanguageTemp(lang)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Get lang[%s] temp fail: %s", logPrefix, lang, err))
		return err
	}
	langDefine, ok := Setting.System.FindProgramLanguage(lang)
	if !ok {
		err = errors.New(fmt.Sprintf("[%s] -lang error at %s: lang undefined!", logPrefix, lang))
		return err
	}

	//Logger.Infoln(fmt.Sprintf("[%s][SheetName=%s, Context=%s]", logPrefix, sheet.SheetName, constCtx))
	outEle, ok := Setting.Excel.Const.GetOutputInfo(constCtx.RangeName)
	if !ok {
		err = errors.New(fmt.Sprintf("[%s] -field error at \"%s\": output file name!", logPrefix, constCtx.RangeName))
		return err
	}
	if outEle.FileAxis == "" {
		return nil
	}
	fileName, err := sheet.ValueAtAxis(outEle.FileAxis)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] GetTitleFileName Error: {Err=%s,FileName=%s}", logPrefix, err, fileName))
		return err
	}
	if strings.TrimSpace(fileName) == "" { // 导出文件如果为空，认为忽略导出
		Logger.Traceln(fmt.Sprintf("[%s] Ignore export because the file name is empty. ", logPrefix))
		return nil
	}

	clsName, err := sheet.ValueAtAxis(outEle.ClassAxis)
	if nil != err || strings.TrimSpace(clsName) == "" {
		err = errors.New(fmt.Sprintf("[%s] GetTitleClassName Error: {Err=%s,ClassName=%s}", logPrefix, err, clsName))
		return err
	}
	namespace, err := sheet.ValueAtAxis(outEle.NamespaceAxis)
	if nil != err || strings.TrimSpace(namespace) == "" {
		err = errors.New(fmt.Sprintf("[%s] GetTitleClassName Error: {Err=%s,ClassName=%s}", logPrefix, err, namespace))
		return err
	}
	targetDir := Setting.Project.Target.GetConstDir(constCtx.RangeName)
	if !filex.IsExist(targetDir) {
		os.MkdirAll(targetDir, os.ModePerm)
	}
	extendName := langDefine.ExtendName
	filePath := filex.Combine(targetDir, fileName+"."+extendName)

	// 创建模板数据代理
	startRow := Setting.Excel.Const.DataStartRow
	endRow := len(sheet.Rows) + 1
	tempConstProxy := &TempConstProxy{Sheet: sheet, Excel: excel, ConstCtx: constCtx,
		FileName: fileName, ClassName: clsName, Namespace: namespace,
		StartRow: startRow, EndRow: endRow}

	buff := bytes.NewBuffer(nil)
	err = temp.Execute(buff, tempConstProxy, false)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Execute Template error: %s ", logPrefix, err))
		return err
	}
	filex.WriteFile(filePath, buff.Bytes(), os.ModePerm)
	Logger.Infoln(fmt.Sprintf("[%s] \t file => %s", logPrefix, filePath))
	return nil
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
	return nil, errors.New(fmt.Sprintf("[core.getConstLanguageTemp] Undefined Program Lanaguage for Const: %s", lang))
}
