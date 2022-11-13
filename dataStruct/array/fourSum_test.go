package array_test

import (
	"algorithm/dataStruct/array"
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := array.FourSum(nums, target)
	fmt.Println(result)
}
