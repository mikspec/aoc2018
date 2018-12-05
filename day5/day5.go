package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// File loading generates sorted activity log
func loadFile(name string) string {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

// Check byte array for matching letters
func processData(data []byte) []byte {
	for i := 1; i < len(data); {
		a := data[i-1]
		b := data[i]
		if a < b {
			a, b = data[i], data[i-1]
		}
		if a-b == 32 {
			data = append(data[:i-1], data[i+1:]...)
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	return data
}

func optimisePolymer(sequence string) (int, byte) {
	minlen := len(sequence)
	minchar := byte(0)
	// Scan ascii characters
	for i := 65; i < 91; i++ {
		str := strings.Replace(sequence, string(i), "", -1)
		str = strings.Replace(str, string(i+32), "", -1)
		polymer := processData([]byte(str))
		if minlen > len(polymer) {
			minlen = len(polymer)
			minchar = byte(i)
		}
	}
	return minlen, minchar
}

func main() {
	str := loadFile("input.txt")
	result := processData([]byte(str))
	strVal := string(result)
	fmt.Printf("Part 1: %d\n", len(strVal))

	minlen, minchar := optimisePolymer(str)
	fmt.Println("Optimised polymer: ", minlen, string(minchar))
}
