package core

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/ExcelExporter/src/core/data"
	"github.com/xuzhuoxi/ExcelExporter/src/core/excel"
	"github.com/xuzhuoxi/infra-go/filex"
	"os"
	"strings"
)

func execExcelDataContext(excel *excel.ExcelProxy, dataCtx *DataContext) error {
	sheets := excel.GetSheets(dataCtx.EnablePrefix)
	if len(sheets) == 0 {
		return nil
	}

	logPrefix := "core.execExcelTitleContext"
	Logger.Infoln(fmt.Sprintf("[%s][--Start DataContext]: %s", logPrefix, dataCtx))
	for _, sheet := range sheets {
		err := execSheetDataContext(excel, sheet, dataCtx)
		if nil != err {
			return err
		}
	}
	Logger.Infoln(fmt.Sprintf("[%s][--Finish DataContext]: %s", logPrefix, dataCtx))
	return nil
}

func execSheetDataContext(excel *excel.ExcelProxy, sheet *excel.ExcelSheet, dataCtx *DataContext) (err error) {
	// 过滤Sheet的命名
	if strings.Index(sheet.SheetName, dataCtx.EnablePrefix) != 0 {
		return nil
	}
	logPrefix := "core.execSheetDataContext"
	//Logger.Infoln(fmt.Sprintf("[%s][SheetName=%s, FileName=%s]", logPrefix, sheet.SheetName, sheet.FileName()))
	outEle, ok := Setting.Excel.TitleData.GetOutputInfo(dataCtx.RangeName)
	if !ok {
		err = errors.New(fmt.Sprintf("[%s] -field error at \"%s\": output file name!", logPrefix, dataCtx.RangeName))
		return err
	}
	size := getControlSize(sheet)
	fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
	if nil == fieldRangeRow || fieldRangeRow.Empty() {
		Logger.Warnln(fmt.Sprintf("[%s] Sheet execute pass at '%s' with filed type empty! ", logPrefix, sheet.SheetName))
		return nil
	}
	selects, selectNames, err := parseRangeRow(sheet, fieldRangeRow, uint(dataCtx.RangeType)-1, dataCtx.StartColIndex, size)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Parse Range Row error: %s ", logPrefix, err))
		return err
	}
	if len(selects) == 0 {
		return nil
	}

	fileName, err := sheet.ValueAtAxis(outEle.DataFileAxis)
	if nil != err || strings.TrimSpace(fileName) == "" {
		err = errors.New(fmt.Sprintf("[%s] GetTitleFileName Error: {Err=%s,FileName=%s}", logPrefix, err, fileName))
		return err
	}
	keyRowNum := Setting.Excel.TitleData.GetFileKeyRow(dataCtx.DataFileFormat)
	if -1 == keyRowNum {
		Logger.Warnln(fmt.Sprintf("[%s] Parse file format: %s ", logPrefix, dataCtx.DataFileFormat))
		return nil
	}
	keyRow := sheet.GetRowAt(keyRowNum - 1)
	//typeRow := sheet.GetRowAt(Setting.Excel.Title.FieldFormatRow - 1)
	typeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldFormatRow - 1)
	//startRow := Setting.Excel.Data.StartRow
	startRow := dataCtx.StartRowNum
	builder := data.GenBuilder(dataCtx.DataFileFormat)
	builder.StartWriteData()
	for startRow > 0 {
		dataRow := sheet.GetRowAt(startRow - 1)
		if nil == dataRow || len(dataRow.Cell) == 0 || strings.TrimSpace(dataRow.Cell[0]) == "" { // 到达 表尾、空白头
			break
		}
		ktvRow := getRowData(keyRow, typeRow, dataRow, selects, selectNames, startRow)
		err = builder.WriteRow(ktvRow)
		if nil != err {
			err = errors.New(fmt.Sprintf("[%s] Builder WriteRow error: %s ", logPrefix, err))
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
	err = builder.WriteDataToFile(filePath)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] WriteDataFile error: %s ", logPrefix, err))
		return err
	}
	Logger.Infoln(fmt.Sprintf("[%s] \t file => %s", logPrefix, filePath))
	return nil
}
