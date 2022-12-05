package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
	"strconv"
	"strings"
)

func testInitialise() [][]string {
	stacks := make([][]string, 0)
	stacks = append(stacks, []string{"Z", "N"})
	stacks = append(stacks, []string{"M", "C", "D"})
	stacks = append(stacks, []string{"P"})
	return stacks
}

func initialise() [][]string {
	stacks := make([][]string, 0)
	stacks = append(stacks, []string{"B", "G", "S", "C"})
	stacks = append(stacks, []string{"T", "M", "W", "H", "J", "N", "V", "G"})
	stacks = append(stacks, []string{"M", "Q", "S"})
	stacks = append(stacks, []string{"B", "S", "L", "T", "W", "N", "M"})
	stacks = append(stacks, []string{"J", "Z", "F", "T", "V", "G", "W", "P"})
	stacks = append(stacks, []string{"C", "T", "B", "G", "Q", "H", "S"})
	stacks = append(stacks, []string{"T", "J", "P", "B", "W"})
	stacks = append(stacks, []string{"G", "D", "C", "Z", "F", "T", "Q", "M"})
	stacks = append(stacks, []string{"N", "S", "H", "B", "P", "F"})
	return stacks
}

func problem(lines []string, partTwo bool) (string, error) {

	stacks := initialise()
	//stacks := testInitialise()

	for _, line := range lines {
		elems := strings.Split(line, " ")
		numMove, _ := strconv.Atoi(elems[1])
		from, _ := strconv.Atoi(elems[3])
		to, _ := strconv.Atoi(elems[5])

		from--
		to--

		fromStack := stacks[from]
		if partTwo {
			for i := len(fromStack) - numMove; i <= len(fromStack)-1; i++ {
				stacks[to] = append(stacks[to], fromStack[i])
			}
		} else {
			for i := len(fromStack) - 1; i >= len(fromStack)-numMove; i-- {
				stacks[to] = append(stacks[to], fromStack[i])
			}
		}

		fromStack = fromStack[:len(fromStack)-numMove]
		stacks[from] = fromStack
	}

	r := ""
	for _, stack := range stacks {
		r += string(stack[len(stack)-1])
	}
	return r, nil
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
