package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var stacks = map[int][]string{
	1: {"H", "C", "R"},
	2: {"B", "J", "H", "L", "S", "F"},
	3: {"R", "M", "D", "H", "J", "T", "Q"},
	4: {"S", "G", "R", "H", "Z", "B", "J"},
	5: {"R", "P", "F", "Z", "T", "D", "C", "B"},
	6: {"T", "H", "C", "G"},
	7: {"S", "N", "V", "Z", "B", "P", "W", "L"},
	8: {"R", "J", "Q", "G", "C"},
	9: {"L", "D", "T", "R", "H", "P", "F", "S"},
}

func main() {
	file, err := os.Open("../../data/day-five/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip input lines until we reach move instructions
		if !strings.Contains(line, "move") {
			continue
		}

		numOfCratesToMove, fromStack, toStack := getInstructions(line)
		// log.Println("Move", cratesToMove, "from", fromStack, "to", toStack)

		moveBatchOfCrates(stacks, numOfCratesToMove, fromStack, toStack)
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
