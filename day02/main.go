package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
)

const (
	Rock     int = 1
	Paper        = 2
	Scissors     = 3
	Invalid      = 4
	Loss         = 0
	Draw         = 3
	Win          = 6
)

func p2Play(r rune) int {
	switch r {
	case 'A':
		return Rock
	case 'X':
		return Loss
	case 'B':
		return Paper
	case 'Y':
		return Draw
	case 'C':
		return Scissors
	case 'Z':
		return Win
	}
	return Invalid
}

func p1Play(r rune) int {
	switch r {
	case 'A':
		return Rock
	case 'X':
		return Rock
	case 'B':
		return Paper
	case 'Y':
		return Paper
	case 'C':
		return Scissors
	case 'Z':
		return Scissors
	}
	return Invalid
}

func p1Match(p rune, o rune) int {
	op := p1Play(o)
	switch p1Play(p) {
	case Rock:
		switch op {
		case Rock:
			return Rock + Draw
		case Scissors:
			return Rock + Win
		case Paper:
			return Rock + Loss
		}
	case Paper:
		switch op {
		case Rock:
			return Paper + Win
		case Scissors:
			return Paper + Loss
		case Paper:
			return Paper + Draw
		}
	case Scissors:
		switch op {
		case Rock:
			return Scissors + Loss
		case Scissors:
			return Scissors + Draw
		case Paper:
			return Scissors + Win
		}
	}
	return 0
}

func p2Match(p rune, o rune) int {
	op := p2Play(o)
	switch p2Play(p) {
	case Win:
		switch op {
		case Rock:
			return Win + Paper
		case Scissors:
			return Win + Rock
		case Paper:
			return Win + Scissors
		}
	case Loss:
		switch op {
		case Rock:
			return Loss + Scissors
		case Scissors:
			return Loss + Paper
		case Paper:
			return Loss + Rock
		}
	case Draw:
		return Draw + op
	}
	return 0
}

func problem(lines []string, partTwo bool) (int, error) {

	score := 0
	for _, line := range lines {
		if len(line) < 3 {
			return 0, fmt.Errorf("invalid line: %s", line)
		}
		if partTwo {
			score += p2Match(rune(line[2]), rune(line[0]))
		} else {
			score += p1Match(rune(line[2]), rune(line[0]))
		}
	}
	return score, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %v\n", ans)
}
