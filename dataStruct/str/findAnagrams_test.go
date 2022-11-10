package str_test

import (
	"algorithm/dataStruct/str"
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	s := "abab"
	p := "ab"
	result := str.FindAnagrams(s, p)
	fmt.Println(result)
}
