package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	lg := len(nums)
	for i := 0; i < lg-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := lg - 1

		for r > l {
			n := nums[i] + nums[l] + nums[r]
			if n == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l > r && nums[r] == nums[r+1] {
					r++
				}

			} else if n > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return ans
}