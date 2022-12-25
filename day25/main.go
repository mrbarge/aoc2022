package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
)

var snafuLookup = map[string]int64{
	"2": 2,
	"1": 1,
	"0": 0,
	"-": -1,
	"=": -2,
}
var snafuReverseLookup = map[int64]string{
	2:  "2",
	1:  "1",
	0:  "0",
	-1: "-",
	-2: "=",
}

func snafuToInt(s string) int64 {
	r := int64(0)
	l := len(s) - 1
	for i := l; i >= 0; i-- {
		r += (int64(math.Pow(float64(5), float64(i))) * snafuLookup[string(s[l-i])])
	}
	return r
}

func intToSnafu(i int64) string {
	r := make([]string, 0)

	for i > 0 {
		m := i % 5
		i = i / 5
		if m > 2 {
			m -= 5
			i += 1
		}
		r = append(r, snafuReverseLookup[m])
	}

	ret := ""
	for i := len(r) - 1; i >= 0; i-- {
		ret += r[i]
	}
	return ret
}

func problem(lines []string) (string, error) {
	sum := int64(0)
	for _, line := range lines {
		sum += snafuToInt(line)
	}
	fmt.Printf("Sum is: %v\n", sum)
	r := intToSnafu(sum)
	return r, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

}
