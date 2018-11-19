package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	// Prompt and read
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Day 1: %s", text)
}
