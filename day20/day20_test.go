package main

import (
	"testing"
)

func TestDay20(t *testing.T) {
	data, plan := loadFile("test1.txt")
	_, plan = processPath(0, 0, 0, data, plan)
	minX, minY, planArray := createPlan(plan)
	maxDist := calcDoors(-minX, -minY, 0, planArray)
	if maxDist != 10 {
		t.Error("Expected max distance: 10, got:", maxDist)
	}
	roomsCnt := roomsAboveLimit(9, planArray)
	if roomsCnt != 4 {
		t.Error("Expected rooms cnt: 4, got:", roomsCnt)
	}
}
