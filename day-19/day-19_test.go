package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	. "utils"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			2,
		},
		{
			"./test-input-2.txt",
			8,
		},
		{
			"./test-input-3.txt",
			2,
		},
	}

	for _, test := range tests {
		a := puzzle1(getInput(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(getInput("./input.txt")))
}

func getInput(path string) string {
	file, err := ioutil.ReadFile(path)
	Check(err)
	return string(file)
}
