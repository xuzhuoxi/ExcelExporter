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

func execExcelTitleContext(excel *excel.ExcelProxy, titleCtx *TitleContext) error {
	sheets := excel.GetSheets(titleCtx.EnablePrefix)
	if len(sheets) == 0 {
		return nil
	}
	logPrefix := "core.execExcelTitleContext"
	Logger.Infoln(fmt.Sprintf("[%s][--Start TitleContext]: %s", logPrefix, titleCtx))
	for _, sheet := range sheets {
		err := execSheetTitleContext(excel, sheet, titleCtx)
		if nil != err {
			return err
		}
	}
	Logger.Infoln(fmt.Sprintf("[%s][--Finish TitleContext]: %s", logPrefix, titleCtx))
	return nil
}

func execSheetTitleContext(excel *excel.ExcelProxy, sheet *excel.ExcelSheet, titleCtx *TitleContext) error {
	// 过滤Sheet的命名
	if strings.Index(sheet.SheetName, titleCtx.EnablePrefix) != 0 {
		return nil
	}
	logPrefix := "core.execSheetTitleContext"
	lang := titleCtx.Language
	temp, err := getTitleLanguageTemp(lang)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Get lang[%s] temp fail: %s", logPrefix, lang, err))
		return err
	}

	langDefine, ok := Setting.System.FindProgramLanguage(lang)
	if !ok {
		err = errors.New(fmt.Sprintf("[%s] -lang error at %s: lang undefined!", logPrefix, lang))
		return err
	}

	//Logger.Infoln(fmt.Sprintf("[%s][SheetName=%s, FileName=%s]", logPrefix, sheet.SheetName, sheet.FileName()))
	outEle, ok := Setting.Excel.TitleData.GetOutputInfo(titleCtx.RangeName)
	if !ok {
		err = errors.New(fmt.Sprintf("[%s] -field error at \"%s\": output file name!", logPrefix, titleCtx.RangeName))
		return err
	}

	size := getControlSize(sheet)
	fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
	if nil == fieldRangeRow || fieldRangeRow.Empty() { // 忽略
		Logger.Warnln(fmt.Sprintf("[%s] Ignore[%s] execution for filed type empty!", logPrefix, sheet.SheetName))
		return nil
	}
	selects, _, err := parseRangeRow(sheet, fieldRangeRow, uint(titleCtx.RangeType)-1, titleCtx.StartColIndex, size)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Parse Range Row error: %s ", logPrefix, err))
		return err
	}
	if len(selects) == 0 {
		return nil
	}

	fileName, err := sheet.ValueAtAxis(outEle.TitleFileAxis)
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
		err = errors.New(fmt.Sprintf("[%s] GetTitleNamespace Error: {Err=%s,Namespace=%s}", logPrefix, err, namespace))
		return err
	}
	targetDir := Setting.Project.Target.GetTitleDir(titleCtx.RangeName)
	if !filex.IsExist(targetDir) {
		_ = os.MkdirAll(targetDir, os.ModePerm)
	}
	extendName := langDefine.ExtendName
	filePath := filex.Combine(targetDir, fileName+"."+extendName)

	// 创建模板数据代理
	tempDataProxy := &TempTitleProxy{Sheet: sheet, Excel: excel, TitleCtx: titleCtx,
		FileName: fileName, FieldIndex: selects, ClassName: clsName, Namespace: namespace}
	buff := bytes.NewBuffer(nil)
	err = temp.Execute(buff, tempDataProxy, false)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Execute Template error: %s ", logPrefix, err))
		return err
	}
	_ = filex.WriteFile(filePath, buff.Bytes(), os.ModePerm)
	Logger.Infoln(fmt.Sprintf("[%s] \t file => %s", logPrefix, filePath))
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
