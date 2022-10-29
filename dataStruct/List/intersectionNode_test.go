package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	list1 := list.ListNode{
		Val: 4,
	}
	list1.Next = &list.ListNode{
		Val: 1,
	}

	list2 := list.ListNode{
		Val: 5,
	}
	list2.Next = &list.ListNode{
		Val: 6,
	}
	list2.Next.Next = &list.ListNode{
		Val: 1,
	}

	list3 := list.ListNode{
		Val: 8,
	}
	list3.Next = &list.ListNode{
		Val: 4,
	}
	list3.Next.Next = &list.ListNode{
		Val: 5,
	}
	list1.Next.Next = &list3
	list2.Next.Next.Next = &list3
	intersectionNode := list.GetIntersectionNodeV2(&list1, &list2)
	if *intersectionNode != list3 {
		t.Fatal()
	}
}
