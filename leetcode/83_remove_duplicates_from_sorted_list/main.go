package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list5 := &ListNode{
		Val:  3,
		Next: nil,
	}
	list4 := &ListNode{
		Val:  3,
		Next: list5,
	}
	list3 := &ListNode{
		Val:  2,
		Next: list4,
	}
	list2 := &ListNode{
		Val:  1,
		Next: list3,
	}
	list1 := &ListNode{
		Val:  1,
		Next: list2,
	}

	showList(deleteDuplicates(list1))
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func showList(list *ListNode) {
	current := list

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}
