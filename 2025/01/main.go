package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type State struct {
	number int
	count  int
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func part1(direction string, amount int, state *State) {
	if direction == "L" {
		state.number = mod(state.number-amount, 100)
	} else {
		state.number = mod(state.number+amount, 100)
	}

	if state.number == 0 {
		state.count++
	}
}

func part2(direction string, amount int, state *State) {
	var temp int
	if direction == "L" {
		temp = state.number - amount
	} else {
		temp = state.number + amount
	}

	if amount >= 100 {
		state.count += amount / 100
	} else if temp <= 0 || temp > 99 {
		state.count++
	}

	state.number = mod(temp, 100)
}

func process(filename string, updateState func(string, int, *State)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := &State{
		number: 50,
		count:  0,
	}

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Error parsing number: %v\n", err)
		}
		updateState(direction, amount, state)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(state.count)
}

func main() {
	process("input.txt", part1)
	process("input.txt", part2)
}
