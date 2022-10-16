package gexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

// OpenAsColumn2Slice TODO
// 打开Excel
func (xlsx GExcel) OpenAsColumn2Slice() [][]string {
	f, err := excelize.OpenFile(xlsx.FileName)
	if err != nil {
		panic(err)
	}

	return f.GetRows(xlsx.SheetName)
}

// OpenAsColumn2Map TODO
// 打开Excel，需要先执行Save的demo保存一波
/*
	rr := gexcel.GExcel{
		FileName:  "test.xlsx",
		SheetName: "Sheet1",
	}.OpenAsColumn2Map()

	bt, _ := json.Marshal(rr)
	fmt.Println(string(bt))
*/
func (xlsx GExcel) OpenAsColumn2Map() []map[string]string {
	f, err := excelize.OpenFile(xlsx.FileName)
	if err != nil {
		panic(err)
	}

	rows := f.GetRows(xlsx.SheetName)

	result := make([]map[string]string, len(rows))
	for indexRow, row := range rows {
		result[indexRow] = make(map[string]string)
		for col := 0; col < len(row); col++ {
			colName := excelize.ToAlphaString(col + startCol)
			result[indexRow][colName] = row[col]
		}
	}

	return result
}
