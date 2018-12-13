package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type directionType struct {
	horizontal int
	vertical   int
	image      rune
}

// LEFT move
var LEFT = directionType{-1, 0, '<'}

// RIGHT move
var RIGHT = directionType{1, 0, '>'}

// UP move
var UP = directionType{0, -1, '^'}

// DOWN move
var DOWN = directionType{0, 1, 'v'}

type cartType struct {
	cartID    int
	direction directionType
	turnCnt   int
	moveID    int
}

type planItemType struct {
	cart *cartType
	path rune
}

type turnItemType struct {
	direction directionType
	turnCnt   int
	image     rune
}

type turnMapType map[turnItemType]directionType

type roadPlanType [][]planItemType

var turnMap = turnMapType{
	// Corners
	turnItemType{LEFT, -1, '/'}:   DOWN,
	turnItemType{LEFT, -1, '\\'}:  UP,
	turnItemType{RIGHT, -1, '/'}:  UP,
	turnItemType{RIGHT, -1, '\\'}: DOWN,
	turnItemType{UP, -1, '/'}:     RIGHT,
	turnItemType{UP, -1, '\\'}:    LEFT,
	turnItemType{DOWN, -1, '/'}:   LEFT,
	turnItemType{DOWN, -1, '\\'}:  RIGHT,
	// Crosses
	turnItemType{LEFT, 0, '+'}:  DOWN,
	turnItemType{LEFT, 1, '+'}:  LEFT,
	turnItemType{LEFT, 2, '+'}:  UP,
	turnItemType{RIGHT, 0, '+'}: UP,
	turnItemType{RIGHT, 1, '+'}: RIGHT,
	turnItemType{RIGHT, 2, '+'}: DOWN,
	turnItemType{UP, 0, '+'}:    LEFT,
	turnItemType{UP, 1, '+'}:    UP,
	turnItemType{UP, 2, '+'}:    RIGHT,
	turnItemType{DOWN, 0, '+'}:  RIGHT,
	turnItemType{DOWN, 1, '+'}:  DOWN,
	turnItemType{DOWN, 2, '+'}:  LEFT,
}

// File loading generates road plan
func loadFile(name string) (roadPlan roadPlanType, cartCnt int) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	roadPlan = make(roadPlanType, 0)
	scanner := bufio.NewScanner(file)
	cartCnt = 0
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			row := make([]planItemType, len(line))
			roadPlan = append(roadPlan, row)
			for i, v := range line {
				if v == '>' {
					cartCnt++
					row[i] = planItemType{&cartType{cartCnt, RIGHT, 0, 0}, '-'}
				} else if v == '<' {
					cartCnt++
					row[i] = planItemType{&cartType{cartCnt, LEFT, 0, 0}, '-'}
				} else if v == '^' {
					cartCnt++
					row[i] = planItemType{&cartType{cartCnt, UP, 0, 0}, '|'}
				} else if v == 'v' {
					cartCnt++
					row[i] = planItemType{&cartType{cartCnt, DOWN, 0, 0}, '|'}
				} else {
					row[i] = planItemType{nil, v}
				}
			}
		}
	}
	return
}

// Display road map
func displayMap(roadPlan roadPlanType) {
	for y := 0; y < len(roadPlan); y++ {
		for x := 0; x < len(roadPlan[y]); x++ {
			if roadPlan[y][x].cart != nil {
				fmt.Printf("%s", string(roadPlan[y][x].cart.direction.image))
			} else {
				fmt.Printf("%s", string(roadPlan[y][x].path))
			}
		}
		fmt.Print("\n")
	}
}

// Run tick
func tick(roadPlan roadPlanType, moveID int, cardCnt *int) (crashX, crashY int) {
	crashX, crashY = -1, -1
	for y := 0; y < len(roadPlan); y++ {
		for x := 0; x < len(roadPlan[y]); x++ {
			if roadPlan[y][x].cart != nil && roadPlan[y][x].cart.moveID != moveID {
				newX := x + roadPlan[y][x].cart.direction.horizontal
				newY := y + roadPlan[y][x].cart.direction.vertical
				if roadPlan[newY][newX].cart != nil {
					// First collision
					if crashX == -1 {
						crashX = newX
						crashY = newY
					}
					// Remove carts
					roadPlan[y][x].cart = nil
					roadPlan[newY][newX].cart = nil
					*cardCnt -= 2
					continue
				}
				cart := roadPlan[y][x].cart
				cart.moveID = moveID
				roadPlan[y][x].cart = nil
				roadPlan[newY][newX].cart = cart
				turnCnt := -1
				if roadPlan[newY][newX].path == '+' {
					turnCnt = roadPlan[newY][newX].cart.turnCnt
					cart.turnCnt = (turnCnt + 1) % 3
				}
				newDirection, ok := turnMap[turnItemType{cart.direction, turnCnt, roadPlan[newY][newX].path}]
				if ok {
					cart.direction = newDirection
				}
			}
		}
	}
	return
}

// Map processing
func processPlan(roadPlan roadPlanType, cartCnt int) (crashX, crashY, lastX, lastY int) {
	crashX, crashY, lastX, lastY = -1, -1, -1, -1
	x, y := -1, -1
	var i int
	for i = 1; ; i++ {
		x, y = tick(roadPlan, i, &cartCnt)
		if x != -1 && crashX == -1 {
			crashX = x
			crashY = y
		}
		if cartCnt == 0 {
			break
		}
		if cartCnt == 1 {
			for y := 0; y < len(roadPlan); y++ {
				for x := 0; x < len(roadPlan[y]); x++ {
					if roadPlan[y][x].cart != nil {
						lastX, lastY = x, y
					}
				}
			}
			break
		}
	}
	return
}

func main() {
	roadPlan, cartCnt := loadFile("input.txt")
	crashX, crashY, lastX, lastY := processPlan(roadPlan, cartCnt)
	fmt.Println("Crash at [", crashX, crashY, "] last cart position [", lastX, lastY, "]")
}
