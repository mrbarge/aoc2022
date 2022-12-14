package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	CAVE_SIZE = 1000
	ROCK      = 1
	AIR       = 0
	SAND      = 2
)

type Dimensions struct {
	minx int
	miny int
	maxx int
	maxy int
}

func newDimension() Dimensions {
	return Dimensions{
		minx: math.MaxInt,
		maxx: math.MinInt,
		miny: math.MaxInt,
		maxy: math.MinInt,
	}
}

func (d *Dimensions) Set(x int, y int) {
	if x < d.minx {
		d.minx = x
	} else if x > d.maxx {
		d.maxx = x
	}
	if y < d.miny {
		d.miny = y
	} else if y > d.maxy {
		d.maxy = y
	}
}

func readCave(lines []string, partTwo bool) ([][]int, Dimensions) {
	r := make([][]int, CAVE_SIZE)
	for i, _ := range r {
		r[i] = make([]int, CAVE_SIZE)
	}

	d := newDimension()

	for _, line := range lines {
		paths := strings.Split(line, " -> ")
		xp := -1
		yp := -1
		for _, path := range paths {
			x, _ := strconv.Atoi(strings.Split(path, ",")[0])
			y, _ := strconv.Atoi(strings.Split(path, ",")[1])

			r[y][x] = ROCK

			if xp >= 0 && yp >= 0 {
				if xp == x {
					// must be horizontal
					if y < yp {
						// going left
						for i := y; i <= yp; i++ {
							r[i][x] = ROCK
						}
					} else {
						for i := yp; i <= y; i++ {
							r[i][x] = ROCK
						}
					}
				} else {
					// must be vertical
					// must be horizontal
					if x < xp {
						// going left
						for i := x; i <= xp; i++ {
							r[y][i] = ROCK
						}
					} else {
						for i := xp; i <= x; i++ {
							r[y][i] = ROCK
						}
					}
				}
			}
			d.Set(x, y)
			xp = x
			yp = y
		}
	}

	for x := 0; x < CAVE_SIZE; x++ {
		r[d.maxy+2][x] = ROCK
	}
	return r, d
}

func printCave(cave [][]int, d Dimensions) {
	for y := d.miny; y <= d.maxy; y++ {
		for x := d.minx; x <= d.maxx; x++ {
			v := cave[y][x]
			if v == ROCK {
				fmt.Printf("#")
			} else if v == AIR {
				fmt.Printf(".")
			} else {
				fmt.Printf("o")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func sand(cave [][]int, d Dimensions, partTwo bool) ([][]int, bool) {
	sand := helper.Coord{X: 500, Y: 0}
	rest := false

	for !rest {

		if cave[sand.Y+1][sand.X] == AIR {
			sand.Y = sand.Y + 1
			continue
		} else if cave[sand.Y+1][sand.X-1] == AIR {
			sand.Y = sand.Y + 1
			sand.X = sand.X - 1
		} else if cave[sand.Y+1][sand.X+1] == AIR {
			sand.Y = sand.Y + 1
			sand.X = sand.X + 1
		} else {
			rest = true
		}

		if !partTwo && sand.Y > d.maxy {
			// overflow
			return cave, true
		}

		if rest && partTwo && sand.X == 500 && sand.Y == 0 {
			return cave, true
		}
	}

	cave[sand.Y][sand.X] = SAND
	return cave, false
}

func problem(lines []string, partTwo bool) (int, error) {

	cave, d := readCave(lines, partTwo)
	done := false
	steps := 0
	for !done {
		cave, done = sand(cave, d, partTwo)
		steps++
		//printCave(cave, d)
	}

	if !partTwo {
		steps--
	}
	return steps, nil
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
