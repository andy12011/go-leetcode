package dynamicprograming

// Base case: F(0) = F(1) = 1
// Recursive case: F(n) = F(n-1) + F(n-2)

var dp map[int]int

func init() {
	dp = make(map[int]int)
}

func fibonacciSequence(len int) int {
	return recursiveCal(len)
}

func recursiveCal(n int) int {
	// another solution is define default in dp map
	// dp[0] = 1
	// dp[1] = 1
	if n <= 1 {
		if _, ok := dp[n]; !ok {
			dp[n] = 1
		}
		return dp[n]
	}

	if _, ok := dp[n]; !ok {
		dp[n] = recursiveCal(n-1) + recursiveCal(n-2)
	}

	return dp[n]
}
