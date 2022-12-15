package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
	"sort"
)

type boundary struct {
	minx int
	maxx int
}

func problem(lines []string, rowOfInterest int) (int, error) {

	nobeacons := make(map[int]map[int]bool, 0)
	beacons := make([]helper.Coord, 0)

	for _, line := range lines {
		var sx, sy, cx, cy int

		r, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &cx, &cy)
		if err != nil || r != 4 {
			return 0, err
		}

		scrd := helper.Coord{X: sx, Y: sy}
		ccrd := helper.Coord{X: cx, Y: cy}
		beacons = append(beacons, ccrd)
		md := helper.ManhattanDistance(scrd, ccrd)

		for y := sy + md; y >= sy-md; y-- {
			if y != rowOfInterest {
				continue
			}

			xw := md - int(math.Abs(float64(y)-float64(sy)))
			for x := sx - xw; x <= sx+xw; x++ {
				if _, ok := nobeacons[y]; !ok {
					nobeacons[y] = make(map[int]bool, 0)
				}
				nobeacons[y][x] = true
			}
		}
	}

	for _, beacon := range beacons {
		delete(nobeacons[beacon.Y], beacon.X)
	}
	return len(nobeacons[rowOfInterest]), nil
}

func compareSlices(left interface{}, right interface{}) int {
	l := left.(boundary)
	r := right.(boundary)
	if l.minx < r.minx {
		return -1
	} else if l.minx == r.minx {
		return 0
	} else {
		return 1
	}
}

func problemTwo(lines []string, bLimit int) (uint64, error) {
	nobeacons := make(map[int]map[int]bool, 0)
	beacons := make([]helper.Coord, 0)

	boundaries := make([][]boundary, bLimit)
	for x := 0; x < bLimit; x++ {
		boundaries[x] = make([]boundary, 0)
	}

	for _, line := range lines {
		var sx, sy, cx, cy int

		r, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &cx, &cy)
		if err != nil || r != 4 {
			return 0, err
		}

		scrd := helper.Coord{X: sx, Y: sy}
		ccrd := helper.Coord{X: cx, Y: cy}
		beacons = append(beacons, ccrd)
		md := helper.ManhattanDistance(scrd, ccrd)

		for y := sy + md; y >= sy-md; y-- {
			if y < 0 || y >= bLimit {
				continue
			}
			xw := md - int(math.Abs(float64(y)-float64(sy)))
			xmn := sx - xw
			xmx := sx + xw

			if xmn < 0 {
				xmn = 0
			}
			if xmx > bLimit {
				xmx = bLimit
			}
			//fmt.Printf("Adding %d boundary %v to %v\n", y, xmn, xmx)
			boundaries[y] = append(boundaries[y], boundary{minx: xmn, maxx: xmx})
		}
	}

	for y, b := range boundaries {

		sort.Slice(b, func(i, j int) bool {
			if b[i].minx < b[j].minx {
				return true
			}
			return false
		})

		if len(b) <= 1 {
			continue
		}
		mnx := b[0].minx
		mxm := b[0].maxx
		for _, boundary := range b {
			if boundary.minx > mxm+1 {
				// another range starting
				val := uint64((boundary.minx-1)*4000000) + uint64(y)
				return val, nil
			}
			if boundary.minx < mnx {
				mnx = boundary.minx
			}
			if boundary.maxx > mxm {
				mxm = boundary.maxx
			}
		}
	}
	fmt.Printf("%v\n", nobeacons)
	fmt.Printf("%v\n", len(nobeacons))
	return 0, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, 2000000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans2, err := problemTwo(lines, 4000000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans2)

}
