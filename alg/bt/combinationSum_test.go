package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := bt.CombinationSum(candidates, target)
	fmt.Println(result)
}
