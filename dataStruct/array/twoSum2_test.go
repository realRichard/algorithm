package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := array.TwoSum2(numbers, target)
	values := [2]int{}
	for index, value := range result {
		values[index] = value
	}
	expected := [2]int{1, 2}
	if values != expected {
		t.Errorf("expected %v, result %v", expected, result)
	}

}
