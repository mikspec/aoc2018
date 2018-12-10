package main

import (
	"fmt"
	"testing"
)

func TestDay10(t *testing.T) {
	inputArray := loadFile("test1.txt")
	minSize, minSizeCnt := getMinSize(inputArray)
	fmt.Println(minSize, minSizeCnt)
	inputArray = loadFile("test1.txt")
	drawPicture(minSizeCnt, inputArray)
}
