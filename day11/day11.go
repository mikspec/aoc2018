package main

import "fmt"

const gridSerialNum = 2694
const gridSize = 300

type gridType [gridSize][gridSize]int

// Populate cell grid
func populateGrid(grid *gridType, gridSerialNum int) {
	for y := range grid {
		for x := range grid[y] {
			rackID := x + 1 + 10
			powerLevel := rackID * (y + 1)
			powerLevel += gridSerialNum
			powerLevel *= rackID
			powerLevel /= 100
			powerLevel %= 10
			powerLevel -= 5
			grid[y][x] = powerLevel
		}
	}
}

// Find max 3x3 area
func findMaxTotalPower(grid *gridType) (maxX, maxY, maxSum int) {
	maxSum = -100
	maxX, maxY = -1, -1
	for y := 0; y < len(grid)-2; y++ {
		for x := 0; x < len(grid[y])-2; x++ {
			sum := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					sum += grid[y+i][x+j]
				}
			}
			if sum > maxSum {
				maxSum = sum
				maxX = x + 1
				maxY = y + 1
			}
		}
	}
	return
}

// Find any max area
func findMaxTotalPower2(grid *gridType) (maxX, maxY, maxSum, maxLen int) {
	maxSum = -1000
	maxX, maxY = -1, -1
	maxLen = 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			lenLimit := len(grid) - y
			if len(grid[y])-x < lenLimit {
				lenLimit = len(grid[y]) - x
			}
			for len := 1; len <= lenLimit; len++ {
				sum := 0
				for i := 0; i < len; i++ {
					for j := 0; j < len; j++ {
						sum += grid[y+i][x+j]
					}
				}
				if sum > maxSum {
					maxSum = sum
					maxX = x + 1
					maxY = y + 1
					maxLen = len
				}
			}
		}
	}
	return
}

func main() {
	grid := gridType{}
	populateGrid(&grid, gridSerialNum)
	maxX, maxY, maxSum := findMaxTotalPower(&grid)
	fmt.Println("X =", maxX, "Y =", maxY, "Sum =", maxSum)
	var maxLen int
	maxX, maxY, maxSum, maxLen = findMaxTotalPower2(&grid)
	fmt.Println("X =", maxX, "Y =", maxY, "Sum =", maxSum, "Max len =", maxLen)
}
