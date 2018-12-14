package main

import "fmt"

const nrOfRecipts = 825401

var pattern = reciptType{8, 2, 5, 4, 0, 1}

type reciptType []byte

type labType struct {
	elf1Pos int
	elf2Pos int
	recits  reciptType
}

func initLab() (lab labType) {
	lab = labType{0, 1, make(reciptType, 0, nrOfRecipts+11)}
	lab.recits = append(lab.recits, 3, 7)
	return
}

func generateRecipts(lab *labType) {
	sum := lab.recits[lab.elf1Pos] + lab.recits[lab.elf2Pos]
	if sum > 9 {
		lab.recits = append(lab.recits, sum/10)
	}
	lab.recits = append(lab.recits, sum%10)
	lab.elf1Pos = (lab.elf1Pos + int(lab.recits[lab.elf1Pos]) + 1) % len(lab.recits)
	lab.elf2Pos = (lab.elf2Pos + int(lab.recits[lab.elf2Pos]) + 1) % len(lab.recits)
}

func findScore(lab *labType, pattern reciptType) (countLeft int) {
	var i int
	countLeft = -1
	offset := 0
	for {
		if len(lab.recits) < len(pattern)+1 {
			generateRecipts(lab)
			continue
		}
		for i = 0; i < len(pattern); i++ {
			if lab.recits[len(lab.recits)-1-offset-i] != pattern[len(pattern)-1-i] {
				break
			}
		}
		if i == len(pattern) {
			countLeft = len(lab.recits) - len(pattern) - offset
			return
		}
		if offset == 0 {
			offset = 1
			continue
		}
		generateRecipts(lab)
		offset = 0
	}
}

func main() {
	lab := initLab()
	for len(lab.recits) < nrOfRecipts+10 {
		generateRecipts(&lab)
	}
	fmt.Println("Scores ", lab.recits[nrOfRecipts:nrOfRecipts+10])
	lab = initLab()
	counterLeft := findScore(&lab, pattern)
	fmt.Println("Scores before", pattern, " = ", counterLeft)
}
