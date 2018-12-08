package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// File loading generates array of steps
func loadFile(name string) []int {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	dataStrArray := strings.Split(string(data), " ")
	dataIntArray := make([]int, len(dataStrArray))
	for i, v := range dataStrArray {
		n, _ := strconv.Atoi(v)
		dataIntArray[i] = n
	}
	return dataIntArray
}

func processData(inputArray []int) (nodeLen, crcSum int) {
	// No child nodes
	crcSum = 0
	nodeLen = 2 + inputArray[1]
	if inputArray[0] == 0 {
		for _, v := range inputArray[2 : 2+inputArray[1]] {
			crcSum += v
		}
		return nodeLen, crcSum
	}
	nodeIndex := 2
	// Iterate over all child nodes
	for i := 0; i < inputArray[0]; i++ {
		childLen, childSum := processData(inputArray[nodeIndex:])
		crcSum += childSum
		nodeIndex += childLen
		nodeLen += childLen
	}
	// Add crc of current node
	for _, v := range inputArray[nodeIndex : nodeIndex+inputArray[1]] {
		crcSum += v
	}
	return nodeLen, crcSum
}

func processData2(inputArray []int) (nodeLen, crcSum int) {
	// No child nodes
	crcSum = 0
	nodeLen = 2 + inputArray[1]
	if inputArray[0] == 0 {
		for _, v := range inputArray[2 : 2+inputArray[1]] {
			crcSum += v
		}
		return nodeLen, crcSum
	}
	nodeIndex := 2
	childSumArray := make([]int, 0)
	// Iterate over all child nodes
	for i := 0; i < inputArray[0]; i++ {
		childLen, childSum := processData2(inputArray[nodeIndex:])
		childSumArray = append(childSumArray, childSum)
		nodeIndex += childLen
		nodeLen += childLen
	}
	// Add crc of current node
	for _, v := range inputArray[nodeIndex : nodeIndex+inputArray[1]] {
		if v > 0 && v <= len(childSumArray) {
			crcSum += childSumArray[v-1]
		}
	}
	return nodeLen, crcSum
}

func main() {
	inputArray := loadFile("input.txt")
	nodeLen, crcSum := processData(inputArray)
	fmt.Println("Part 1 - Node length:", nodeLen, "CRC sum:", crcSum)
	nodeLen, crcSum = processData2(inputArray)
	fmt.Println("Part 2 - Node length:", nodeLen, "CRC sum:", crcSum)
}
