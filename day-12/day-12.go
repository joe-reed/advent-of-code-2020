package main

import (
	"math"
	"regexp"
	. "utils"
)

func puzzle1(input []string) int {
	instructions := getInstructions(input)
	f := Ferry{x: 0, y: 0, d: 90}
	for _, instruction := range instructions {
		instruction.Execute(&f)
	}
	return f.Distance()
}

func puzzle2(input []string) int {
	instructions := getInstructions(input)
	f := Ferry{x: 0, y: 0, d: 90}
	w := Waypoint{x: 10, y: 1}
	for _, instruction := range instructions {
		instruction.ExecuteWithWaypoint(&f, &w)
	}
	return f.Distance()
}

func getInstructions(input []string) []Instruction {
	var r []Instruction
	for _, line := range input {
		r = append(r, getInstruction(line))
	}
	return r
}

func getInstruction(line string) Instruction {
	r := regexp.MustCompile(`(.{1})(\d+)`)
	m := r.FindStringSubmatch(line)
	val := ConvertToInt(m[2])
	switch m[1] {
	case "N":
		return Move{axis: "y", amount: val}
	case "S":
		return Move{axis: "y", amount: -val}
	case "E":
		return Move{axis: "x", amount: val}
	case "W":
		return Move{axis: "x", amount: -val}
	case "R":
		return Rotate{degrees: val}
	case "L":
		return Rotate{degrees: 360 - val}
	default:
		return Forward{amount: val}
	}
}

type Ferry struct {
	x int
	y int
	d int
}

type Waypoint struct {
	x int
	y int
}

func (f *Ferry) Distance() int {
	return int(math.Abs(float64(f.x)) + math.Abs(float64(f.y)))
}

type Instruction interface {
	Execute(*Ferry)
	ExecuteWithWaypoint(*Ferry, *Waypoint)
}

type Move struct {
	axis   string
	amount int
}

type Rotate struct {
	degrees int
}

type Forward struct {
	amount int
}

func (m Move) Execute(f *Ferry) {
	switch m.axis {
	case "x":
		f.x += m.amount
	case "y":
		f.y += m.amount
	}
}

func (r Rotate) Execute(f *Ferry) {
	d := (f.d + r.degrees) % 360
	if d < 0 {
		d += 360
	}
	f.d = d
}

func (fo Forward) Execute(f *Ferry) {
	switch f.d {
	case 0:
		f.y += fo.amount
	case 90:
		f.x += fo.amount
	case 180:
		f.y -= fo.amount
	case 270:
		f.x -= fo.amount
	}
}

func (m Move) ExecuteWithWaypoint(f *Ferry, w *Waypoint) {
	switch m.axis {
	case "x":
		w.x += m.amount
	case "y":
		w.y += m.amount
	}
}

func (r Rotate) ExecuteWithWaypoint(f *Ferry, w *Waypoint) {
	switch r.degrees {
	case 90:
		x := w.x
		w.x = w.y
		w.y = -x
	case 180:
		w.x = -w.x
		w.y = -w.y
	case 270:
		x := w.x
		w.x = -w.y
		w.y = x
	}
}

func (fo Forward) ExecuteWithWaypoint(f *Ferry, w *Waypoint) {
	f.x += fo.amount * w.x
	f.y += fo.amount * w.y
}
