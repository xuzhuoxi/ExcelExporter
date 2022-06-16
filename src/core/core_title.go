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

	Logger.Infoln(fmt.Sprintf("[core.executeTitleContext][Start Execute TitleContext]: %s", titleCtx))
	//prefix := Setting.Excel.Prefix.Data
	prefix := Setting.Excel.TitleData.Prefix
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}

		outEle, ok := Setting.Excel.TitleData.GetOutputInfo(titleCtx.RangeName)
		if !ok {
			err = errors.New(fmt.Sprintf("-field error at \"%s\": output file name!", titleCtx.RangeName))
			//Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Error At %s ", err))
			return err
		}

		clsEle, ok := Setting.Excel.TitleData.GetClassInfo(titleCtx.RangeName)
		if !ok {
			err = errors.New(fmt.Sprintf("-field error at \"%s\": output class name!", titleCtx.RangeName))
			//Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Error At %s ", err))
			return err
		}

		//fieldRangeRow := sheet.GetRowAt(Setting.Excel.Title.FieldRangeRow - 1)
		size := getControlSize(sheet)
		fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
		if nil == fieldRangeRow || fieldRangeRow.Empty() {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Sheet execute pass at '%s' with filed type empty! ", sheet.SheetName))
			continue
		}
		selects, err := parseRangeRow(sheet, fieldRangeRow, uint(titleCtx.RangeType)-1, size)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Parse file type error: %s ", err))
			return err
		}
		if len(selects) == 0 {
			continue
		}

		fileName, err := sheet.ValueAtAxis(outEle.TitleFileName)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] GetTitleFileName error: %s ", err))
			return err
		}
		className, err := sheet.ValueAtAxis(clsEle.Value)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] GetTitleClassName error: %s ", err))
			return err
		}
		targetDir := Setting.Project.Target.GetTitleDir(titleCtx.RangeName)
		if !filex.IsExist(targetDir) {
			os.MkdirAll(targetDir, fs.ModePerm)
		}
		extendName := langDefine.ExtendName
		filePath := filex.Combine(targetDir, fileName+"."+extendName)

		// 创建模板数据代理
		tempDataProxy := &TempTitleProxy{Sheet: sheet, Excel: excel, TitleCtx: titleCtx, FileName: fileName, Index: selects, ClassName: className, Language: titleCtx.ProgramLanguage}

		//fileName, err := sheet.ValueAtAxis(outEle.ClassName)
		buff := bytes.NewBuffer(nil)
		err = temp.Execute(buff, tempDataProxy, false)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeTitleContext] Execute Template error: %s ", err))
			return err
		}
		os.WriteFile(filePath, buff.Bytes(), fs.ModePerm)
		Logger.Infoln(fmt.Sprintf("[core.executeTitleContext] [%s]Generate file: %s", sheet.SheetName, filePath))
	}
	Logger.Infoln(fmt.Sprintf("[core.executeTitleContext][Finish Execute TitleContext]: %s", titleCtx))
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
