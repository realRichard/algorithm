package find_test

import (
	"algorithm/find"
	"testing"
)

func TestLeftBound(t *testing.T) {
	nums := []int{1, 2, 2, 2, 3}
	target := 2
	expected := 1
	result := find.LeftBound(nums, target)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}

}
