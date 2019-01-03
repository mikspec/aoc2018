package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	attackPower  = 3
)

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

func displayBattleField(battleField *battleFieldType, goblins *objectMap, elves *objectMap) {
	for _, v := range *goblins {
		fmt.Print(v.power, " ")
	}
	fmt.Println("")
	for _, v := range *elves {
		fmt.Print(v.power, " ")
	}
	fmt.Println("")
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

func populateDistance(posX, posY, distance int, battleField *[][]int) {
	if posX > 0 && (*battleField)[posY][posX-1] > distance {
		(*battleField)[posY][posX-1] = distance
		populateDistance(posX-1, posY, distance+1, battleField)
	}
	if posX < len((*battleField)[posY])-1 && (*battleField)[posY][posX+1] > distance {
		(*battleField)[posY][posX+1] = distance
		populateDistance(posX+1, posY, distance+1, battleField)
	}
	if posY > 0 && (*battleField)[posY-1][posX] > distance {
		(*battleField)[posY-1][posX] = distance
		populateDistance(posX, posY-1, distance+1, battleField)
	}
	if posY < len(*battleField)-1 && (*battleField)[posY+1][posX] > distance {
		(*battleField)[posY+1][posX] = distance
		populateDistance(posX, posY+1, distance+1, battleField)
	}
}

func minDistance(object *objectType, battleField *battleFieldType, enemies *objectMap, maxDistance int) (minDistance int, minEnemy *objectType, dirX, dirY int) {
	dirX, dirY = 0, 0
	minEnemy = nil
	minPower := initialPower
	battleTmp := make([][]int, len(*battleField))
	for y := range *battleField {
		battleTmp[y] = make([]int, len((*battleField)[y]))
		for x := range (*battleField)[y] {
			if (*battleField)[y][x] != nil && !(x == object.posX && y == object.posY) {
				battleTmp[y][x] = -1
			} else {
				battleTmp[y][x] = maxDistance
			}
		}
	}
	for _, enemy := range *enemies {
		battleTmp[enemy.posY][enemy.posX] = 0
		populateDistance(enemy.posX, enemy.posY, 1, &battleTmp)
	}
	minDistance = battleTmp[object.posY][object.posX]
	if minDistance == 1 {
		if object.posY < len(battleTmp)-1 && battleTmp[object.posY+1][object.posX] == 0 {
			minEnemy = (*battleField)[object.posY+1][object.posX]
			minPower = minEnemy.power
		}
		if object.posX < len(battleTmp[object.posY])-1 && battleTmp[object.posY][object.posX+1] == 0 {
			enemy := (*battleField)[object.posY][object.posX+1]
			if enemy.power <= minPower {
				minEnemy = enemy
				minPower = enemy.power
			}
		}
		if object.posX > 0 && battleTmp[object.posY][object.posX-1] == 0 {
			enemy := (*battleField)[object.posY][object.posX-1]
			if enemy.power <= minPower {
				minEnemy = enemy
				minPower = enemy.power
			}
		}
		if object.posY > 0 && battleTmp[object.posY-1][object.posX] == 0 {
			enemy := (*battleField)[object.posY-1][object.posX]
			if enemy.power <= minPower {
				minEnemy = enemy
				minPower = enemy.power
			}
		}
		return
	}
	for y := range *battleField {
		for x := range (*battleField)[y] {
			if (*battleField)[y][x] != nil {
				battleTmp[y][x] = -1
			} else {
				battleTmp[y][x] = maxDistance
			}
		}
	}
	battleTmp[object.posY][object.posX] = 0
	populateDistance(object.posX, object.posY, 1, &battleTmp)
	y, x := 0, 0
	distance := minDistance - 1
	enemyType := objectEnum(ELF)
	if object.objectTyp == ELF {
		enemyType = objectEnum(GOBLIN)
	}
loopBattle:
	for y = range battleTmp {
		for x = range battleTmp[y] {
			if battleTmp[y][x] == distance {
				if y < len(battleTmp)-1 && (*battleField)[y+1][x] != nil && (*battleField)[y+1][x].objectTyp == enemyType {
					break loopBattle
				} else if x < len(battleTmp[y])-1 && (*battleField)[y][x+1] != nil && (*battleField)[y][x+1].objectTyp == enemyType {
					break loopBattle
				} else if x > 0 && (*battleField)[y][x-1] != nil && (*battleField)[y][x-1].objectTyp == enemyType {
					break loopBattle
				} else if y > 0 && (*battleField)[y-1][x] != nil && (*battleField)[y-1][x].objectTyp == enemyType {
					break loopBattle
				}
			}
		}
	}
	for {
		if distance == 1 {
			break
		}
		if y < len(battleTmp)-1 && battleTmp[y+1][x] == distance-1 {
			y++
		} else if x < len(battleTmp[y])-1 && battleTmp[y][x+1] == distance-1 {
			x++
		} else if x > 0 && battleTmp[y][x-1] == distance-1 {
			x--
		} else if y > 0 && battleTmp[y-1][x] == distance-1 {
			y--
		}
		distance--
	}
	dirX = x - object.posX
	dirY = y - object.posY
	return
}

func canMove(object *objectType, battleField *battleFieldType, goblins *objectMap, elves *objectMap) (moveFlg bool, dirX, dirY int) {
	maxDistance := len(*battleField) + len((*battleField)[0]) + 2
	minDist := -1
	enemies := goblins
	if object.objectTyp == GOBLIN {
		enemies = elves
	}
	minDist, _, dirX, dirY = minDistance(object, battleField, enemies, maxDistance)
	return (minDist < maxDistance) && (minDist > 1), dirX, dirY
}

func canAttack(object *objectType, battleField *battleFieldType, goblins *objectMap, elves *objectMap) (attackFlg bool, enemy *objectType) {
	maxDistance := len(*battleField) + len((*battleField)[0]) + 2
	minDist := -1
	enemies := goblins
	enemy = nil
	if object.objectTyp == GOBLIN {
		enemies = elves
	}
	minDist, enemy, _, _ = minDistance(object, battleField, enemies, maxDistance)
	return minDist == 1, enemy
}

func battleTick(battleID int, battleField *battleFieldType, goblins *objectMap, elves *objectMap) (endBattle, fullBattle bool) {
	for y := range *battleField {
		for x := range (*battleField)[y] {
			if len(*goblins) == 0 || len(*elves) == 0 {
				return true, false
			}
			object := (*battleField)[y][x]
			if object != nil && object.objectTyp != WALL {
				moveFlg, dirX, dirY := canMove(object, battleField, goblins, elves)
				if moveFlg && object.moveCnt != battleID {
					(*battleField)[y+dirY][x+dirX] = object
					(*battleField)[y][x] = nil
					object.posX += dirX
					object.posY += dirY
				}
				attackFlg, enemy := canAttack(object, battleField, goblins, elves)
				if attackFlg && object.moveCnt != battleID {
					enemy.power -= attackPower
					if enemy.power <= 0 {
						(*battleField)[enemy.posY][enemy.posX] = nil
						if enemy.objectTyp == GOBLIN {
							delete(*goblins, enemy.objectID)
						} else {
							delete(*elves, enemy.objectID)
						}
					}
				}
				object.moveCnt = battleID
			}
		}
	}
	return false, false
}

func sumObjects(battleField *battleFieldType) (sumGoblins, sumElves int) {
	for y := range *battleField {
		for x := range (*battleField)[y] {
			if (*battleField)[y][x] != nil && (*battleField)[y][x].objectTyp != WALL {
				if (*battleField)[y][x].objectTyp == GOBLIN {
					sumGoblins++
				}
				if (*battleField)[y][x].objectTyp == ELF {
					sumElves++
				}
			}
		}
	}
	return
}

func battle(battleField *battleFieldType, goblins *objectMap, elves *objectMap) (battleResult, battleID int) {
	battleID = 0
	battleResult = 0
	for {
		battleID++
		if endBattle, fullBattle := battleTick(battleID, battleField, goblins, elves); endBattle {
			sum := 0
			for _, v := range *goblins {
				sum += v.power
			}
			for _, v := range *elves {
				sum += v.power
			}
			if !fullBattle {
				battleID--
			}
			return sum * battleID, battleID
		}
	}
}

func main() {
	battleField, goblins, elves := loadFile("input.txt")
	result, battleID := battle(&battleField, &goblins, &elves)
	fmt.Println("Battle result ", result, battleID)
}
