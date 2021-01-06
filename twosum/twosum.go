package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15} //nums 基本範例
	target := 9 //
	ans := twoSum(nums, target)
	fmt.Printf("%v", ans)
}

//TwoSum leetcode
func twoSum(nums []int, target int) []int {
	var targetNum int
	var ans []int
	prevNums := make(map[int]int, len(nums))
	for i, num := range nums {
		if targetIdx, ok := prevNums[num]; ok { //代表有
			ans = []int{targetIdx, i}
			break
		} else { //開始放進去 map
			targetNum = target - num //目標號碼 - 當前號碼 = 需要號碼
			prevNums[targetNum] = i  //紀錄 需要號碼 => index
		}
	}
	return ans
}