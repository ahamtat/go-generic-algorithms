package main

import (
	"fmt"
	"math/rand"
)

// Direction abstraction
type Direction int

const (
	N Direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
	NotAvailable
)

func (d Direction) String() string {
	switch d {
	case N:
		return "north"
	case NE:
		return "north-east"
	case E:
		return "east"
	case SE:
		return "south-east"
	case S:
		return "south"
	case SW:
		return "south-west"
	case W:
		return "west"
	case NW:
		return "north-west"
	case NotAvailable:
		return "not available"
	}
	return "unknown"
}

func (d Direction) PrintDirection() {
	fmt.Println("direction: ", d)
}

// Point abstraction
type Point struct {
	x, y int
}

func (p Point) Equals(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func (p Point) PrintPoint() {
	fmt.Printf("<%d, %d>\n", p.x, p.y)
}

// Path abstraction
type Path struct {
	point          Point
	move           Direction
	movesAvailable []Direction
}

func NewPath(point Point) Path {
	path := Path{point, Direction(NotAvailable), []Direction{}}
	path.move = NotAvailable

	// Initially all directions available
	path.movesAvailable = []Direction{N, NE, E, SE, S, SW, W, NW}
	return path
}

func (path *Path) RandomMove() Direction {
	// Returns value of move and changes the receiver
	var indicesAvailable []int
	for index := 0; index < 8; index++ {
		if path.movesAvailable[index] != NotAvailable {
			indicesAvailable = append(indicesAvailable, index)
		}
	}

	if count := len(indicesAvailable); count > 0 {
		randomIndex := rand.Intn(count)
		path.move = path.movesAvailable[indicesAvailable[randomIndex]]
		path.movesAvailable[indicesAvailable[randomIndex]] = NotAvailable
		return path.move
	} else {
		return NotAvailable
	}
}
