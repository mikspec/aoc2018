package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type stateType [6]int

type commandType struct {
	command string
	valA    int
	valB    int
	valC    int
}

type commandFuncType func(stateType, commandType) stateType

type commandMapType map[string]commandFuncType

var commandMap commandMapType

func init() {
	commandMap = commandMapType{
		"addr": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] + state[command.valB]
			return state
		},
		"addi": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] + command.valB
			return state
		},
		"mulr": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] * state[command.valB]
			return state
		},
		"muli": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] * command.valB
			return state
		},
		"banr": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] & state[command.valB]
			return state
		},
		"bani": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] & command.valB
			return state
		},
		"borr": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] | state[command.valB]
			return state
		},
		"bori": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA] | command.valB
			return state
		},
		"setr": func(state stateType, command commandType) stateType {
			state[command.valC] = state[command.valA]
			return state
		},
		"seti": func(state stateType, command commandType) stateType {
			state[command.valC] = command.valA
			return state
		},
		"gtir": func(state stateType, command commandType) stateType {
			if command.valA > state[command.valB] {
				state[command.valC] = 1
			} else {
				state[command.valC] = 0
			}
			return state
		},
		"gtri": func(state stateType, command commandType) stateType {
			if state[command.valA] > command.valB {
				state[command.valC] = 1
			} else {
				state[command.valC] = 0
			}
			return state
		},
		"gtrr": func(state stateType, command commandType) stateType {
			if state[command.valA] > state[command.valB] {
				state[command.valC] = 1
			} else {
				state[command.valC] = 0
			}
			return state
		},
		"eqir": func(state stateType, command commandType) stateType {
			if command.valA == state[command.valB] {
				state[command.valC] = 1
			} else {
				state[command.valC] = 0
			}
			return state
		},
		"eqri": func(state stateType, command commandType) stateType {
			if state[command.valA] == command.valB {
				state[command.valC] = 1
			} else {
				state[command.valC] = 0
			}
			return state
		},
		"eqrr": func(state stateType, command commandType) stateType {
			if state[command.valA] == state[command.valB] {
				state[command.valC] = 1
			} else {
				state[command.valC] = 0
			}
			return state
		},
	}
}

// File loading generates road plan
func loadFile(name string) (ip int, commands []commandType) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reRegister := regexp.MustCompile("#ip (\\d+)")
	reCommand := regexp.MustCompile("([a-z]+) (\\d+) (\\d+) (\\d+)")
	scanner := bufio.NewScanner(file)
	commands = make([]commandType, 0)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			state := reRegister.FindStringSubmatch(line)
			if len(state) > 0 {
				ip, _ = strconv.Atoi(state[1])
			} else {
				commandStr := reCommand.FindStringSubmatch(line)
				command := commandType{}
				command.command = commandStr[1]
				command.valA, _ = strconv.Atoi(commandStr[2])
				command.valB, _ = strconv.Atoi(commandStr[3])
				command.valC, _ = strconv.Atoi(commandStr[4])
				commands = append(commands, command)
			}
		}
	}
	return
}

// Process input commands
func processCommands(commandMap *commandMapType, ip int, commands []commandType, initState stateType) (state stateType) {
	ipVal := 0
	state = initState
	cnt := 0
	for {
		state[ip] = ipVal
		command := commands[ipVal]
		state = (*commandMap)[command.command](state, command)
		ipVal = state[ip]
		cnt++
		if cnt > 6427000 {
			return
		}
		ipVal++
		if ipVal >= len(commands) {
			return
		}
	}
}

func main() {
	ip, commands := loadFile("input.txt")
	fmt.Println(ip, commands)
	state := processCommands(&commandMap, ip, commands, stateType{})
	fmt.Println("End state:", state)
	state = processCommands(&commandMap, ip, commands, stateType{1, 0, 0, 0, 0, 0})
	fmt.Println("End state after init:", state)
	sum := 0
	for i := 1; i <= state[5]; i++ {
		if state[5]%i == 0 {
			sum += i
		}
	}
	fmt.Println("Sum register 0", sum)
}
