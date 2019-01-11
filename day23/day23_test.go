package main

import (
	"testing"
)

func TestDay23(t *testing.T) {
	inputArray := loadFile("test1.txt")
	ind, cnt := findStrongest(inputArray)
	if cnt != 7 || ind != 0 {
		t.Error("Expected 0 7, got", ind, cnt)
	}
	inputArray = loadFile("test2.txt")
	bestPos := findBestPosition(inputArray)
	dist := distance(pointType{0, 0, 0}, bestPos)
	if dist != 36 {
		t.Error("Expected distance 36, got", dist)
	}
}
