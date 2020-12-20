package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	. "utils"
)

func TestPiece(t *testing.T) {
	e1 := edge("#..#...#")
	e2 := edge("#..##..#")
	e3 := edge(".....##.")
	e4 := edge("#..#...#")
	e5 := edge("........")
	e6 := edge("#.......")
	e7 := edge("#......#")
	e8 := edge("##.....#")

	p1 := piece{edges: [4]edge{e1, e2, e3, e4}}
	p2 := piece{edges: [4]edge{e5, e2, e6, e7}}
	p3 := piece{edges: [4]edge{e5, e6, e7, e8}}

	index, _ := p1.getMatchingEdge(p2)
	if index != 1 {
		t.Errorf("expected: 1, actual: %d", index)
	}

	edge, err := p1.getMatchingEdge(p3)
	if err == nil {
		t.Errorf("received unexpected matching edge %d", edge)
	}
}

func TestRotate90(t *testing.T) {
	// ..#..
	// #.###
	// ..#.#
	// ##.#.
	// #..#.
	e := [4]edge{
		edge("..#.."),
		edge(".##.."),
		edge(".#..#"),
		edge("##.#."),
	}
	i := []string{".##", ".#.", "#.#"}
	p := piece{edges: e, inner: i}
	rotated := p.rotate90()

	// ##.#.
	// .#...
	// ..###
	// ##.#.
	// .##..
	expectedEdges := [4]edge{
		edge("##.#."),
		edge("..#.."),
		edge(".##.."),
		edge(".#..#"),
	}
	expectedInner := []string{"#..", ".##", "#.#"}
	expectedPiece := piece{edges: expectedEdges, inner: expectedInner}
	checkPiecesAreEqual(expectedPiece, rotated, t)

	rotated = p.rotate90().rotate90().rotate90().rotate90()
	checkPiecesAreEqual(p, rotated, t)
}

func TestFlipAroundX(t *testing.T) {
	// ..#..
	// #.###
	// ..#.#
	// ##.#.
	// #..#.
	e := [4]edge{
		edge("..#.."),
		edge(".##.."),
		edge(".#..#"),
		edge("##.#."),
	}
	i := []string{".##", ".#.", "#.#"}
	p := piece{edges: e, inner: i}
	flipped := p.flipAroundX()

	// #..#.
	// ##.#.
	// ..#.#
	// #.###
	// ..#..
	expectedEdges := [4]edge{
		edge("#..#."),
		edge("..##."),
		edge("..#.."),
		edge(".#.##"),
	}
	expectedInner := []string{"#.#", ".#.", ".##"}
	expectedPiece := piece{edges: expectedEdges, inner: expectedInner}
	checkPiecesAreEqual(expectedPiece, flipped, t)

	flipped = p.flipAroundX().flipAroundX()
	checkPiecesAreEqual(p, flipped, t)
}

func TestFlipAndRotate(t *testing.T) {
	// ..#..
	// #.###
	// ..#.#
	// ##.#.
	// #..#.
	e := [4]edge{
		edge("..#.."),
		edge(".##.."),
		edge(".#..#"),
		edge("##.#."),
	}
	i := []string{".##", ".#.", "#.#"}
	p := piece{edges: e, inner: i}

	flippedAndRotated := p.flipAroundX().rotate90().rotate90()

	// ..#..
	// ###.#
	// #.#..
	// .#.##
	// .#..#
	expectedEdges := [4]edge{
		edge("..#.."),
		edge(".#.##"),
		edge("#..#."),
		edge("..##."),
	}
	expectedInner := []string{"##.", ".#.", "#.#"}
	expectedPiece := piece{edges: expectedEdges, inner: expectedInner}
	checkPiecesAreEqual(expectedPiece, flippedAndRotated, t)
}

func checkPiecesAreEqual(expectedPiece piece, actualPiece piece, t *testing.T) {
	for i, expectedEdge := range expectedPiece.edges {
		if actualPiece.edges[i] != expectedEdge {
			t.Errorf("expected edge: %s, actual: %s (row %d)", expectedEdge, actualPiece.edges[i], i)
		}
	}
	for i, expectedInnerRow := range expectedPiece.inner {
		if actualPiece.inner[i] != expectedInnerRow {
			t.Errorf("expected inner row: %s, actual: %s, (row %d)", expectedInnerRow, actualPiece.inner[i], i)
		}
	}
}

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			20899048083289,
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

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			273,
		},
	}

	for _, test := range tests {
		a := puzzle2(getInput(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(getInput("./input.txt")))
}

func getInput(path string) string {
	file, err := ioutil.ReadFile(path)
	Check(err)
	return string(file)
}
