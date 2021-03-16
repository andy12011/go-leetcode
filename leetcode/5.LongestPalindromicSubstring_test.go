package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test_01", args{"bbbbbbbbbbaaaaaaaabbbbbbbbbb"}, "bbbbbbbbbbaaaaaaaabbbbbbbbbb"},
		{"test_02", args{"cbbd"}, "bb"},
		{"test_03", args{"a"}, "a"},
		{"test_04", args{"ac"}, "a"},
		{"test_05", args{"cbbc"}, "cbbc"},
		{"test_06", args{"cccccc"}, "cccccc"},
		{"test_07", args{"ccc"}, "ccc"},
		{"test_08", args{"aacabdkacaa"}, "aca"},
		{"test_09", args{"abbcccbbbcaaccbababcbcabca"},"bbcccbb"},
		{"test_10", args{"abb"}, "bb"},
		{"test_10", args{"aba"}, "aba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
