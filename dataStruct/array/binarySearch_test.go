package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	index := array.BinarySearch(nums, target)
	if index != 4 {
		t.Fatal()
	}
}
