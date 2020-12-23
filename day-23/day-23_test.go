package main

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    []int
		moves    int
		expected int
	}{
		{
			[]int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			10,
			92658374,
		},
		{
			[]int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			100,
			67384529,
		},
	}

	for _, test := range tests {
		a := puzzle1(test.input, test.moves)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:",
		puzzle1(
			[]int{3, 6, 4, 2, 8, 9, 7, 1, 5},
			100,
		),
	)
}
