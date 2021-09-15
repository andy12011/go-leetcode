package dynamicprograming

var cutRodDP map[int]int

func init() {
	cutRodDP = make(map[int]int)
	cutRodDP[0] = 0
}

//Bottom-Up 
func cutRod(price [1000]int, len int) int {
	if len <= 0 {
		return 0
	}

	if val, ok := cutRodDP[len]; ok {
		return val
	}

	maxVal := -1

	for i := 0; i < len; i++ {
		maxVal = max(maxVal, price[i]+cutRod(price, len-i-1))
	}

	cutRodDP[len] = maxVal

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
