package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../../data/day-four/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		if doesOverlap(ranges[0], ranges[1]) {
			total++
		}
	}

	log.Printf("Total: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isFullyContained(range1 string, range2 string) bool {
	rangeOneLowerBound, rangeOneUpperBound := getRangeBounds(range1)
	rangeTwoLowerBound, rangeTwoUpperBound := getRangeBounds(range2)

	if rangeOneLowerBound <= rangeTwoLowerBound && rangeOneUpperBound >= rangeTwoUpperBound {
		return true
	}
	if rangeTwoLowerBound <= rangeOneLowerBound && rangeTwoUpperBound >= rangeOneUpperBound {
		return true
	}

	return false
}

func doesOverlap(range1 string, range2 string) bool {
	rangeOneLowerBound, rangeOneUpperBound := getRangeBounds(range1)
	rangeTwoLowerBound, rangeTwoUpperBound := getRangeBounds(range2)

	if rangeOneUpperBound >= rangeTwoLowerBound && rangeOneLowerBound <= rangeTwoUpperBound {
		return true
	}
	if rangeTwoUpperBound >= rangeOneLowerBound && rangeTwoLowerBound <= rangeOneUpperBound {
		return true
	}

	return false
}

func getRangeBounds(rangeVal string) (int, int) {
	bounds := strings.Split(rangeVal, "-")
	lowerBound, err := strconv.Atoi(bounds[0])
	if err != nil {
		log.Fatal(err)
	}

	upperBound, err := strconv.Atoi(bounds[1])
	if err != nil {
		log.Fatal(err)
	}

	return lowerBound, upperBound
}
