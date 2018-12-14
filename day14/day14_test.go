package main

import (
	"reflect"
	"testing"
)

func TestDay14(t *testing.T) {
	lab := initLab()
	for i := 0; i < 15; i++ {
		generateRecipts(&lab)
	}
	score := lab.recits[9:19]
	expected := reciptType{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}
	if !reflect.DeepEqual(score, expected) {
		t.Error("Expected", expected, "got", score)
	}
}

func TestFindScore(t *testing.T) {
	lab := initLab()
	countLeft := findScore(&lab, reciptType{5, 1, 5, 8, 9})
	if countLeft != 9 {
		t.Error("Expected 9 got", countLeft)
	}
	lab = initLab()
	countLeft = findScore(&lab, reciptType{5, 9, 4, 1, 4})
	if countLeft != 2018 {
		t.Error("Expected 9 got", countLeft)
	}

}
