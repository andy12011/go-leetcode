package leetcode

import (
	"strings"
)

// leetcode #3 Longest Substring Without Repeating Characters
// created by 2021-03-09 by Andy

//Node 節點
type Node struct {
	C    string
	Next *Node
}

var tmp = make(map[string]*Node)

func lengthOfLongestSubstring(s string) int {
	max := 0
	tmp = make(map[string]*Node)
	var left *Node
	var set *Node
	var right *Node
	for idx, c := range s {
		if idx == 0 {
			left = &Node{
				C: string(c),
			}
			set = left
			right = left
		} else {
			set.Next = &Node{
				C: string(c),
			}
			right = set
			set = set.Next
		}

		if duplicate, ok := tmp[string(c)]; ok {
			//重複
			if count := countNode(left, right); count > max {
				max = count
			}
			deleteTmp(left, duplicate)
			left = duplicate.Next
			tmp[string(c)] = set
		} else {
			tmp[string(c)] = set
			if idx+1 == strings.Count(s, "")-1 {
				if count := countNode(left, set); count > max {
					max = count
				}
			}
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
