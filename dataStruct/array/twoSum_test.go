package array_test

import (
	"algorithm/dataStruct/array"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{3, 2, 4}
	target := 6
	result := array.TwoSum(numbers, target)
	values := [2]int{}
	for index, value := range result {
		values[index] = value
	}
	expected := [2]int{1, 2}
	if values != expected {
		t.Errorf("expected %v, result %v", expected, result)
	}

}
