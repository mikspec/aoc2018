package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Plan item types
const (
	WALL = iota
	ROOM
	DOORHOR
	DOORVER
	START
)

type coordsType struct {
	x    int
	y    int
	item int
}

type planItemType struct {
	object   int
	distance int
}

// Load regular expression
func loadFile(name string) ([]byte, []coordsType) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	plan := make([]coordsType, 0)
	plan = append(plan, coordsType{0, 0, START})
	return data, plan
}

func move(posX, posY, dirX, dirY int, plan []coordsType) (int, int, []coordsType) {
	posX += dirX
	posY += dirY
	doorType := DOORHOR
	if dirY != 0 {
		doorType = DOORVER
	}
	plan = append(plan, coordsType{posX, posY, doorType})
	posX += dirX
	posY += dirY
	plan = append(plan, coordsType{posX, posY, ROOM})
	return posX, posY, plan
}

// Create coordinate grid
func processPath(posX, posY, index int, exp []byte, plan []coordsType) (int, []coordsType) {
	var dirX, dirY int
mainLoop:
	for index < len(exp) {
		dirX, dirY = 0, 0
		if exp[index] == '^' || exp[index] == '$' {
			index++
			continue
		}
		if exp[index] == '(' {
			index++
			for {
				index, plan = processPath(posX, posY, index, exp, plan)
				if exp[index] == ')' {
					index++
					continue mainLoop
				}
			}
		}
		if exp[index] == '|' {
			index++
			return index, plan
		}
		if exp[index] == ')' {
			return index, plan
		}
		if exp[index] == 'W' {
			dirX = -1
		}
		if exp[index] == 'E' {
			dirX = 1
		}
		if exp[index] == 'N' {
			dirY = -1
		}
		if exp[index] == 'S' {
			dirY = 1
		}
		posX, posY, plan = move(posX, posY, dirX, dirY, plan)
		index++
	}
	return index, plan
}

// Make a plan from coords list
func createPlan(plan []coordsType) (minX, minY int, planArray [][]planItemType) {
	minX, minY = 0, 0
	maxX, maxY := 0, 0
	for _, v := range plan {
		if v.x > maxX {
			maxX = v.x
		}
		if v.x < minX {
			minX = v.x
		}
		if v.y > maxY {
			maxY = v.y
		}
		if v.y < minY {
			minY = v.y
		}
	}
	minX--
	maxX++
	minY--
	maxY++
	planArray = make([][]planItemType, 0)
	for y := minY; y <= maxY; y++ {
		planArray = append(planArray, make([]planItemType, maxX-minX+1))
	}
	for _, v := range plan {
		planArray[v.y-minY][v.x-minX].object = v.item
		planArray[v.y-minY][v.x-minX].distance = len(planArray) * len(planArray[0])
	}
	return
}

// Displays area table
func displayArea(area [][]planItemType) {
	var c rune
	for y := range area {
		for x := range area[y] {
			if area[y][x].object == ROOM {
				c = '.'
			}
			if area[y][x].object == WALL {
				c = '#'
			}
			if area[y][x].object == DOORHOR {
				c = '|'
			}
			if area[y][x].object == DOORVER {
				c = '-'
			}
			if area[y][x].object == START {
				c = 'X'
			}
			fmt.Printf("%s", string(c))
		}
		fmt.Println()
	}
}

// Calculate max distance
func calcDoors(posX, posY, distance int, planArray [][]planItemType) (maxDist int) {
	planArray[posY][posX].distance = distance
	distLeft, distRight, distUp, distDown := 0, 0, 0, 0
	if planArray[posY][posX-1].object != 0 && planArray[posY][posX-2].distance > distance+1 {
		distLeft = calcDoors(posX-2, posY, distance+1, planArray)
	}
	if planArray[posY][posX+1].object != 0 && planArray[posY][posX+2].distance > distance+1 {
		distRight = calcDoors(posX+2, posY, distance+1, planArray)
	}
	if planArray[posY-1][posX].object != 0 && planArray[posY-2][posX].distance > distance+1 {
		distUp = calcDoors(posX, posY-2, distance+1, planArray)
	}
	if planArray[posY+1][posX].object != 0 && planArray[posY+2][posX].distance > distance+1 {
		distDown = calcDoors(posX, posY+2, distance+1, planArray)
	}
	maxDist = distance
	if maxDist < distLeft {
		maxDist = distLeft
	}
	if maxDist < distRight {
		maxDist = distRight
	}
	if maxDist < distUp {
		maxDist = distUp
	}
	if maxDist < distDown {
		maxDist = distDown
	}
	return
}

// How many rooms above limit
func roomsAboveLimit(limit int, plan [][]planItemType) (cnt int) {
	for y := range plan {
		for x := range plan[y] {
			if plan[y][x].object == ROOM && plan[y][x].distance >= limit {
				cnt++
			}
		}
	}
	return
}

func main() {
	data, plan := loadFile("input.txt")
	_, plan = processPath(0, 0, 0, data, plan)
	minX, minY, planArray := createPlan(plan)
	displayArea(planArray)
	maxDist := calcDoors(-minX, -minY, 0, planArray)
	fmt.Println("Max distance:", maxDist)
	roomsCnt := roomsAboveLimit(1000, planArray)
	fmt.Println("Rooms counter:", roomsCnt)
}
