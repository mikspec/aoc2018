package main

import (
	"testing"
)

func TestDay25(t *testing.T) {

	cords := loadFile("test1.txt")
	cnt := processCords(cords)
	if cnt != 2 {
		t.Error("Expected constellations 2, got", cnt)
	}
}
