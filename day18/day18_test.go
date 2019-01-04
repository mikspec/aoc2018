package main

import (
	"reflect"
	"testing"
)

func TestDay18(t *testing.T) {
	area := loadFile("test1.txt")
	for i := 0; i < 10; i++ {
		area = processMinute(area)
	}
	result := getCoverage(area)
	expected := 1147
	if result != expected {
		t.Error("Expected", expected, "got", result)
	}
}

func TestGetAdjacentCounters(t *testing.T) {
	area := loadFile("test1.txt")
	stat := getAdjacentCounters(area, 0, 0)
	expected := map[rune]int{'.': 2, '#': 1, '|': 0}
	if !reflect.DeepEqual(stat, expected) {
		t.Error("Expected", expected, "got", stat)
	}
	stat = getAdjacentCounters(area, 9, 9)
	expected = map[rune]int{'.': 1, '#': 0, '|': 2}
	if !reflect.DeepEqual(stat, expected) {
		t.Error("Expected", expected, "got", stat)
	}
	stat = getAdjacentCounters(area, 3, 3)
	expected = map[rune]int{'.': 3, '#': 1, '|': 4}
	if !reflect.DeepEqual(stat, expected) {
		t.Error("Expected", expected, "got", stat)
	}
}
