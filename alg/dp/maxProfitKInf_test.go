package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestMaxProfitKInf(t *testing.T) {
	prices := []int{6, 1, 3, 2, 4, 7}
	expected := 7
	result := dp.MaxProfitKInfV2(prices)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}

}
