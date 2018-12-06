package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cordsType struct {
	x int
	y int
}

type areaPoint struct {
	id   int // occupied by
	dist int // distance
}

type areaType [][]areaPoint

// File loading generates array of cords and maxs
func loadFile(name string) (cords []cordsType, maxX int, maxY int) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArray := make([]cordsType, 0)
	scanner := bufio.NewScanner(file)
	maxX = -1
	maxY = -1

	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			cordsStr := strings.Split(i, ",")
			x, _ := strconv.Atoi(cordsStr[0])
			y, _ := strconv.Atoi(strings.Trim(cordsStr[1], " "))
			inputArray = append(inputArray, cordsType{x, y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	return inputArray, maxX, maxY
}

// Manhatan distance
func distance(p1 cordsType, p2 cordsType) int {
	distX := p2.x - p1.x
	if distX < 0 {
		distX *= -1
	}
	distY := p2.y - p1.y
	if distY < 0 {
		distY *= -1
	}
	return distX + distY
}

// Populate ara table by given cords
func populateArea(inputArray []cordsType, maxX int, maxY int) areaType {
	area := make(areaType, maxY+1, maxY+1)
	for i := range area {
		area[i] = make([]areaPoint, maxX+1, maxX+1)
		for j := range area[i] {
			area[i][j].dist = maxX + maxY + 2
		}
	}
	for i, point := range inputArray {
		for y := 0; y < len(area); y++ {
			for x := 0; x < len(area[y]); x++ {
				dist := distance(cordsType{x, y}, point)
				if area[y][x].dist > dist {
					area[y][x].dist = dist
					area[y][x].id = i
				} else if area[y][x].dist == dist {
					area[y][x].id = -1
				}
			}
		}
	}
	return area
}

// Find max finite area
func findMaxArea(inputArray []cordsType, area areaType) (int, int) {
	maxSum := 0
	maxID := -1
AreaLoop:
	for i := range inputArray {
		sum := 0
		for y := 0; y < len(area); y++ {
			for x := 0; x < len(area[y]); x++ {
				if area[y][x].id == i {
					// Found infinite area
					if x == 0 || y == 0 || x == len(area[y])-1 || y == len(area)-1 {
						continue AreaLoop
					} else {
						sum++
					}
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
			maxID = i
		}
	}
	return maxSum, maxID
}

// Find max area with total distance less than X
func findMaxArea2(inputArray []cordsType, area areaType, totalDist int) int {
	maxArea := 0
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			sum := 0
			for _, point := range inputArray {
				dist := distance(cordsType{x, y}, point)
				sum += dist
			}
			if sum < totalDist {
				maxArea++
			}
		}
	}
	return maxArea
}

func main() {
	inputArray, maxX, maxY := loadFile("input.txt")
	area := populateArea(inputArray, maxX, maxY)
	maxArea, maxID := findMaxArea(inputArray, area)
	maxArea2 := findMaxArea2(inputArray, area, 10000)
	fmt.Println("Max area: ", maxArea, "for id", maxID)
	fmt.Println("Max area2: ", maxArea2)
}
