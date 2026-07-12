package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list13 := &ListNode{
		Val:  4,
		Next: nil,
	}
	list12 := &ListNode{
		Val:  2,
		Next: list13,
	}
	list11 := &ListNode{
		Val:  1,
		Next: list12,
	}

	list23 := &ListNode{
		Val:  4,
		Next: nil,
	}
	list22 := &ListNode{
		Val:  3,
		Next: list23,
	}
	list21 := &ListNode{
		Val:  1,
		Next: list22,
	}

	showList(mergeTwoLists(list11, list21))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return dummyHead.Next
}

func showList(list *ListNode) {
	current := list

	for current.Next != nil {
		fmt.Println(current)
		current = current.Next
	}
}
