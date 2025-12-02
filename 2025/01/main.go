package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	number := 50
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		var temp int
		if direction == "L" {
			temp = number - amount
		} else {
			temp = number + amount
		}

		if amount >= 100 {
			count += amount / 100
		} else if temp <= 0 || temp > 99 {
			count++
		}

		number = mod(temp, 100)
		fmt.Println(temp, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(count)
}
