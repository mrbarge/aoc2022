package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
	"strconv"
	"strings"
)

func problem(lines []string, ropelen int) (int, error) {
	seenCoords := make(map[string]bool, 0)
	seenCoords["0,0"] = true

	tails := make([]helper.Coord, 0)
	for i := 0; i < ropelen; i++ {
		tails = append(tails, helper.Coord{0, 0})
	}

	for _, line := range lines {
		dir := strings.Split(line, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(line, " ")[1])

		for i := 0; i < steps; i++ {
			// move the head
			tails[ropelen-1] = moveCoord(dir, tails[ropelen-1])

			for i := len(tails) - 2; i >= 0; i-- {
				if !touching(tails[i+1], tails[i]) {
					// move towards neighbour
					tails[i] = moveTail(tails[i+1], tails[i])
				}
			}
			tps := fmt.Sprintf("%v,%v", tails[0].X, tails[0].Y)
			seenCoords[tps] = true
		}
	}

	return len(seenCoords), nil
}

func moveCoord(dir string, c helper.Coord) helper.Coord {
	switch dir {
	case "R":
		return helper.Coord{c.X + 1, c.Y}
	case "L":
		return helper.Coord{c.X - 1, c.Y}
	case "U":
		return helper.Coord{c.X, c.Y - 1}
	case "D":
		return helper.Coord{c.X, c.Y + 1}
	}
	return c
}

func moveTail(head helper.Coord, tail helper.Coord) helper.Coord {

	nc := helper.Coord{tail.X, tail.Y}
	if head.X > tail.X {
		nc.X++
	}
	if head.X < tail.X {
		nc.X--
	}
	if head.Y > tail.Y {
		nc.Y++
	}
	if head.Y < tail.Y {
		nc.Y--
	}
	return nc
}

func touching(a helper.Coord, b helper.Coord) bool {
	xd := math.Abs(float64(a.X) - float64(b.X))
	yd := math.Abs(float64(a.Y) - float64(b.Y))
	return !(xd > 1 || yd > 1)
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
