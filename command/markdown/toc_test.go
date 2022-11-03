package markdown

import "testing"

func Test_listFiles(t *testing.T) {
	type args struct {
		curPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "curPath",
			args: args{
				curPath: "/Users/yuanzhengliang/Documents/智影/腾讯学院",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listFiles(tt.args.curPath)
		})
	}
}
