package main

import (
	"reflect"
	"testing"
)

func TestDay6(t *testing.T) {

	testCords := []cordsType{
		{1, 1},
		{1, 6},
		{8, 3},
		{3, 4},
		{5, 5},
		{8, 9},
	}
	inputArray, maxX, maxY := loadFile("test1.txt")
	if !reflect.DeepEqual(inputArray, testCords) {
		t.Error("Expected", testCords, "got", inputArray)
	}
	if maxX != 8 || maxY != 9 {
		t.Error("Expected max:", 8, 9, "got", maxX, maxY)
	}
	area := populateArea(inputArray, maxX, maxY)
	maxArea, maxID := findMaxArea(inputArray, area)
	if maxArea != 17 || maxID != 4 {
		t.Error("Expected area:", 17, 4, "got", maxArea, maxID)
	}
	maxArea2 := findMaxArea2(inputArray, area, 32)
	if maxArea2 != 16 {
		t.Error("Expected area2:", 16, "got", maxArea2)
	}
}

func TestDistance(t *testing.T) {
	p1 := cordsType{3, 2}
	p2 := cordsType{1, 1}
	dist := distance(p1, p2)
	if dist != 3 {
		t.Error("Expected dist:", 3, "got", dist)
	}
}
