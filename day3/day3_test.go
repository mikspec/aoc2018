package main

import (
	"reflect"
	"testing"
)

func TestLoadFile(t *testing.T) {
	inputArray := loadFile("test1.txt")
	if !reflect.DeepEqual(inputArray, []claimType{{1, 100, 366, 24, 27}}) {
		t.Error("Expected [[{1,100,366,24,27}], got ", inputArray)
	}
}

func TestInsertClaim(t *testing.T) {
	fabric := new(fabricType)
	claim := claimType{1, 2, 3, 4, 5}
	var value int
	insertClaim(fabric, &claim)
	for row := 0; row < len(fabric); row++ {
		for col := 0; col < len(fabric[0]); col++ {
			if row < claim.y || row >= claim.y+claim.height || col < claim.x || col >= claim.x+claim.width {
				value = 0
			} else {
				value = 1
			}
			if value == 1 {
				if fabric[row][col] == nil || len(fabric[row][col]) != 1 {
					t.Errorf("Expected [%d][%d] = %d, got %d", row, col, value, fabric[row][col])
				}
			} else {
				if fabric[row][col] != nil {
					t.Errorf("Expected [%d][%d] = %d, got %d", row, col, value, fabric[row][col])
				}
			}
		}
	}
	insertClaim(fabric, &claimType{2, 3, 4, 1, 1})
	if len(fabric[4][3]) != 2 {
		t.Errorf("Expected [%d][%d] = %d, got %d", 4, 3, 2, fabric[4][3])
	}
	count := countSharedArea(fabric)
	if count != 1 {
		t.Errorf("Expected %d, got %d", 1, count)
	}
	goodClaim := findGoodClaim(fabric)
	if goodClaim != -1 {
		t.Errorf("Expected %d, got %d", -1, goodClaim)
	}
	insertClaim(fabric, &claimType{3, 100, 100, 2, 2})
	goodClaim = findGoodClaim(fabric)
	if goodClaim != 3 {
		t.Errorf("Expected %d, got %d", 3, goodClaim)
	}
}
