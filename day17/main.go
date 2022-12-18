package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
)

const (
	WellWidth = 7
)

type FallingRock struct {
	c        helper.Coord
	rockType int
}

type CycleState struct {
	tick    int64
	hd      int
	rockIdx int
	windIdx int
}

type Rock [][]bool

var (
	heightDiff = []int{
		-4,
		-6,
		-6,
		-7,
		-5,
	}
	minHeightDiff = []int{
		-1,
		-3,
		-3,
		-4,
		-2,
	}
	rock1 = Rock{
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
		{true, true, true, true},
	}
	rock2 = Rock{
		{false, false, false, false},
		{false, true, false, false},
		{true, true, true, false},
		{false, true, false, false},
	}
	rock3 = Rock{
		{false, false, false, false},
		{false, false, true, false},
		{false, false, true, false},
		{true, true, true, false},
	}
	rock4 = Rock{
		{true, false, false, false},
		{true, false, false, false},
		{true, false, false, false},
		{true, false, false, false},
	}
	rock5 = Rock{
		{false, false, false, false},
		{false, false, false, false},
		{true, true, false, false},
		{true, true, false, false},
	}
	rocks = []Rock{
		rock1,
		rock2,
		rock3,
		rock4,
		rock5,
	}
	cycleCache = make(map[string]bool)
)

func (r FallingRock) HitLeftEdge() bool {
	return r.c.X == 0
}

func (r FallingRock) HitRightEdge() bool {
	rt := rocks[r.rockType]
	for _, row := range rt {
		for x, val := range row {
			xc := r.c.X + x
			if val == true && xc >= WellWidth-1 {
				return true
			}
		}
	}
	return false
}

func (r FallingRock) Overlap(vs []FallingRock) (bool, int) {
	cmap := make(map[int]map[int]bool)
	rt := rocks[r.rockType]
	for y, row := range rt {
		for x, val := range row {
			yc := r.c.Y - len(rt) + y
			xc := r.c.X + x
			if _, ok := cmap[yc]; !ok {
				cmap[yc] = make(map[int]bool, 0)
			}
			cmap[yc][xc] = val
		}
	}
	for vi := len(vs) - 1; vi >= 0; vi-- {
		v := vs[vi]
		vt := rocks[v.rockType]
		for y, row := range vt {
			for x, val := range row {
				yc := v.c.Y - len(vt) + y
				xc := v.c.X + x
				if _, ok := cmap[yc]; !ok {
					continue
				}
				if val && cmap[yc][xc] {
					return true, vi
				}
			}
		}
	}
	return false, -1
}

func (r FallingRock) HitFloor(floor int) bool {
	return r.c.Y >= floor
}

func problem2(line string, limit int64, partTwo bool, debug bool) (int64, error) {

	numRocks := int64(0)
	rockIdx := 0
	windIdx := 0
	allRocks := []FallingRock{}
	upperLimit := -3
	upperRockLimit := 0
	floor := upperLimit + 3

	cycles := make([]CycleState, 0)
	seenHash := make(map[string]int64)
	rockHash := make(map[string]int64)

	for numRocks < limit {
		rock := FallingRock{
			rockType: rockIdx,
			c: helper.Coord{X: 2,
				Y: upperLimit,
			},
		}

		// see where it lands
		rest := false
		for !rest {

			nextCoord := helper.Coord{X: rock.c.X, Y: rock.c.Y}

			// see where wind blows
			windDir := line[windIdx]
			windIdx = (windIdx + 1) % len(line)

			if windDir == '<' {
				if !rock.HitLeftEdge() {
					nextCoord.X -= 1
					nextRock := FallingRock{
						rockType: rock.rockType,
						c:        nextCoord,
					}
					overlap, _ := nextRock.Overlap(allRocks)
					if !overlap {
						rock.c = nextCoord
					} else {
						nextCoord.X += 1
					}
				}
			} else if windDir == '>' {
				if !rock.HitRightEdge() {
					nextCoord.X += 1
					nextRock := FallingRock{
						rockType: rock.rockType,
						c:        nextCoord,
					}
					overlap, _ := nextRock.Overlap(allRocks)
					if !overlap {
						rock.c = nextCoord
					} else {
						nextCoord.X -= 1
					}
				}
			}

			if rock.HitFloor(floor) {
				rest = true
			} else {
				nextCoord.Y++
				nextRock := FallingRock{
					rockType: rock.rockType,
					c:        nextCoord,
				}
				overlap, _ := nextRock.Overlap(allRocks)
				if overlap {
					rest = true
				} else {
					rock.c = nextCoord
				}
			}
		}

		allRocks = append(allRocks, rock)
		heighestRock := findHeighest(allRocks)
		upperLimit = heighestRock.c.Y + heightDiff[heighestRock.rockType]
		upperRockLimit = int(math.Abs(float64(heighestRock.c.Y + minHeightDiff[heighestRock.rockType])))

		if partTwo && numRocks > 2000 {
			// enough time for a loop to commence
			hd := math.Abs(float64(rock.c.Y+minHeightDiff[rock.rockType])) - math.Abs(float64(upperRockLimit))
			cs := CycleState{
				tick:    numRocks,
				hd:      int(hd),
				rockIdx: rockIdx,
				windIdx: windIdx,
			}

			cycles = append(cycles, cs)
			if len(cycles) > 10 {
				cycles = cycles[1:]
				chash := hashCycles(cycles)
				if _, ok := seenHash[chash]; ok {
					cycleRockDiff := numRocks - rockHash[chash]
					rocksLeft := limit - numRocks

					// look for a clean divisor
					if rocksLeft%cycleRockDiff == 0 {
						cycleHeightDiff := int64(upperRockLimit) - seenHash[chash]
						forecastHeight := int64(upperRockLimit)
						v := int64(0)
						for ; v < rocksLeft; v += cycleRockDiff {
							forecastHeight += cycleHeightDiff
						}
						return forecastHeight - 1, nil
					}
				} else {
					seenHash[chash] = int64(upperRockLimit)
					rockHash[chash] = numRocks
				}
			}
		}

		rockIdx = (rockIdx + 1) % len(rocks)
		numRocks++

	}
	return int64(math.Abs(float64(upperRockLimit))), nil
}

func hashCycles(c []CycleState) string {
	r := ""
	for _, v := range c {
		s := fmt.Sprintf("%d:%d:%d", v.hd, v.rockIdx, v.windIdx)
		r += s
	}
	return r
}

func findHeighest(fr []FallingRock) FallingRock {
	smallest := 0
	smallestRock := FallingRock{}
	for _, rock := range fr {
		hr := rock.c.Y + minHeightDiff[rock.rockType]
		if hr <= smallest {
			smallest = hr
			smallestRock = rock
		}
	}
	return smallestRock
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem2(lines[0], 2022, false, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem2(lines[0], 1000000000000, true, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
