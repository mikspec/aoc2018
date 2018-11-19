package main

import "testing"

func TestSum(t *testing.T) {
	var v int
	v = Sum([]int{1, 2, 3})
	if v != 6 {
		t.Error("Expected 6, got ", v)
	}
}
