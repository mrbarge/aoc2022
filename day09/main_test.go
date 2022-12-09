package main

import (
	"github.com/mrbarge/aoc2022/helper"
	"testing"
)

func TestTouching(t *testing.T) {
	tests := []struct {
		name   string
		c1     helper.Coord
		c2     helper.Coord
		result bool
	}{
		{
			name: "neighbour of x",
			c1: helper.Coord{
				X: 1,
				Y: 0,
			},
			c2: helper.Coord{
				X: 0,
				Y: 0,
			},
			result: true,
		},
		{
			name: "separated from x",
			c1: helper.Coord{
				X: 1,
				Y: 0,
			},
			c2: helper.Coord{
				X: 3,
				Y: 0,
			},
			result: false,
		},
		{
			name: "diagonal",
			c1: helper.Coord{
				X: 1,
				Y: 0,
			},
			c2: helper.Coord{
				X: 0,
				Y: 1,
			},
			result: true,
		},
		{
			name: "same space",
			c1: helper.Coord{
				X: 0,
				Y: 0,
			},
			c2: helper.Coord{
				X: 0,
				Y: 0,
			},
			result: true,
		},
		{
			name: "neighbour of y",
			c1: helper.Coord{
				X: 0,
				Y: 1,
			},
			c2: helper.Coord{
				X: 0,
				Y: 0,
			},
			result: true,
		},
	}
	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			r := touching(test.c1, test.c2)
			if r != test.result {
				t.Errorf("%v, expected %v but got %v", test.name, test.result, r)
			}
		})
	}
}
