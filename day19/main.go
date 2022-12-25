package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"os"
)

type Mineral int

const (
	Ore Mineral = iota
	Clay
	Obsidian
	Geode
)

type Blueprint map[Mineral]Robot

type Robot struct {
	cost map[Mineral]int
}

type State struct {
	// map of minute
	min int
	// map of robot type to num of robots
	robots map[Mineral]int
	// map of mineral type to num collected
	total map[Mineral]int
	// what's in the queue to build
	buildqueue []Mineral
}

func makeBlueprint(s string) (Blueprint, error) {
	var blueprint, oreOre, clayOre, obsidianOre, obsidianClay, geodeOre, geodeObsidian int

	_, err := fmt.Sscanf(s, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. "+
		"Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&blueprint, &oreOre, &clayOre, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)

	if err != nil {
		return nil, err
	}

	bp := make(map[Mineral]Robot, 0)
	bp[Ore] = Robot{
		cost: map[Mineral]int{
			Ore: oreOre,
		},
	}
	bp[Clay] = Robot{
		cost: map[Mineral]int{
			Ore: clayOre,
		},
	}
	bp[Obsidian] = Robot{
		cost: map[Mineral]int{
			Ore:  obsidianOre,
			Clay: obsidianClay,
		},
	}
	bp[Geode] = Robot{
		cost: map[Mineral]int{
			Ore:      geodeOre,
			Obsidian: geodeObsidian,
		},
	}

	return bp, nil
}

func problem(lines []string, limit int, partTwo bool) (int, error) {
	count := 0

	blueprints := make([]Blueprint, 0)
	for _, line := range lines {
		bp, err := makeBlueprint(line)
		if err != nil {
			return 0, err
		}
		blueprints = append(blueprints, bp)
	}

	if partTwo {
		count = 1
		for i, blueprint := range blueprints {
			geodes := run(blueprint, limit)
			fmt.Printf("Blueprint %v produced %v\n", i+1, geodes)
			count *= geodes
		}
	} else {
		for i, blueprint := range blueprints {
			geodes := run(blueprint, limit)
			fmt.Printf("Blueprint %v produced %v\n", i+1, geodes)
			count += geodes * (i + 1)
		}
	}
	return count, nil
}

func run(blueprint Blueprint, limit int) int {
	state := State{
		min: 1,
		robots: map[Mineral]int{
			Ore:      1,
			Geode:    0,
			Clay:     0,
			Obsidian: 0,
		},
		total: map[Mineral]int{
			Ore:      0,
			Geode:    0,
			Clay:     0,
			Obsidian: 0,
		},
		buildqueue: make([]Mineral, 0),
	}

	states := make([]State, 0)
	states = append(states, state)

	seenStates := make(map[string]bool)
	maxGeodes := 0

	for len(states) > 0 {
		s := states[0]
		states = states[1:]

		if maxGeodes < s.total[Geode] {
			// this state has made the most geodes
			maxGeodes = s.total[Geode]
		}

		if s.min > limit {
			// exceeded time, quit
			continue
		}

		if _, ok := seenStates[h]; ok {
			// been here, ignore
			continue
		}
		seenStates[s.hash()] = true

		//fmt.Printf("Min:%v ", s.min-1)
		//fmt.Printf("Built: O:%d,C:%d,OB:%d,G:%d  Mined: O:%d,C:%d,OB:%d,G:%d ", s.robots[Ore],
		//	s.robots[Clay], s.robots[Obsidian], s.robots[Geode], s.total[Ore], s.total[Clay], s.total[Obsidian], s.total[Geode])
		//fmt.Println()

		// what robots could we make next?
		nextStates := s.nextStates(blueprint)

		for _, nextState := range nextStates {

			// update totals that robots have mined
			for _, v := range []Mineral{Ore, Clay, Obsidian, Geode} {
				nextState.total[v] += nextState.robots[v]
			}

			// produce a robot if we built one
			if len(nextState.buildqueue) >= 1 {
				nextState.robots[nextState.buildqueue[0]] += 1
				nextState.buildqueue = nextState.buildqueue[1:]
			}

			states = append(states, nextState)
		}

	}

	return maxGeodes
}

func (s State) nextStates(blueprint Blueprint) []State {
	r := make([]State, 0)

	for _, mineral := range []Mineral{Geode, Obsidian, Clay, Ore} {
		// build robot of that type
		canbuild := true
		botToBuild := blueprint[mineral]
		for minNeeded, mincost := range botToBuild.cost {
			if s.total[minNeeded] < mincost {
				canbuild = false
				break
			}
		}

		if canbuild {
			ns := s.copy()
			ns.min += 1
			ns.buildqueue = append(ns.buildqueue, mineral)

			// update totals based on building the bot
			for minNeeded, minCost := range botToBuild.cost {
				ns.total[minNeeded] -= minCost
			}
			r = append(r, ns)
		}
	}

	// Also add a state where we do nothing
	ns := s.copy()
	ns.min += 1
	r = append(r, ns)

	return r
}

func (s State) copy() State {
	ns := State{
		min:        s.min,
		robots:     make(map[Mineral]int),
		total:      make(map[Mineral]int),
		buildqueue: make([]Mineral, 0),
	}
	for k, v := range s.robots {
		ns.robots[k] = v
	}
	for k, v := range s.total {
		ns.total[k] = v
	}
	for _, v := range s.buildqueue {
		ns.buildqueue = append(ns.buildqueue, v)
	}
	return ns
}

func (s State) hash() string {
	r := fmt.Sprintf("%d:", s.min)
	r += fmt.Sprintf("%d:%d:%d:%d:%d:%d:%d:%d ",
		Ore, s.total[Ore], Clay, s.total[Clay],
		Obsidian, s.total[Obsidian], Geode, s.total[Geode])
	for k, v := range s.buildqueue {
		r += fmt.Sprintf("%d:%d:", k, v)
	}
	return r
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, 24, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines[0:3], 32, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
