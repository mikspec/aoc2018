package main

import (
	"fmt"
	"testing"
)

func TestDay24(t *testing.T) {
	immune, infection := loadFile("test1.txt")
	sumImmune, sumInfection := part1(&immune, &infection)
	if sumInfection != 5216 || sumImmune != 0 {
		t.Error("Expected sum infection 5216, got", sumInfection)
	}
	boost := 0
	sumImmune, sumInfection, boost = part2("test1.txt")
	fmt.Println("Boost", boost)
	if sumImmune != 51 {
		t.Error("Expected boost 51, got", boost)
	}
}
