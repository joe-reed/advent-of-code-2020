package main

import (
	"fmt"
	"testing"
	. "utils"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			25,
		},
		{
			[]string{
				"R360",
				"F10",
			},
			10,
		},
		{
			[]string{
				"F10",
				"L180",
				"F10",
			},
			0,
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

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			286,
		},
		{
			[]string{
				"R360",
				"F10",
			},
			110,
		},
		{
			[]string{
				"F10",
				"L180",
				"F10",
			},
			0,
		},
	}

	for _, test := range tests {
		a := puzzle2(test.input)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(getInput()))
}

func getInput() []string {
	return LoadFile("./input.txt")
}
