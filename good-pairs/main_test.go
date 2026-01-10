package main

import "testing"

func TestGoodPairs(t *testing.T) {
	result1 := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	result2 := numIdenticalPairs([]int{1, 1, 1, 1})
	result3 := numIdenticalPairs([]int{1, 2, 3})

	expect1 := 4
	expect2 := 6
	expect3 := 0

	if (result1 != expect1) || (result2 != expect2) || (result3 != expect3) {
		t.Errorf("Error")
	}

}
