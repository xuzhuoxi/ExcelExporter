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

var (
	sqlMergeBuff      = bytes.NewBuffer(nil)
	sqlMergeBuffExist = false
)

func writeMergedSql(sqlCtx *SqlContext) {
	if nil != sqlMergeBuff && sqlMergeBuffExist {
		targetDir := Setting.Project.Target.GetSqlDir(sqlCtx.RangeName)
		fileName := "all_merge.sql"
		filePath := filex.Combine(targetDir, fileName)
		logPrefix := "core.writeMergedSql"
		err := filex.WriteFile(filePath, sqlMergeBuff.Bytes(), os.ModePerm)
		Logger.Println()
		if nil != err {
			Logger.Errorln(fmt.Sprintf("[%s] WriteSqlFile error: %s ", logPrefix, err))
		} else {
			Logger.Infoln(fmt.Sprintf("[%s] \t file => %s", logPrefix, filePath))
		}
	}
}

func execExcelSqlContext(excel *excel.ExcelProxy, sqlCtx *SqlContext) (err error) {
	if !sqlCtx.TitleOn && !sqlCtx.DataOn {
		return nil
	}
	logPrefix := "core.execExcelConstContext"
	sheets := excel.GetSheets(sqlCtx.EnablePrefix)
	if len(sheets) == 0 {
		return nil
	}
	Logger.Infoln(fmt.Sprintf("[%s][--Start SqlContext]: %s", logPrefix, sqlCtx))
	for _, sheet := range sheets {
		err := execSheetSqlContext(excel, sheet, sqlCtx)
		if nil != err {
			return err
		}
	}

	Logger.Infoln(fmt.Sprintf("[%s][--Finish SqlContext]: %s", logPrefix, sqlCtx))
	return nil
}

func execSheetSqlContext(excel *excel.ExcelProxy, sheet *excel.ExcelSheet, sqlCtx *SqlContext) (err error) {
	// 过滤Sheet的命名
	if strings.Index(sheet.SheetName, sqlCtx.EnablePrefix) != 0 {
		return nil
	}
	logPrefix := "core.execSheetSqlContext"
	//Logger.Infoln(fmt.Sprintf("[%s][SheetName=%s, FileName=%s]", logPrefix, sheet.SheetName, sheet.FileName()))
	size := getControlSize(sheet)
	fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
	if nil == fieldRangeRow || fieldRangeRow.Empty() {
		Logger.Warnln(fmt.Sprintf("[%s] Ignore[%s] execution for filed type empty!", logPrefix, sheet.SheetName))
		return nil
	}
	selects, _, err := parseRangeRow(sheet, fieldRangeRow, uint(sqlCtx.RangeType)-1, sqlCtx.StartColIndex, size)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Parse Range Row error: %s ", logPrefix, err))
		return err
	}
	if len(selects) == 0 {
		return nil
	}
	targetDir := Setting.Project.Target.GetSqlDir(sqlCtx.RangeName)
	if !filex.IsExist(targetDir) {
		os.MkdirAll(targetDir, os.ModePerm)
	}
	sql := Setting.Excel.TitleData.GetSqlInfo()
	tableName, _ := sheet.ValueAtAxis(sql.TableNameAxis)
	endRow := getSqlDataEndRow(sheet, sqlCtx.StartRowNum)
	tempSqlProxy := &TempSqlProxy{Sheet: sheet, Excel: excel, SqlCtx: sqlCtx,
		TableName: tableName, FieldIndex: selects,
		StartRow: sqlCtx.StartRowNum, EndRow: endRow, StartColIndex: sqlCtx.StartColIndex}
	err1 := execSqlTable(tempSqlProxy, sheet, sqlCtx, targetDir)
	if nil != err1 {
		return err1
	}
	err2 := execSqlData(tempSqlProxy, sheet, sqlCtx, targetDir)
	if nil != err2 {
		return err2
	}
	return nil
}

func execSqlTable(sqlProxy *TempSqlProxy, sheet *excel.ExcelSheet, sqlCtx *SqlContext, targetDir string) (err error) {
	if !sqlCtx.TitleOn {
		return nil
	}
	logPrefix := "core.execSqlTable"
	sql := Setting.Excel.TitleData.GetSqlInfo()
	fileName, err := sheet.ValueAtAxis(sql.FileNameAxis)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Read Table Name Error At [%s]", logPrefix, sql.FileNameAxis))
		return err
	}
	if strings.TrimSpace(fileName) == "" { // 导出文件如果为空，认为忽略导出
		Logger.Traceln(fmt.Sprintf("[%s] Ignore export because the file name is empty. ", logPrefix))
		return nil
	}
	temp, err := getSqlTableTemps()
	if nil != err {
		return err
	}
	extendName := "table.sql"
	filePath := filex.Combine(targetDir, fileName+"."+extendName)
	buff := bytes.NewBuffer(nil)
	err = temp.Execute(buff, sqlProxy, false)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Execute Template error: %s ", logPrefix, err))
		return err
	}
	if sqlCtx.SqlMerge {
		appendDataToMergeBuff(buff.Bytes())
		return nil
	} else {
		return writeDataToFile(buff.Bytes(), filePath)
	}
}

func execSqlData(sqlProxy *TempSqlProxy, sheet *excel.ExcelSheet, sqlCtx *SqlContext, targetDir string) (err error) {
	if !sqlCtx.DataOn {
		return nil
	}
	logPrefix := "core.execSqlData"
	sql := Setting.Excel.TitleData.GetSqlInfo()
	fileName, err := sheet.ValueAtAxis(sql.FileNameAxis)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Read Table Name Error At [%s]", logPrefix, sql.FileNameAxis))
		return err
	}
	if strings.TrimSpace(fileName) == "" { // 导出文件如果为空，认为忽略导出
		Logger.Traceln(fmt.Sprintf("[%s] Ignore export because the file name is empty. ", logPrefix))
		return nil
	}
	temp, err := getSqlDataTemps()
	if nil != err {
		return err
	}
	extendName := "data.sql"
	filePath := filex.Combine(targetDir, fileName+"."+extendName)
	buff := bytes.NewBuffer(nil)
	err = temp.Execute(buff, sqlProxy, false)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] Execute Template error: %s ", logPrefix, err))
		return err
	}
	if sqlCtx.SqlMerge {
		appendDataToMergeBuff(buff.Bytes())
		return nil
	} else {
		return writeDataToFile(buff.Bytes(), filePath)
	}
}

func appendDataToMergeBuff(data []byte) {
	sqlMergeBuff.Write(data)
	sqlMergeBuff.WriteString(fmt.Sprintln())
	sqlMergeBuffExist = true
	logPrefix := "core.appendDataToMergeBuff"
	Logger.Infoln(fmt.Sprintf("[%s] size=%d", logPrefix, len(data)))
}

func writeDataToFile(data []byte, filePath string) error {
	logPrefix := "core.writeDataToFile"
	err := filex.WriteFile(filePath, data, os.ModePerm)
	if nil != err {
		err = errors.New(fmt.Sprintf("[%s] WriteSqlFile error: %s ", logPrefix, err))
		return err
	}
	Logger.Infoln(fmt.Sprintf("[%s] \t file => %s", logPrefix, filePath))
	return nil
}

func getSqlTableTemps() (t *temps.TemplateProxy, err error) {
	if nil != SqlTableTemps {
		return SqlTableTemps, nil
	}
	db, ok := Setting.System.Databases.GetDefaultDatabase()
	if !ok {
		return nil, errors.New("[core.getSqlTableTemps] Default Database Undefined! ")
	}
	temp, err := temps.LoadTemplates(db.GetTempsTablePath())
	if nil != err {
		return nil, err
	}
	SqlTableTemps = temp
	return temp, nil
}

func getSqlDataTemps() (t *temps.TemplateProxy, err error) {
	if nil != SqlDataTemps {
		return SqlDataTemps, nil
	}
	db, ok := Setting.System.Databases.GetDefaultDatabase()
	if !ok {
		return nil, errors.New("[core.getSqlDataTemps] Default Database Undefined! ")
	}
	temp, err := temps.LoadTemplates(db.GetTempsDataPath())
	if nil != err {
		return nil, err
	}
	SqlDataTemps = temp
	return temp, nil
}

func getSqlDataEndRow(sheet *excel.ExcelSheet, startRow int) int {
	row := startRow
	for row > 0 {
		dataRow := sheet.GetRowAt(row - 1)
		if nil == dataRow || len(dataRow.Cell) == 0 || strings.TrimSpace(dataRow.Cell[0]) == "" { // 到达 表尾、空白头
			return row
		}
		row += 1
	}
	return row
}
