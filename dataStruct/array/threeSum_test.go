package array_test

import (
	"algorithm/dataStruct/array"
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := array.ThreeSum(nums)
	fmt.Println(result)
}
