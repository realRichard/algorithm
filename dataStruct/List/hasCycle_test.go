package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestHasCycle(t *testing.T) {
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
	list1.Next.Next.Next.Next.Next = list1.Next.Next
	if !list.HasCycle(&list1) {
		t.Error("expected true, result false")
	}
}

func TestDetectCycle(t *testing.T) {
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
	list1.Next.Next.Next.Next.Next = list1.Next.Next
	node := list.DetectCycle(&list1)
	if node != list1.Next.Next {
		t.Errorf("expected %v, result %v", list1.Next.Next, node)
	}
}
