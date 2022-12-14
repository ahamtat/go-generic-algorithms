package main

var (
	rows int
	cols int
)

type Grid [][]bool

var grid Grid
var newGrid Grid

func (g *Grid) initializeGrid(r, c int) {
	rows = r
	cols = c
	*g = make([][]bool, rows)
	for row := 0; row < rows; row++ {
		(*g)[row] = make([]bool, cols)
	}
}

func Copy(target [][]bool, source [][]bool) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			target[row][col] = source[row][col]
		}
	}
}

func (g Grid) bringAlive(row, col int) {
	g[row][col] = true
}

func (g Grid) kill(row, col int) {
	g[row][col] = false
}

func (g Grid) numberLiveNeighbors(row, col int) int {
	result := 0
	if row > 0 && g[row-1][col] {
		result++
	}
	if row > 0 && col < cols-1 && g[row-1][col+1] {
		result += 1
	}
	if col < cols-1 && g[row][col+1] {
		result += 1
	}
	if row < rows-1 && col < cols-1 && g[row+1][col+1] {
		result += 1
	}
	if row < rows-1 && g[row+1][col] {
		result += 1
	}
	if row < rows-1 && col > 0 && g[row+1][col-1] {
		result += 1
	}
	if col > 0 && g[row][col-1] {
		result += 1
	}
	if row > 0 && col > 0 && g[row-1][col-1] {
		result += 1
	}
	return result
}

func (g Grid) evolveGrid() {
	Copy(newGrid, g)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveN := g.numberLiveNeighbors(row, col)
			// Rules 1 and 2
			if g[row][col] && (liveN < 2 || liveN >= 4) {
				newGrid[row][col] = false
			}
			// Rule 4
			if !g[row][col] && liveN == 3 {
				newGrid[row][col] = true
			}
		}
	}
	Copy(g, newGrid)
}
