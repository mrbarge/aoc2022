package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
)

func checkVisible(lines [][]int, x int, y int) (bool, int) {
	scores := make([]int, 4)
	visibles := []bool{true, true, true, true}
	for i := y - 1; i >= 0; i-- {
		scores[0] += 1
		if lines[i][x] >= lines[y][x] {
			visibles[0] = false
			break
		}
	}
	for i := y + 1; i < len(lines); i++ {
		scores[1] += 1
		if lines[i][x] >= lines[y][x] {
			visibles[1] = false
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		scores[2] += 1
		if lines[y][i] >= lines[y][x] {
			visibles[2] = false
			break
		}
	}
	for i := x + 1; i < len(lines[y]); i++ {
		scores[3] += 1
		if lines[y][i] >= lines[y][x] {
			visibles[3] = false
			break
		}
	}

	sumScores := 1
	for _, v := range scores {
		sumScores *= v
	}
	isVisible := false
	for _, v := range visibles {
		isVisible = v || isVisible
	}
	return isVisible, sumScores
}

func problem(lines [][]int, partTwo bool) (int, int, error) {

	shorter := 0
	maxScore := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			visible, score := checkVisible(lines, x, y)
			if visible {
				shorter++
			}
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return shorter, maxScore, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLinesAsIntArray(fh)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans1, ans2, err := problem(lines, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans1)
	fmt.Printf("Part two: %v\n", ans2)

}
