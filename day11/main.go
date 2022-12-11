package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type Monkey struct {
	items       []int64
	addOp       int
	multOp      int
	testMod     int
	testDest    map[bool]int
	square      bool
	inspections int
}

func (m *Monkey) Inspect(i int) int64 {
	if m.square {
		return m.items[i] * m.items[i]
	}
	return m.items[i]*int64(m.multOp) + int64(m.addOp)
}

func (m *Monkey) Test(i int64) int {
	r := i%int64(m.testMod) == 0
	return m.testDest[r]
}

func readMonkey(m string) (*Monkey, error) {
	var items, opSign, opNum string
	var trueDest, falseDest, testMod int

	for _, line := range strings.Split(m, "\n") {
		if strings.Contains(line, "Starting items: ") {
			items = strings.Split(line, "Starting items: ")[1]
		}
		if strings.Contains(line, "Operation:") {
			fmt.Sscanf(line, "  Operation: new = old %s %s", &opSign, &opNum)
		}
		if strings.Contains(line, "Test: divisible") {
			fmt.Sscanf(line, "  Test: divisible by %d", &testMod)
		}
		if strings.Contains(line, "If true") {
			fmt.Sscanf(line, "    If true: throw to monkey %d", &trueDest)
		}
		if strings.Contains(line, "If false") {
			fmt.Sscanf(line, "    If false: throw to monkey %d", &falseDest)
		}
	}

	monkey := Monkey{items: make([]int64, 0), addOp: 0, multOp: 1}
	it := strings.ReplaceAll(items, " ", "")
	itx := strings.Split(it, ",")
	for _, i := range itx {
		in, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		monkey.items = append(monkey.items, int64(in))
	}

	if opNum == "old" {
		monkey.square = true
	} else {
		oi, err := strconv.Atoi(opNum)
		if err != nil {
			return nil, err
		}
		if opSign == "*" {
			monkey.multOp = oi
		}
		if opSign == "+" {
			monkey.addOp = oi
		}
	}

	monkey.testDest = make(map[bool]int, 0)
	monkey.testDest[true] = trueDest
	monkey.testDest[false] = falseDest
	monkey.testMod = testMod

	return &monkey, nil
}

func readInput(input string) ([]*Monkey, error) {
	monkeys := make([]*Monkey, 0)

	monkeyBlocks := strings.Split(input, "\n\n")
	for _, mb := range monkeyBlocks {
		monkey, err := readMonkey(mb)
		if err != nil {
			return nil, err
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys, nil
}

func calcModulus(monkeys []*Monkey) int64 {
	r := int64(1)
	for _, m := range monkeys {
		r *= int64(m.testMod)
	}
	return r
}

func problem(monkeys []*Monkey, rounds int, partTwo bool) (int, error) {

	p2Mod := calcModulus(monkeys)
	for i := 0; i < rounds; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for x := 0; x < len(monkey.items); x++ {
				// Inspect item
				ni := monkey.Inspect(x)

				// Increase inspection count
				monkey.inspections++

				// Update worry level
				nw := int64(0)
				if partTwo {
					nw = ni % p2Mod
				} else {
					nw = int64(math.Floor(float64(ni) / 3))
				}

				// Test worry level
				nm := monkey.Test(nw)

				// Send to next monkey
				nextMonkey := monkeys[nm]
				nextMonkey.items = append(nextMonkey.items, nw)
			}

			// All items dealt with, so clear them
			monkey.items = make([]int64, 0)
		}
	}

	inspections := make([]int, 0)
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	sort.Ints(inspections)

	if len(inspections) < 2 {
		return 0, fmt.Errorf("Not enough monkey inspections recorded")
	}

	monkeyBusiness := inspections[len(inspections)-1] * inspections[len(inspections)-2]
	return monkeyBusiness, nil
}

func main() {

	monkeys, err := readInput(input)
	if err != nil {
		panic(err)
	}

	ans, err := problem(monkeys, 20, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	monkeys, err = readInput(input)
	if err != nil {
		panic(err)
	}

	ans, err = problem(monkeys, 10000, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
