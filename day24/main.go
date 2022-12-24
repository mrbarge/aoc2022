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
	EAST
	SOUTH
	WEST
)

type Player struct {
	c     helper.Coord
	turns int
}

type State struct {
	p    Player
	tick int
}

type Range struct {
	minx int
	miny int
	maxx int
	maxy int
}

type Blizzard struct {
	dir   Direction
	coord helper.Coord
}

var cycles map[int]map[helper.Coord][]Blizzard

func readMap(lines []string) (r map[helper.Coord][]Blizzard, start helper.Coord, end helper.Coord) {
	var mnx, mny, mxx, mxy int
	r = make(map[helper.Coord][]Blizzard)
	for y, row := range lines {
		for x, v := range row {
			c := helper.Coord{X: x, Y: y}
			if x > mnx {
				mnx = x
			}
			if mny > y {
				mny = y
			}
			if mxx < x {
				mxx = x
			}
			if mxy < y {
				mxy = y
			}
			if v != '#' {
				r[c] = make([]Blizzard, 0)
			}
			switch v {
			case 'v':
				r[c] = append(r[c], Blizzard{coord: c, dir: SOUTH})
			case '<':
				r[c] = append(r[c], Blizzard{coord: c, dir: WEST})
			case '^':
				r[c] = append(r[c], Blizzard{coord: c, dir: NORTH})
			case '>':
				r[c] = append(r[c], Blizzard{coord: c, dir: EAST})
			}
		}
	}
	for i, v := range lines[0] {
		if v == '.' {
			start = helper.Coord{X: i, Y: 0}
			break
		}
	}
	for i, v := range lines[len(lines)-1] {
		if v == '.' {
			end = helper.Coord{X: i, Y: len(lines) - 1}
			break
		}
	}
	return r, start, end
}

func makeCycles(m map[helper.Coord][]Blizzard) map[int]map[helper.Coord][]Blizzard {
	r := ranges(m)

	cycles := make(map[int]map[helper.Coord][]Blizzard)
	cycles[0] = m

	seen := make(map[string]bool)
	h := hash(m, 1, r)
	seen[h] = true
	done := false
	i := 1
	for !done {
		ns := nextState(cycles[i-1], r)
		cycles[i] = ns
		h := hash(ns, 1, r)
		if _, ok := seen[h]; ok {
			done = true
		} else {
			seen[h] = true
		}
		i++
	}

	return cycles
}

func makeNextStates(s State, cycleLength int, r Range) []State {

	nextStates := make([]State, 0)

	cycleTick := (s.tick + 1) % len(cycles)
	nextState := nextState(cycles[cycleTick], r)

	//does the player need to move?
	mustMove := false
	if len(nextState[s.p.c]) > 0 {
		// there are blizzards moving onto the player's position, so yes
		mustMove = true
	}

	neighbours := s.p.c.GetNeighbours(false)
	for _, neighbour := range neighbours {
		if blizzards, seen := nextState[neighbour]; seen {
			if len(blizzards) == 0 {
				// the player can move here
				validState := State{
					p: Player{
						c:     neighbour,
						turns: s.p.turns + 1,
					},
					tick: (s.tick + 1) % len(cycles),
				}
				nextStates = append(nextStates, validState)
			}
		}
	}
	if mustMove && len(nextStates) == 0 {
		// this is a failed state, we can't wait
		return nextStates
	}

	if !mustMove {
		// add waiting
		waitState := State{
			p: Player{
				c:     s.p.c,
				turns: s.p.turns + 1,
			},
			tick: (s.tick + 1) % len(cycles),
		}
		nextStates = append(nextStates, waitState)
	}

	return nextStates
}

func copy(m []helper.Coord) []helper.Coord {
	r := make([]helper.Coord, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func problem(lines []string, partTwo bool) (int, error) {
	m, pstart, end := readMap(lines)
	r := ranges(m)
	cycleLength := r.maxx - r.minx + 1
	p := Player{c: pstart, turns: 0}
	cycles = makeCycles(m)

	states := make([]State, 0)
	firstState := State{
		p:    p,
		tick: 0,
	}
	states = append(states, firstState)
	minTurns := math.MaxInt

	seenStates := make(map[string]int)

	var stage1Mins, stage2Mins, stage3Mins int
	var stage1 State
	var stage2 State

	for len(states) > 0 {
		state := states[0]
		states = states[1:]

		// ignore states where the player has already been at the same cycle, with more turns
		sn := fmt.Sprintf("%v,%v:%v", state.p.c.X, state.p.c.Y, state.tick)
		if _, seen := seenStates[sn]; seen {
			continue
		}
		seenStates[sn] = state.p.turns

		// Are we already exceeding our minimum turns?
		if state.p.turns > minTurns {
			continue
		}

		// Has the player reached the end?
		if state.p.c == end {
			if state.p.turns < minTurns {
				// Was it the least turns?
				minTurns = state.p.turns
				stage1 = state
			}
			continue
		}

		nextstates := makeNextStates(state, cycleLength, r)
		for _, ns := range nextstates {
			states = append(states, ns)
		}
	}

	stage1Mins = minTurns
	fmt.Printf("Stage 1 done: %v\n", stage1Mins)

	// Now stage 2
	seenStates = make(map[string]int)
	states = make([]State, 0)
	minTurns = math.MaxInt
	stage1.p.turns = 0
	states = append(states, stage1)

	for len(states) > 0 {
		state := states[0]
		states = states[1:]

		// ignore states where the player has already been at the same cycle, with more turns
		sn := fmt.Sprintf("%v,%v:%v", state.p.c.X, state.p.c.Y, state.tick)
		if _, seen := seenStates[sn]; seen {
			continue
		}
		seenStates[sn] = state.p.turns

		// Are we already exceeding our minimum turns?
		if state.p.turns > minTurns {
			continue
		}

		// Has the player reached the end?
		if state.p.c == pstart {
			if state.p.turns < minTurns {
				// Was it the least turns?
				minTurns = state.p.turns
				stage2 = state
			}
			continue
		}

		nextstates := makeNextStates(state, cycleLength, r)
		for _, ns := range nextstates {
			states = append(states, ns)
		}
	}

	stage2Mins = minTurns
	fmt.Printf("Stage 2 done: %v\n", stage2Mins)

	// Now stage 3
	seenStates = make(map[string]int)
	states = make([]State, 0)
	minTurns = math.MaxInt
	stage2.p.turns = 0
	states = append(states, stage2)

	for len(states) > 0 {
		state := states[0]
		states = states[1:]

		// ignore states where the player has already been at the same cycle, with more turns
		sn := fmt.Sprintf("%v,%v:%v", state.p.c.X, state.p.c.Y, state.tick)
		if _, seen := seenStates[sn]; seen {
			continue
		}
		seenStates[sn] = state.p.turns

		// Are we already exceeding our minimum turns?
		if state.p.turns > minTurns {
			continue
		}

		// Has the player reached the end?
		if state.p.c == end {
			if state.p.turns < minTurns {
				// Was it the least turns?
				minTurns = state.p.turns
			}
			continue
		}

		nextstates := makeNextStates(state, cycleLength, r)
		for _, ns := range nextstates {
			states = append(states, ns)
		}
	}
	stage3Mins = minTurns
	fmt.Printf("Stage 3 done: %v\n", stage3Mins)

	return stage1Mins + stage2Mins + stage3Mins, nil
}

func nextState(m map[helper.Coord][]Blizzard, r Range) map[helper.Coord][]Blizzard {
	grid := make(map[helper.Coord][]Blizzard)
	for c, blizzards := range m {
		if _, seen := grid[c]; !seen {
			grid[c] = make([]Blizzard, 0)
		}
		if len(blizzards) == 0 {
			continue
		}
		for _, blizzard := range blizzards {
			var nc helper.Coord
			switch blizzard.dir {
			case NORTH:
				nc = helper.Coord{X: c.X, Y: c.Y - 1}
				if _, seen := m[nc]; !seen {
					// move to bottom
					nc.Y = r.maxy - 1
				}
			case SOUTH:
				nc = helper.Coord{X: c.X, Y: c.Y + 1}
				if _, seen := m[nc]; !seen {
					// move to bottom
					nc.Y = r.miny + 1
				}
			case EAST:
				nc = helper.Coord{X: c.X + 1, Y: c.Y}
				if _, seen := m[nc]; !seen {
					// move to west
					nc.X = r.minx
				}
			case WEST:
				nc = helper.Coord{X: c.X - 1, Y: c.Y}
				if _, seen := m[nc]; !seen {
					// move to east
					nc.X = r.maxx
				}
			}
			grid[nc] = append(grid[nc], Blizzard{coord: nc, dir: blizzard.dir})
		}
	}
	return grid
}

func printWithPath(m map[helper.Coord][]Blizzard, path []helper.Coord, r Range) {
	for y := r.miny; y <= r.maxy; y++ {
		for x := r.minx; x <= r.maxx; x++ {
			c := helper.Coord{X: x, Y: y}

			printPlayer := false
			for _, v := range path {
				if c == v {
					fmt.Printf("P")
					printPlayer = true
				}
			}
			if printPlayer {
				continue
			}

			if _, seen := m[c]; !seen {
				fmt.Printf("#")
				continue
			}
			if len(m[c]) == 0 {
				fmt.Printf(".")
			} else if len(m[c]) > 1 {
				fmt.Printf("%v", len(m[c]))
			} else {
				switch m[c][0].dir {
				case NORTH:
					fmt.Printf("^")
				case SOUTH:
					fmt.Printf("v")
				case EAST:
					fmt.Printf(">")
				case WEST:
					fmt.Printf("<")
				}
			}
		}
		fmt.Println()
	}
}

func print(m map[helper.Coord][]Blizzard, r Range) {
	for y := r.miny; y <= r.maxy; y++ {
		fmt.Printf("%s\n", hash(m, y, r))
	}
}

func hash(m map[helper.Coord][]Blizzard, row int, r Range) string {
	s := ""
	for x := r.minx; x <= r.maxx; x++ {
		c := helper.Coord{X: x, Y: row}
		if _, seen := m[c]; !seen {
			s += "#"
			continue
		}
		if len(m[c]) == 0 {
			s += "."
		} else if len(m[c]) > 1 {
			s += fmt.Sprintf("%v", len(m[c]))
		} else {
			switch m[c][0].dir {
			case NORTH:
				s += "^"
			case SOUTH:
				s += "v"
			case EAST:
				s += ">"
			case WEST:
				s += "<"
			}
		}
	}
	return s
}

func ranges(m map[helper.Coord][]Blizzard) Range {
	minx := math.MaxInt
	miny := math.MaxInt
	maxx := math.MinInt
	maxy := math.MinInt

	for b, _ := range m {
		if b.X < minx {
			minx = b.X
		}
		if b.Y < miny {
			miny = b.Y
		}
		if b.X > maxx {
			maxx = b.X
		}
		if b.Y > maxy {
			maxy = b.Y
		}
	}

	return Range{
		minx: minx,
		miny: miny,
		maxx: maxx,
		maxy: maxy,
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
	
}
