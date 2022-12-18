package main

import (
	"github.com/mrbarge/aoc2022/helper"
	"testing"
)

func TestOverlap(t *testing.T) {
	tests := []struct {
		name   string
		r1     FallingRock
		r2     FallingRock
		result bool
	}{
		{
			name: "1",
			r1: FallingRock{
				rockType: 0,
				c:        helper.Coord{X: 0, Y: 4},
			},
			r2: FallingRock{
				rockType: 1,
				c:        helper.Coord{X: 0, Y: 4},
			},
			result: true,
		},
		{
			name: "2",
			r1: FallingRock{
				rockType: 0,
				c:        helper.Coord{X: 0, Y: 4},
			},
			r2: FallingRock{
				rockType: 1,
				c:        helper.Coord{X: 0, Y: 3},
			},
			result: false,
		},
		{
			name: "3",
			r1: FallingRock{
				rockType: 2,
				c:        helper.Coord{X: 0, Y: 4},
			},
			r2: FallingRock{
				rockType: 4,
				c:        helper.Coord{X: 0, Y: 3},
			},
			result: false,
		},
	}
	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			res := test.r1.Overlap([]FallingRock{test.r2})
			res2 := test.r2.Overlap([]FallingRock{test.r1})
			if res != test.result || res2 != test.result {
				t.Errorf("%v, expected %v but got %v/%v", test.name, test.result, res, res2)
			}
		})
	}
}

func TestAtEdge(t *testing.T) {
	tests := []struct {
		name   string
		r1     FallingRock
		result bool
	}{
		{
			name: "2",
			r1: FallingRock{
				rockType: 0,
				c:        helper.Coord{X: 1, Y: 4},
			},
			result: false,
		},
		{
			name: "4",
			r1: FallingRock{
				rockType: 4,
				c:        helper.Coord{X: 5, Y: 0},
			},
			result: true,
		},
		{
			name: "5",
			r1: FallingRock{
				rockType: 4,
				c:        helper.Coord{X: 4, Y: 0},
			},
			result: false,
		},
	}
	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			res := test.r1.HitRightEdge()
			if res != test.result {
				t.Errorf("%v, expected %v but got %v", test.name, test.result, res)
			}
		})
	}
}
