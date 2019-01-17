package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type objectEnum int

// Object type
const (
	_ = iota
	ELF
	GOBLIN
	WALL
)

const (
	initialPower = 200
	attackPowerG = 3
)

var attackPowerE = 3
var dirs = []struct{ dirX, dirY int }{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

type objectMap map[int]*objectType

type objectType struct {
	objectID  int
	objectTyp objectEnum
	power     int
	moveCnt   int
	posX      int
	posY      int
}

type battleFieldType [][]*objectType
type graphType struct {
	dist int
	posX int
	posY int
	prev *graphType
}

type graphSort []*graphType

func (s graphSort) Len() int {
	return len(s)
}
func (s graphSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s graphSort) Less(i, j int) bool {
	return s[i].dist < s[j].dist || s[i].dist == s[j].dist && s[i].posY < s[j].posY || s[i].dist == s[j].dist && s[i].posY == s[j].posY && s[i].posX < s[j].posX
}

// File loading generates road plan
func loadFile(name string) (battleField battleFieldType, goblins objectMap, elves objectMap) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	battleField = make(battleFieldType, 0)
	goblins = make(map[int]*objectType)
	elves = make(map[int]*objectType)
	scanner := bufio.NewScanner(file)
	objectCnt := 0
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			row := make([]*objectType, len(line))
			for i, v := range line {
				if v == '#' {
					row[i] = &objectType{-1, WALL, -1, 0, i, len(battleField)}
				}
				if v == 'G' {
					objectCnt++
					goblin := objectType{objectCnt, GOBLIN, initialPower, 0, i, len(battleField)}
					row[i] = &goblin
					goblins[objectCnt] = &goblin
				}
				if v == 'E' {
					objectCnt++
					elf := objectType{objectCnt, ELF, initialPower, 0, i, len(battleField)}
					row[i] = &elf
					elves[objectCnt] = &elf
				}
			}
			battleField = append(battleField, row)
		}
	}
	return
}

// Battle image
func displayBattleField(battleID int, battleField *battleFieldType, goblins *objectMap, elves *objectMap) {
	fmt.Println("Step", battleID)
	for y := range *battleField {
		for x := range (*battleField)[y] {
			object := "    "
			if (*battleField)[y][x] != nil {
				if (*battleField)[y][x].objectTyp == WALL {
					object = "####"
				}
				if (*battleField)[y][x].objectTyp == ELF {
					object = fmt.Sprintf("E%3d", (*battleField)[y][x].power)

				}
				if (*battleField)[y][x].objectTyp == GOBLIN {
					object = fmt.Sprintf("G%3d", (*battleField)[y][x].power)
				}
			}
			fmt.Print(object)
		}
		fmt.Println("")
	}
}

// Adjacent points check
func addPointsToReach(posX, posY int, battleField *battleFieldType, dist int, prev *graphType, object *objectType) []*graphType {
	posList := make([]*graphType, 0)
	for _, v := range dirs {
		newX := posX + v.dirX
		newY := posY + v.dirY
		if newX >= len((*battleField)[posY]) || newX < 0 {
			continue
		}
		if newY >= len(*battleField) || newY < 0 {
			continue
		}
		if (*battleField)[newY][newX] != nil {
			if prev != nil {
				continue
			}
			if !(prev == nil && newX == object.posX && newY == object.posY) {
				continue
			}
		}
		posList = append(posList, &graphType{dist, newX, newY, prev})
	}
	return posList
}

// BFS search for shortest distance
func populateDistance(object *objectType, battleField *battleFieldType, maxDistance int) [][]*graphType {
	distMatrix := make([][]*graphType, len(*battleField))
	for y := range *battleField {
		row := make([]*graphType, len((*battleField)[y]))
		for x := range row {
			row[x] = &graphType{maxDistance, x, y, nil}
		}
		distMatrix[y] = row
	}
	start := graphType{0, object.posX, object.posY, nil}
	distMatrix[object.posY][object.posX] = &start
	distQueue := make([]*graphType, 0)
	distQueue = append(distQueue, &start)
	for len(distQueue) > 0 {
		v := distQueue[0]
		distQueue = distQueue[1:]
		posList := addPointsToReach(v.posX, v.posY, battleField, v.dist+1, v, object)
		for i := range posList {
			if distMatrix[posList[i].posY][posList[i].posX].dist == maxDistance {
				distMatrix[posList[i].posY][posList[i].posX] = posList[i]
				distQueue = append(distQueue, posList[i])
			}
		}
	}
	return distMatrix
}

// Find min distance and enemy to attack is distance == 0
func minDistance(object *objectType, battleField *battleFieldType, enemies *objectMap, maxDistance int) (minDistance int, minEnemy *objectType, dirX, dirY int) {
	dirX, dirY = 0, 0
	minEnemy = nil
	minDistance = maxDistance
	graph := make([]*graphType, 0)
	for _, v := range *enemies {
		graph = append(graph, addPointsToReach(v.posX, v.posY, battleField, maxDistance, nil, object)...)
	}
	distMatrix := populateDistance(object, battleField, maxDistance)
	for i := range graph {
		graph[i] = distMatrix[graph[i].posY][graph[i].posX]
	}

	if len(graph) == 0 {
		return
	}
	sort.Sort(graphSort(graph))
	aim := graph[0]
	if aim.dist == maxDistance {
		return
	}
	minDistance = aim.dist
	if aim.dist == 0 {
		for _, v := range dirs {
			enemy := (*battleField)[aim.posY+v.dirY][aim.posX+v.dirX]
			if enemy != nil && enemy.objectTyp != WALL && enemy.objectTyp != object.objectTyp {
				if minEnemy == nil || enemy.power < minEnemy.power {
					minEnemy = enemy
					dirX = enemy.posX - object.posX
					dirY = enemy.posY - object.posY
				}
			}
		}
	} else {
		prev := aim
		for v := aim; v.prev != nil; v = v.prev {
			prev = v
		}
		dirX = prev.posX - prev.prev.posX
		dirY = prev.posY - prev.prev.posY
	}
	return
}

// Can object move
func canMove(object *objectType, battleField *battleFieldType, goblins *objectMap, elves *objectMap) (moveFlg bool, dirX, dirY int) {
	minDist := -1
	maxDistance := len(*battleField) * len((*battleField)[0])
	enemies := goblins
	if object.objectTyp == GOBLIN {
		enemies = elves
	}
	minDist, _, dirX, dirY = minDistance(object, battleField, enemies, maxDistance)
	return minDist < maxDistance && minDist > 0, dirX, dirY
}

// Can object attack
func canAttack(object *objectType, battleField *battleFieldType, goblins *objectMap, elves *objectMap) (attackFlg bool, enemy *objectType) {
	minDist := -1
	maxDistance := len(*battleField) * len((*battleField)[0])
	enemies := goblins
	enemy = nil
	if object.objectTyp == GOBLIN {
		enemies = elves
	}
	minDist, enemy, _, _ = minDistance(object, battleField, enemies, maxDistance)
	return minDist == 0, enemy
}

// Attack procedure
func attack(enemy *objectType, battleField *battleFieldType, goblins *objectMap, elves *objectMap) {
	if enemy.objectTyp == GOBLIN {
		enemy.power -= attackPowerE
	} else {
		enemy.power -= attackPowerG
	}
	if enemy.power <= 0 {
		(*battleField)[enemy.posY][enemy.posX] = nil
		if enemy.objectTyp == GOBLIN {
			delete(*goblins, enemy.objectID)
		} else {
			delete(*elves, enemy.objectID)
		}
	}
}

// One round of battle
func battleTick(battleID int, battleField *battleFieldType, goblins *objectMap, elves *objectMap) (endBattle bool) {
	for y := range *battleField {
		for x := range (*battleField)[y] {
			object := (*battleField)[y][x]
			if object != nil && object.objectTyp != WALL {
				if object.moveCnt == battleID {
					continue
				}
				if attackFlg, enemy := canAttack(object, battleField, goblins, elves); attackFlg {
					attack(enemy, battleField, goblins, elves)
				} else if moveFlg, dirX, dirY := canMove(object, battleField, goblins, elves); moveFlg {
					(*battleField)[y+dirY][x+dirX] = object
					(*battleField)[y][x] = nil
					object.posX += dirX
					object.posY += dirY
					if attackFlg, enemy = canAttack(object, battleField, goblins, elves); attackFlg {
						attack(enemy, battleField, goblins, elves)
					}
				} else {
					if len(*goblins) == 0 || len(*elves) == 0 {
						return true
					}
				}
				object.moveCnt = battleID
			}
		}
	}
	return false
}

// Battle simulation
func battle(battleField *battleFieldType, goblins *objectMap, elves *objectMap) (battleResult, battleID int) {
	battleID = 0
	battleResult = 0
	//displayBattleField(battleID, battleField, goblins, elves)
	for {
		battleID++
		if endBattle := battleTick(battleID, battleField, goblins, elves); endBattle {
			sum := 0
			for _, v := range *goblins {
				sum += v.power
			}
			for _, v := range *elves {
				sum += v.power
			}
			battleID--
			return sum * battleID, battleID
		}
		//displayBattleField(battleID, battleField, goblins, elves)
	}
}

// Battle simulation with growing elves power attack value
func part2(file string) (int, int, int) {
	for attackPowerE = 4; ; attackPowerE++ {
		battleField, goblins, elves := loadFile(file)
		preElvesCnt := len(elves)
		result, battleID := battle(&battleField, &goblins, &elves)
		if preElvesCnt == len(elves) {
			return result, battleID, attackPowerE
		}
	}
}

func main() {
	file := "input.txt"
	battleField, goblins, elves := loadFile(file)
	result, battleID := battle(&battleField, &goblins, &elves)
	fmt.Println("Battle result ", result, battleID)
	attackPowerE = 3
	result, battleID, attackPowerE = part2(file)
	fmt.Println("Battle 2 result ", result, battleID, attackPowerE)
}
