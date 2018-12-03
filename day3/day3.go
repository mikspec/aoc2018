package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type claimType struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

type fabricType [1000][1000][]int

// File loading generates array of claims
func loadFile(name string) []claimType {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArray := make([]claimType, 0)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")

	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			claimStr := re.FindStringSubmatch(i)
			if claimStr != nil {
				newClaim := claimType{}
				newClaim.id, _ = strconv.Atoi(claimStr[1])
				newClaim.x, _ = strconv.Atoi(claimStr[2])
				newClaim.y, _ = strconv.Atoi(claimStr[3])
				newClaim.width, _ = strconv.Atoi(claimStr[4])
				newClaim.height, _ = strconv.Atoi(claimStr[5])
				inputArray = append(inputArray, newClaim)
			}
		}
	}

	return inputArray
}

// Populate single claim into fabric matrix
func insertClaim(fabric *fabricType, claim *claimType) {
	for row := claim.y; row < claim.height+claim.y; row++ {
		for col := claim.x; col < claim.width+claim.x; col++ {
			if fabric[row][col] == nil {
				fabric[row][col] = make([]int, 0, 1)
			}
			fabric[row][col] = append(fabric[row][col], claim.id)
		}
	}
}

// Fabric represented by array of arrays
func populateFabricMatrix(inputArray []claimType) *fabricType {
	fabric := new(fabricType)
	for _, v := range inputArray {
		insertClaim(fabric, &v)
	}
	return fabric
}

// Count shared area of fabric
func countSharedArea(fabric *fabricType) int {
	count := 0
	for row := 0; row < len(fabric); row++ {
		for col := 0; col < len(fabric[0]); col++ {
			if fabric[row][col] != nil && len(fabric[row][col]) > 1 {
				count++
			}
		}
	}
	return count
}

// Find claim with not overlaped area
func findGoodClaim(fabric *fabricType) int {

	claimMap := make(map[int]int)
	for row := 0; row < len(fabric); row++ {
		for col := 0; col < len(fabric[0]); col++ {
			if fabric[row][col] != nil {
				for _, v := range fabric[row][col] {
					if len(fabric[row][col]) > claimMap[v] {
						claimMap[v] = len(fabric[row][col])
					}
				}
			}
		}
	}
	for k, v := range claimMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	inputArray := loadFile("input.txt")
	fabric := populateFabricMatrix(inputArray)
	count := countSharedArea(fabric)
	fmt.Println("Shared area: ", count)
	claim := findGoodClaim(fabric)
	fmt.Println("Good claim: ", claim)
}
