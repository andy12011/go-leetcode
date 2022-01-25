package leetcode

func twoSum(nums []int, target int) []int {
	var ans []int
	prevNums := make(map[int]int, len(nums))
	for i, num := range nums {
		if targetIdx, ok := prevNums[num]; ok {
			ans = []int{targetIdx, i}
			break
		} else {
			prevNums[target - num] = i
		}
	}
	return ans
}