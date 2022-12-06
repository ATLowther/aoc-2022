package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../../data/day-five/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	instructionsReached := false
	stacks := make(map[int][]string)
	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)

		// Skip input lines until we reach move instructions
		if strings.Contains(line, "move") {
			instructionsReached = true
		}

		if !instructionsReached && len(line) > 0 {
			for i := 0; i < 9; i++ {
				// If a crate exists on a stack, it will be at an index that is a multiple of 4,
				// offset by 1, for a possible `[` character.
				mapIndex := (i * 4) + 1
				buildMap(bytes[mapIndex], i+1, stacks)
			}
		}

		if instructionsReached {
			numOfCratesToMove, fromStack, toStack := getInstructions(line)
			moveBatchOfCrates(stacks, numOfCratesToMove, fromStack, toStack)
		}

	}

	log.Printf("Top Crates: %s\n", getTopCratesFromStacks(stacks))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// We want to split the string on any field that is not a number
func getNumbers(possibleNum rune) bool {
	if !unicode.IsNumber(possibleNum) {
		return true
	}
	return false
}

func moveCrates(stacks map[int][]string, cratesToMove int, fromStack int, toStack int) {
	for i := 0; i < cratesToMove; i++ {
		indexOfCrateToMove := len(stacks[fromStack]) - 1
		crateToMove := stacks[fromStack][indexOfCrateToMove]
		stacks[fromStack] = stacks[fromStack][:indexOfCrateToMove]
		stacks[toStack] = append(stacks[toStack], crateToMove)
	}
	return
}

func moveBatchOfCrates(stacks map[int][]string, numOfCratesToMove int, fromStack int, toStack int) {
	lowerBoundOfCratesToMove := len(stacks[fromStack]) - numOfCratesToMove

	// We want our starting index to be the lower bound and grab to the end of the slice
	cratesToMove := stacks[fromStack][lowerBoundOfCratesToMove:]
	stacks[toStack] = append(stacks[toStack], cratesToMove...)

	// We want to keep all crates until the lower bound of the crates we moved.
	stacks[fromStack] = stacks[fromStack][:lowerBoundOfCratesToMove]

	return
}

func getInstructions(instructions string) (int, int, int) {
	// Possibly unsafe implementation, docs say that the order that fields are passed
	// are not guaranteed; https://pkg.go.dev/strings#FieldsFunc
	// We are splitting the fields on word characters, which leaves us just with the numbers
	fields := strings.FieldsFunc(instructions, isWord)

	cratesToMove, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal(err)
	}

	fromStack, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatal(err)
	}

	toStack, err := strconv.Atoi(fields[2])
	if err != nil {
		log.Fatal(err)
	}

	return cratesToMove, fromStack, toStack
}

func getTopCratesFromStacks(stacks map[int][]string) string {
	topCratesString := ""
	// Start from a 1 index as the first key in the map is 1
	for i := 1; i <= len(stacks); i++ {
		lastCrateIndex := len(stacks[i]) - 1
		topCratesString += stacks[i][lastCrateIndex]
	}
	return topCratesString
}

func isWord(field rune) bool {
	if !unicode.IsNumber(field) {
		return true
	}
	return false
}

func buildMap(char byte, index int, stack map[int][]string) {
	if unicode.IsLetter(rune(char)) {
		// Prepend character, as we are building our stack from top to bottom
		stack[index] = append([]string{string(char)}, stack[index]...)
	}
}
