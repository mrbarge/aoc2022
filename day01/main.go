package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/mrbarge/aoc2022/helper"
)

func problem(lines []string, partTwo bool) (int, error) {

	calorieCounts := make([]int, 0)
	calorieCount := 0

	for _, line := range lines {
		if line == "" {
			calorieCounts = append(calorieCounts, calorieCount)
			calorieCount = 0
			continue
		}
		calorie, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		calorieCount += calorie
	}

	sort.Ints(calorieCounts)

	elves := len(calorieCounts)
	if partTwo {
		if elves < 3 {
			return 0, fmt.Errorf("not enough elves for part two")
		}
		return calorieCounts[elves-1] + calorieCounts[elves-2] + calorieCounts[elves-3], nil
	}
	return calorieCounts[elves-1], nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, true)
	fmt.Printf("Part one: %v\n", ans)
}
