package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_01", args{"42"}, 42},
		{"test_02", args{"   -42"}, -42},
		{"test_03", args{"4193 with words"}, 4193},
		{"test_04", args{"words and 987"}, 0},
		{"test_05", args{"-91283472332"}, -2147483648},
		{"test_06", args{"+-12"}, 0},
		{"test_07", args{"00000-42a1234"}, 0},
		{"test_08", args{"+1"}, 1},
		{"test_09", args{"0000000000012345678"}, 12345678},
		{"test_10", args{" +1"}, 1},
		{"test_11", args{"  0000000000012345678"}, 12345678},
		{"test_12", args{"  +  413"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
