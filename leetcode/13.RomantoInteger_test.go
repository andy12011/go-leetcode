package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{s: "MMM"}, want: 3000},
		{name: "case2", args: args{s: "III"}, want: 3},
		{name: "case3", args: args{s: "MCM"}, want: 1900},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
