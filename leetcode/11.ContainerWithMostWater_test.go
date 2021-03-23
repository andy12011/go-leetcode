package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"test02", args{[]int{1, 1}}, 1},
		{"test03", args{[]int{4, 3, 2, 1, 4}}, 16},
		{"test04", args{[]int{2, 3, 4, 5, 18, 17, 6}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
