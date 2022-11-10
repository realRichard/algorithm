package str_test

import (
	"algorithm/dataStruct/str"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	target := "ABC"
	result := str.MinWindow(s, target)
	expected := "BANC"
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
