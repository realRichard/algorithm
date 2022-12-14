package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	candidates := []int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34, 16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16, 8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12}
	result := bt.CombinationSum2(candidates, 27)
	fmt.Println(result)
}
