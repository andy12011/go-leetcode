package leetcode

var params = []int{2, 3, 5}

func isUgly(n int) bool {
	if n < 1 {
		return false
	}

	for _, p := range params {
		for (n % p) == 0 {
			n /= p
		}
	}

	return n == 1
}
