package find_test

import (
	"algorithm/find"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4
	result := find.Search(nums, target)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
