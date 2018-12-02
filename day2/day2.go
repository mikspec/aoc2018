package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

// File loading generates array of integers and sum of elements
func loadFile(name string) []string {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArray := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			inputArray = append(inputArray, i)
		}
	}

	return inputArray
}

// CRC calculation for input array
func inputHash(inputArray []string) int {

	cntTwo := 0
	cntThree := 0

	for _, valStr := range inputArray {
		hasTwo, hasThree := processString(valStr)
		if hasTwo {
			cntTwo++
		}
		if hasThree {
			cntThree++
		}

	}

	return cntTwo * cntThree
}

// Returns flags has two and has three
func processString(valStr string) (hasTwo bool, hasThree bool) {

	hasTwo = false
	hasThree = false
	charMap := make(map[rune]int)

	for _, char := range valStr {
		charMap[char]++
	}

	for _, cnt := range charMap {
		if cnt == 2 {
			hasTwo = true
		}
		if cnt == 3 {
			hasThree = true
		}
	}

	return
}

func findDiff(a, b string) (found bool, diff string) {

	found = false
	diff = ""
	cntDiff := 0
	var buffer bytes.Buffer

	if len(a) != len(b) {
		return
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cntDiff++
			if cntDiff > 1 {
				return
			}
		} else {
			buffer.WriteByte(a[i])
		}
	}

	diff = buffer.String()
	found = true

	return
}

func correctBoxes(inputArray []string) []string {
	correctB := make([]string, 0)
	for i := 0; i < len(inputArray)-1; i++ {
		for j := i + 1; j < len(inputArray); j++ {
			if found, diff := findDiff(inputArray[i], inputArray[j]); found {
				correctB = append(correctB, diff)
			}
		}
	}
	return correctB
}

func main() {
	inputArray := loadFile("input.txt")
	fmt.Println("Input = ", len(inputArray))
	hash := inputHash(inputArray)
	fmt.Println("Hash = ", hash)
	fmt.Println("Correct Boxes = ", correctBoxes(inputArray))
}
