package leetcode

import "testing"

//
func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test_01", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"test_02", args{"A", 2}, "A"},
		{"test_03", args{"", 2}, ""},
		{"test_02", args{"AB", 3}, "AB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
