package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	array.ReverseString(s)
	if string(s) != "olleh" {
		t.Fatal()
	}
}
