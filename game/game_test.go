package game

import (
	"testing"

	"github.com/avelinoschz/conways-game-of-life/pattern"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return correctly a new game instance When New,5,5)", func(t *testing.T) {
		expectedGame := &Game{
			Width:     5,
			Height:    5,
			CellsGrid: pattern.New(5, 5),
		}

		actualGame, err := New(5, 5)

		assert.NoError(err)
		assert.Equal(expectedGame, actualGame)
	})

	t.Run("Should return `ErrScreenWidthRequired` When width is 0", func(t *testing.T) {
		_, err := New(5, 0)

		assert.EqualError(err, ErrCellsGridWidthRequired.Error())
	})

	t.Run("Should return `ErrScreenHeigthRequired` When heigth is 0", func(t *testing.T) {
		_, err := New(0, 5)

		assert.EqualError(err, ErrCellsGridHeigthRequired.Error())
	})
}

func TestPopulate(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return error `nil` When received [][]bool", func(t *testing.T) {
		g := &Game{}
		pattern := [][]bool{}

		err := g.Populate(pattern)

		assert.NoError(err)
	})

	t.Run("Should return `ErrPopulateNil` When received nil", func(t *testing.T) {
		g := &Game{}

		err := g.Populate(nil)

		assert.Error(err)
		assert.EqualError(err, ErrPopulateNil.Error())
	})
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return vertical blinker when horizon blinker", func(t *testing.T) {
		g := &Game{
			Width:  5,
			Height: 5,
		}
		initialPattern := [][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, true, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}
		g.Populate(initialPattern)

		g.Update()

		expectedPattern := [][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		}
		assert.Equal(expectedPattern, g.CellsGrid)
	})
}

func TestCountLiveNeighbors(t *testing.T) {
	assert := assert.New(t)

	g := &Game{
		Width:  5,
		Height: 5,
	}
	pattern := [][]bool{
		{false, true, false, true, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, true, false, true, false},
	}
	g.Populate(pattern)

	t.Run("Should return `1` When y is 2 & x is 1", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(2, 1)

		assert.Equal(1, actualCount)
	})

	t.Run("Should return `2` When y is 2 & x is 2", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(2, 2)

		assert.Equal(2, actualCount)
	})

	t.Run("Should return `5` When y is 1 & x is 2", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(1, 2)

		assert.Equal(5, actualCount)
	})

	t.Run("Should return `1` When y is 0 & x is 0", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(0, 0)

		assert.Equal(1, actualCount)
	})

	t.Run("Should return `1` When y is 4 & x is 4", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(4, 4)

		assert.Equal(1, actualCount)
	})

	t.Run("Should return `1` When y is 0 & x is 4", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(0, 4)

		assert.Equal(1, actualCount)
	})

	t.Run("Should return `1` When y is 4 & x is 0", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(4, 0)

		assert.Equal(1, actualCount)
	})

	t.Run("Should return `1` When y is 2 & x is 0", func(t *testing.T) {
		actualCount := g.countLiveNeighbors(2, 0)

		assert.Equal(1, actualCount)
	})
}

func TestStayAlive(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return `false` When state is `false` & got less than 2 live neighbors", func(t *testing.T) {
		actualState := stayAlive(false, 1)

		assert.False(actualState)
	})

	t.Run("Should return `true` When state is `false` & got 3 live neighbors", func(t *testing.T) {
		actualState := stayAlive(false, 3)

		assert.True(actualState)
	})

	t.Run("Should return `false` When state is `true` & got more than 3 live neighbors", func(t *testing.T) {
		actualState := stayAlive(true, 5)

		assert.False(actualState)
	})

	t.Run("Should return `false` When state is `true` & got less than 2 live neighbors", func(t *testing.T) {
		actualState := stayAlive(true, 1)

		assert.False(actualState)
	})

	t.Run("Should return `true` When state is `true` & got 2 live neighbors", func(t *testing.T) {
		actualState := stayAlive(true, 2)

		assert.True(actualState)
	})

	t.Run("Should return `true` When state is `true` & got 3 live neighbors", func(t *testing.T) {
		actualState := stayAlive(true, 3)

		assert.True(actualState)
	})

}
