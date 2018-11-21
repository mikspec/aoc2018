package main

import (
	"flag"
	"fmt"
)

var conParam string

func init() {
	const (
		usage = "String to be parsed"
	)
	flag.StringVar(&conParam, "c", "", usage)
}

// Sum adds integers
func Sum(arr []int) int {

	sum := 0
	for _, a := range arr {
		sum += a
	}
	return sum
}

func main() {

	flag.Parse()

	argsWithProg := flag.Args()

	Sum([]int{1, 2, 3})

	fmt.Printf("Day 1: %s\n%s\n", argsWithProg, conParam)
}
