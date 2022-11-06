package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	res := bt.Subsets(nums)
	fmt.Println(res)
}
