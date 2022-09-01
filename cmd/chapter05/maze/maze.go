package main

import (
	"bufio"
	"fmt"
	"github.com/ahamtat/go-generic-algorithms/pkg/stack"
	"log"
	"os"
)

var None = Point{-1, -1}

type Maze struct {
	rows, cols int
	start, end Point
	barriers   [][]bool
	current    Path
	moveCount  int
	pathStack  stack.Sliced[Path]
	gameOver   bool
}

func NewMaze(rows int, cols int, start Point, end Point, mazefile string) (maze Maze) {
	maze.rows = rows
	maze.cols = cols
	maze.start = start
	maze.end = end

	// Initialize maze.barriers
	maze.barriers = make([][]bool, rows)
	for i := range maze.barriers {
		maze.barriers[i] = make([]bool, cols)
	}

	file, err := os.Open(mazefile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textlines []string
	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}
	defer file.Close()

	for row := 0; row < rows; row++ {
		line := textlines[row]
		for col := 0; col < cols; col++ {
			maze.barriers[row][col] = string(line[col]) == "1"
		}
	}

	maze.current = NewPath(start)
	maze.pathStack = stack.Sliced[Path]{}
	maze.pathStack.Push(maze.current)
	maze.barriers[start.x][start.y] = true

	return maze
}

func NewPosition(oldPosition Point, move Direction) Point {
	switch move {
	case N:
		return Point{oldPosition.x, oldPosition.y - 1}
	case NE:
		return Point{oldPosition.x + 1, oldPosition.y - 1}
	case E:
		return Point{oldPosition.x + 1, oldPosition.y}
	case SE:
		return Point{oldPosition.x + 1, oldPosition.y + 1}
	case S:
		return Point{oldPosition.x, oldPosition.y + 1}
	case SW:
		return Point{oldPosition.x - 1, oldPosition.y + 1}
	case W:
		return Point{oldPosition.x - 1, oldPosition.y}
	default:
		return Point{oldPosition.x - 1, oldPosition.y - 1}
	}
}

func (m *Maze) StepAhead() (Point, Point) {
	var (
		backTrackPoint = None
		newPos         = None
		validMove      = false
	)

	for {
		if m.gameOver || validMove || m.pathStack.IsEmpty() {
			break
		}

		validMove = false
		m.current, _ = m.pathStack.Pop()
		m.moveCount += 1
		nextMove := m.current.RandomMove()

		for {
			if validMove || nextMove == NotAvailable {
				break
			}
			newPos = NewPosition(m.current.point, m.current.move)
			if !m.barriers[newPos.y][newPos.x] {
				validMove = true

				if newPos.Equals(m.end) {
					for {
						if m.pathStack.IsEmpty() {
							break
						}
						_, _ = m.pathStack.Pop()
					}
					m.gameOver = true
				}

				m.barriers[newPos.y][newPos.x] = true
				m.pathStack.Push(m.current)
				newPathObject := NewPath(newPos)
				m.pathStack.Push(newPathObject)
			} else {
				nextMove = m.current.RandomMove()
			}
		}

		if !validMove && !m.pathStack.IsEmpty() {
			fmt.Printf("\nBacktrack from %v to %v\n", m.current.point, m.pathStack.Top().point)
			backTrackPoint = m.pathStack.Top().point
		}
	}

	if m.pathStack.IsEmpty() {
		fmt.Println("No solution is possible")
		return None, None
	}

	return newPos, backTrackPoint
}

func (m *Maze) GetMoveCount() int {
	return m.moveCount
}
