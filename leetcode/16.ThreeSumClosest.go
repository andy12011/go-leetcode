package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	var closestNum int
	sort.Ints(nums)
	lg := len(nums)

	if lg >2 {
		closestNum = nums[0] + nums[1] + nums[2]
	}

	for i := 0; i < lg-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := lg - 1

		for r > l {
			n := nums[i] + nums[l] + nums[r]
			if n == target {
				closestNum = n
				break

			} else if n > target {
				x := (target - n)
				x2 := (target - closestNum)
				if x*x < x2*x2 {
					closestNum = n
				}
				r--
			} else {
				x := (target - n)
				x2 := (target - closestNum)
				if x*x < x2*x2 {
					closestNum = n
				}
				l++
			}
		}
	}

	return closestNum
}