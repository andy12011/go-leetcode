package leetcode

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	lg := len(nums)
	if lg < 4 {
		return ans
	}

	sort.Ints(nums)

	for idx := 0; idx < lg-3; idx++ {
		if idx != 0 && nums[idx] == nums[idx-1] {
			continue
		}
		getThreeSum(nums[idx+1:], target-nums[idx], &nums[idx], &ans)
	}
	return ans
}

func getThreeSum(nums []int, target int, start *int, ans *[][]int) {
	lg := len(nums)
	for i := 0; i < lg-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := lg - 1

		for r > l {
			n := nums[i] + nums[l] + nums[r]
			if n == target {
				*ans = append(*ans, []int{*start, nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l > r && nums[r] == nums[r+1] {
					r++
				}

			} else if n > target {
				r--
			} else {
				l++
			}
		}
	}
}