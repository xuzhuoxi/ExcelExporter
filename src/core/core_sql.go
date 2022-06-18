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
	sqlBuffMerge      = bytes.NewBuffer(nil)
	sqlBuffMergeExist = false
)

func writeSqlMergeContext(sqlCtx *SqlContext) {
	if nil != sqlBuffMerge && sqlBuffMergeExist {
		targetDir := Setting.Project.Target.GetSqlDir(sqlCtx.RangeName)
		fileName := "all_merge.sql"
		filePath := filex.Combine(targetDir, fileName)
		filex.WriteFile(filePath, sqlBuffMerge.Bytes(), os.ModePerm)
	}
}

func executeSqlContext(excel *excel.ExcelProxy, sqlCtx *SqlContext) (err error) {
	if !sqlCtx.TitleOn && !sqlCtx.DataOn {
		return nil
	}
	Logger.Infoln(fmt.Sprintf("[core.executeSqlContext][Start Execute SqlContext]: %s", sqlCtx))
	prefix := Setting.Excel.TitleData.Prefix
	for _, sheet := range excel.Sheets {
		// 过滤Sheet的命名
		if strings.Index(sheet.SheetName, prefix) != 0 {
			continue
		}
		size := getControlSize(sheet)
		fieldRangeRow := sheet.GetRowAt(Setting.Excel.TitleData.FieldRangeRow - 1)
		if nil == fieldRangeRow || fieldRangeRow.Empty() {
			Logger.Warnln(fmt.Sprintf("[core.executeSqlContext] Sheet execute pass at '%s' with filed type empty! ", sheet.SheetName))
			continue
		}
		selects, err := parseRangeRow(sheet, fieldRangeRow, uint(sqlCtx.RangeType)-1, size)
		if nil != err {
			Logger.Warnln(fmt.Sprintf("[core.executeSqlContext] Parse file type error: %s ", err))
			return err
		}
		if len(selects) == 0 {
			continue
		}
		targetDir := Setting.Project.Target.GetSqlDir(sqlCtx.RangeName)
		if !filex.IsExist(targetDir) {
			os.MkdirAll(targetDir, os.ModePerm)
		}
		sql := Setting.Excel.TitleData.GetSqlInfo()
		tableName, _ := sheet.ValueAtAxis(sql.TableNameAxis)
		startRow := Setting.Excel.TitleData.DataStartRow
		endRow := getSqlDataEndRow(sheet, startRow)
		tempSqlProxy := &TempSqlProxy{Sheet: sheet, Excel: excel, SqlCtx: sqlCtx,
			TableName: tableName, FieldIndex: selects,
			StartRow: startRow, EndRow: endRow}
		execSqlTable(tempSqlProxy, sheet, sqlCtx, targetDir)
		execSqlData(tempSqlProxy, sheet, sqlCtx, targetDir)
	}

	Logger.Infoln(fmt.Sprintf("[core.executeSqlContext][Finish Execute SqlContext]: %s", sqlCtx))
	return nil
}

func execSqlTable(sqlProxy *TempSqlProxy, sheet *excel.ExcelSheet, sqlCtx *SqlContext, targetDir string) (err error) {
	if !sqlCtx.TitleOn {
		return nil
	}
	sql := Setting.Excel.TitleData.GetSqlInfo()
	fileName, err := sheet.ValueAtAxis(sql.FileNameAxis)
	if nil != err {
		return err
	}
	if len(fileName) == 0 {
		err = errors.New(fmt.Sprintf("Table Name Empty At [%s]", sql.FileNameAxis))
		return err
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
		Logger.Warnln(fmt.Sprintf("[core.executeSqlContext] Execute Template error: %s ", err))
		return err
	}
	if sqlCtx.SqlMerge {
		appendDataToBuffMerge(buff.Bytes())
	} else {
		filex.WriteFile(filePath, buff.Bytes(), os.ModePerm)
	}
	return nil
}

func execSqlData(sqlProxy *TempSqlProxy, sheet *excel.ExcelSheet, sqlCtx *SqlContext, targetDir string) (err error) {
	if !sqlCtx.DataOn {
		return nil
	}
	sql := Setting.Excel.TitleData.GetSqlInfo()
	fileName, err := sheet.ValueAtAxis(sql.FileNameAxis)
	if nil != err {
		return err
	}
	if len(fileName) == 0 {
		err = errors.New(fmt.Sprintf("Data File Name Empty At [%s]", sql.FileNameAxis))
		return err
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
		Logger.Warnln(fmt.Sprintf("[core.executeSqlContext] Execute Template error: %s ", err))
		return err
	}
	if sqlCtx.SqlMerge {
		appendDataToBuffMerge(buff.Bytes())
	} else {
		filex.WriteFile(filePath, buff.Bytes(), os.ModePerm)
	}
	return nil
}

func appendDataToBuffMerge(data []byte) {
	sqlBuffMerge.Write(data)
	sqlBuffMergeExist = true
}

func getSqlTableTemps() (t *temps.TemplateProxy, err error) {
	if nil != SqlTableTemps {
		return SqlTableTemps, nil
	}
	db, ok := Setting.System.Databases.GetDefaultDatabase()
	if !ok {
		return nil, errors.New("Default Database Undefined! ")
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
		return nil, errors.New("Default Database Undefined! ")
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
