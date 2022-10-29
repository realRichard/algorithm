package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	list1 := list.ListNode{
		Val: 1,
	}
	list1.Next = &list.ListNode{
		Val: 1,
	}
	list1.Next.Next = &list.ListNode{
		Val: 2,
	}
	head := list.DeleteDuplicates(&list1)
	values := [2]int{}
	index := 0
	for head := head; head != nil; head = head.Next {
		values[index] = head.Val
		index++
	}
	expected := [2]int{1, 2}
	if values != expected {
		t.Errorf("expected %v, result %v\n", expected, values)
	}
}
