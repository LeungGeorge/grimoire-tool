package gexcel

import (
	"testing"
)

func TestGExcel_Save(t *testing.T) {
	type fields struct {
		FileName  string
		SheetName string
		Header    []string
		Data      [][]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "save xlsx",
			fields: fields{
				FileName:  "a.xlsx",
				SheetName: defaultSheetName,
				Header:    []string{"a", "b", "c"},
				Data: [][]string{
					{
						"a1",
						"b1",
						"c1",
					},
					{
						"a2",
						"b2",
						"c2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xlsx := GExcel{
				FileName:  tt.fields.FileName,
				SheetName: tt.fields.SheetName,
				Header:    tt.fields.Header,
				Data:      tt.fields.Data,
			}
			xlsx.Save()
		})
	}
}
