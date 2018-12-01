package main

import (
	"reflect"
	"testing"
)

func TestProcessArray(t *testing.T) {
	var v int
	v = processArray([]int{+1, -1})
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
	v = processArray([]int{+3, +3, +4, -2, -4})
	if v != 10 {
		t.Error("Expected 10, got ", v)
	}
	v = processArray([]int{-6, +3, +8, +5, -6})
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
	v = processArray([]int{+7, +7, -2, -7, -4})
	if v != 14 {
		t.Error("Expected 14, got ", v)
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
