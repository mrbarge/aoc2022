package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
)

type Direction int

const (
	NORTH = iota
	SOUTH
	WEST
	EAST
)

type Elf struct {
	id    int
	coord helper.Coord
}

func readElves(lines []string) (map[helper.Coord]*Elf, []*Elf) {
	id := 0
	r := make(map[helper.Coord]*Elf)
	re := make([]*Elf, 0)
	for y, row := range lines {
		for x, v := range row {
			if v == '#' {
				c := helper.Coord{X: x, Y: y}
				elf := Elf{
					id:    id,
					coord: c,
				}
				r[c] = &elf
				re = append(re, &elf)
				id++
			}
		}
	}
	return r, re
}

func round(m map[helper.Coord]*Elf, elves []*Elf, dir Direction) (map[helper.Coord]*Elf, bool) {

	proposed := makeProposed(m, elves, dir)
	next, moved := makeMoved(proposed, m)
	return next, moved
}

func makeMoved(proposed map[int]helper.Coord, m map[helper.Coord]*Elf) (map[helper.Coord]*Elf, bool) {
	r := make(map[helper.Coord]*Elf, 0)
	moved := false
	for _, elf := range m {
		pc := proposed[elf.id]
		cannotMove := false
		for id, e2 := range proposed {
			if id != elf.id && e2 == pc {
				// cannot move, stay put
				r[elf.coord] = elf
				cannotMove = true
				break
			}
		}
		if !cannotMove {
			if elf.coord != pc {
				moved = true
			}
			r[pc] = elf
			elf.coord = pc
		}
	}
	return r, moved
}

func checkMove(d Direction, e *Elf, m map[helper.Coord]*Elf) (helper.Coord, bool) {
	switch d {
	case NORTH:
		coords := []helper.Coord{{X: e.coord.X - 1, Y: e.coord.Y - 1}, {X: e.coord.X, Y: e.coord.Y - 1}, {X: e.coord.X + 1, Y: e.coord.Y - 1}}
		for _, c := range coords {
			if _, seen := m[c]; seen {
				return e.coord, false
			}
		}
		return helper.Coord{X: e.coord.X, Y: e.coord.Y - 1}, true
	case SOUTH:
		coords := []helper.Coord{{X: e.coord.X - 1, Y: e.coord.Y + 1}, {X: e.coord.X, Y: e.coord.Y + 1}, {X: e.coord.X + 1, Y: e.coord.Y + 1}}
		for _, c := range coords {
			if _, seen := m[c]; seen {
				return e.coord, false
			}
		}
		return helper.Coord{X: e.coord.X, Y: e.coord.Y + 1}, true
	case EAST:
		coords := []helper.Coord{{X: e.coord.X + 1, Y: e.coord.Y - 1}, {X: e.coord.X + 1, Y: e.coord.Y}, {X: e.coord.X + 1, Y: e.coord.Y + 1}}
		for _, c := range coords {
			if _, seen := m[c]; seen {
				return e.coord, false
			}
		}
		return helper.Coord{X: e.coord.X + 1, Y: e.coord.Y}, true
	case WEST:
		coords := []helper.Coord{{X: e.coord.X - 1, Y: e.coord.Y - 1}, {X: e.coord.X - 1, Y: e.coord.Y}, {X: e.coord.X - 1, Y: e.coord.Y + 1}}
		for _, c := range coords {
			if _, seen := m[c]; seen {
				return e.coord, false
			}
		}
		return helper.Coord{X: e.coord.X - 1, Y: e.coord.Y}, true
	}
	return e.coord, false
}

func makeProposed(m map[helper.Coord]*Elf, elves []*Elf, dir Direction) map[int]helper.Coord {
	r := make(map[int]helper.Coord, 0)

	for c, elf := range m {

		// first check that every space around is empty
		neighbourHasElf := false
		neighbours := c.GetNeighbours(true)
		for _, n := range neighbours {
			if _, seen := m[n]; seen {
				neighbourHasElf = true
				break
			}
		}
		if !neighbourHasElf {
			// should stay put
			r[elf.id] = c
			continue
		}

		// then check where to move
		foundProposal := false
		for i := 0; i < 4; i++ {
			d := (int(dir) + i) % 4
			if nc, ok := checkMove(Direction(d), elf, m); ok {
				foundProposal = true
				r[elf.id] = nc
				break
			}
		}
		if !foundProposal {
			r[elf.id] = c
		}
	}

	return r
}

func problem(lines []string, partTwo bool) (int, error) {
	ans := 0

	dir := Direction(NORTH)

	if !partTwo {
		rounds := 10
		ec, el := readElves(lines)
		for x := 0; x < rounds; x++ {
			ec, _ = round(ec, el, dir)
			dir = (dir + 1) % 4
		}
		mnx, mny, mxx, mxy := ranges(ec)
		ans = findEmpty(ec, mnx, mny, mxx, mxy)
	} else {
		ans = 0
		moving := true
		ec, el := readElves(lines)
		for moving {
			ec, moving = round(ec, el, dir)
			dir = (dir + 1) % 4
			ans++
		}
	}

	return ans, nil
}

func findEmpty(m map[helper.Coord]*Elf, mnx, mny, mxx, mxy int) int {
	r := 0
	for x := mnx; x <= mxx; x++ {
		for y := mny; y <= mxy; y++ {
			c := helper.Coord{X: x, Y: y}
			if _, seen := m[c]; !seen {
				r++
			}
		}
	}
	return r
}
func ranges(m map[helper.Coord]*Elf) (minx, miny, maxx, maxy int) {
	minx = math.MaxInt
	miny = math.MaxInt
	maxx = math.MinInt
	maxy = math.MinInt

	for _, elf := range m {
		if elf.coord.X < minx {
			minx = elf.coord.X
		}
		if elf.coord.Y < miny {
			miny = elf.coord.Y
		}
		if elf.coord.X > maxx {
			maxx = elf.coord.X
		}
		if elf.coord.Y > maxy {
			maxy = elf.coord.Y
		}
	}

	return minx, miny, maxx, maxy
}

func print(m map[helper.Coord]*Elf, mnx, mny, mxx, mxy int) {
	for y := mny; y <= mxy; y++ {
		for x := mnx; x <= mxx; x++ {
			c := helper.Coord{X: x, Y: y}
			if _, seen := m[c]; !seen {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
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
