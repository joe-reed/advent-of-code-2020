package main

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"0,3,6",
			436,
		},
		{
			"1,3,2",
			1,
		},
		{
			"2,1,3",
			10,
		},
		{
			"1,2,3",
			27,
		},
		{
			"2,3,1",
			78,
		},
		{
			"3,2,1",
			438,
		},
		{
			"3,1,2",
			1836,
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
			"0,3,6",
			175594,
		},
		{
			"1,3,2",
			2578,
		},
		{
			"2,1,3",
			3544142,
		},
		{
			"1,2,3",
			261214,
		},
		{
			"2,3,1",
			6895259,
		},
		{
			"3,2,1",
			18,
		},
		{
			"3,1,2",
			362,
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
	return "6,3,15,13,1,0"
}
