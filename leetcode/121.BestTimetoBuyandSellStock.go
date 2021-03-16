package leetcode

//min紀錄value 跟 idx 結果差不多
func maxProfit(prices []int) int {
	ans, midIdx := 0, 0

	for i := 1; i < len(prices); i++ {

		if prices[midIdx] > prices[i] {
			midIdx = i
			continue
		} else if r := prices[i] - prices[midIdx]; r > ans {
			ans = r
		}
	}

	return ans
}
