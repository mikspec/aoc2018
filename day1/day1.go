package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// File loading generates array of integers and sum of elements
func loadFile(name string) (int, []int) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	inputArray := make([]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			inputArray = append(inputArray, i)
			sum += i
		}
	}
	return sum, inputArray
}

// ProcessArray returns doubled frequency
func processArray(inputArray []int) int {
	sum := 0
	freqset := make(map[int]bool)
	freqset[sum] = true

	for i := 0; ; i++ {
		sum += inputArray[i%len(inputArray)]
		if _, found := freqset[sum]; found == false {
			freqset[sum] = true
		} else {
			return sum
		}
	}
}

func main() {
	sum, inputArray := loadFile("input.txt")
	fmt.Println("Sum = ", sum)

	freq := processArray(inputArray)
	fmt.Println("First freq duplicated = ", freq)
}
