package main

import (
	"fmt"
	"math"
	"os"

	"github.com/albertorestifo/dijkstra"
	"github.com/mrbarge/aoc2022/helper"
)

func problem(lines []string, partTwo bool) (int, error) {

	g := dijkstra.Graph{}

	var startPos, endPos string
	p2StartPos := make([]string, 0)

	for y, row := range lines {
		for x, _ := range row {
			c := helper.Coord{X: x, Y: y}
			nodeVal := fmt.Sprintf("%v,%v", x, y)
			g[nodeVal] = make(map[string]int, 0)

			switch lines[y][x] {
			case 'S':
				startPos = nodeVal
			case 'E':
				endPos = nodeVal
			case 'a':
				p2StartPos = append(p2StartPos, nodeVal)
			}

			neighbours := c.GetNeighboursPos(false)
			for _, n := range neighbours {
				if n.X >= len(lines[0]) || n.Y >= len(lines) {
					continue
				}
				neighVal := fmt.Sprintf("%v,%v", n.X, n.Y)

				weight := 1
				if lines[y][x] == 'S' {
					if lines[n.Y][n.X] != 'a' && lines[n.Y][n.X] != 'b' {
						continue
					}
				} else if lines[n.Y][n.X] == 'E' {
					if lines[y][x] != 'z' && lines[y][x] != 'y' {
						continue
					}
				} else if (int(lines[n.Y][n.X]) - int(lines[y][x])) > 1 {
					continue
				}
				g[nodeVal][neighVal] = weight
			}
		}
	}

	var score int
	if partTwo {
		score = math.MaxInt
		for _, spos := range p2StartPos {
			_, s, err := g.Path(spos, endPos)
			if err != nil {
				continue
			}
			if s > 0 && s < score {
				score = s
			}
		}
	} else {
		_, s, err := g.Path(startPos, endPos)
		if err != nil {
			return 0, err
		}
		score = s
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
