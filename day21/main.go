package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
	"strings"
)

type Monkey struct {
	name      string
	monkey1   string
	monkey2   string
	operation string
	value     int64
	isSet     bool
}

func (m *Monkey) Known() bool {
	return m.isSet
}

func readMonkeys(lines []string) (map[string]*Monkey, error) {
	monkeys := make(map[string]*Monkey)

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		var name, m1, m2, op string
		var value int
		var isSet bool
		if len(tokens) > 2 {
			_, err := fmt.Sscanf(line, "%s %s %s %s", &name, &m1, &op, &m2)
			if err != nil {
				return nil, err
			}
		} else {
			_, err := fmt.Sscanf(line, "%s %d", &name, &value)
			if err != nil {
				return nil, err
			}
			isSet = true
		}

		name = strings.Split(name, ":")[0]
		monkey := Monkey{
			name:      name,
			value:     int64(value),
			monkey1:   m1,
			monkey2:   m2,
			operation: op,
			isSet:     isSet,
		}

		monkeys[name] = &monkey
	}

	return monkeys, nil
}

func problem(lines []string, partTwo bool) (int64, error) {
	done := false
	ans := int64(0)
	humn := int64(3343167719000)
	for !done {

		monkeys, err := readMonkeys(lines)
		if err != nil {
			return 0, err
		}

		rootMonkey := monkeys["root"]
		if partTwo {
			rootMonkey.operation = "="
			monkeys["humn"].value = int64(humn)
		}

		testing := true
		for testing {

			for _, monkey := range monkeys {
				if !monkey.Known() {
					if !(monkeys[monkey.monkey1].Known() && monkeys[monkey.monkey2].Known()) {
						continue
					}

					s1 := monkeys[monkey.monkey1].value
					s2 := monkeys[monkey.monkey2].value
					switch monkey.operation {
					case "+":
						monkey.value = s1 + s2
						monkey.isSet = true
					case "-":
						monkey.value = s1 - s2
						monkey.isSet = true
					case "*":
						monkey.value = s1 * s2
						monkey.isSet = true
					case "/":
						monkey.value = s1 / s2
						monkey.isSet = true
					case "=":
						fmt.Printf("Comparing %v and %v\n", s1, s2)
						if s1 != s2 {
							testing = false
						} else {
							monkey.value = s1
							monkey.isSet = true
						}
					}
				}
				if rootMonkey.Known() {
					testing = false
					break
				}
				if !testing {
					break
				}
			}
		}
		if rootMonkey.Known() {
			done = true
			if partTwo {
				ans = humn
			} else {
				ans = rootMonkey.value
			}
		}
		humn++
	}
	return ans, nil
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
