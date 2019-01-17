package main

import (
	"testing"
)

func TestDay15(t *testing.T) {
	battleField, goblins, elves := loadFile("test1.txt")
	result, battleID := battle(&battleField, &goblins, &elves)
	if result != 27730 {
		t.Error("Expected 27730, 47, got", result, battleID)
	}
	battleField, goblins, elves = loadFile("test3.txt")
	result, battleID = battle(&battleField, &goblins, &elves)
	if result != 28944 {
		t.Error("Expected 28944, 54, got", result, battleID)
	}
}
