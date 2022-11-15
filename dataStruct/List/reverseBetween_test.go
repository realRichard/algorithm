package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestReverseBetween(t *testing.T) {
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
	left, right := 2, 4
	node := list.ReverseBetween(&list1, left, right)
	values := [5]int{}
	index := 0
	for node := node; node != nil; node = node.Next {
		values[index] = node.Val
		index++
	}
	expected := [5]int{1, 4, 3, 2, 5}
	if values != expected {
		t.Errorf("expected %v, result %v", expected, values)
	}
}
