package main

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{
				5764801,
				17807724,
			},
			14897079,
		},
	}

	for _, test := range tests {
		a := puzzle1(test.input)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}
func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(getInput()))
}

func getInput() []int {
	return []int{
		11239946,
		10464955,
	}
}
