package main

import (
	"reflect"
	"testing"
)

func TestLoadFile(t *testing.T) {
	log := []string{
		"[1518-10-12 23:58] Guard #421 begins shift",
		"[1518-10-13 00:19] falls asleep",
		"[1518-10-14 00:50] wakes up",
	}
	inputArray := loadFile("test1.txt")
	if !reflect.DeepEqual(inputArray, log) {
		t.Errorf("Expected sorted %s, got %s", log, inputArray)
	}
}

func TestPupulateSchedule(t *testing.T) {
	testSchedule := &[60]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 2, 2, 1, 1, 1,
		1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	activityMap := make(activityType)
	guardID := 7
	activityMap[guardID] = &[60]int{}
	pupulateSchedule(activityMap, guardID, 10, 25)
	pupulateSchedule(activityMap, guardID, 15, 17)
	schedule := activityMap[guardID]
	if !reflect.DeepEqual(schedule, testSchedule) {
		t.Error("Expected ", testSchedule, " got ", schedule)
	}
}

func TestCreateStats(t *testing.T) {

	testMap := make(activityType)
	testSchedule7 := &[60]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 2, 2, 1, 1, 1,
		1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
	}
	testSchedule8 := &[60]int{
		0, 0, 0, 0, 0, 3, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	testMap[7] = testSchedule7
	testMap[8] = testSchedule8

	inputArray := []string{
		"[1518-03-01 23:58] Guard #7 begins shift",
		"[1518-03-02 00:10] falls asleep",
		"[1518-03-02 00:25] wakes up",
		"[1518-03-03 00:00] Guard #7 begins shift",
		"[1518-03-03 00:15] falls asleep",
		"[1518-03-03 00:17] wakes up",
		"[1518-03-03 00:55] falls asleep",
		"[1518-03-01 23:58] Guard #8 begins shift",
		"[1518-03-02 00:05] falls asleep",
		"[1518-03-02 00:06] wakes up",
		"[1518-03-03 00:00] Guard #8 begins shift",
		"[1518-03-03 00:05] falls asleep",
		"[1518-03-03 00:06] wakes up",
		"[1518-03-04 00:00] Guard #8 begins shift",
		"[1518-03-04 00:05] falls asleep",
		"[1518-03-04 00:06] wakes up",
	}

	activityMap := createStats(inputArray)
	for k, v := range testMap {
		for i := 0; i < len(v); i++ {
			if testMap[k][i] != activityMap[k][i] {
				t.Error("Expected ", testMap[k], " got ", activityMap[k])
				break
			}
		}
	}

	crc1 := foundGuardCRC1(activityMap)
	if crc1 != 7*15 {
		t.Error("CRC1 Expected ", 7*15, " got ", crc1)
	}
	crc2 := foundGuardCRC2(activityMap)
	if crc2 != 8*5 {
		t.Error("CRC2 Expected ", 8*5, " got ", crc2)
	}
}
