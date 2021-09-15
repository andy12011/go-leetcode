package dynamicprograming

// Base case: F(0) = F(1) = 1
// Recursive case: F(n) = F(n-1) + F(n-2)

var fibonacciDP map[int]int

func init() {
	fibonacciDP = make(map[int]int)
}

func fibonacciSequence(len int) int {
	return recursiveCal(len)
}

func recursiveCal(n int) int {
	// another solution is define default in dp map
	// dp[0] = 1
	// dp[1] = 1
	if n <= 1 {
		if _, ok := fibonacciDP[n]; !ok {
			fibonacciDP[n] = 1
		}
		return fibonacciDP[n]
	}

	if _, ok := fibonacciDP[n]; !ok {
		fibonacciDP[n] = recursiveCal(n-1) + recursiveCal(n-2)
	}

	return fibonacciDP[n]
}
