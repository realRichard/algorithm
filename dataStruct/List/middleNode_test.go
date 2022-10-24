package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	list1 := list.ListNode{
		Val: 1,
	}
	list1.Next = &list.ListNode{
		Val: 2,
	}
	list1.Next.Next = &list.ListNode{
		Val: 3,
	}
	list1.Next.Next.Next = &list.ListNode{
		Val: 4,
	}
	list1.Next.Next.Next.Next = &list.ListNode{
		Val: 5,
	}
	// list1.Next.Next.Next.Next.Next = &list.ListNode{
	// 	Val: 6,
	// }
	node := list.MiddleNode(&list1)
	values := [3]int{}
	index := 0
	for node := node; node != nil; node = node.Next {
		values[index] = node.Val
		index++
	}
	expected := [3]int{3, 4, 5}
	if values != expected {
		t.Errorf("expected %v, result %v", expected, values)
	}
}
