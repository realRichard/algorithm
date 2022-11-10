package str_test

import (
	"algorithm/dataStruct/str"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	result := str.CheckInclusion(s1, s2)
	expected := true
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
