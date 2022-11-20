package list_test

import (
	list "algorithm/dataStruct/List"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	list1 := list.ListNode{
		Val: 1,
	}
	list1.Next = &list.ListNode{
		Val: 1,
	}
	list1.Next.Next = &list.ListNode{
		Val: 2,
	}
	list1.Next.Next.Next = &list.ListNode{
		Val: 1,
	}
	result := list.IsPalindromeV2(&list1)
	expected := false
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
