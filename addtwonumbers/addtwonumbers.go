package main

import "fmt"

//ListNode test
type ListNode struct {
	Val  int
	Next *ListNode
}

var carry = 0

func main() {

	node1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	node2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ans := addTwoNumbers(&node1, &node2)
	var array []int
	hasNext := true
	currentNode := ans
	for hasNext {
		array = append(array, currentNode.Val)
		if currentNode.Next != nil {
			currentNode = currentNode.Next
		} else {
			hasNext = false
		}
	}
	fmt.Printf("%v", reverse(array))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	tmpVal := l1.Val + l2.Val + carry
	if tmpVal >= 10 {
		tmpVal -= 10
		carry = 1
	} else {
		carry = 0
	}
	ans := new(ListNode)
	ans.Val = tmpVal

	if (l1.Next != nil) || (l2.Next != nil) || carry != 0 {
		ans.Next = addTwoNumbers(l1.Next, l2.Next)
	} else {
		ans.Next = nil
	}

	return ans
}

func reverse(arr []int) []int {

	var reverseArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, arr[i])
	}
	return reverseArr
}
