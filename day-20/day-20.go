package main

import (
	"errors"
	"math"
	"regexp"
	"sort"
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

func puzzle2(input string) int {
	pieces := mapInputToPieces(input)
	puzzle := buildPuzzle(pieces)
	marked := markSeaMonsters(puzzle)
	s := ""
	for _, m := range marked {
		s += m
	}
	return strings.Count(s, "#")
}

type position struct {
	x int
	y int
}

func markSeaMonsters(puzzle []string) []string {
	rotations := 0
	hsm := hasSeaMonsters(puzzle)
	for !hsm {
		if rotations < 3 {
			rotations++
			puzzle = rotate90(puzzle)
			hsm = hasSeaMonsters(puzzle)
			continue
		}
		rotations = 0
		puzzle = flipAroundX(rotate90(puzzle))
		hsm = hasSeaMonsters(puzzle)
	}

	result := make([]string, len(puzzle))
	copy(result, puzzle)
	earliestHeadPosition := 18
	latestHeadPosition := len(puzzle[0]) - 2
	for i := 0; i < len(result)-1; i++ {
		for j := earliestHeadPosition; j <= latestHeadPosition; j++ {
			seaMonster := []position{
				{x: i, y: j},
				{x: i + 1, y: j + 1},
				{x: i + 1, y: j},
				{x: i + 1, y: j - 1},
				{x: i + 1, y: j - 6},
				{x: i + 1, y: j - 7},
				{x: i + 1, y: j - 12},
				{x: i + 1, y: j - 13},
				{x: i + 1, y: j - 18},
				{x: i + 2, y: j - 17},
				{x: i + 2, y: j - 14},
				{x: i + 2, y: j - 11},
				{x: i + 2, y: j - 8},
				{x: i + 2, y: j - 5},
				{x: i + 2, y: j - 2},
			}
			if checkSeaMonster(result, seaMonster) {
				result = replaceSeaMonster(result, seaMonster)
			}
		}
	}
	return result
}

func hasSeaMonsters(puzzle []string) bool {
	earliestHeadPosition := 18
	latestHeadPosition := len(puzzle[0]) - 2
	for i := 0; i < len(puzzle)-1; i++ {
		for j := earliestHeadPosition; j <= latestHeadPosition; j++ {
			seaMonster := []position{
				{x: i, y: j},
				{x: i + 1, y: j + 1},
				{x: i + 1, y: j},
				{x: i + 1, y: j - 1},
				{x: i + 1, y: j - 6},
				{x: i + 1, y: j - 7},
				{x: i + 1, y: j - 12},
				{x: i + 1, y: j - 13},
				{x: i + 1, y: j - 18},
				{x: i + 2, y: j - 17},
				{x: i + 2, y: j - 14},
				{x: i + 2, y: j - 11},
				{x: i + 2, y: j - 8},
				{x: i + 2, y: j - 5},
				{x: i + 2, y: j - 2},
			}
			if checkSeaMonster(puzzle, seaMonster) {
				return true
			}
		}
	}
	return false
}

func checkSeaMonster(puzzle []string, seaMonster []position) bool {
	for _, position := range seaMonster {
		if string(puzzle[position.x][position.y]) != "#" {
			return false
		}
	}
	return true
}

func replaceSeaMonster(puzzle []string, seaMonster []position) []string {
	result := make([]string, len(puzzle))
	copy(result, puzzle)

	for _, position := range seaMonster {
		row := result[position.x]
		result[position.x] = row[:position.y] + "0" + row[position.y+1:]
	}
	return result
}

func buildPuzzle(pieces map[int]piece) []string {
	sortedPieces := make(map[int][]piece)
	for _, piece := range pieces {
		optionCount := len(piece.getMatchingEdges(pieces))
		sortedPieces[optionCount] = append(sortedPieces[optionCount], piece)
	}

	piece1 := rotateTopLeftPiece(sortedPieces[2][0], pieces)
	occupiedSlots := occupySlots(piece1, sortedPieces, pieces)
	if !areAllOccupied(occupiedSlots) {
		piece1 := rotateTopLeftPiece(sortedPieces[2][0].flipAroundX(), pieces)
		occupiedSlots = occupySlots(piece1, sortedPieces, pieces)
	}
	return combinePieces(occupiedSlots)
}

func combinePieces(slots [][]slot) (result []string) {
	width := len(slots)
	for y := 0; y < width; y++ {
		for i := 0; i < len(slots[0][0].piece.inner); i++ {
			row := ""
			for x := 0; x < width; x++ {
				piece := slots[y][x].piece
				row += piece.inner[i]
			}
			result = append(result, row)
		}
	}
	return
}

func rotateTopLeftPiece(p piece, pieces map[int]piece) piece {
	result := p
	indices := result.getMatchingEdges(pieces)
	for indices[0] != 1 || indices[1] != 2 {
		result = result.rotate90()
		indices = result.getMatchingEdges(pieces)
	}
	return result
}

func occupySlots(startingPiece piece, sortedPieces map[int][]piece, pieces map[int]piece) [][]slot {
	width := int(math.Sqrt(float64(len(pieces))))
	slots := make([][]slot, width)
	for y := range slots {
		slots[y] = make([]slot, width)
	}
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			slots[y][x] = slot{x: x, y: y}
		}
	}
	slots[0][0].occupyWith(startingPiece)

	remainingPieces := make(map[int][]piece)
	for k, v := range sortedPieces {
		newV := make([]piece, len(v))
		copy(newV, v)
		remainingPieces[k] = newV
	}
	remainingPieces[2] = remove(remainingPieces[2], 0)
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			if x == 0 && y == 0 {
				continue
			}
			slot := slots[y][x]
			connectedEdges := getNumberOfConnectedEdges(x, y, width)
			options := remainingPieces[connectedEdges]
			for i, option := range options {
				accepts, orientatedOption := slot.accepts(option, slots)
				if accepts {
					slots[y][x].occupyWith(orientatedOption)
					remainingPieces[connectedEdges] = remove(remainingPieces[connectedEdges], i)
					break
				}
			}
		}
	}
	return slots
}

func getNumberOfConnectedEdges(x, y, width int) int {
	if (y == 0 && x == width-1) || (x == 0 && y == width-1) || (x == width-1 && y == width-1) {
		return 2
	} else if y == 0 || x == 0 || x == width-1 || y == width-1 {
		return 3
	}
	return 4
}

func areAllOccupied(slots [][]slot) bool {
	width := len(slots)
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			slot := slots[x][y]
			if !slot.isOccupied {
				return false
			}
		}
	}
	return true
}

func mapInputToPieces(input string) map[int]piece {
	result := make(map[int]piece)
	split := strings.Split(input, "\n\n")
	for _, part := range split {
		lines := strings.Split(part, "\n")

		edgeWidth := len(lines[1])

		edge1 := edge("")
		edge3 := edge("")

		inner := []string{}
		for i := 1; i <= edgeWidth; i++ {
			edge1 += edge(lines[i][edgeWidth-1])
			edge3 += edge(lines[i][0])
			if i == 1 || i == edgeWidth {
				continue
			}
			inner = append(inner, lines[i][1:edgeWidth-1])
		}

		edge0 := edge(lines[1])
		edge2 := edge(lines[edgeWidth])

		r := regexp.MustCompile(`\d+`)
		id := ConvertToInt(r.FindStringSubmatch(lines[0])[0])
		result[id] = piece{id: id, edges: [4]edge{edge0, edge1, reverse(edge2), reverse(edge3)}, inner: inner}
	}
	return result
}

type slot struct {
	isOccupied bool
	piece      piece
	x          int
	y          int
}

type piece struct {
	id    int
	edges [4]edge
	inner []string
}

type edge string

func (p piece) isCornerPiece(pieces map[int]piece) bool {
	return len(p.getMatchingEdges(pieces)) == 2
}

func (p piece) getMatchingEdge(p2 piece) (int, error) {
	for i, e := range p.edges {
		for _, e2 := range p2.edges {
			if e == e2 || e == reverse(e2) {
				return i, nil
			}
		}
	}
	return 0, errors.New("No matching edge found")
}

func (p piece) rotate90() piece {
	edges := [4]edge{
		p.edges[3],
		p.edges[0],
		p.edges[1],
		p.edges[2],
	}

	inner := rotate90(p.inner)

	return piece{id: p.id, edges: edges, inner: inner}
}

func rotate90(s []string) []string {
	result := make([]string, len(s))
	for i := range s {
		for _, row := range s {
			result[i] = string(row[i]) + result[i]
		}
	}
	return result
}

func (p piece) flipAroundX() piece {
	edges := [4]edge{
		reverse(p.edges[2]),
		reverse(p.edges[1]),
		reverse(p.edges[0]),
		reverse(p.edges[3]),
	}

	inner := flipAroundX(p.inner)

	return piece{id: p.id, edges: edges, inner: inner}
}

func flipAroundX(s []string) []string {
	return reverseSlice(s)
}

func (p piece) getMatchingEdges(pieces map[int]piece) []int {
	result := []int{}
	for _, piece := range pieces {
		if piece.id == p.id {
			continue
		}
		index, err := p.getMatchingEdge(piece)
		if err == nil && !contains(result, index) {
			result = append(result, index)
		}
	}
	sort.Ints(result)
	return result
}

func (s *slot) occupyWith(p piece) {
	s.piece = p
	s.isOccupied = true
}

func (s *slot) accepts(p piece, slots [][]slot) (bool, piece) {
	result := p
	occupiedAdjacents := []slot{}
	if s.x != 0 {
		occupiedAdjacents = append(occupiedAdjacents, slots[s.y][s.x-1])
	}
	if s.y != 0 {
		occupiedAdjacents = append(occupiedAdjacents, slots[s.y-1][s.x])
	}

	accepts := false
	for i, adjacent := range occupiedAdjacents {
		if i == 0 {
			index, err := result.getMatchingEdge(adjacent.piece)
			if err == nil {
				if adjacent.x != s.x {
					for index != 3 {
						result = result.rotate90()
						index, _ = result.getMatchingEdge(adjacent.piece)
					}
					if result.edges[3] == adjacent.piece.edges[1] {
						result = result.flipAroundX()
					}
				} else {
					for index != 0 {
						result = result.rotate90()
						index, _ = result.getMatchingEdge(adjacent.piece)
					}
					if result.edges[0] == adjacent.piece.edges[2] {
						result = result.flipAroundX().rotate90().rotate90()
					}
				}
				accepts = true
				continue
			}
		} else {
			index, err := result.getMatchingEdge(adjacent.piece)
			accepts = accepts && err == nil && index == 0
		}
	}
	return accepts, result
}

func reverse(e edge) edge {
	runes := []rune(e)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return edge(runes)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(s []piece, i int) []piece {
	result := make([]piece, len(s))
	copy(result, s)
	result[i] = result[len(s)-1]
	return result[:len(s)-1]
}

func reverseSlice(s []string) []string {
	result := make([]string, len(s))
	copy(result, s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
