package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestMoveZores(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	array.MoveZores(nums)
	result := [5]int{}
	for i := 0; i < len(result); i++ {
		result[i] = nums[i]
	}
	expected := [5]int{1, 3, 12, 0, 0}
	if expected != result {
		t.Errorf("expected %v, result %v", expected, result)
	}
}
