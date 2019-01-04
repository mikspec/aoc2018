package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rowType []rune
type areaType []rowType

// File loading generates road plan
func loadFile(name string) areaType {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	area := make(areaType, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			row := make(rowType, len(line))
			area = append(area, row)
			for i, v := range line {
				row[i] = v
			}
		}
	}
	return area
}

// Displays area table
func displayArea(area areaType) {
	for y := range area {
		for x := range area[y] {
			fmt.Printf("%s", string(area[y][x]))
		}
		fmt.Println()
	}
}

// Get counters for adjacent acres
func getAdjacentCounters(area areaType, y, x int) (stat map[rune]int) {
	stat = make(map[rune]int)
	stat['.'] = 0
	stat['|'] = 0
	stat['#'] = 0
	if y > 0 {
		if x > 0 {
			stat[area[y-1][x-1]]++
		}
		stat[area[y-1][x]]++
		if x < len(area[y-1])-1 {
			stat[area[y-1][x+1]]++
		}
	}
	if x > 0 {
		stat[area[y][x-1]]++
	}
	if x < len(area[y])-1 {
		stat[area[y][x+1]]++
	}
	if y < len(area)-1 {
		if x > 0 {
			stat[area[y+1][x-1]]++
		}
		stat[area[y+1][x]]++
		if x < len(area[y+1])-1 {
			stat[area[y+1][x+1]]++
		}
	}
	return
}

func getValue(area areaType, y, x int) (newVal rune) {
	newVal = area[y][x]
	stat := getAdjacentCounters(area, y, x)
	if newVal == '.' {
		if stat['|'] >= 3 {
			newVal = '|'
		}
		return
	}
	if newVal == '|' {
		if stat['#'] >= 3 {
			newVal = '#'
		}
		return
	}
	if newVal == '#' {
		if !(stat['|'] > 0 && stat['#'] > 0) {
			newVal = '.'
		}
		return
	}
	return
}

// Process one minute
func processMinute(area areaType) (newArea areaType) {
	newArea = make(areaType, len(area))
	for y := range area {
		newArea[y] = make(rowType, len(area[y]))
		for x := range area[y] {
			newArea[y][x] = getValue(area, y, x)
		}
	}
	return
}

// Gets coverage statistics
func getCoverage(area areaType) int {
	trees, lumberyards := 0, 0
	for y := 0; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			if area[y][x] == '|' {
				trees++
			}
			if area[y][x] == '#' {
				lumberyards++
			}
		}
	}
	return trees * lumberyards
}

func main() {
	area := loadFile("input.txt")
	for i := 0; i < 10; i++ {
		area = processMinute(area)
	}
	displayArea(area)
	fmt.Println("Coverage after 10 minutes:", getCoverage(area))
	area = loadFile("input.txt")
	for i := 0; i < 1000; i++ {
		area = processMinute(area)
	}
	displayArea(area)
	fmt.Println("Coverage after 1000000000 minutes:", getCoverage(area))
}
