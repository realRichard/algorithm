package bt_test

import (
	"algorithm/alg/bt"
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	res := bt.SubsetsWithDup(nums)
	fmt.Println(res)
}
