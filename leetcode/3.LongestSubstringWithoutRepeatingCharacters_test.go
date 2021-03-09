package leetcode

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_01", args{"abcabcbb"}, 3},
		{"test_02", args{"bbbbb"}, 1},
		{"test_03", args{"pwwkew"}, 3},
		{"test_04", args{"gsqygebs"}, 6},
		{"test_05", args{"tmmzuxt"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
