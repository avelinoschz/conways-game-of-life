package pattern

import (
	"errors"
	"math/rand"
)

// ErrRandomPatternRows
var ErrRandomPatternRows = errors.New("pattern rows is required")

// ErrRandomPatternCols
var ErrRandomPatternCols = errors.New("pattern cols is required")

type pattern [][]bool

// Preset patterns names
const (
	Blinker = "blinker"
)

// Actual preset patterns, identified by name
// This presets are for usage only on 5x5 game grids as samples.
var Patterns = map[string]pattern{
	Blinker: [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	},
}

// Random generates a random values matrix with the given values of rows and cols.
func Random(rowsNum, colsNum int) (pattern, error) {
	if rowsNum == 0 {
		return nil, ErrRandomPatternRows
	}

	if colsNum == 0 {
		return nil, ErrRandomPatternCols
	}

	matrix := New(rowsNum, colsNum)

	for _, row := range matrix {
		for i := range row {
			if rand.Intn(2) == 0 {
				row[i] = true
			} else {
				row[i] = false
			}
		}
	}

	return matrix, nil
}

// New generates new matrix with all false values.
func New(rowsNum, colsNum int) [][]bool {
	matrix := make([][]bool, rowsNum)
	for i := range matrix {
		matrix[i] = make([]bool, colsNum)
	}
	return matrix
}
