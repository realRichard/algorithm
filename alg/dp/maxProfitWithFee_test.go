package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestMaxProfitWithFee(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	expected := 8
	result := dp.MaxProfitWithFee(prices, fee)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}

}
