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
			getTestInput(),
			10,
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
		rounds   int
		expected int
	}{
		{
			getTestInput(),
			1,
			15,
		},
		{
			getTestInput(),
			2,
			12,
		},
		{
			getTestInput(),
			3,
			25,
		},
		{
			getTestInput(),
			4,
			14,
		},
		{
			getTestInput(),
			100,
			2208,
		},
	}

	for _, test := range tests {
		a := puzzle2(test.input, test.rounds)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}
func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(getInput(), 100))
}

func getInput() []string {
	return LoadFile("./input.txt")
}

func getTestInput() []string {
	return []string{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew",
	}
}
