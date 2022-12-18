package main

import (
	"fmt"
	"github.com/albertorestifo/dijkstra"
	"github.com/mrbarge/aoc2022/helper"
	"os"
	"sort"
	"strings"
)

const (
	p1MaxTime = 30
	p2MaxTime = 26
)

var (
	routes = make(map[string]int)
)

type StepState struct {
	tunnel    string
	pressure  int
	flow      int
	tick      int
	running   int
	visited   []string
	unvisited []string
}

func problem(lines []string, partTwo bool, limit int) (int, error) {

	g := dijkstra.Graph{}

	routes = make(map[string]int)

	tcon := make(map[string][]string)
	flowrates := make(map[string]int)
	for _, line := range lines {
		var valve string
		var fr int

		_, err := fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnels lead to valves", &valve, &fr)
		if err != nil {
			return 0, err
		}
		tcon[valve] = strings.Split(strings.ReplaceAll(strings.Split(line, "valves")[1], " ", ""), ",")
		flowrates[valve] = fr
	}

	for t, v := range tcon {
		g[t] = map[string]int{}
		for _, tv := range v {
			g[t][tv] = 1
		}
	}

	distances := make(map[string]map[string]int)
	for t, _ := range tcon {
		distances[t] = make(map[string]int)
		for t2, _ := range tcon {
			_, score, err := g.Path(t, t2)
			if err != nil {
				return 0, err
			}
			distances[t][t2] = score
		}
	}

	unvisited := make([]string, 0)
	for k, v := range flowrates {
		if v != 0 {
			unvisited = append(unvisited, k)
		}
	}
	startingState := StepState{
		tunnel:    "AA",
		pressure:  0,
		flow:      0,
		tick:      0,
		unvisited: unvisited,
	}

	pressure := step(distances, flowrates, startingState, limit)
	//fmt.Printf("%v\n", routes)

	ans := 0
	if partTwo {
		for path1, score1 := range routes {
			for path2, score2 := range routes {
				if unique(path1, path2) {
					s := score1 + score2
					if ans < s {
						fmt.Printf("%v and %v are unique, setting max %d\n", path1, path2, s)
						ans = s
					}
				}
			}
		}
	} else {
		ans = pressure
	}
	return ans, nil
}

func unique(a string, b string) bool {
	al := strings.Split(a, ",")
	bl := strings.Split(b, ",")
	for _, v := range al {
		for _, v2 := range bl {
			if v == v2 {
				return false
			}
		}
	}
	return true
}

func remove(v string, l []string) []string {
	r := make([]string, 0)
	for _, e := range l {
		if e != v {
			r = append(r, e)
		}
	}
	return r
}

func step(distances map[string]map[string]int, rates map[string]int, state StepState, limit int) int {
	max := state.pressure + (limit-state.tick)*state.flow

	for _, nextTunnel := range state.unvisited {
		travelTime := distances[state.tunnel][nextTunnel]
		totalAction := travelTime + 1
		if state.tick+totalAction < limit {
			// can do this in time
			nextState := StepState{
				tick:      state.tick + totalAction,
				pressure:  state.pressure + (totalAction * state.flow),
				flow:      state.flow + rates[nextTunnel],
				visited:   visit(state.tunnel, state.visited),
				unvisited: remove(nextTunnel, state.unvisited),
				running:   state.running + (limit-state.tick-travelTime-1)*rates[nextTunnel],
				tunnel:    nextTunnel,
			}

			path := strings.Join(nextState.visited, ",")
			if _, ok := routes[path]; ok {
				if routes[path] < state.running {
					routes[path] = state.running
				}
			} else {
				routes[path] = state.running
			}

			//fmt.Printf("%v to %v Calculating (%d - %d - 1) * %d\n", state.tunnel, nextState.tunnel, (limit - state.tick), travelTime, rates[nextTunnel])

			ns := step(distances, rates, nextState, limit)
			if ns > max {
				max = ns
			}
		}
	}

	return max
}

func visit(t string, visited []string) []string {
	r := make([]string, 0)
	for _, v := range visited {
		r = append(r, v)
	}
	if t != "AA" {
		// don't record starting step
		r = append(r, t)
	}
	sort.Strings(r)
	return r
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false, p1MaxTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true, p2MaxTime)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
