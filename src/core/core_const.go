package core

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/ExcelExporter/src/core/temps"
	"github.com/xuzhuoxi/infra-go/filex"
	"io/fs"
	"os"
	"strings"
)

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
			err := errors.New(fmt.Sprintf("-field error at \"%s\": output file name!", constCtx.RangeName))
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] Error at %s ", err))
			return err
		}
		if outEle.Value == "" {
			continue
		}

		clsEle, ok := Setting.Excel.Const.GetClassInfo(constCtx.RangeName)
		if !ok {
			err := errors.New(fmt.Sprintf("-field error at \"%s\": output class name!", constCtx.RangeName))
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] Error at %s ", err))
			return err
		}

		fileName, err := sheet.ValueAtAxis(outEle.Value)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] Get file name error: %s ", err))
			return err
		}
		clsName, err := sheet.ValueAtAxis(clsEle.Value)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] Get class name error: %s ", err))
			return err
		}
		targetDir := Setting.Project.Target.GetConstDir(constCtx.RangeName)
		if !filex.IsExist(targetDir) {
			os.MkdirAll(targetDir, fs.ModePerm)
		}
		extendName := langDefine.ExtendName
		filePath := filex.Combine(targetDir, fileName+"."+extendName)

		// 创建模板数据代理
		tempDataProxy := &TempConstProxy{Sheet: sheet, Excel: excel, ConstCtx: constCtx, FileName: fileName, ClassName: clsName, Language: constCtx.ProgramLanguage,
			StartRow: Setting.Excel.Const.DataStartRow, EndRow: len(sheet.Rows)}

		buff := bytes.NewBuffer(nil)
		err = temp.Execute(buff, tempDataProxy, false)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeConstContext] Execute Template error: %s ", err))
			return err
		}
		os.WriteFile(filePath, buff.Bytes(), fs.ModePerm)
		Logger.Infoln(fmt.Sprintf("[core.executeConstContext] Generate file : %s", filePath))
	}
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
	return nil, errors.New(fmt.Sprintf("Undefined Program Lanaguage for Const: %s", lang))
}
