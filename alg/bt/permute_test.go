package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	result := bt.Permute(nums)
	fmt.Println(result)
}
