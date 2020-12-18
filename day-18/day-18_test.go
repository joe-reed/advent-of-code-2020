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
			"1 + 2 * 3 + 4 * 5 + 6",
			71,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			26,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			437,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			12240,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			13632,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))\n((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			25872,
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
		input    string
		expected int
	}{
		{
			"1 + 2 * 3 + 4 * 5 + 6",
			231,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			46,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			1445,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			669060,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			23340,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))\n((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			692400,
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

func BenchmarkPuzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		puzzle2(getInput())
	}
}

func getInput() string {
	file, err := ioutil.ReadFile("./input.txt")
	Check(err)
	return string(file)
}
