package main

import (
	"reflect"
	"testing"
)

func TestDay8(t *testing.T) {
	testData := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
	inputArray := loadFile("test1.txt")
	if !reflect.DeepEqual(inputArray, testData) {
		t.Error("Expected", testData, "got", inputArray)
	}

	nodeLen, crcSum := processData(testData[2:])
	if nodeLen != 5 || crcSum != 33 {
		t.Error("Expected node len and crc 5 33 got", nodeLen, crcSum)
	}

	nodeLen, crcSum = processData(testData[7:])
	if nodeLen != 6 || crcSum != 101 {
		t.Error("Expected node len and crc 6 101 got", nodeLen, crcSum)
	}
	nodeLen, crcSum = processData(testData)
	if nodeLen != 16 || crcSum != 138 {
		t.Error("Expected node len and crc 16 138 got", nodeLen, crcSum)
	}
	nodeLen, crcSum = processData2(testData[7:])
	if nodeLen != 6 || crcSum != 0 {
		t.Error("Expected node len and crc 16 66 got", nodeLen, crcSum)
	}
	nodeLen, crcSum = processData2(testData)
	if nodeLen != 16 || crcSum != 66 {
		t.Error("Expected node len and crc 16 66 got", nodeLen, crcSum)
	}
}
