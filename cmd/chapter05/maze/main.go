package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := Point{1, 1}
	end := Point{38, 38}
	maze := NewMaze(40, 40, start, end, "maze.txt")

	newPos, _ := maze.StepAhead()
	time.Sleep(1 * time.Second)

	if newPos != None {
		fmt.Println(newPos)
	}

	for {
		if newPos == None || newPos.Equals(end) {
			break
		}
		newPos, _ = maze.StepAhead()

		time.Sleep(100 * time.Millisecond)
		if newPos != None {
			fmt.Println(newPos, maze.GetMoveCount())
		}
	}

	if newPos.Equals(end) {
		fmt.Println("SUCCESS!Reached ", end)
	}
}
