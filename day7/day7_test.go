package main

import (
	"reflect"
	"testing"
)

func TestDay7(t *testing.T) {

	testGraph := []stepType{
		{"C", "A"},
		{"C", "F"},
		{"A", "B"},
		{"A", "D"},
		{"B", "E"},
		{"D", "E"},
		{"F", "E"},
	}
	inputArray := loadFile("test1.txt")
	if !reflect.DeepEqual(inputArray, testGraph) {
		t.Error("Expected", testGraph, "got", inputArray)
	}

	graph := createMap(inputArray)
	path := processGraph(graph)
	if path != "CABDFE" {
		t.Error("Expected", "CABDFE", "got", path)
	}

	graph = createMap(inputArray)
	path, time := processGraphWithWorkers(graph, 2, 0)
	if path != "CABFDE" || time != 15 {
		t.Error("Expected", "CABFDE 15", "got", path, time)
	}
}
