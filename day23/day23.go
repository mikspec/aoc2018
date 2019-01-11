package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pointType struct {
	x, y, z int
}

type nanobootType struct {
	pos     pointType
	radius  int
	counter int
}

type cubeType struct {
	min     pointType
	max     pointType
	counter int
	index   int
}

// Set cube intersecrion counter with all boots
func (q *cubeType) setIntersection(nanoboots []nanobootType) {
	for _, bot := range nanoboots {
		if q.botInside(bot.pos) {
			q.counter++
			continue
		}
		wallInter := q.getProjectionToWall(bot.pos)
		dist := distance(wallInter, bot.pos)
		if dist <= bot.radius {
			q.counter++
			continue
		}
	}
}

// Get bot position projection towards nearest wall
func (q *cubeType) getProjectionToWall(pos pointType) (wall pointType) {
	wall = pos
	if pos.x < q.min.x {
		wall.x = q.min.x
	}
	if pos.x > q.max.x {
		wall.x = q.max.x
	}
	if pos.y < q.min.y {
		wall.y = q.min.y
	}
	if pos.y > q.max.y {
		wall.y = q.max.y
	}
	if pos.z < q.min.z {
		wall.z = q.min.z
	}
	if pos.z > q.max.z {
		wall.z = q.max.z
	}
	return
}

// Is boot inside cube
func (q *cubeType) botInside(pos pointType) bool {
	return q.min.x <= pos.x && q.max.x >= pos.x && q.min.y <= pos.y && q.max.y >= pos.y && q.min.z <= pos.z && q.max.z >= pos.z
}

// Get cube center
func (q *cubeType) getDistToCenter() int {
	cubeCenter := pointType{(q.min.x + q.max.x) / 2, (q.min.y + q.max.y) / 2, (q.min.z + q.max.z) / 2}
	return distance(cubeCenter, pointType{0, 0, 0})
}

// Get cube size
func (q *cubeType) getSize() int {
	return (q.max.x - q.min.x + 1) * (q.max.y - q.min.y + 1) * (q.max.z - q.min.z + 1)
}

// Cube split to eight subcubes
func (q *cubeType) splitQube() []cubeType {
	cubes := make([]cubeType, 8)
	cubes[0] = cubeType{
		pointType{q.min.x, q.min.y, q.min.z},
		pointType{(q.min.x + q.max.x) / 2, (q.min.y + q.max.y) / 2, (q.min.z + q.max.z) / 2}, 0, -1}
	cubes[1] = cubeType{
		pointType{(q.min.x+q.max.x)/2 + 1, q.min.y, q.min.z},
		pointType{q.max.x, (q.min.y + q.max.y) / 2, (q.min.z + q.max.z) / 2}, 0, -1}
	cubes[2] = cubeType{
		pointType{q.min.x, (q.min.y+q.max.y)/2 + 1, q.min.z},
		pointType{(q.min.x + q.max.x) / 2, q.max.y, (q.min.z + q.max.z) / 2}, 0, -1}
	cubes[3] = cubeType{
		pointType{(q.min.x+q.max.x)/2 + 1, (q.min.y+q.max.y)/2 + 1, q.min.z},
		pointType{q.max.x, q.max.y, (q.min.z + q.max.z) / 2}, 0, -1}
	cubes[4] = cubeType{
		pointType{q.min.x, q.min.y, (q.min.z+q.max.z)/2 + 1},
		pointType{(q.min.x + q.max.x) / 2, (q.min.y + q.max.y) / 2, q.max.z}, 0, -1}
	cubes[5] = cubeType{
		pointType{(q.min.x+q.max.x)/2 + 1, q.min.y, (q.min.z+q.max.z)/2 + 1},
		pointType{q.max.x, (q.min.y + q.max.y) / 2, q.max.z}, 0, -1}
	cubes[6] = cubeType{
		pointType{q.min.x, (q.min.y+q.max.y)/2 + 1, (q.min.z+q.max.z)/2 + 1},
		pointType{(q.min.x + q.max.x) / 2, q.max.y, q.max.z}, 0, -1}
	cubes[7] = cubeType{
		pointType{(q.min.x+q.max.x)/2 + 1, (q.min.y+q.max.y)/2 + 1, (q.min.z+q.max.z)/2 + 1},
		pointType{q.max.x, q.max.y, q.max.z}, 0, -1}
	return cubes
}

type priorityQueue []*cubeType

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].counter > pq[j].counter {
		return true
	}
	if pq[i].counter == pq[j].counter {
		return pq[i].getDistToCenter() < pq[j].getDistToCenter()
	}
	return false
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*cubeType)
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

// File loading generates array of steps
func loadFile(name string) []nanobootType {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArray := make([]nanobootType, 0)
	scanner := bufio.NewScanner(file)
	//pos=<-19978367,16381145,16826174>, r=99768722
	re := regexp.MustCompile("pos=<(-?\\d+),(-?\\d+),(-?\\d+)>, r=(\\d+)")
	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			command := re.FindStringSubmatch(i)
			posX, _ := strconv.Atoi(strings.Trim(command[1], " "))
			posY, _ := strconv.Atoi(strings.Trim(command[2], " "))
			posZ, _ := strconv.Atoi(strings.Trim(command[3], " "))
			radius, _ := strconv.Atoi(strings.Trim(command[4], " "))
			inputArray = append(inputArray, nanobootType{pointType{posX, posY, posZ}, radius, 0})
		}
	}
	return inputArray
}

// Calculate distance
func distance(point1 pointType, point2 pointType) (dist int) {
	distX := point1.x - point2.x
	if distX < 0 {
		distX *= -1
	}
	distY := point1.y - point2.y
	if distY < 0 {
		distY *= -1
	}
	distZ := point1.z - point2.z
	if distZ < 0 {
		distZ *= -1
	}
	dist = distX + distY + distZ
	return
}

// Find strongest boot
func findStrongest(inputArray []nanobootType) (int, int) {
	maxRadius := 0
	maxInd := -1
	for i, v := range inputArray {
		if v.radius > maxRadius {
			maxRadius = v.radius
			maxInd = i
		}
	}
	for j := range inputArray {
		dist := distance(inputArray[maxInd].pos, inputArray[j].pos)
		if dist <= maxRadius {
			inputArray[maxInd].counter++
		}
	}
	return maxInd, inputArray[maxInd].counter
}

// Get space limits
func getSpaceCorners(inputArray []nanobootType) (cornerMin, cornerMax pointType) {
	cornerMin.x, cornerMax.x = inputArray[0].pos.x, inputArray[0].pos.x
	cornerMin.y, cornerMax.y = inputArray[0].pos.y, inputArray[0].pos.y
	cornerMin.z, cornerMax.y = inputArray[0].pos.z, inputArray[0].pos.z
	for _, v := range inputArray {
		if cornerMin.x > v.pos.x {
			cornerMin.x = v.pos.x
		}
		if cornerMax.x < v.pos.x {
			cornerMax.x = v.pos.x
		}
		if cornerMin.y > v.pos.y {
			cornerMin.y = v.pos.y
		}
		if cornerMax.y < v.pos.y {
			cornerMax.y = v.pos.y
		}
		if cornerMin.z > v.pos.z {
			cornerMin.z = v.pos.z
		}
		if cornerMax.z < v.pos.z {
			cornerMax.z = v.pos.z
		}
	}
	return
}

// Find best position in range most boots
func findBestPosition(nanoboots []nanobootType) (bestPos pointType) {
	cornerMin, cornerMax := getSpaceCorners(nanoboots)
	startCube := cubeType{cornerMin, cornerMax, 0, 0}
	startCube.setIntersection(nanoboots)
	pq := make(priorityQueue, 1)
	pq[0] = &startCube
	heap.Init(&pq)

	for len(pq) > 0 {
		// Get cube with maximum intersection number
		cube := heap.Pop(&pq).(*cubeType)
		// First cube with size 1 is the answer
		if cube.getSize() == 1 {
			bestPos = cube.min
			return
		}
		// Skip cube with zero intersections
		if cube.counter == 0 {
			continue
		}
		// Divide and conquer
		newCubes := cube.splitQube()
		for i := range newCubes {
			newQube := newCubes[i]
			newQube.setIntersection(nanoboots)
			if newQube.counter > 0 {
				heap.Push(&pq, &newQube)
			}
		}
	}
	return
}

func main() {
	inputArray := loadFile("input.txt")
	ind, cnt := findStrongest(inputArray)
	fmt.Println("Strongest nanoboot", ind, "has", cnt, "nanoboots in range")
	inputArray = loadFile("input.txt")
	bestPos := findBestPosition(inputArray)
	fmt.Println("Best position", bestPos, "distance", distance(pointType{0, 0, 0}, bestPos))
}
