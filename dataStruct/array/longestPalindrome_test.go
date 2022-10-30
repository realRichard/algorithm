package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "aba"
	result := array.IsPalindrome(s)
	if result == false {
		t.Fatal()
	}
}

func TestLongestPalindrome(t *testing.T) {
	s := "cbbd"
	result := array.LongestPalindrome(s)
	expected := "bb"
	if result != expected {
		t.Fatalf("expected %v, result %v", expected, result)
	}
}
