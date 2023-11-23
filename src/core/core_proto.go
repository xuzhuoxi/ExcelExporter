// Package core
// Create on 2023/5/21
// @author xuzhuoxi
package core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/core/temps"
	"github.com/xuzhuoxi/ExcelExporter/src/setting"
	"github.com/xuzhuoxi/infra-go/filex"
	"os"
	"strings"
)

func execExcelProtoContext(excel *excel.ExcelProxy, protoCtx *ProtoContext) error {
	sheets := excel.GetSheets(protoCtx.EnablePrefix)
	if len(sheets) == 0 {
		return nil
	}
	logPrefix := "core.execExcelProtoContext"
	Logger.Infoln(fmt.Sprintf("[%s][--Start ProtoContext]: %s", logPrefix, protoCtx))
	for _, sheet := range sheets {
		err := execSheetProtoContext(excel, sheet, protoCtx)
		if nil != err {
			return err
		}
	}
	Logger.Infoln(fmt.Sprintf("[%s][--Finish ProtoContext]: %s", logPrefix, protoCtx))
	return nil
}

func execSheetProtoContext(excel *excel.ExcelProxy, sheet *excel.ExcelSheet, protoCtx *ProtoContext) error {
	// 过滤Sheet的命名
	if strings.Index(sheet.SheetName, protoCtx.EnablePrefix) != 0 {
		return nil
	}
	logPrefix := "core.execSheetProtoContext"
	lang := protoCtx.Language
	temp, err := getProtoLanguageTemp(lang)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Get lang[%s] temp fail: %s", logPrefix, lang, err))
		return err
	}
	langDefine, ok := Setting.System.FindProgramLanguage(lang)
	if !ok {
		err = errors.New(fmt.Sprintf("[%s] -lang error at %s: lang undefined!", logPrefix, lang))
		return err
	}
	extendName := langDefine.ExtendName
	sheetTitle := getProtoTitle(sheet, Setting.Excel.Proto)
	if !sheetTitle.MatchRange(protoCtx.RangeName) {
		return nil
	}

	tempSheetProxy := &ProtoSheetProxy{Excel: excel, Sheet: sheet, ProtoCtx: protoCtx, Title: sheetTitle}
	items, err := tempSheetProxy.GetItems()
	if nil != err {
		return err
	}
	targetDir := filex.Combine(Setting.Project.Target.GetProtoDir(protoCtx.RangeName), sheetTitle.ExportSubDir)
	if !filex.IsExist(targetDir) {
		_ = os.MkdirAll(targetDir, os.ModePerm)
	}
	for _, item := range items {
		tempProtoProxy := &TempProtoProxy{ProtoItem: item, SheetProxy: tempSheetProxy}
		buff := bytes.NewBuffer(nil)
		err = temp.Execute(buff, tempProtoProxy, false)
		if nil != err {
			err = errors.New(fmt.Sprintf("[%s] Execute Template error: %s ", logPrefix, err))
			return err
		}

		filePath := filex.Combine(targetDir, item.File+"."+extendName)
		_ = filex.WriteFile(filePath, buff.Bytes(), os.ModePerm)
		Logger.Infoln(fmt.Sprintf("[%s] \t file => %s", logPrefix, filePath))
	}
	return nil
}

func getProtoLanguageTemp(lang string) (t *temps.TemplateProxy, err error) {
	if _, ok := ProtoLanguageTemps[lang]; ok {
		return ProtoLanguageTemps[lang], nil
	}
	if l, ok := Setting.System.FindProgramLanguage(lang); ok {
		temp, err := temps.LoadTemplates(l.GetTempsProtoPath())
		if nil != err {
			return nil, err
		}
		ProtoLanguageTemps[lang] = temp

		return temp, nil
	}
	return nil, errors.New(fmt.Sprintf("[core.getProtoLanguageTemp] Undefined Program Lanaguage for Proto: %s", lang))
}

func getProtoTitle(sheet *excel.ExcelSheet, settings setting.Proto) ProtoSheetTitle {
	idDataType, _ := sheet.ValueAtAxis(settings.IdDataType)
	rangeName, _ := sheet.ValueAtAxis(settings.RangeName)
	namespace, _ := sheet.ValueAtAxis(settings.Namespace)
	exportSub, _ := sheet.ValueAtAxis(settings.Export)
	ranges := strings.Split(strings.ToLower(strings.TrimSpace(rangeName)), setting.ParamsSep)
	return ProtoSheetTitle{
		IdDataType:   strings.ToLower(strings.TrimSpace(idDataType)),
		RangeName:    ranges,
		Namespace:    strings.TrimSpace(namespace),
		ExportSubDir: strings.TrimSpace(exportSub),
	}
}
