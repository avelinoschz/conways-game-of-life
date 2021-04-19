package main

import (
	"fmt"
	"os"

	"github.com/avelinoschz/conways-game-of-life/game"
	"github.com/avelinoschz/conways-game-of-life/pattern"
	"github.com/avelinoschz/conways-game-of-life/window"
)

const TITLE = "Game of Life"
const SCREEN_WIDTH = 900
const SCREEN_HEIGHT = 900
const FPS = 1

const BOARD_WIDTH = 30
const BOARD_HEIGHT = 30

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

	game, err := game.New(BOARD_WIDTH, BOARD_HEIGHT)
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
