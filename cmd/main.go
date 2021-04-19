package main

import (
	"fmt"
	"os"

	"github.com/avelinoschz/conways-game-of-life/game"
	"github.com/avelinoschz/conways-game-of-life/pattern"
	"github.com/avelinoschz/conways-game-of-life/window"
)

// Default window values
const (
	TITLE         = "Game of Life"
	SCREEN_WIDTH  = 900
	SCREEN_HEIGHT = 900
	FPS           = 1
)

// Default game grid values
const (
	GAME_GRID_WIDTH  = 30
	GAME_GRID_HEIGHT = 30
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Initializing Game Of Life")

	window, err := window.New(SCREEN_WIDTH, SCREEN_HEIGHT, FPS, TITLE)
	if err != nil {
		return err
	}
	window.Setup()

	game, err := game.New(GAME_GRID_WIDTH, GAME_GRID_HEIGHT)
	if err != nil {
		return err
	}

	randomPattern, err := pattern.Random(game.Height, game.Width)
	if err != nil {
		return err
	}
	game.Populate(randomPattern)

	window.Init(int32(game.Width), int32(game.Height))
	for !window.Refresh() {
		window.Draw(game.CellsGrid)
		game.Update()
	}
	window.Close()

	return nil
}
