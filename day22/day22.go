package main

import (
	"container/heap"
	"fmt"
)

// Constans
const (
	CaveDepth  = 7305
	TargetX    = 13
	TargetY    = 734
	MulX       = 16807
	MulY       = 48271
	CaveModulo = 20183
)

// Tools
const (
	NEITHER = iota
	TORCH
	GEAR
)

type cavePlanItemType struct {
	geoIndex int
	erosion  int
}
type cavePlanType [][]cavePlanItemType

type graphItemKeyType struct {
	x, y int
	tool int
}

type graphItemType struct {
	*graphItemKeyType
	dist      int
	index     int
	prev      *graphItemType
	processed bool
}

type graphMapType map[graphItemKeyType]*graphItemType

type priorityQueue []*graphItemType

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*graphItemType)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) update(item *graphItemType, distance int) {
	item.dist = distance
	heap.Fix(pq, item.index)
}

// Geological index
func calculateIndex(caveDepth, x, y, targetX, targetY int, cavePlan cavePlanType) (index int) {
	if (x == 0 && y == 0) || (x == targetX && y == targetY) {
		index = caveDepth % CaveModulo
		return
	}
	if x > 0 && y == 0 {
		index = (x*MulX + caveDepth) % CaveModulo
		return
	}
	if x == 0 && y > 0 {
		index = (y*MulY + caveDepth) % CaveModulo
		return
	}
	index = (cavePlan[y-1][x].geoIndex*cavePlan[y][x-1].geoIndex + caveDepth) % CaveModulo
	return
}

// Cave plan creation
func createPlan(caveDepth, targetX, targetY int) cavePlanType {
	cavePlan := make(cavePlanType, targetY+1)
	for y := range cavePlan {
		cavePlan[y] = make([]cavePlanItemType, targetX+1)
		for x := range cavePlan[y] {
			cavePlan[y][x].geoIndex = calculateIndex(caveDepth, x, y, targetX, targetY, cavePlan)
			cavePlan[y][x].erosion = cavePlan[y][x].geoIndex % 3
		}
	}
	return cavePlan
}

// Sum calculation
func erosionSum(cavePlan cavePlanType) (sum int) {
	for y := range cavePlan {
		for x := range cavePlan[y] {
			sum += cavePlan[y][x].erosion
		}
	}
	return
}

// Generate cave image
func displayCave(targetX, targetY int, cavePlan cavePlanType) {
	for y := range cavePlan {
		for x := range cavePlan[y] {
			if x == 0 && y == 0 {
				fmt.Print("M")
				continue
			}
			if x == targetX && y == targetY {
				fmt.Print("T")
				continue
			}
			if cavePlan[y][x].erosion == 0 {
				fmt.Print(".")
			}
			if cavePlan[y][x].erosion == 1 {
				fmt.Print("=")
			}
			if cavePlan[y][x].erosion == 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Add new row to the cave
func addRow(caveDepth, targetX, targetY int, cavePlan cavePlanType) cavePlanType {
	cavePlan = append(cavePlan, make([]cavePlanItemType, len(cavePlan[0])))
	y := len(cavePlan) - 1
	for x := range cavePlan[y] {
		cavePlan[y][x].geoIndex = calculateIndex(caveDepth, x, y, targetX, targetY, cavePlan)
		cavePlan[y][x].erosion = cavePlan[y][x].geoIndex % 3
	}
	return cavePlan
}

// Add new column to the cave
func addColumn(caveDepth, targetX, targetY int, cavePlan cavePlanType) cavePlanType {
	for y := range cavePlan {
		cavePlan[y] = append(cavePlan[y], cavePlanItemType{})
		x := len(cavePlan[y]) - 1
		cavePlan[y][x].geoIndex = calculateIndex(caveDepth, x, y, targetX, targetY, cavePlan)
		cavePlan[y][x].erosion = cavePlan[y][x].geoIndex % 3
	}
	return cavePlan
}

// Create new item if not exists or update distance
func createGraphItem(x, y, tool, dist int, pq *priorityQueue, graphMap *graphMapType, head *graphItemType, cavePlan *cavePlanType) {
	key := graphItemKeyType{x, y, tool}
	if item, ok := (*graphMap)[key]; !ok {
		item = &graphItemType{&key, dist, -1, head, false}
		(*graphMap)[key] = item
		heap.Push(pq, item)
	} else {
		if !item.processed && dist < item.dist {
			item.dist = dist
			item.prev = head
			heap.Fix(pq, item.index)
		}
	}
}

// Find rescue path from 0,0 to target
func findRescuePathLength(caveDepth, targetX, targetY int, cavePlan cavePlanType) (int, cavePlanType) {
	graphMap := make(graphMapType)
	startTorchKey := graphItemKeyType{0, 0, TORCH}
	startTorch := graphItemType{&startTorchKey, 0, -1, nil, false}
	pq := make(priorityQueue, 1)
	pq[0] = &startTorch
	graphMap[startTorchKey] = &startTorch
	heap.Init(&pq)

	for len(pq) > 0 {
		head := heap.Pop(&pq).(*graphItemType)
		head.processed = true
		// If target reached return distance
		if head.x == targetX && head.y == targetY && head.tool == TORCH {
			return head.dist, cavePlan
		}
		// Left
		if head.x > 0 && cavePlan[head.y][head.x-1].erosion != head.tool {
			createGraphItem(head.x-1, head.y, head.tool, head.dist+1, &pq, &graphMap, head, &cavePlan)
		}
		// Up
		if head.y > 0 && cavePlan[head.y-1][head.x].erosion != head.tool {
			createGraphItem(head.x, head.y-1, head.tool, head.dist+1, &pq, &graphMap, head, &cavePlan)
		}
		// Generate additional cave plan if required
		if head.x == len(cavePlan[head.y])-1 {
			cavePlan = addColumn(caveDepth, targetX, targetY, cavePlan)
		}
		if head.y == len(cavePlan)-1 {
			cavePlan = addRow(caveDepth, targetX, targetY, cavePlan)
		}
		// Right
		if cavePlan[head.y][head.x+1].erosion != head.tool {
			createGraphItem(head.x+1, head.y, head.tool, head.dist+1, &pq, &graphMap, head, &cavePlan)
		}
		// Down
		if cavePlan[head.y+1][head.x].erosion != head.tool {
			createGraphItem(head.x, head.y+1, head.tool, head.dist+1, &pq, &graphMap, head, &cavePlan)
		}
		// Change the tool operation is also one of options
		newTool := (head.tool + 1) % 3
		if newTool == cavePlan[head.y][head.x].erosion {
			newTool = (newTool + 1) % 3
		}
		createGraphItem(head.x, head.y, newTool, head.dist+7, &pq, &graphMap, head, &cavePlan)
	}
	return -1, cavePlan
}

func main() {
	cavePlan := createPlan(CaveDepth, TargetX, TargetY)
	sum := erosionSum(cavePlan)
	dist, newPlan := findRescuePathLength(CaveDepth, TargetX, TargetY, cavePlan)
	displayCave(TargetX, TargetY, newPlan)
	fmt.Println("Erosion sum", sum)
	fmt.Println("Distance ", dist)
}
