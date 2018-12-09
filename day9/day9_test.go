package main

import (
	"testing"
)

func TestDay9(t *testing.T) {
	score := runGame(9, 25)
	if score != 32 {
		t.Error("Expected 32 got", score)
	}
}
