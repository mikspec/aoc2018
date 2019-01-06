package main

import (
	"testing"
)

func TestDay17(t *testing.T) {
	minX, _, minY, maxY, area := loadFile("test1.txt", sourceX, sourceY)
	simulation(area, sourceX, sourceY, minX, maxY)
	drainWater, waterOnRest := calculateWater(area, minY)
	if drainWater+waterOnRest != 57 {
		t.Error("All expected water 57 got", drainWater+waterOnRest)
	}
	if waterOnRest != 29 {
		t.Error("Water on rest expected 29 got", waterOnRest)
	}
}
