package t21

import (
	"fmt"
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
		fmt.Printf("l1.v = %d, l2.v = %d\n", list1.Val, list2.Val)
		if list1.Val != list2.Val {
			fmt.Println(list1.Val, list2.Val)
			return false
		}
		list1 = list1.Next
		list2 = list2.Next

	}
	return list1 == nil && list2 == nil
}

func TestMergeTwoLists(t *testing.T) {
	data := []struct {
		name     string
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			"Ch.1",
			createLinkedList([]int{1, 2, 4}),
			createLinkedList([]int{1, 3, 4}),
			createLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			"Ch.2",
			createLinkedList([]int{2}),
			createLinkedList([]int{1}),
			createLinkedList([]int{1, 2}),
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := mergeTwoLists(d.list1, d.list2)
			if !comapareLists(d.expected, resp) {
				t.Errorf("⚠️ Test `%s` failed!", d.name)
			}
		})
	}
}
