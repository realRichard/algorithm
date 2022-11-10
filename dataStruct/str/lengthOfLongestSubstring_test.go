package str_test

import (
	"algorithm/dataStruct/str"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	expected := 3
	result := str.LengthOfLongestSubstring(s)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
