package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestPermuteRepeat(t *testing.T) {
	nums := []int{1, 2}
	result := bt.PermuteRepeat(nums)
	fmt.Println(result)
}
