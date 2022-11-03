package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	n := dp.CoinChangeV3(coins, amount)
	expected := 3
	if n != expected {
		t.Errorf("expected %v, result %v", expected, n)
	}

}
