package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	. "utils"
)

func TestEdge(t *testing.T) {
	e1 := edge("#..#...#")
	e2 := edge("#..#...#")
	if !e1.equals(e2) {
		t.Errorf("expeted %s to equal %s", e1, e2)
	}

	e3 := edge(".....##.")
	e4 := edge("#..#...#")
	if e3.equals(e4) {
		t.Errorf("expeted %s not to equal %s", e3, e4)
	}
}

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
		t.Errorf("expeted matching edge %d to equal 1", index)
	}

	_, err := p1.getMatchingEdge(p3)
	if err == nil {
		t.Errorf("received unexpected matching edge")
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

func getInput(path string) string {
	file, err := ioutil.ReadFile(path)
	Check(err)
	return string(file)
}
