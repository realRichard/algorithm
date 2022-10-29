package array_test

import (
	"algorithm/dataStruct/array"
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	length := array.RemoveElement(nums, val)
	fmt.Println(nums, length)
	for i := 0; i < length-1; i++ {
		if nums[i] == val {
			t.Fatal()
		}
	}
}
