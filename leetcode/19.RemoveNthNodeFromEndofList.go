package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	right := head
	deletePointer := head
	lastPointer := head
	i := n
	for i > 0 {
		right = right.Next
		i--
	}

	idx := 0
	for {
		if right == nil {
			break
		}
		if idx != 0 {
			lastPointer = lastPointer.Next
		}
		right = right.Next
		deletePointer = deletePointer.Next
		idx++
	}
	if idx == 0 {
		head = head.Next
	} else {
		lastPointer.Next = deletePointer.Next
	}

	return head
}