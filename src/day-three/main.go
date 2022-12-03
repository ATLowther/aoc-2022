package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var m = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func main() {
	file, err := os.Open("../../data/day-three/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	total := 0
	lineNum := 1
	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()

		// When line is blank, we've reach EOF
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)

		// Process 3 lines at a time
		if lineNum%3 == 0 {
			sackOne, sackTwo, sackThree := lines[0], lines[1], lines[2]
			dupeValue := findBadge(sackOne, sackTwo, sackThree)
			total += dupeValue

			lines = nil
		}
		lineNum++
	}

	fmt.Printf("Total: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findBadge(sackOne string, sackTwo string, sackThree string) int {
	oneSplice, twoSplice, threeSplice := []byte(sackOne), []byte(sackTwo), []byte(sackThree)

	for _, charInOne := range oneSplice {
		for _, charInTwo := range twoSplice {
			for _, charInThree := range threeSplice {
				if (charInOne == charInTwo) && (charInOne == charInThree) {
					return m[string(charInOne)]
				}
			}
		}
	}
	return 0
}
