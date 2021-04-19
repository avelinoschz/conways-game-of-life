// game package contains the main game logic regarding
// Conway's Game of Life.

package game

import (
	"errors"

	"github.com/avelinoschz/conways-game-of-life/pattern"
)

// ErrCellsGridWidthRequired thrown when missing game grid width is missing.
var ErrCellsGridWidthRequired = errors.New("cells grid width is required")

// ErrCellsGridWidthRequired thrown when missing game grid height is missing.
var ErrCellsGridHeigthRequired = errors.New("cells grid heigth is required")

// ErrPopulateNil thrown when trying to populate with a nil pattern.
var ErrPopulateNil = errors.New("trying to populate with nil")

type Game struct {
	// Width represents the actual cells matrix width.
	Width int

	// Height represents the actual cells matrix height.
	Height int

	// CellsGrid represents the main grid of cells life statuses.
	CellsGrid [][]bool
}

// New returns a newly created Game instance.
func New(height, width int) (*Game, error) {
	if width <= 0 {
		return nil, ErrCellsGridWidthRequired
	}

	if height <= 0 {
		return nil, ErrCellsGridHeigthRequired
	}

	return &Game{
		Width:     width,
		Height:    height,
		CellsGrid: pattern.New(height, width),
	}, nil
}

// Populate fills the cells grid with the received pattern.
// In case the cells grid already exists, it got replaced entirely.
func (g *Game) Populate(pattern [][]bool) error {
	if pattern == nil {
		return ErrPopulateNil
	}
	g.CellsGrid = pattern
	return nil
}

// Update is the main function of the game's engine.
// It iterate through all the cells and verifies the Conway's Game Of Life rule set.
// Updates the whole cells grid to the next generationerrn of cells statuses.
func (g *Game) Update() {
	newCellsGeneration := pattern.New(g.Height, g.Width)

	for y, rows := range g.CellsGrid {
		for x := range rows {
			cellState := g.CellsGrid[y][x]
			liveNeighbors := g.countLiveNeighbors(y, x)
			if stayAlive(cellState, liveNeighbors) {
				newCellsGeneration[y][x] = true
			}
		}
	}

	g.CellsGrid = newCellsGeneration
}

// countNeighbors counts all the live neighbors of a current cell.
// The 8 adjacent positions to the cell are considered the neighbors.
// Returns the number of neighbors that are alive.
func (g *Game) countLiveNeighbors(y, x int) int {
	count := 0

	initialrow := -1
	initialcol := -1

	limitrow := 1
	limitcol := 1

	if y == 0 {
		initialrow = 0
	}

	if x == 0 {
		initialcol = 0
	}

	if y == g.Height-1 {
		limitrow = 0
	}

	if x == g.Width-1 {
		limitcol = 0
	}

	for row := initialrow; row <= limitrow; row++ {
		for col := initialcol; col <= limitcol; col++ {
			if g.CellsGrid[y+row][x+col] {
				count++
			}
		}
	}

	if g.CellsGrid[y][x] {
		count--
	}
	return count
}

// stayAlive contains validates if a cell will live, keep alive, die or stay dead in the next generation.
// Receives the current state of a cell and all the live neighbors to determinate the
// cell's next generation state.
// Conway's Game Of Life ruleset work as it follows:
// 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
// 2. Any live cell with two or three live neighbours lives on to the next generation.
// 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
// 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func stayAlive(currentCellState bool, liveNeighbors int) bool {

	var nextState bool

	switch {
	case !currentCellState && liveNeighbors == 3:
		nextState = true

	case currentCellState && (liveNeighbors < 2 || liveNeighbors > 3):
		nextState = false

	case currentCellState && (liveNeighbors == 2 || liveNeighbors == 3):
		nextState = true
	}

	return nextState
}
