package main

import "fmt"

// Number of players
const numOfPlayers = 468

// Last marble worth
const lastMarble = 71843

// Player scores
type scoreTableType = []int

type node struct {
	prev *node
	next *node
	val  int
}

type list struct {
	current *node
}

func (l *list) insert(val int) {
	if l.current == nil { // add first elem
		item := &node{
			val: val,
		}
		item.next = item
		item.prev = item
		l.current = item
	} else {
		right := l.current.next
		item := &node{
			val:  val,
			next: right,
			prev: l.current,
		}
		right.prev = item
		l.current.next = item
		l.current = item
	}
}

func (l *list) remove() int {
	if l.current == nil {
		return 0
	}
	l.current.prev.next = l.current.next
	l.current.next.prev = l.current.prev
	removed := l.current
	l.current = l.current.next
	return removed.val
}

func (l *list) moveRight(step int) {
	for i := 0; i < step; i++ {
		l.current = l.current.next
	}
}

func (l *list) moveLeft(step int) {
	for i := 0; i < step; i++ {
		l.current = l.current.prev
	}
}

func (l *list) display() {
	index := l.current
	for {
		fmt.Print(index.val, " ")
		index = index.next
		if index == l.current {
			break
		}
	}
	fmt.Print("\n")
}

func runGame(numOfPlayers, lastMarble int) (maxScore int) {
	maxScore = 0
	scoreTable := make(scoreTableType, numOfPlayers)
	circle := list{}
	circle.insert(0)
	for i := 1; i <= lastMarble; i++ {
		if i%23 != 0 { // clockwise move
			circle.moveRight(1)
			circle.insert(i)
		} else {
			player := (i - 1) % numOfPlayers
			circle.moveLeft(7)
			score := i + circle.remove()
			scoreTable[player] += score
		}
	}
	for _, v := range scoreTable {
		if v > maxScore {
			maxScore = v
		}
	}
	return maxScore
}

func main() {
	score := runGame(numOfPlayers, lastMarble)
	fmt.Println("Score 1:", score)
	score = runGame(numOfPlayers, lastMarble*100)
	fmt.Println("Score 2:", score)
}
