package main

import (
	"fmt"
	"testing"
	. "utils"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		startTime int
		schedule  string
		expected  int
	}{
		{
			939,
			"7,13,x,x,59,x,31,19",
			295,
		},
	}

	for _, test := range tests {
		a := puzzle1(test.startTime, test.schedule)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	startTime, schedule := getInput()
	fmt.Println("Puzzle 1:", puzzle1(startTime, schedule))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		schedule string
		expected int
	}{
		{
			"7,13,x,x,59,x,31,19",
			1068781,
		},
		{
			"17,x,13,19",
			3417,
		},
		{
			"67,7,59,61",
			754018,
		},
		{
			"67,x,7,59,61",
			779210,
		},
		{
			"67,7,x,59,61",
			1261476,
		},
		{
			"1789,37,47,1889",
			1202161486,
		},
	}

	for _, test := range tests {
		a := puzzle2(test.schedule)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	_, schedule := getInput()
	fmt.Println("Puzzle 2:", puzzle2(schedule))
}

func BenchmarkPuzzle2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, schedule := getInput()
		puzzle2(schedule)
	}
}

func TestModInverse(t *testing.T) {
	a := modInverse(40, 7)
	if 3 != a {
		t.Errorf("expected: %d, actual: %d", 3, a)
	}
}

func getInput() (int, string) {
	input := LoadFile("./input.txt")
	return ConvertToInt(input[0]), input[1]
}
