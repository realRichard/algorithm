package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestPartition(t *testing.T) {
	list1 := list.ListNode{
		Val: 1,
	}
	list1.Next = &list.ListNode{
		Val: 4,
	}
	list1.Next.Next = &list.ListNode{
		Val: 3,
	}
	list1.Next.Next.Next = &list.ListNode{
		Val: 2,
	}
	list1.Next.Next.Next.Next = &list.ListNode{
		Val: 5,
	}
	list1.Next.Next.Next.Next.Next = &list.ListNode{
		Val: 2,
	}

	header := list.Partition(&list1, 3)
	vals := [6]int{}
	index := 0
	for header := header; header != nil; header = header.Next {
		vals[index] = header.Val
		index++
	}
	expected := [6]int{1, 2, 2, 4, 3, 5}
	if vals != expected {
		t.Errorf("expected: %+v, result: %+v", expected, vals)
	}
}
