package leetcode

import (
	"testing"
)

func Test_isUgly(t *testing.T) {

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{5}, want:true},
		{name: "case2", args: args{7}, want:false},
		{name: "case3", args: args{14}, want:false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
