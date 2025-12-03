package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func toInt(x string) int {
	val, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return val
}

func process(filename string, isInvalidID func(id int) bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read string")
	}
	res := 0
	for idRange := range strings.SplitSeq(strings.Replace(line, "\n", "", 1), ",") {
		parts := strings.Split(idRange, "-")
		start, end := toInt(parts[0]), toInt(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidID(i) {
				res += i
			}
		}
	}

	fmt.Println(res)
}

func part1(id int) bool {
	s := strconv.Itoa(id)
	size := len(s)
	if size%2 != 0 {
		return false
	}
	first, second := s[:size/2], s[size/2:]
	return first == second
}

func part2(id int) bool {
	s := strconv.Itoa(id)
	size := len(s)
	repeatLength := size / 2

	for repeatLength > 0 {
		if size%repeatLength != 0 {
			repeatLength--
			continue
		}
		repeatAmount := size / repeatLength

		repeats := strings.Repeat(s[:repeatLength], repeatAmount)
		if repeats == s {
			return true
		}
		repeatLength--
	}
	return false
}

func main() {
	process("input.txt", part1)
	process("input.txt", part2)
}
