package main

import (
	"reflect"
	"testing"
)

func TestCommand(t *testing.T) {
	state := stateType{3, 2, 1, 1}
	stateNew := commandMap["mulr"](state, commandType{"mulr", 2, 1, 2})
	expected := stateType{3, 2, 2, 1}
	if !reflect.DeepEqual(stateNew, expected) {
		t.Error("Expected mulr", expected, "got", stateNew)
	}
	stateNew = commandMap["addi"](state, commandType{"addi", 2, 1, 2})
	expected = stateType{3, 2, 2, 1}
	if !reflect.DeepEqual(stateNew, expected) {
		t.Error("Expected addi", expected, "got", stateNew)
	}
	stateNew = commandMap["seti"](state, commandType{"seti", 2, 1, 2})
	expected = stateType{3, 2, 2, 1}
	if !reflect.DeepEqual(stateNew, expected) {
		t.Error("Expected seti", expected, "got", stateNew)
	}
}

func TestProcessStack(t *testing.T) {
	ip, commands := loadFile("test1.txt")
	_, _, state := processCommands(&commandMap, ip, commands, stateType{})
	expected := stateType{6, 5, 6, 0, 0, 9}
	if !reflect.DeepEqual(state, expected) {
		t.Error("Expected", expected, "got", state)
	}
}
