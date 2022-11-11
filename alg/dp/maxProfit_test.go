package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{2, 4, 1}
	result := dp.MaxProfit(prices)
	expected := 2
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
