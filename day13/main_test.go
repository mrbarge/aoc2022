package main

import (
	"testing"
)

func TestLastBracketPos(t *testing.T) {
	tests := []struct {
		name   string
		val    string
		result int
	}{
		{
			name:   "t1",
			val:    "[xx]",
			result: 3,
		},
		{
			name:   "t2",
			val:    "[[]",
			result: -1,
		},
		{
			name:   "t3",
			val:    "[[]]",
			result: 3,
		},
		{
			name:   "t4",
			val:    "[x[x]x]",
			result: 6,
		},
		{
			name:   "t5",
			val:    "[[]][]",
			result: 3,
		},
	}
	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			r := lastBracketPos(test.val)
			if r != test.result {
				t.Errorf("%v, expected %v but got %v", test.name, test.result, r)
			}
		})
	}
}
