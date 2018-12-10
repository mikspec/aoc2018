package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pointType struct {
	posX, posY, speedX, speedY int
}

// File loading generates array of steps
func loadFile(name string) []pointType {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArray := make([]pointType, 0)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("position=<([ -]\\d+), ([ -]\\d+)> velocity=<([ -]\\d+), ([ -]\\d+)>")
	for scanner.Scan() {

		if i := scanner.Text(); len(i) > 0 {
			command := re.FindStringSubmatch(i)
			posX, _ := strconv.Atoi(strings.Trim(command[1], " "))
			posY, _ := strconv.Atoi(strings.Trim(command[2], " "))
			speedX, _ := strconv.Atoi(strings.Trim(command[3], " "))
			speedY, _ := strconv.Atoi(strings.Trim(command[4], " "))
			inputArray = append(inputArray, pointType{posX, posY, speedX, speedY})
		}
	}
	return inputArray
}

func getBorder(inputArray []pointType) (left, top, right, down int) {
	left = inputArray[0].posX
	top = inputArray[0].posY
	right = inputArray[0].posX
	down = inputArray[0].posY
	for _, v := range inputArray {
		if v.posX < left {
			left = v.posX
		}
		if v.posX > right {
			right = v.posX
		}
		if v.posY < top {
			top = v.posY
		}
		if v.posY > down {
			down = v.posY
		}
	}
	return
}

func processImage(inputArray []pointType) {
	for i, v := range inputArray {
		v.posX += v.speedX
		v.posY += v.speedY
		inputArray[i] = v
	}
}

func getMinSize(inputArray []pointType) (int, int) {
	left, top, right, down := getBorder(inputArray)
	minSize := (right - left) * (down - top)
	i := 0
	for ; ; i++ {
		processImage(inputArray)
		left, top, right, down = getBorder(inputArray)
		size := (right - left) * (down - top)
		if size >= minSize {
			break
		} else {
			minSize = size
		}
	}
	return minSize, i
}

func drawPicture(minSizeCnt int, inputArray []pointType) {
	for i := 0; i < minSizeCnt; i++ {
		processImage(inputArray)
	}
	left, top, right, down := getBorder(inputArray)
	picture := make([][]bool, down-top+1, down-top+1)
	for y := 0; y < down-top+1; y++ {
		picture[y] = make([]bool, right-left+1, right-left+1)
	}
	for _, v := range inputArray {
		picture[v.posY-top][v.posX-left] = true
	}
	for i := range picture {
		for _, v := range picture[i] {
			if v {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	inputArray := loadFile("input.txt")
	minSize, minSizeCnt := getMinSize(inputArray)
	fmt.Println(minSize, minSizeCnt)
	inputArray = loadFile("input.txt")
	drawPicture(minSizeCnt, inputArray)
}
