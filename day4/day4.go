package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type activityType map[int]*[60]int

// File loading generates sorted activity log
func loadFile(name string) []string {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArray := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			inputArray = append(inputArray, i)
		}
	}

	sort.Strings(inputArray)

	return inputArray
}

// Generate sleep statistics
func createStats(inputArray []string) activityType {
	activityMap := make(activityType)

	re := regexp.MustCompile("(\\[\\d{4}-\\d{2}-\\d{2} \\d{2}:(\\d{2})\\]) (Guard #(\\d+) begins shift|falls asleep|wakes up)")
	guardID := -1
	sleep := -1
	wake := -1
	for _, v := range inputArray {
		log := re.FindStringSubmatch(v)
		if log != nil {
			if strings.HasPrefix(log[3], "Guard") {
				if sleep != -1 {
					pupulateSchedule(activityMap, guardID, sleep, wake)
				}
				guardID, _ = strconv.Atoi(log[4])
				if _, ok := activityMap[guardID]; !ok {
					activityMap[guardID] = &[60]int{}
				}
				sleep = -1
				wake = -1
			}
			if strings.HasPrefix(log[3], "falls") {
				sleep, _ = strconv.Atoi(log[2])
			}
			if strings.HasPrefix(log[3], "wakes") {
				wake, _ = strconv.Atoi(log[2])
				pupulateSchedule(activityMap, guardID, sleep, wake)
				sleep = -1
				wake = -1
			}
		}
	}
	if sleep != -1 {
		pupulateSchedule(activityMap, guardID, sleep, wake)
	}

	return activityMap
}

// Populate schedule for given time window
func pupulateSchedule(activityMap activityType, guardID int, sleep int, wake int) {
	endTime := len(activityMap[guardID])
	if wake != -1 {
		endTime = wake
	}
	schedule := activityMap[guardID]
	for i := sleep; i < endTime; i++ {
		schedule[i]++
	}
}

// Strategy 1
func foundGuardCRC1(activityMap activityType) int {
	guard := -1
	maxsum := -1
	minute := -1
	counter := -1
	for id, schedule := range activityMap {
		sum := 0
		for _, tim := range schedule {
			if tim > 0 {
				sum += tim
			}
		}
		if sum > maxsum {
			maxsum = sum
			guard = id
		}
	}
	for i, val := range activityMap[guard] {
		if val > counter {
			minute = i
			counter = val
		}
	}
	return guard * minute
}

// Strategy 2
func foundGuardCRC2(activityMap activityType) int {
	guard := -1
	minute := -1
	counter := -1
	for id, schedule := range activityMap {
		for i, tim := range schedule {
			if tim > counter {
				guard = id
				minute = i
				counter = tim
			}
		}
	}
	return guard * minute
}

func main() {
	inputArray := loadFile("input.txt")
	activityMap := createStats(inputArray)
	fmt.Println("Strategy 1:", foundGuardCRC1(activityMap))
	fmt.Println("Strategy 2:", foundGuardCRC2(activityMap))
}
