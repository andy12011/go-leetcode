package main

import "fmt"

//ListNode test
type ListNode struct {
	Val  int
	Next *ListNode
}

var carry = 0

func main() {

	// node1 := ListNode{
	// 	Val: 9,
	// 	Next: &ListNode{
	// 		Val: 9,
	// 		Next: &ListNode{
	// 			Val:  9,
	// 			Next: nil,
	// 		},
	// 	},
	// }

	node1 := ListNode{
		Val: 0,
		Next: nil,
	}

	node2 := ListNode{
		Val: 1,
		Next: nil,
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
	fmt.Printf("%v", array)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 != nil && l2 != nil {
		ans := new(ListNode)
		ans.Val = checkCarry(l1.Val + l2.Val + carry)
		ans.Next = addTwoNumbers(l1.Next, l2.Next)
		return ans
	} else if l1 != nil {
		ans := new(ListNode)
		ans.Val = checkCarry(l1.Val + carry)
		ans.Next = addTwoNumbers(l1.Next, nil)
		return ans
	} else if l2 != nil {
		ans := new(ListNode)
		ans.Val = checkCarry(l2.Val + carry)
		ans.Next = addTwoNumbers(nil, l2.Next)
		return ans
	} else if carry != 0 {
		ans := new(ListNode)
		ans.Val = carry
		ans.Next = nil
		return ans
	} else {
		return nil
	}
}

func reverse(arr []int) []int {

	var reverseArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, arr[i])
	}
	return reverseArr
}

func checkCarry(num int) int {
	if num >= 10 {
		carry = 1
		return num - 10
	} else {
		carry = 0
		return num
	}
}
