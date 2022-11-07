package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	nums := []int{2, 2, 2}
	result := bt.PermuteUnique(nums)
	fmt.Println(result)
}
