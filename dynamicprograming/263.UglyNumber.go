package dynamicprograming

func nthUglyNumber(n int) int {
	if n < 1 {
		return 1
	}

	dp := make([]int, n)
	//第一個醜數，即1
	dp[0] = 1
	//三指標初始指向陣列第一個元素
	p2, p3, p5 := 0, 0, 0
	cur := 1

	for cur < n {
		//計算醜數，取最小值
		minVal := min(dp[p2]*2, dp[p3]*3, dp[p5]*5)
		dp[cur] = minVal
		// 每次誰計算出來，誰的指標就後移一位
		if minVal == dp[p2]*2 {
			p2++
		}
		if minVal == dp[p3]*3 {
			p3++
		}
		if minVal == dp[p5]*5 {
			p5++
		}
		cur++
	}
	return dp[cur-1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}

	if b < c {
		return b
	}
	return c
}