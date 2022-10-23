package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := list.ListNode{
		Val: 1,
	}
	list1.Next = &list.ListNode{
		Val: 2,
	}
	list1.Next.Next = &list.ListNode{
		Val: 4,
	}

	list2 := list.ListNode{
		Val: 1,
	}
	list2.Next = &list.ListNode{
		Val: 3,
	}
	list2.Next.Next = &list.ListNode{
		Val: 4,
	}
	header := list.MergeTwoListV2(&list1, &list2)
	vals := [6]int{}
	index := 0
	for header := header; header != nil; header = header.Next {
		vals[index] = header.Val
		index++
	}
	expected := [6]int{1, 1, 2, 3, 4, 4}
	if vals != expected {
		t.Errorf("expected: %+v, result: %+v", expected, vals)
	}
}
