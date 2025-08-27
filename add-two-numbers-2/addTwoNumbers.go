/*
Package addtwonumbers2

You are given two non-empty linked lists representing two non-negative integers.
The most significant digit comes first and each of their nodes contains a single
digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number
0 itself.
*/
package addtwonumbers2

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// showLinkedList(head)
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	// showLinkedList(prev)
	return prev
}

func showLinkedList(l *ListNode) {
	for l.Next != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(">")
		}
		l = l.Next
	}
	fmt.Println(l.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Next == nil && l2.Next == nil {
		head := &ListNode{Val: l1.Val + l2.Val}
		if head.Val >= 10 {
			tens, digits := head.Val/10, head.Val%10
			head.Val = digits
			return &ListNode{Val: tens, Next: head}
		}
	}

	l1Reverse := reverseList(l1)
	l2Reverse := reverseList(l2)

	head := &ListNode{Val: 0}
	current := head
	for l1Reverse != nil || l2Reverse != nil {
		_s := 0
		if l1Reverse == nil && l2Reverse != nil {
			_s = l2Reverse.Val
		} else if l1Reverse != nil && l2Reverse == nil {
			_s = l1Reverse.Val
		} else {
			_s = l1Reverse.Val + l2Reverse.Val
		}
		tens, digits := _s/10, _s%10

		current.Val += digits
		if current.Val > 9 {
			t, d := current.Val/10, current.Val%10
			tens += t
			current.Val = d
		}
		current.Next = &ListNode{Val: tens}
		// fmt.Printf("%#v\n", l1Reverse)
		// fmt.Printf("%#v\n", l2Reverse)
		// fmt.Printf("%#v\n", current)
		current = current.Next
		if l1Reverse != nil {
			l1Reverse = l1Reverse.Next
		}
		if l2Reverse != nil {
			l2Reverse = l2Reverse.Next
		}

	}
	_r := reverseList(head)
	for _r.Val == 0 && _r.Next != nil {
		_r = _r.Next
	}
	if _r.Val >= 10 {
		tens, digits := _r.Val/10, _r.Val%10
		_r.Val = digits
		return &ListNode{Val: tens, Next: _r}
	}
	showLinkedList(_r)
	return _r
}
