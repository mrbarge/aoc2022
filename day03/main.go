package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
	"unicode"
)

func partOne(lines []string) (int, error) {

	sum := 0
	for i, line := range lines {
		l := len(line)
		if l%2 != 0 {
			return 0, fmt.Errorf("unbalanced line %d: %s", i, line)
		}
		seen := make(map[rune]bool)
		for j, r := range line {
			if j < l/2 {
				seen[r] = true
				continue
			}
			if _, ok := seen[r]; ok {
				sum += priority(r)
				break
			}
		}
	}
	return sum, nil
}

func processGroup(group []string) int {
	seen := make(map[rune]int, 0)
	for _, gline := range group {
		seenline := make(map[rune]bool, 0)
		for _, r := range gline {
			if _, ok := seenline[r]; !ok {
				seen[r] += 1
				seenline[r] = true
			}
		}
	}
	for r, v := range seen {
		if v == 3 {
			return priority(r)
		}
	}
	return 0
}

func partTwo(lines []string) (int, error) {
	sum := 0
	group := make([]string, 0)
	for i, line := range lines {
		if i > 0 && i%3 == 0 {
			// new group, process the existing one
			sum += processGroup(group)
			group = make([]string, 0)
		}
		group = append(group, line)
	}
	sum += processGroup(group)
	return sum, nil
}

func priority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 'A' + 27
	} else {
		return int(r) - 'a' + 1
	}
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = partTwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}
