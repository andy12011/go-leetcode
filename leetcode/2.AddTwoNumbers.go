package leetcode

import "fmt"

//ListNode Node結構
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode
	var head *ListNode
	carry := 0
	list1 := true
	list2 := true
	ans = &ListNode{Val: 0}
	for true {
		if head == nil {
			head = ans
		}
		if list1 {
			ans.Val += l1.Val
		}
		if list2 {
			ans.Val += l2.Val
		}

		if carry == 1 {
			ans.Val += carry
			carry = 0
		}

		if ans.Val >= 10 {
			carry = 1
			ans.Val -= 10
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			list1 = false
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			list2 = false
		}
		if !list2 && !list1 {
			if carry == 1 {
				ans.Next = &ListNode{Val: 1}
			}
			break;
		}
		ans.Next = &ListNode{Val: 0}
		ans = ans.Next
	}

	return head
}

func show(head *ListNode)  {
	str := ""
	for true {
		str += fmt.Sprintf("%d", head.Val)
		if head.Next != nil {
			head = head.Next
		} else {
			break;
		}
	}
	fmt.Println(str)
}

// func buildList(nums []int) *ListNode {
// 	var result *ListNode
// 	var currentNode *ListNode
// 	for idx, val := range nums {
// 		tempNode := ListNode{Val: val, Next: nil}
// 		if idx == 0 {
// 			result = &tempNode
// 			currentNode = &tempNode
// 		} else {
// 			currentNode.Next = &tempNode
// 			currentNode = currentNode.Next
// 		}
// 	}
// 	return result
// }
