package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const (
	LT    = -1
	EQUAL = 0
	GT    = 1
)

var (
	//go:embed input.txt
	input string
)

type Packet interface{}

func compareSlices(left interface{}, right interface{}) int {

	if left == nil {
		return GT
	}
	if right == nil {
		return LT
	}
	lv := left.([]interface{})
	rv := right.([]interface{})
	i := 0

	for ; i < len(lv) && i < len(rv); i++ {

		lv := lv[i]
		rv := rv[i]

		res := EQUAL

		switch lvt := lv.(type) {
		case []interface{}:
			switch rvt := rv.(type) {
			case []interface{}:
				res = compareSlices(lvt, rvt)
			case float64:
				ri := make([]interface{}, 1)
				ri[0] = rv
				res = compareSlices(lv, ri)
			default:
				fmt.Printf("%T received\n", rvt)
			}
		case float64:
			switch rvt := rv.(type) {
			case []interface{}:
				li := make([]interface{}, 1)
				li[0] = lv
				res = compareSlices(li, rvt)
			case float64:
				if lv.(float64) < rv.(float64) {
					return LT
				} else if lv.(float64) > rv.(float64) {
					return GT
				}
			default:
				fmt.Printf("%T received\n", rvt)
			}
		default:
			fmt.Printf("%T received, %v", lvt)
		}

		if res != EQUAL {
			return res
		}
	}

	if i == len(lv) && i == len(rv) {
		return EQUAL
	}
	if i < len(rv) {
		return LT
	}
	return GT
}

func problem(input string, partTwo bool) (int, error) {

	pairBlocks := strings.Split(input, "\n\n")
	sum := 0
	for i, p := range pairBlocks {
		p1 := strings.Split(p, "\n")[0]
		p2 := strings.Split(p, "\n")[1]

		var p1json interface{}
		var p2json interface{}

		json.Unmarshal([]byte(p1), &p1json)
		json.Unmarshal([]byte(p2), &p2json)

		r := compareSlices(p1json, p2json)
		if r == LT {
			sum += i + 1
			fmt.Printf("Right order: %v\n", i+1)
		}
	}
	return sum, nil
}

func partTwo(input string) (int, error) {

	ni := input + "\n[[2]]\n[[6]]"
	lines := strings.Split(ni, "\n")

	packets := make([]interface{}, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		var p1json interface{}
		json.Unmarshal([]byte(line), &p1json)

		packets = append(packets, p1json)
	}

	sort.Slice(packets, func(i, j int) bool {
		r := compareSlices(packets[i], packets[j])
		return r == LT
	})

	for i, v := range packets {
		fmt.Printf("%d: %v\n", i, v)
	}
	return 0, nil
}

func main() {
	ans, err := problem(input, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = partTwo(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
