package main

import (
	"strings"
)

var deadRune = '.'
var aliveRune = '*'

type World struct {
	dy, dx int
	grid   [][]bool
}

func NewWorld(dy, dx int) World {
	grid := makeGrid(dy, dx)
	return World{dy, dx, grid}
}

func (w World) Awaken(y, x int) {
	if w.isWithinBounds(y, x) {
		w.grid[y][x] = true
	}
}

func (w *World) Advance() {
	nextGen := makeGrid(w.dy, w.dx)
	for y := range nextGen {
		for x := range nextGen[y] {
			nextGen[y][x] = w.nextGen(y, x)
		}
	}
	w.grid = nextGen
}

func makeGrid(dy, dx int) [][]bool {
	grid := make([][]bool, dy)
	for y := range grid {
		grid[y] = make([]bool, dx)
	}
	return grid
}

// 1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
// 2. Any live cell with more than three live neighbours dies, as if by overcrowding.
// 3. Any live cell with two or three live neighbours lives on to the next generation.
// 4. Any dead cell with exactly three live neighbours becomes a live cell.
func (w World) nextGen(y, x int) bool {
	liveNeighbours := w.liveNeighboursOf(y, x)
	isAlive := w.isAlive(y, x)
	if isAlive && (liveNeighbours < 2 || liveNeighbours > 3) {
		return false
	} else if isAlive && (liveNeighbours == 2 || liveNeighbours == 3) {
		return true
	} else if !isAlive && liveNeighbours == 3 {
		return true
	}
	return false
}

func (w World) liveNeighboursOf(y, x int) int {
	liveNeighbours := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if !(i == y && j == x) && w.isAlive(i, j) {
				liveNeighbours++
			}
		}
	}
	return liveNeighbours
}

func (w World) String() string {
	var str strings.Builder
	for _, row := range w.grid {
		str.WriteRune('\n')
		for _, cell := range row {
			str.WriteRune(stateRune(cell))
		}
	}
	return str.String()
}

func (w World) isAlive(y, x int) bool {
	if w.isWithinBounds(y, x) {
		return w.grid[y][x]
	}
	return false
}

func (w World) isWithinBounds(y, x int) bool {
	return y >= 0 && x >= 0 && y < w.dy && x < w.dx
}

func stateRune(isAlive bool) rune {
	if isAlive {
		return aliveRune
	}
	return deadRune
}
