package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestKThNodeFromEnd(t *testing.T) {
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
	// node := list.KthNodeFromEnd(&list1, 5)
	// if node.Val != 1 {
	// 	t.Errorf("expected 1, got %v", node.Val)
	// }

	list2 := list.RemoveNthFromEnd(&list1, 2)
	values := [4]int{}
	index := 0
	for list2 := list2; list2 != nil; list2 = list2.Next {
		values[index] = list2.Val
		index++
	}
	expected := [4]int{1, 2, 3, 5}
	if values != expected {
		t.Errorf("expected %v, result %v", expected, values)
	}
}
