package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestMaxProfitWithCold(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	expected := 3
	result := dp.MaxProfitWithCold(prices)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}

}
