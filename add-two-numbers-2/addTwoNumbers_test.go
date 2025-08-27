package addtwonumbers2

import (
	"testing"
)

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func comapareLists(list1, list2 *ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next

	}
	return list1 == nil && list2 == nil
}

func TestAddTwoNumbers2(t *testing.T) {
	data := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			"Ch.1112",
			createLinkedList([]int{8, 9, 9}),
			createLinkedList([]int{2}),
			createLinkedList([]int{9, 0, 1}),
		},
		{
			"Ch.1062",
			createLinkedList([]int{1}),
			createLinkedList([]int{9, 9}),
			createLinkedList([]int{1, 0, 0}),
		},
		{
			"Ch.1059",
			createLinkedList([]int{5}),
			createLinkedList([]int{5}),
			createLinkedList([]int{1, 0}),
		},
		{
			"Ch.1",
			createLinkedList([]int{7, 2, 4, 3}),
			createLinkedList([]int{5, 6, 4}),
			createLinkedList([]int{7, 8, 0, 7}),
		},
		{
			"Ch.2",
			createLinkedList([]int{2, 4, 3}),
			createLinkedList([]int{5, 6, 4}),
			createLinkedList([]int{8, 0, 7}),
		},
		{
			"Ch.3",
			createLinkedList([]int{0}),
			createLinkedList([]int{0}),
			createLinkedList([]int{0}),
		},
		{
			"Ch.3",
			createLinkedList([]int{0}),
			createLinkedList([]int{1}),
			createLinkedList([]int{1}),
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := addTwoNumbers(d.l1, d.l2)
			if !comapareLists(d.expected, resp) {
				t.Errorf("⚠️ Test `%s` failed!", d.name)
			}
		})
	}
}
