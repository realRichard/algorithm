package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := array.Constructor(nums)
	left, right := 0, 2
	result := numArray.SumRange(left, right)
	expected := 1
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
