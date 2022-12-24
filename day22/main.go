package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
	"strconv"
)

type Direction int

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

type Player struct {
	coord helper.Coord
	dir   Direction
}

type Sector struct {
	coord helper.Coord
	dir   map[Direction]*Sector
	turn  map[Direction]Direction
	wall  bool
}

func readMapPartTwo(lines []string) (map[helper.Coord]*Sector, *Sector) {
	var mx, my int

	r := make(map[helper.Coord]*Sector, 0)

	for i, row := range lines {
		if my < i {
			my = i
		}
		for j, col := range row {
			c := helper.Coord{X: j, Y: i}
			if mx < j {
				mx = j
			}

			switch col {
			case '.':
				s := Sector{
					coord: c,
					dir:   make(map[Direction]*Sector),
					turn:  make(map[Direction]Direction),
					wall:  false,
				}
				r[c] = &s
			case '#':
				s := Sector{
					coord: c,
					dir:   make(map[Direction]*Sector),
					turn:  make(map[Direction]Direction),
					wall:  true,
				}
				r[c] = &s
			}
		}
	}

	/*
	   a a 1 1 2 2   1w -> 4w    1w -> x:50,y:0:49, 1n -> x:50:99,y:0
	   a a 1 1 2 2   1n -> 6w    3w -> x:50,y:50:99, 3e -> x:99, y:50:99
	   b b 3 3 c c   2n -> 6s    4w -> x:0, y:100->149, 4n -> x:0:49, y:100
	   b b 3 3 c c   2e -> 5e    5e -> x:99, y:100->149   5s -> x:50:99, y:149
	   4 4 5 5 d d   3w -> 4n    6w -> x:0, y:150:199   6s -> x:0:49, y:199   6e -> x:49, y:150:199
	   4 4 5 5 d d   3e -> 2s    2n -> x:100:149, y:0   2e -> x:149, y:0:49   2s -> x:100:149, y:49
	   6 6 e e f f
	   6 6 e e f f
	                 6e -> 5s
	*/

	for y := 0; y <= my; y++ {
		for x := 0; x <= mx; x++ {
			c := helper.Coord{X: x, Y: y}
			if z, seen := r[c]; seen {

				// do all the connections we know about
				north := helper.Coord{X: c.X, Y: c.Y - 1}
				south := helper.Coord{X: c.X, Y: c.Y + 1}
				east := helper.Coord{X: c.X + 1, Y: c.Y}
				west := helper.Coord{X: c.X - 1, Y: c.Y}
				if v, seen := r[north]; seen {
					r[c].dir[NORTH] = v
					r[c].turn[NORTH] = NORTH
				}
				if v, seen := r[east]; seen {
					r[c].dir[EAST] = v
					r[c].turn[EAST] = EAST
				}
				if v, seen := r[west]; seen {
					r[c].dir[WEST] = v
					r[c].turn[WEST] = WEST
				}
				if v, seen := r[south]; seen {
					r[c].dir[SOUTH] = v
					r[c].turn[SOUTH] = SOUTH
				}

				// now the cube stuff..

				// connect 1 west to 4 west
				if c.X == 50 && c.Y >= 0 && c.Y < 50 {
					nc := helper.Coord{X: 0, Y: 149 - c.Y}
					r[c].dir[WEST] = r[nc]
					r[nc].dir[WEST] = z
					r[c].turn[WEST] = EAST
					r[nc].turn[WEST] = EAST
				}
				// connect 1 north to 6 west
				if c.X >= 50 && c.X < 100 && c.Y == 0 {
					nc := helper.Coord{X: 0, Y: c.X + 100}
					r[c].dir[NORTH] = r[nc]
					r[nc].dir[WEST] = z
					r[c].turn[NORTH] = EAST
					r[nc].turn[WEST] = SOUTH
				}
				// connect 2 north to 6 south
				if c.Y == 0 && c.X >= 100 && c.X <= 149 {
					nc := helper.Coord{X: c.X - 100, Y: 199}
					r[c].dir[NORTH] = r[nc]
					r[nc].dir[SOUTH] = z
					r[c].turn[NORTH] = NORTH
					r[nc].turn[SOUTH] = SOUTH
				}
				// connect 2 east to 5 east
				if c.X == 149 && c.Y >= 0 && c.Y <= 49 {
					nc := helper.Coord{X: 99, Y: 149 - c.Y}
					r[c].dir[EAST] = r[nc]
					r[nc].dir[EAST] = z
					r[c].turn[EAST] = WEST
					r[nc].turn[EAST] = WEST
				}
				// connect 3 west to 4 north
				if c.X == 50 && c.Y >= 50 && c.Y <= 99 {
					nc := helper.Coord{X: c.Y - 50, Y: 100}
					r[c].dir[WEST] = r[nc]
					r[nc].dir[NORTH] = z
					r[c].turn[WEST] = SOUTH
					r[nc].turn[NORTH] = EAST
				}
				// connect 3 east to 2 south
				if c.X == 99 && c.Y >= 50 && c.Y <= 99 {
					nc := helper.Coord{X: c.Y + 50, Y: 49}
					r[c].dir[EAST] = r[nc]
					r[nc].dir[SOUTH] = z
					r[c].turn[EAST] = NORTH
					r[nc].turn[SOUTH] = WEST
				}
				// connect 6 east to 5 south
				if c.X == 49 && c.Y >= 150 && c.Y <= 199 {
					nc := helper.Coord{X: c.Y - 100, Y: 149}
					r[c].dir[EAST] = r[nc]
					r[nc].dir[SOUTH] = z
					r[c].turn[EAST] = NORTH
					r[nc].turn[SOUTH] = WEST
				}

			}
		}
	}

	var start *Sector
	for j, col := range lines[0] {
		if col == '.' {
			c := helper.Coord{X: j, Y: 0}
			start = r[c]
			break
		}
	}

	return r, start
}

func readMap(lines []string) (map[helper.Coord]*Sector, *Sector) {
	var mx, my int

	r := make(map[helper.Coord]*Sector, 0)

	for i, row := range lines {
		if my < i {
			my = i
		}
		for j, col := range row {
			c := helper.Coord{X: j, Y: i}
			if mx < j {
				mx = j
			}

			switch col {
			case '.':
				s := Sector{
					coord: c,
					dir:   make(map[Direction]*Sector),
					wall:  false,
				}
				r[c] = &s
			case '#':
				s := Sector{
					coord: c,
					dir:   make(map[Direction]*Sector),
					wall:  true,
				}
				r[c] = &s
			}
		}
	}

	for y := 0; y <= my; y++ {
		for x := 0; x <= mx; x++ {
			c := helper.Coord{X: x, Y: y}
			if _, seen := r[c]; seen {
				north := helper.Coord{X: c.X, Y: c.Y - 1}
				south := helper.Coord{X: c.X, Y: c.Y + 1}
				east := helper.Coord{X: c.X + 1, Y: c.Y}
				west := helper.Coord{X: c.X - 1, Y: c.Y}
				if v, seen := r[north]; !seen {
					r[c].dir[NORTH] = getWrappedSector(c, r, SOUTH)
				} else {
					r[c].dir[NORTH] = v
				}
				if v, seen := r[east]; !seen {
					r[c].dir[EAST] = getWrappedSector(c, r, WEST)
				} else {
					r[c].dir[EAST] = v
				}
				if v, seen := r[west]; !seen {
					r[c].dir[WEST] = getWrappedSector(c, r, EAST)
				} else {
					r[c].dir[WEST] = v
				}
				if v, seen := r[south]; !seen {
					r[c].dir[SOUTH] = getWrappedSector(c, r, NORTH)
				} else {
					r[c].dir[SOUTH] = v
				}
			}
		}
	}

	var start *Sector
	for j, col := range lines[0] {
		if col == '.' {
			c := helper.Coord{X: j, Y: 0}
			start = r[c]
			break
		}
	}

	return r, start
}

func getWrappedSector(start helper.Coord, s map[helper.Coord]*Sector, dir Direction) *Sector {
	var r *Sector
	switch dir {
	case NORTH:
		for y := start.Y - 1; ; y-- {
			c := helper.Coord{X: start.X, Y: y}
			if _, seen := s[c]; seen {
				r = s[c]
			} else {
				break
			}
		}
	case SOUTH:
		for y := start.Y + 1; ; y++ {
			c := helper.Coord{X: start.X, Y: y}
			if _, seen := s[c]; seen {
				r = s[c]
			} else {
				break
			}
		}
	case EAST:
		for x := start.X + 1; ; x++ {
			c := helper.Coord{X: x, Y: start.Y}
			if _, seen := s[c]; seen {
				r = s[c]
			} else {
				break
			}
		}
	case WEST:
		for x := start.X - 1; ; x-- {
			c := helper.Coord{X: x, Y: start.Y}
			if _, seen := s[c]; seen {
				r = s[c]
			} else {
				break
			}
		}
	}

	return r
}

func (p *Player) move(steps int, grid map[helper.Coord]*Sector, partTwo bool) {

	for i := 0; i < steps; i++ {
		switch p.dir {
		case NORTH:
			ng := grid[p.coord].dir[NORTH]
			if ng.wall {
				return
			}
			if partTwo {
				p.dir = grid[p.coord].turn[NORTH]
			}
			p.coord = ng.coord
		case SOUTH:
			ng := grid[p.coord].dir[SOUTH]
			if ng.wall {
				return
			}
			if partTwo {
				p.dir = grid[p.coord].turn[SOUTH]
			}
			p.coord = ng.coord
		case EAST:
			ng := grid[p.coord].dir[EAST]
			if ng.wall {
				return
			}
			if partTwo {
				p.dir = grid[p.coord].turn[EAST]
			}
			p.coord = ng.coord
		case WEST:
			ng := grid[p.coord].dir[WEST]
			if ng.wall {
				return
			}
			if partTwo {
				p.dir = grid[p.coord].turn[WEST]
			}
			p.coord = ng.coord
		}
	}
}

func (p *Player) turn(turn string) {
	switch p.dir {
	case NORTH:
		switch turn {
		case "L":
			p.dir = WEST
		case "R":
			p.dir = EAST
		}
	case SOUTH:
		switch turn {
		case "L":
			p.dir = EAST
		case "R":
			p.dir = WEST
		}
	case EAST:
		switch turn {
		case "L":
			p.dir = NORTH
		case "R":
			p.dir = SOUTH
		}
	case WEST:
		switch turn {
		case "L":
			p.dir = SOUTH
		case "R":
			p.dir = NORTH
		}
	}
}

func problem(mapdata []string, dirdata string, partTwo bool) (int, error) {

	m, start := readMap(mapdata)
	if partTwo {
		m, start = readMapPartTwo(mapdata)
	}
	dirs := splitInput(dirdata)

	p := Player{
		dir:   EAST,
		coord: start.coord,
	}

	cc := m[helper.Coord{X: 50, Y: 0}]
	cc = cc
	for _, dir := range dirs {
		//sc := p.coord
		if dir == "L" || dir == "R" {
			p.turn(dir)
			continue
		}
		steps, err := strconv.Atoi(dir)
		if err != nil {
			return 0, err
		}
		p.move(steps, m, partTwo)
		//ec := p.coord
		//fmt.Printf("Player moved from (%v,%v) to (%v,%v)\n", sc.X, sc.Y, ec.X, ec.Y)
	}

	return (1000 * (p.coord.Y + 1)) + (4 * (p.coord.X + 1)) + int(p.dir), nil
}

func splitInput(s string) []string {
	r := make([]string, 0)
	var buf string
	for _, c := range s {
		if c >= '0' && c <= '9' {
			buf += string(c)
		} else if c == 'L' || c == 'R' {
			r = append(r, buf)
			r = append(r, string(c))
			buf = ""
		}
	}
	if buf != "" {
		r = append(r, buf)
	}
	return r
}

func main() {
	fh, _ := os.Open("map.txt")
	maplines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}
	fh.Close()

	fh, _ = os.Open("dir.txt")
	dirlines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}
	fh.Close()

	ans, err := problem(maplines, dirlines[0], false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(maplines, dirlines[0], true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
