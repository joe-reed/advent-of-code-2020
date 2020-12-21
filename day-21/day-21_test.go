package main

import (
	"fmt"
	"testing"
	. "utils"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		foodList []string
		expected int
	}{
		{
			[]string{
				"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
				"trh fvjkl sbzzf mxmxvkd (contains dairy)",
				"sqjhc fvjkl (contains soy)",
				"sqjhc mxmxvkd sbzzf (contains fish)",
			},
			5,
		},
	}

	for _, test := range tests {
		a := puzzle1(test.foodList)
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
		foodList []string
		expected string
	}{
		{
			[]string{
				"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
				"trh fvjkl sbzzf mxmxvkd (contains dairy)",
				"sqjhc fvjkl (contains soy)",
				"sqjhc mxmxvkd sbzzf (contains fish)",
			},
			"mxmxvkd,sqjhc,fvjkl",
		},
	}

	for _, test := range tests {
		a := puzzle2(test.foodList)
		if a != test.expected {
			t.Errorf("expected: %s, actual: %s", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(getInput()))
}

func getInput() []string {
	return LoadFile("./input.txt")
}
