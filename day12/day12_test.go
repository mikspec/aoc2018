package main

import (
	"testing"
)

func TestDay12(t *testing.T) {
	rules, state := loadFile("test1.txt")
	for i := 0; i < 20; i++ {
		state = nextGeneration(&rules, state)
	}
	crc := calculateCRC(state)
	if crc != 325 {
		t.Error("Expected 325, got", crc)
	}
}

func TestDecodeRule(t *testing.T) {
	ruleStr := []string{"..#.# => #", "..#.#", "#"}
	key, val := decodeRule(ruleStr)
	if key != 5 || val != true {
		t.Error("Expected key 5 , val true , got key", key, ", val", val)
	}
}
