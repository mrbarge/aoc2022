package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	From int
	To   int
}

func (s *Section) ContainsAll(t *Section) bool {
	return t.From >= s.From && t.To <= s.To
}

func (s *Section) ContainsPartial(t *Section) bool {
	for i := t.From; i <= t.To; i++ {
		if i >= s.From && i <= s.To {
			return true
		}
	}
	return false
}

func parseSection(s string) (*Section, error) {
	n := strings.Split(s, "-")
	if len(n) < 2 {
		return nil, fmt.Errorf("invalid section: %s", s)
	}
	n1, err := strconv.Atoi(n[0])
	if err != nil {
		return nil, fmt.Errorf("invalid section segment %s", n[0])
	}
	n2, err := strconv.Atoi(n[1])
	if err != nil {
		return nil, fmt.Errorf("invalid section segment %s", n[1])
	}
	return &Section{
		From: n1,
		To:   n2,
	}, nil
}

func problem(lines []string, partTwo bool) (int, error) {
	count := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		if len(sections) < 2 {
			return 0, fmt.Errorf("invalid line: %s", line)
		}
		e1, err := parseSection(sections[0])
		if err != nil {
			return 0, err
		}
		e2, err := parseSection(sections[1])
		if err != nil {
			return 0, err
		}
		if partTwo {
			if e1.ContainsPartial(e2) || e2.ContainsPartial(e1) {
				count++
			}
		} else {
			if e1.ContainsAll(e2) || e2.ContainsAll(e1) {
				count++
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

	ans, err = problem(lines, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
