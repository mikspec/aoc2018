package main

import (
	"testing"
)

func TestDay5(t *testing.T) {
	const TESTLEN = 10
	inputArray := loadFile("test1.txt")
	res := processData([]byte(inputArray))
	actualLen := len(res)
	if actualLen != TESTLEN {
		t.Errorf("Expected length %d, got %d", TESTLEN, actualLen)
	}
}

func TestProcessData(t *testing.T) {
	tests := [4][2]string{
		{"aAbcd", "bcd"},
		{"acCb", "ab"},
		{"aA", ""},
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
	}
	for i := range tests {
		res := processData([]byte(tests[i][0]))
		strOut := string(res)
		if strOut != tests[i][1] {
			t.Errorf("Expected value %s, got %s", tests[i][1], strOut)
		}
	}
}

func TestOptymisePolymer(t *testing.T) {
	minlen, _ := optimisePolymer("dabAcCaCBAcCcaDA")
	if minlen != 4 {
		t.Errorf("Expected value %d, got %d", 4, minlen)
	}
}
