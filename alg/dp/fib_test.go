package dp_test

import (
	"algorithm/alg/dp"
	"testing"
)

func TestFib(t *testing.T) {
	result := dp.FibV4(4)
	expected := 3
	if result != expected {
		t.Errorf("expected %v, result %v", expected, result)
	}
}
