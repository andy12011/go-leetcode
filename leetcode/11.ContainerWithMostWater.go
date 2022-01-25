package leetcode

// 題目
// https://leetcode.com/problems/container-with-most-water/submissions/
func maxArea(height []int) int {
	ans := 0
	lx := 0
	rx := len(height) - 1

	if rx < 1 {
		return 0
	}
	var tmp int
	for {
		if lx == rx {
			break
		}
		if height[lx] > height[rx] {
			tmp = height[rx] * (rx - lx)
			rx--
		} else {
			tmp = height[lx] * (rx - lx)
			lx++
		}
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}
