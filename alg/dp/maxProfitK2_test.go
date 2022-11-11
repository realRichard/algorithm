package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestMaxProfitK2(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	result := dp.MaxProfitK2(prices)
	expected := 6
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
