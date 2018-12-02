package main

import (
	"reflect"
	"testing"
)

func TestProcessArray(t *testing.T) {

	var tests = []struct {
		n        []int // input
		expected int   // expected result
	}{
		{[]int{+1, -1}, 0},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}

	for _, tt := range tests {
		actual := processArray(tt.n)
		if actual != tt.expected {
			t.Errorf("Input(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
func TestLoadFile(t *testing.T) {
	sum, inputArray := loadFile("test1.txt")
	if sum != 0 {
		t.Error("Expected 0, got ", sum)
	}
	if !reflect.DeepEqual(inputArray, []int{1, 1, -2}) {
		t.Error("Expected [1,1,-2], got ", inputArray)
	}
}
