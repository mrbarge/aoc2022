package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
)

func problem(lines []string, partTwo bool) (int, error) {

	count := 0

	coords := make([]helper.Coord3D, 0)
	for _, line := range lines {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			return 0, err
		}
		coord := helper.Coord3D{X: x, Y: y, Z: z}
		coords = append(coords, coord)
	}

	for _, c := range coords {
		sides := c.EmptySides(coords)
		count += sides
	}
	return count, nil
}

func makeSeenCoords(s []helper.Coord3D) map[string]bool {
	r := make(map[string]bool)
	for _, v := range s {
		r[v.AsString()] = true
	}
	return r
}

func unique(s1 map[string]bool, s2 map[string]bool) map[string]bool {
	r := make(map[string]bool)
	for k, _ := range s1 {
		if _, seen := s2[k]; !seen {
			r[k] = true
		}
	}
	return r
}

func problem2(lines []string, partTwo bool) (int, error) {

	count := 0

	coords := make([]helper.Coord3D, 0)
	for _, line := range lines {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			return 0, err
		}
		coord := helper.Coord3D{X: x, Y: y, Z: z}
		coords = append(coords, coord)
	}

	xmin, xmax, ymin, ymax, zmin, zmax := helper.Ranges(coords)
	allCoords := makeSeenCoords(coords)

	perimeter := make(map[string]bool, 0)

	for _, c := range coords {
		neighbours := c.AdjNeighbours()
		for _, n := range neighbours {
			outside := false

			seencubes := make(map[string]bool)
			queue := make([]helper.Coord3D, 0)
			queue = append(queue, n)

			for len(queue) > 0 {
				cv := queue[0]
				queue = queue[1:]
				if _, seen := seencubes[cv.AsString()]; seen {
					continue
				}
				seencubes[cv.AsString()] = true

				if _, seen := perimeter[cv.AsString()]; seen {
					un := unique(seencubes, allCoords)
					for k, _ := range un {
						perimeter[k] = true
					}
					outside = true
					break
				}

				if cv.X < xmin || cv.X > xmax || cv.Y < ymin || cv.Y > ymax || cv.Z < zmin || cv.Z > zmax {
					un := unique(seencubes, allCoords)
					for k, _ := range un {
						perimeter[k] = true
					}
					outside = true
					break
				}

				if _, seen := allCoords[cv.AsString()]; !seen {
					cvn := cv.AdjNeighbours()
					for _, nn := range cvn {
						queue = append(queue, nn)
					}
				}
			}

			if outside {
				count += 1
			}
		}
	}
	return count, nil
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

	ans, err = problem2(lines, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
