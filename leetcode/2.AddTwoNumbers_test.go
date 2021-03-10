package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test_01",
			args{BuildList([]int{0}), BuildList([]int{0})},
			BuildList([]int{0}),
		},
		{
			"test_02",
			args{BuildList([]int{2, 4, 3}), BuildList([]int{5, 6, 4})},
			BuildList([]int{7, 0, 8}),
		},
		{
			"test_02",
			args{BuildList([]int{9, 9, 9, 9, 9, 9, 9}), BuildList([]int{9, 9, 9, 9})},
			BuildList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BuildList(nums []int) *ListNode {
	var result *ListNode
	var currentNode *ListNode
	for idx, val := range nums {
		tempNode := ListNode{Val: val, Next: nil}
		if idx == 0 {
			result = &tempNode
			currentNode = &tempNode
		} else {
			currentNode.Next = &tempNode
			currentNode = currentNode.Next
		}
	}
	return result
}
