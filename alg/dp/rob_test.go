package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := 4
	result := dp.RobV4(nums)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
