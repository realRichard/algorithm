package array_test

import (
	"algorithm/dataStruct/array"
	"fmt"
	"testing"
)

func TestNSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	n := 4
	target := 0
	result := array.NSum(nums, n, target)
	fmt.Println(result)
}
