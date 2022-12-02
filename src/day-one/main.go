package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../../data/day-one/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	current, one, two, three := 0, 0, 0, 0
	for scanner.Scan() {
		temp := 0
		value := scanner.Text()
		if value == "" {
			if current >= three {
				three = current
			}
			if three > two {
				temp = two
				two = three
				three = temp
			}
			if two > one {
				temp = one
				one = two
				two = temp
			}
			current = 0
			continue
		}

		intVal, intErr := strconv.Atoi(value)
		if intErr != nil {
			log.Fatal(err)
		}
		current += intVal
	}

	fmt.Printf("Total: %d\n", one+two+three)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}