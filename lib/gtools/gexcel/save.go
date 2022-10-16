package gexcel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Save TODO
// 保存Excel
/*
	header := []string{}
	data := [][]string{}
	rowData := []string{}
	for i := 1; i < 10; i++ {
		header = append(header, fmt.Sprintf("第 %d 列", i))
		rowData = append(rowData, fmt.Sprintf("col value %d", i))
	}

	for i := 1; i < 10; i++ {
		data = append(data, rowData)
	}

	xlsx := gexcel.GExcel{
		FileName:  "test.xlsx",
		SheetName: "Sheet1",
		Header:    header,
		Data:      data,
	}
	xlsx.Save()
*/
func (xlsx GExcel) Save() {
	// Create a new sheet.
	sheetName := xlsx.getSheetName()

	f := excelize.NewFile()
	index := f.NewSheet(sheetName)
	rowIndex := startRow
	// save header
	for col, value := range xlsx.Header {
		colName := excelize.ToAlphaString(col + startCol)
		pos := fmt.Sprintf("%s%d", colName, rowIndex)
		f.SetCellValue(sheetName, pos, value)
	}

	// save data
	for _, rowData := range xlsx.Data {
		rowIndex++
		for col, value := range rowData {
			colName := excelize.ToAlphaString(col + startCol)
			cellName := fmt.Sprintf("%s%d", colName, rowIndex)
			f.SetCellValue(sheetName, cellName, value)
		}
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := f.SaveAs(xlsx.FileName)
	if err != nil {
		panic(err)
	}
}
