package gexcel

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestGExcel_OpenAsColumn2Map(t *testing.T) {
	type fields struct {
		FileName  string
		SheetName string
		Header    []string
		Data      [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   []map[string]string
	}{
		{
			name: "save xlsx",
			fields: fields{
				FileName:  "a.xlsx",
				SheetName: defaultSheetName,
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
			got := xlsx.OpenAsColumn2Map()
			sGot, _ := jsoniter.MarshalToString(got)
			fmt.Println(sGot)
		})
	}
}

func TestGExcel_OpenAsColumn2Slice(t *testing.T) {
	type fields struct {
		FileName  string
		SheetName string
		Header    []string
		Data      [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]string
	}{
		{
			name: "save xlsx",
			fields: fields{
				FileName:  "a.xlsx",
				SheetName: defaultSheetName,
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
			got := xlsx.OpenAsColumn2Slice()
			sGot, _ := jsoniter.MarshalToString(got)
			fmt.Println(sGot)
		})
	}
}
