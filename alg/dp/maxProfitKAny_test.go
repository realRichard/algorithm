package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestMaxProfitKAny(t *testing.T) {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	result := dp.MaxProfitKAny(k, prices)
	expected := 7
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
