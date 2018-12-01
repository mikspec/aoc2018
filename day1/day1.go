package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
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

	fmt.Println("Sum = ", sum)

	sum = 0
	freqdup := make([]int, 0)
	freqset := make(map[int]bool)
	freqset[sum] = true

	for i := 0; len(freqdup) == 0; i++ {
		sum += inputArray[i%len(inputArray)]
		if _, found := freqset[sum]; found == false {
			freqset[sum] = true
		} else {
			freqdup = append(freqdup, sum)
		}
	}

	fmt.Println("First freq duplicated = ", freqdup[0])
}
