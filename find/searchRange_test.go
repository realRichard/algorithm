package find_test

import (
	"algorithm/find"
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	result := find.SearchRange(nums, target)
	fmt.Println(result)
}
