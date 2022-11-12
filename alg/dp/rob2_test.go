package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestRob2(t *testing.T) {
	nums := []int{2, 3, 2}
	expected := 3
	result := dp.Rob2(nums)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
