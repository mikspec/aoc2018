package main

import (
	"testing"
)

func TestDay13(t *testing.T) {
	roadPlan, cartCnt := loadFile("test1.txt")
	crashX, crashY, lastX, lastY := processPlan(roadPlan, cartCnt)
	if crashX != 7 || crashY != 3 {
		t.Error("Expected 7,3 got", crashX, crashY, lastX, lastY)
	}
}

func TestMapComparission(t *testing.T) {
	v, ok := turnMap[turnItemType{LEFT, -1, '/'}]
	if !ok || v != DOWN {
		t.Error("Expected ", DOWN, "got", v)
	}
}
