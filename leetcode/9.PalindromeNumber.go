package leetcode

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		s := fmt.Sprintf("%d", x)
		head := 0
		tail := len(s) - 1

		for true {
			if s[tail] == s[head] {
				if tail == head || tail+1 == head {
					return true
				}
				tail--
				head++
			} else {
				return false
			}
		}

		return true
	}
}
