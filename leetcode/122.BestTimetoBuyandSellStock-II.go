package leetcode

import "fmt"

// 題目
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit2(prices []int) int {
	ans, midIdx := 0, 0
//7, 1, 5, 3, 6, 4
	for i := 1; i < len(prices); i++ {

		if prices[midIdx] > prices[i] {
			midIdx = i
			continue
			//3, 3, 5, 0, 0, 3, 1, 4
		} else if r := prices[i] - prices[midIdx]; r > 0 {
			fmt.Println(ans, prices[i], prices[midIdx])
			ans += r
			midIdx += 1
		} else if r == 0{
			midIdx += 1
		}
	}

	return ans
}
