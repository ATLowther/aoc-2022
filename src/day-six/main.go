package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const UNIQUE_LENGTH int = 14

// Initalize our slice to the length of unique sequential
// values that we want to track
var chars = make([]string, UNIQUE_LENGTH)

func main() {
	file, err := os.Open("../../data/day-six/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	byteCount := 1
	for scanner.Scan() {
		byte := scanner.Text()

		chars = updateAggregate(chars, byte)
		done := areUnique(chars)
		if done {
			break
		}
		byteCount++
	}

	log.Printf("Indexes to unique: %d\n", byteCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func updateAggregate(chars []string, char string) []string {
	for i, value := range chars {
		if i > 0 {
			chars[i-1] = value
		}
	}
	chars[len(chars)-1] = char

	return chars
}

func areUnique(chars []string) bool {
	currentState := ""
	for _, value := range chars {
		if strings.Contains(currentState, value) {
			return false
		}
		currentState += value
	}

	return true
}
