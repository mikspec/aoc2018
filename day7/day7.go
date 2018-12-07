package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type stepType struct {
	step    string
	follows string
}

type stepMapItem struct {
	prior []string
	next  []string
}

type graphType map[string]*stepMapItem

type workerType struct {
	step string
	time int
}

// File loading generates array of steps
func loadFile(name string) (cords []stepType) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArray := make([]stepType, 0)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("Step ([A-Z]{1}) must be finished before step ([A-Z]{1}) can begin.")
	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			command := re.FindStringSubmatch(i)
			inputArray = append(inputArray, stepType{command[1], command[2]})
		}
	}
	return inputArray
}

// Create graph
func createMap(inputArray []stepType) graphType {
	graph := make(graphType)
	for _, v := range inputArray {
		step, ok := graph[v.step]
		if !ok {
			step = &stepMapItem{[]string{}, []string{}}
			graph[v.step] = step
		}
		step.next = append(step.next, v.follows)
		prev, ok := graph[v.follows]
		if !ok {
			prev = &stepMapItem{[]string{}, []string{}}
			graph[v.follows] = prev
		}
		prev.prior = append(prev.prior, v.step)
	}
	return graph
}

// Remove step from graph
func updateGraph(graph graphType, step string) {
	for _, follows := range graph[step].next {
		priorArray := graph[follows].prior
		tmp := priorArray[:0]
		for _, p := range priorArray {
			if p != step {
				tmp = append(tmp, p)
			}
		}
		graph[follows].prior = tmp
	}
	delete(graph, step)
}

// Find path in graph
func processGraph(graph graphType) string {
	path := make([]string, 0)
	for {
		allowedSteps := make([]string, 0)
		for k, v := range graph {
			if len(v.prior) == 0 {
				allowedSteps = append(allowedSteps, k)
			}
		}
		if len(allowedSteps) == 0 {
			return strings.Join(path, "")
		}
		sort.Strings(allowedSteps)
		step := allowedSteps[0]
		path = append(path, step)
		updateGraph(graph, step)
	}
}

// Update workers queue
func updatePath(graph graphType, path []string, workers []workerType, offset int) ([]string, []workerType, int) {
	minCost := offset + 27
	minID := -1
	minStep := ""
	for i, v := range workers {
		if v.time < minCost {
			minStep = v.step
			minCost = v.time
			minID = i
		}
	}
	for i := range workers {
		workers[i].time -= minCost
	}
	if minID > 0 {
		workers = append(workers[0:minID], workers[minID+1:]...)
	} else {
		workers = workers[1:]
	}
	path = append(path, minStep)
	updateGraph(graph, minStep)

	return path, workers, minCost
}

// Find path with workers
func processGraphWithWorkers(graph graphType, workerCnt int, offset int) (string, int) {
	timeSum := 0
	timeTick := 0
	path := make([]string, 0)
	workers := make([]workerType, 0)
	for {
		allowedSteps := make([]string, 0)
		for k, v := range graph {
			if len(v.prior) == 0 {
				allowedSteps = append(allowedSteps, k)
			}
		}
		if len(allowedSteps) == 0 {
			return strings.Join(path, ""), timeSum
		}
	AllowedSteps:
		for _, step := range allowedSteps {
			for _, v := range workers {
				if v.step == step {
					continue AllowedSteps
				}
			}
			if len(workers) < workerCnt {
				workers = append(workers, workerType{step, offset + int(step[0]) - 64})
			}
		}
		path, workers, timeTick = updatePath(graph, path, workers, offset)
		timeSum += timeTick
	}
}

func main() {
	inputArray := loadFile("input.txt")
	graph := createMap(inputArray)
	path := processGraph(graph)
	fmt.Println("Path:", path)
	graph = createMap(inputArray)
	path2, time := processGraphWithWorkers(graph, 5, 60)
	fmt.Println("Path2:", path2, "time", time)
}
