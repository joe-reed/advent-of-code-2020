package main

import (
	"errors"
	"regexp"
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	result := 1
	pieces := mapInputToPieces(input)
	for _, piece := range pieces {
		if piece.isCornerPiece(pieces) {
			result *= piece.id
		}
	}
	return result
}

func mapInputToPieces(input string) map[int]piece {
	result := make(map[int]piece)
	split := strings.Split(input, "\n\n")
	for _, part := range split {
		lines := strings.Split(part, "\n")

		edgeWidth := len(lines[1])

		edge2 := edge("")
		edge4 := edge("")

		for i := 1; i <= edgeWidth; i++ {
			edge2 += edge(lines[i][edgeWidth-1])
			edge4 += edge(lines[i][0])
		}

		edge1 := edge(lines[1])
		edge3 := edge(lines[edgeWidth])

		r := regexp.MustCompile(`\d+`)
		id := ConvertToInt(r.FindStringSubmatch(lines[0])[0])
		result[id] = piece{id: id, edges: [4]edge{edge1, edge2, edge3, edge4}}
	}
	return result
}

type piece struct {
	id    int
	edges [4]edge
}

type edge string

func (p piece) isCornerPiece(pieces map[int]piece) bool {
	return len(p.getOptions(pieces)) == 2
}

func (p piece) getOptions(pieces map[int]piece) map[int]piece {
	result := make(map[int]piece)
	for id, piece := range pieces {
		if id == p.id {
			continue
		}
		_, err := p.getMatchingEdge(piece)
		if err == nil {
			result[id] = piece
		}
	}
	return result
}

func (p piece) getMatchingEdge(p2 piece) (int, error) {
	for i, e := range p.edges {
		for _, e2 := range p2.edges {
			if e.equals(e2) || e.equals(reverse(e2)) {
				return i, nil
			}
		}
	}
	return 0, errors.New("No matching edge found")
}

func (e edge) equals(e2 edge) bool {
	for i, v := range e {
		if string(e2[i]) != string(v) {
			return false
		}
	}
	return true
}

func reverse(s edge) edge {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return edge(runes)
}
