package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const initialStatePrefix = "initial state: "

type rulesType map[byte]bool

type potsArrayType struct {
	zeroPos int
	plants  []bool
}

// File loading generates map of rules and pots array
func loadFile(name string) (rules rulesType, potsArray *potsArrayType) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	initialState := ""
	potsArray = &potsArrayType{}
	potsArray.plants = make([]bool, 2)
	potsArray.zeroPos = 2
	rules = make(rulesType)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("([.#]{5}) => ([.#]{1})")
	for scanner.Scan() {
		if i := scanner.Text(); len(i) > 0 {
			if strings.HasPrefix(i, initialStatePrefix) {
				initialState = strings.TrimPrefix(i, initialStatePrefix)
			} else {
				rulesStr := re.FindStringSubmatch(i)
				key, val := decodeRule(rulesStr)
				rules[key] = val
			}
		}
	}
	for _, v := range initialState {
		if v == '#' {
			potsArray.plants = append(potsArray.plants, true)
		} else {
			potsArray.plants = append(potsArray.plants, false)
		}
	}
	// Append two dots at the end of array
	potsArray.plants = append(potsArray.plants, false, false)
	return rules, potsArray
}

// Decode rule converts string rule into key - value
func decodeRule(ruleStr []string) (key byte, value bool) {
	key = 0
	value = false
	for i, v := range ruleStr[1] {
		if v == '#' {
			key += 1 << uint(4-i)
		}
	}
	if ruleStr[2][0] == '#' {
		value = true
	}
	return
}

// Next generation processing
func nextGeneration(rules *rulesType, pots *potsArrayType) *potsArrayType {
	nextPots := new(potsArrayType)
	nextPots.zeroPos = pots.zeroPos
	nextPots.plants = make([]bool, 0)
	if pots.plants[0] || pots.plants[1] {
		nextPots.plants = append(nextPots.plants, false, false)
		nextPots.zeroPos += 2
	}
	for i := range pots.plants {
		pattern := byte(0)
		for j := -2; j < 3; j++ {
			if i+j >= 0 && i+j < len(pots.plants) && pots.plants[i+j] {
				pattern += 1 << uint(4-j-2)
			}
		}
		nextPots.plants = append(nextPots.plants, (*rules)[pattern])
	}
	lenPots := len(pots.plants)
	if pots.plants[lenPots-1] || pots.plants[lenPots-2] || pots.plants[lenPots-3] || pots.plants[lenPots-4] {
		nextPots.plants = append(nextPots.plants, false, false)
	}
	return nextPots
}

// Calculate CRC
func calculateCRC(pots *potsArrayType) (crc int) {
	crc = 0
	for i, v := range pots.plants {
		if v {
			crc += i - pots.zeroPos
		}
	}
	return
}

func main() {
	rules, state := loadFile("input.txt")
	for i := 0; i < 20; i++ {
		state = nextGeneration(&rules, state)
	}
	crc := calculateCRC(state)
	fmt.Println("20 generation CRC = ", crc)
	rules, state = loadFile("input.txt")
	for i := 0; i < 500; i++ {
		state = nextGeneration(&rules, state)
	}
	crc = calculateCRC(state)
	fmt.Println("50 bilions generation CRC = ", crc*100000000)
}
