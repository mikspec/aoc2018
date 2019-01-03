package main

import (
	"reflect"
	"testing"
)

func TestDay16(t *testing.T) {
	state := stateType{3, 2, 1, 1}
	stateNew := commandMap["mulr"](state, commandType{9, 2, 1, 2})
	expected := stateType{3, 2, 2, 1}
	if !reflect.DeepEqual(stateNew, expected) {
		t.Error("Expected mulr", expected, "got", stateNew)
	}
	stateNew = commandMap["addi"](state, commandType{9, 2, 1, 2})
	expected = stateType{3, 2, 2, 1}
	if !reflect.DeepEqual(stateNew, expected) {
		t.Error("Expected addi", expected, "got", stateNew)
	}
	stateNew = commandMap["seti"](state, commandType{9, 2, 1, 2})
	expected = stateType{3, 2, 2, 1}
	if !reflect.DeepEqual(stateNew, expected) {
		t.Error("Expected seti", expected, "got", stateNew)
	}
}

func TestProcessStack(t *testing.T) {
	stack, _ := loadFile("input.txt")
	processStack(&commandMap, stack, &mnemonicMap, true)
	if len(mnemonicMap) != 16 {
		t.Error("Mapping error")
	}
}
