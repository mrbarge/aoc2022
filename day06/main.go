package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
)

func problem(line string, partTwo bool) (int, error) {

	m := make([]string, 0)
	mm := make(map[string]bool, 0)

	markerLength := 4
	if partTwo {
		markerLength = 14
	}

	for i, c := range line {
		cs := string(c)
		if _, seen := mm[cs]; !seen {
			if len(mm) == markerLength {
				return i, nil
			}
			m = append(m, cs)
		} else {
			cutpos := len(m) - 1
			for x := len(m) - 1; x >= 0; x-- {
				if m[x] == cs {
					cutpos = x
					break
				}
			}
			m = m[cutpos+1:]
			m = append(m, cs)
			mm = make(map[string]bool)
			for _, c := range m {
				mm[c] = true
			}
		}
		mm[cs] = true
	}
	return 0, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines[0], false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines[0], true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
