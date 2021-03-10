package leetcode

// leetcode #3 Longest Substring Without Repeating Characters
// created by 2021-03-09 by Andy

//Node 節點
type Node struct {
	C    byte
	Next *Node
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)      //字串長度
	m := [256]byte{} //8位元 2^8
	left, right := 0, 0
	counter := 0
	ans := 0
	for right < n {
		idx := int(s[right])
		if m[idx] == 0 {
			counter++
		}
		m[idx]++
		right++

		for counter < (right - left) { //代表重複
			m[int(s[left])]--
			if m[int(s[left])] == 0 {
				counter--
			}
			left++
		}
		if ans < (right - left) {
			ans = right - left
		}
	}
	return ans
}
