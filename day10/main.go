package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
	"strings"
)

func problem(lines []string, partTwo bool) (int, error) {
	x := 1
	queue := make([]int, 0)
	sum := 0
	sumcycles := []int{20, 60, 100, 140, 180, 220}
	for _, line := range lines {
		var op string
		var val int

		queue = append(queue, 0)
		if line != "noop" {
			r := strings.NewReader(line)
			_, err := fmt.Fscanf(r, "%s %d", &op, &val)
			if err != nil {
				return 0, err
			}
			queue = append(queue, val)
		}
	}

	screen := make([][]bool, 0)
	for i := 0; i < 6; i++ {
		screen = append(screen, make([]bool, 40))
	}

	sr := 0
	sc := 0
	for cycle, qv := range queue {
		for _, v := range sumcycles {
			if cycle+1 == v {
				strength := (v * x)
				sum += strength
			}
		}

		if sc >= (x-1) && sc <= (x+1) {
			screen[sr][sc] = true
		}
		x += qv

		// move pixel painter
		sc++
		if sc%40 == 0 {
			sr++
			sc = 0
		}
	}
	printScreen(screen)
	return sum, nil
}

func printScreen(p [][]bool) {
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			if p[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
