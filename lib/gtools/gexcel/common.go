package gexcel

const (
	startRow         = 1
	startCol         = 1
	defaultSheetName = "Sheet1"
)

// GExcel ...
// 自定义Excel
type GExcel struct {
	FileName  string
	SheetName string
	Header    []string
	Data      [][]string
}

// 获取 Excel sheet，默认值为：defaultSheetName
func (xlsx GExcel) getSheetName() string {
	if xlsx.SheetName == "" {
		panic("sheet name is empty.")
	}
	return xlsx.SheetName
}
