package gexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

// OpenAsColumn2Slice ...
// 打开Excel
func (xlsx GExcel) OpenAsColumn2Slice() [][]string {
	f, err := excelize.OpenFile(xlsx.FileName)
	if err != nil {
		panic(err)
	}

	rows, err := f.GetRows(xlsx.SheetName)
	if err != nil {
		panic(err)
	}

	return rows
}

// OpenAsColumn2Map ...
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

	rows, err := f.GetRows(xlsx.SheetName)
	if err != nil {
		panic(err)
	}

	result := make([]map[string]string, len(rows))
	for indexRow, row := range rows {
		result[indexRow] = make(map[string]string)
		for col := 0; col < len(row); col++ {
			colName, err := excelize.ColumnNumberToName(col + startCol)
			if err != nil {
				panic(err)
			}
			result[indexRow][colName] = row[col]
		}
	}

	return result
}
