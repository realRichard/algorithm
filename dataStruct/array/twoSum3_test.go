package array_test

import (
	"algorithm/dataStruct/array"
	"fmt"
	"testing"
)

func TestTwoSumTarget(t *testing.T) {
	nums := []int{1, 3, 1, 2, 2, 3}
	target := 4
	result := array.TwoSumTarget(nums, target)
	fmt.Println(result)
}
