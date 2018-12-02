package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	text      string // input
	expeTwo   bool
	expeThree bool // expected result
}{
	{"abcdef", false, false},
	{"bababc", true, true},
	{"abbcde", true, false},
	{"abcccd", false, true},
	{"aabcdd", true, false},
	{"abcdee", true, false},
	{"ababab", false, true},
}

func TestLoadFile(t *testing.T) {
	inputArray := loadFile("test1.txt")
	if !reflect.DeepEqual(inputArray, []string{"aaa", "bbb", "ccc"}) {
		t.Error("Expected [\"aaa\",\"bbb\",\"ccc\"], got ", inputArray)
	}
}

func TestProcessString(t *testing.T) {

	for _, tt := range tests {
		actualTwo, actualThree := processString(tt.text)
		if actualTwo != tt.expeTwo || actualThree != tt.expeThree {
			t.Errorf("Input(%s): expected {%t,%t}, actual {%t,%t}", tt.text, tt.expeTwo, tt.expeThree, actualTwo, actualThree)
		}
	}
}

func TestInputHash(t *testing.T) {
	testArray := make([]string, 0, len(tests))
	for _, v := range tests {
		testArray = append(testArray, v.text)
	}
	actual := inputHash(testArray)
	if actual != 12 {
		t.Errorf("Input(%s): expected 12, actual %d", testArray, actual)
	}
}

func TestFindDiff(t *testing.T) {

	var testDiff = []struct {
		strA  string
		strB  string
		found bool
		diff  string
	}{
		{"abcd", "abcf", true, "abc"},
		{"abcd", "adcd", true, "acd"},
		{"abcd", "abcd", true, "abcd"},
		{"abcd", "kijd", false, ""},
	}

	for _, v := range testDiff {
		actualFound, actualDiff := findDiff(v.strA, v.strB)
		if actualFound != v.found || actualDiff != v.diff {
			t.Errorf("Input(%s,%s): expected [%t %s], actual [%t %s]", v.strA, v.strB, v.found, v.diff, actualFound, actualDiff)
		}
	}
}

func TestCorrectBoxes(t *testing.T) {
	boxes := correctBoxes([]string{"abcd", "abce", "fdjk", "fgjk", "xznm"})
	if !reflect.DeepEqual(boxes, []string{"abc", "fjk"}) {
		t.Error("Expected [\"abc\",\"fjk\"], got ", boxes)
	}
}
