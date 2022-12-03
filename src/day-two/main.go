package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type RockPaperScissor string
type LoseDrawWin string

const (
	Rock    RockPaperScissor = "A"
	Paper   RockPaperScissor = "B"
	Scissor RockPaperScissor = "C"
)

const (
	Lose LoseDrawWin = "X"
	Draw LoseDrawWin = "Y"
	Win  LoseDrawWin = "Z"
)

func main() {
	file, err := os.Open("../../data/day-two/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		choices := strings.Fields((line))
		opponent, fate := RockPaperScissor(choices[0]), LoseDrawWin(choices[1])
		round := getRoundTotal(opponent, fate)
		total += round
	}

	fmt.Printf("Total: %d\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getRoundTotal(opponentChoice RockPaperScissor, fate LoseDrawWin) int {
	toPlay := battle(opponentChoice, fate)

	participationPoints := pointsForPlaying(toPlay)
	if participationPoints == -1 {
		log.Fatal(toPlay)
	}

	turnoutScore := pointsForTurnout(fate)
	if turnoutScore == -1 {
		log.Fatal(fate)
	}

	return participationPoints + turnoutScore
}

func pointsForPlaying(choice RockPaperScissor) int {
	switch choice {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissor:
		return 3
	default:
		return -1
	}
}

func pointsForTurnout(fate LoseDrawWin) int {
	switch fate {
	case Win:
		return 6
	case Lose:
		return 0
	case Draw:
		return 3
	default:
		return -1
	}
}

func battle(opponent RockPaperScissor, fate LoseDrawWin) RockPaperScissor {
	switch opponent {
	case Rock:
		if fate == Win {
			return Paper
		}
		if fate == Lose {
			return Scissor
		}
		if fate == Draw {
			return Rock
		}
	case Paper:
		if fate == Win {
			return Scissor
		}
		if fate == Lose {
			return Rock
		}
		if fate == Draw {
			return Paper
		}
	case Scissor:
		if fate == Win {
			return Rock
		}
		if fate == Lose {
			return Paper
		}
		if fate == Draw {
			return Scissor
		}
	}
	// We are using known input so this won't ever be reached.
	// This is error prone for unvalidated input
	return Rock
}
