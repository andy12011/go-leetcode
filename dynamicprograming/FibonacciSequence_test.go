package dynamicprograming

import (
	"fmt"
	"testing"
)

func Test_fibonacciSequence(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{len: 0}, want: 1},
		{name: "case2", args: args{len: 1}, want: 1},
		{name: "case3", args: args{len: 2}, want: 2},
		{name: "case4", args: args{len: 3}, want: 3},
		{name: "case5", args: args{len: 4}, want: 5},
		{name: "case6", args: args{len: 5}, want: 8},
		{name: "case7", args: args{len: 6}, want: 13},
		{name: "case8", args: args{len: 7}, want: 21},
		{name: "case9", args: args{len: 8}, want: 34},
		{name: "case10", args: args{len: 9}, want: 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciSequence(tt.args.len); got != tt.want {
				t.Errorf("fibonacciSequence() = %v, want %v", got, tt.want)
			}
		})
	}

	fmt.Println(fibonacciDP)
}
