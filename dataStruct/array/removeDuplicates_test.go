package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	length := array.RemoveDuplicates(nums)
	expected := []int{1, 2}
	if length != len(expected) {
		t.Fatal()
	}
	for i := 0; i < length-1; i++ {
		if nums[i] != expected[i] {
			t.Fatal()
		}
	}
}
