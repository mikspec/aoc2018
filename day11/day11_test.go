package main

import (
	"testing"
)

func TestPopulateGrid(t *testing.T) {
	serialNum := 8
	x := 3
	y := 5
	grid := gridType{}
	populateGrid(&grid, serialNum)
	powerLevel := grid[y-1][x-1]
	if powerLevel != 4 {
		t.Error("Expected power level [", x, ",", y, "] = 4, got ", powerLevel)
	}
	serialNum = 57
	grid = gridType{}
	x = 122
	y = 79
	populateGrid(&grid, serialNum)
	powerLevel = grid[y-1][x-1]
	if powerLevel != -5 {
		t.Error("Expected power level [", x, ",", y, "] = -5, got ", powerLevel)
	}
}

func TestFindMaxTotalPower(t *testing.T) {
	serialNum := 18
	grid := gridType{}
	populateGrid(&grid, serialNum)
	maxX, maxY, _ := findMaxTotalPower(&grid)
	if maxX != 33 || maxY != 45 {
		t.Error("Expected 33, 45 - got", maxX, ",", maxY)
	}
}
