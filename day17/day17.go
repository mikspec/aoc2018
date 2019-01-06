package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const sourceX = 500
const sourceY = 0

type commandType struct {
	minX int
	maxX int
	minY int
	maxY int
}

type rowType []rune
type areaType []rowType

// File loading generates area array
func loadFile(name string, waterSourceX, waterSourceY int) (minX, maxX, minY, maxY int, area areaType) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	commands := make([]commandType, 0)
	reCommand := regexp.MustCompile("(x|y)=(\\d+), (x|y)=(\\d+)\\.\\.(\\d+)")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			commandStr := reCommand.FindStringSubmatch(line)
			command := commandType{}
			if commandStr[1] == "x" {
				command.minX, _ = strconv.Atoi(commandStr[2])
				command.maxX = command.minX
				command.minY, _ = strconv.Atoi(commandStr[4])
				command.maxY, _ = strconv.Atoi(commandStr[5])
			} else {
				command.minY, _ = strconv.Atoi(commandStr[2])
				command.maxY = command.minY
				command.minX, _ = strconv.Atoi(commandStr[4])
				command.maxX, _ = strconv.Atoi(commandStr[5])
			}
			commands = append(commands, command)
		}
	}
	minY = 10000
	minX = 10000
	for _, v := range commands {
		if minY > v.minY {
			minY = v.minY
		}
		if maxY < v.maxY {
			maxY = v.maxY
		}
		if maxX < v.maxX {
			maxX = v.maxX
		}
		if minX > v.minX {
			minX = v.minX
		}
	}
	offset := minX - 1
	area = make(areaType, maxY+1)
	for y := 0; y < len(area); y++ {
		area[y] = make(rowType, maxX-offset+2)
		for x := 0; x < len(area[y]); x++ {
			area[y][x] = '.'
		}
	}
	area[waterSourceY][waterSourceX-offset] = '+'
	for _, v := range commands {
		if v.minX == v.maxX {
			for y := v.minY; y <= v.maxY; y++ {
				area[y][v.maxX-offset] = '#'
			}
		}
		if v.minY == v.maxY {
			for x := v.minX; x <= v.maxX; x++ {
				area[v.minY][x-offset] = '#'
			}
		}
	}
	return
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

// Waterfall simulation
func simulation(area areaType, waterSourceX, waterSourceY, minX, maxY int) (code int) {
	code = goDown(area, waterSourceX-minX+1, waterSourceY+1, maxY)
	return
}

// Mark water on rest
func markWaterOnRest(area areaType, posX, posY int) {
	for i := posX; posX >= 0; i-- {
		if area[posY][i] == '#' {
			break
		}
		area[posY][i] = '~'
	}
	for i := posX; posX < len(area[posY]); i++ {
		if area[posY][i] == '#' {
			break
		}
		area[posY][i] = '~'
	}
}

// Make down move
func goDown(area areaType, posX, posY, maxY int) (code int) {
	// Bottom reached
	if posY > maxY {
		return
	}
	state := area[posY][posX]
	// Sand
	if state == '.' {
		area[posY][posX] = '|'
		code = goDown(area, posX, posY+1, maxY)
		// Barrier found
		if code == 2 {
			codeLeft := goHorizontal(area, posX-1, posY, maxY, -1)
			codeRight := goHorizontal(area, posX+1, posY, maxY, 1)
			code = codeLeft
			if code > codeRight {
				code = codeRight
			}
		}
		if code == 2 {
			markWaterOnRest(area, posX, posY)
		}
	} else if state == '#' || state == '~' {
		code = 2
	} else if state == '|' {
		code = 1
	}
	return
}

// Horizontal move
func goHorizontal(area areaType, posX, posY, maxY, direction int) (code int) {
	state := area[posY][posX]
	if state == '.' {
		area[posY][posX] = '|'
		code = goDown(area, posX, posY+1, maxY)
		if code == 2 {
			code = goHorizontal(area, posX+direction, posY, maxY, direction)
		}
	} else if state == '#' || state == '~' {
		code = 2
	} else if state == '|' {
		code = 1
	}
	return
}

func calculateWater(area areaType, minY int) (drainWater, waterOnRest int) {
	for y := minY; y < len(area); y++ {
		for x := 0; x < len(area[y]); x++ {
			state := area[y][x]
			if state == '~' {
				waterOnRest++
			}
			if state == '|' {
				drainWater++
			}
		}
	}
	return
}

func main() {
	minX, _, minY, maxY, area := loadFile("input.txt", sourceX, sourceY)
	simulation(area, sourceX, sourceY, minX, maxY)
	drainWater, waterOnRest := calculateWater(area, minY)
	fmt.Println("Water: ", drainWater+waterOnRest)
	fmt.Println("Water on rest: ", waterOnRest)

}
