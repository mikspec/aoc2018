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
	w int
	x int
	y int
	z int
}

// File loading generates array of cords and maxs
func loadFile(name string) (cords []cordsType) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cords = make([]cordsType, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			cordsStr := strings.Split(i, ",")
			w, _ := strconv.Atoi(cordsStr[0])
			x, _ := strconv.Atoi(cordsStr[1])
			y, _ := strconv.Atoi(cordsStr[2])
			z, _ := strconv.Atoi(cordsStr[3])
			cords = append(cords, cordsType{w, x, y, z})
		}
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

// Manhatan distance
func distance(p1 cordsType, p2 cordsType) int {
	return abs(p1.w-p2.w) + abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

// Find number of constellations
func processCords(cords []cordsType) (cnt int) {
	// Make set of siblings (dist <= 3) for each point
	cordSet := make(map[int][]int)
	for i, p1 := range cords {
		cordSet[i] = make([]int, 0)
		for j, p2 := range cords {
			dist := distance(p1, p2)
			if dist <= 3 {
				cordSet[i] = append(cordSet[i], j)
			}
		}
	}
	// Set of processed points
	procSet := make(map[int]bool)
	for i := range cords {
		if _, ok := procSet[i]; ok {
			continue
		}
		cnt++
		siblings := make([]int, 0)
		siblings = append(siblings, i)
		// For each point add siblings to the queue, process queue
		for len(siblings) > 0 {
			x := siblings[0]
			siblings = siblings[1:]
			if _, ok := procSet[x]; ok {
				continue
			}
			procSet[x] = true
			for _, k := range cordSet[x] {
				siblings = append(siblings, k)
			}
		}
	}
	return
}

func main() {
	cords := loadFile("input.txt")
	cnt := processCords(cords)
	fmt.Println("Constellations:", cnt)
}
