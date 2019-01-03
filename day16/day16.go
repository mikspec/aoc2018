package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

type stateType [4]int

type commandType struct {
	command int
	valA    int
	valB    int
	valC    int
}

type stackType struct {
	before  stateType
	command commandType
	after   stateType
}

type commandFuncType func(stateType, commandType) stateType

type commandMapType map[string]commandFuncType

type mnemonicMapType map[int]string

var commandMap commandMapType
var mnemonicMap mnemonicMapType

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
	mnemonicMap = make(mnemonicMapType)
}

// File loading generates road plan
func loadFile(name string) (stack []stackType, commands []commandType) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reState := regexp.MustCompile("(Before|After):[ ]{1,2}\\[(\\d+), (\\d+), (\\d+), (\\d+)\\]")
	reCommand := regexp.MustCompile("(\\d+) (\\d+) (\\d+) (\\d+)")
	scanner := bufio.NewScanner(file)
	statFlg := false
	var log stackType
	stack = make([]stackType, 0)
	commands = make([]commandType, 0)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			state := reState.FindStringSubmatch(line)
			if len(state) > 0 {
				if !statFlg {
					log = stackType{}
					log.before[0], _ = strconv.Atoi(state[2])
					log.before[1], _ = strconv.Atoi(state[3])
					log.before[2], _ = strconv.Atoi(state[4])
					log.before[3], _ = strconv.Atoi(state[5])
					statFlg = true
				} else {
					log.after[0], _ = strconv.Atoi(state[2])
					log.after[1], _ = strconv.Atoi(state[3])
					log.after[2], _ = strconv.Atoi(state[4])
					log.after[3], _ = strconv.Atoi(state[5])
					statFlg = false
					stack = append(stack, log)
				}
			} else {
				commandStr := reCommand.FindStringSubmatch(line)
				command := commandType{}
				command.command, _ = strconv.Atoi(commandStr[1])
				command.valA, _ = strconv.Atoi(commandStr[2])
				command.valB, _ = strconv.Atoi(commandStr[3])
				command.valC, _ = strconv.Atoi(commandStr[4])
				if statFlg {
					log.command = command
				} else {
					commands = append(commands, command)
				}
			}
		}
	}
	return
}

// Process input stack
func processStack(commandMap *commandMapType, stack []stackType, mnemonicMap *mnemonicMapType, createMnemMap bool) (counter int) {
	for _, log := range stack {
		n := 0
		commandID := -1
		commandName := ""
		if _, ok := (*mnemonicMap)[log.command.command]; ok {
			continue
		}
	commandLoop:
		for name, command := range *commandMap {
			for _, v := range *mnemonicMap {
				if v == name {
					continue commandLoop
				}
			}
			if reflect.DeepEqual(log.after, command(log.before, log.command)) {
				n++
				commandID = log.command.command
				commandName = name
			}
		}
		if n >= 3 {
			counter++
		}
		if n == 1 && createMnemMap {
			(*mnemonicMap)[commandID] = commandName
		}
		if n == 0 {
			fmt.Println("Error !!!!!!!!")
			return
		}
	}
	return
}

func processCommands(commandMap *commandMapType, mnemonicMap *mnemonicMapType, commands []commandType) (state stateType) {
	for _, command := range commands {
		state = (*commandMap)[(*mnemonicMap)[command.command]](state, command)
	}
	return
}

func main() {
	stack, commands := loadFile("input.txt")
	counter := processStack(&commandMap, stack, &mnemonicMap, false)
	fmt.Println("Number of samples (Three or more opcodes):", counter)
	counter = processStack(&commandMap, stack, &mnemonicMap, true)
	fmt.Println(mnemonicMap)
	state := processCommands(&commandMap, &mnemonicMap, commands)
	fmt.Println("End state:", state)
}
