package leetcode

import (
	"strings"
)

// leetcode #3 Longest Substring Without Repeating Characters
// created by 2021-03-09 by Andy

//Node 節點
type Node struct {
	C    byte
	Next *Node
}

var tmp = make(map[byte]*Node)

func lengthOfLongestSubstring(s string) int {
	max := 0
	tmp = make(map[byte]*Node)
	var left *Node
	var right *Node

	for idx, c := range s {
		if idx == 0 {
			left = &Node{
				C: byte(c),
			}
			right = left
		} else {
			right.Next = &Node{
				C: byte(c),
			}
		}

		if duplicate, ok := tmp[byte(c)]; ok {
			//重複
			if count := countNode(left, right); count > max {
				max = count
			}
			deleteTmp(left, duplicate)
			left = duplicate.Next
			tmp[byte(c)] = right.Next
		} else {
			if idx == 0 {
				tmp[byte(c)] = right
				if idx+1 == strings.Count(s, "")-1 {
					if count := countNode(left, right); count > max {
						max = count
					}
				}
			} else {
				tmp[byte(c)] = right.Next
				if idx+1 == strings.Count(s, "")-1 {
					if count := countNode(left, right.Next); count > max {
						max = count
					}
				}
			}

		}
		if right.Next != nil {
			right = right.Next
		}
	}

	return max
}

func countNode(left *Node, right *Node) int {
	count := 0
	for true {
		count++
		if left.C == right.C {
			break
		}
		left = left.Next
	}

	return count
}

func deleteTmp(left *Node, right *Node) {
	for true {
		delete(tmp, left.C)
		if left.C == right.C {
			break
		}
		left = left.Next
	}
}
