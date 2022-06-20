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

func executeDataContext(excel *excel.ExcelProxy, dataCtx *DataContext) error {
	//prefix := Setting.Excel.Prefix.Data
	prefix := Setting.Excel.TitleData.Prefix
	Logger.Infoln(fmt.Sprintf("[core.executeDataContext][Start Execute DataContext]: %s", dataCtx))
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		//Logger.Infoln(fmt.Sprintf("[core.executeDataContext] Sheet[%s]", sheet.SheetName))
		//outEle, ok := Setting.Excel.Output.GetElement(dataCtx.RangeName)
		outEle, ok := Setting.Excel.TitleData.GetOutputInfo(dataCtx.RangeName)
		if !ok {
			err := errors.New(fmt.Sprintf("-field error at \"%s\"", dataCtx.RangeName))
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Error A %s ", err))
			return err
		}
		//fieldRangeRow := sheet.GetRowAt(Setting.Excel.Title.FieldRangeRow - 1)
		size := getControlSize(sheet)
		fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
		if nil == fieldRangeRow || fieldRangeRow.Empty() {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Sheet execute pass at '%s' with filed type empty! ", sheet.SheetName))
			continue
		}
		selects, err := parseRangeRow(sheet, fieldRangeRow, uint(dataCtx.RangeType)-1, size)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] Parse file type error: %s ", err))
			return err
		}
		if len(selects) == 0 {
			continue
		}

		fileName, err := sheet.ValueAtAxis(outEle.DataFileName)
		if nil != err || strings.TrimSpace(fileName) == "" {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] GetDataFileName Error: {Err=%s,FileName=%s}", err, fileName))
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
			if nil == dataRow || len(dataRow.Cell) == 0 || strings.TrimSpace(dataRow.Cell[0]) == "" { // 到达 表尾、空白头
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
		Logger.Infoln(fmt.Sprintf("[core.executeDataContext] [%s]Generate file: %s", sheet.SheetName, filePath))
		err = builder.WriteDataToFile(filePath)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeDataContext] WriteDataFile error: %s ", err))
		}
	}
	Logger.Infoln(fmt.Sprintf("[core.executeDataContext][Finish Execute DataContext]: %s", dataCtx))
	return nil
}
