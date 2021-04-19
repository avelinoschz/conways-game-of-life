// window package serves as a wrapper to the raylib library.
// raylib is an abstraction of OpenGL to handle graphics control.
// In this way window package abstracts the interaction directly
// to hardware in a simple manner and usage.

package window

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ErrScreenWidthMissing
var ErrScreenWidthRequired = errors.New("screen width is required")

// ErrScreenHeigthMissing
var ErrScreenHeigthRequired = errors.New("screen height is required")

// ErrScreenHeigthMissing
var ErrFPSRequired = errors.New("screen height is required")

const defaultTitle = "Game of Life"

type Window struct {
	// Title is the title rendered into the window title bar.
	Title string

	// Width is the actual width the window will have represented in pixels.
	Width int32

	// Height is the actual height the window will have represented in pixels.
	Height int32

	// FPS is the frames per second the window will refresh the screen.
	FPS int32

	// Canvas holds information needed from raylib to process and draw the window.
	Canvas rl.RenderTexture2D
}

// New returns a newly created Window instance.
func New(width, height, fps int32, title string) (*Window, error) {
	_title := defaultTitle

	if width <= 0 {
		return nil, ErrScreenWidthRequired
	}

	if height <= 0 {
		return nil, ErrScreenHeigthRequired
	}

	if fps <= 0 {
		return nil, ErrFPSRequired
	}

	if title != "" {
		_title = title
	}

	return &Window{
		Width:  width,
		Height: height,
		FPS:    fps,
		Title:  _title,
	}, nil
}

// Setup prepares the configuration for the window.
func (w *Window) Setup() {
	rl.InitWindow(w.Width, w.Height, w.Title)
	rl.SetTargetFPS(w.FPS)
}

// Init loads the screen pixels based on the board to be used for next drawings.
func (w *Window) Init(boardWidth, boardHeight int32) {
	w.Canvas = rl.LoadRenderTexture(boardWidth, boardHeight)
}

// Draw receives a set of pixels and draws them in the window.
func (w *Window) Draw(pixels [][]bool) {
	rl.BeginDrawing()
	rl.BeginTextureMode(w.Canvas)
	rl.ClearBackground(rl.RayWhite)

	for y, rows := range pixels {
		for x, col := range rows {
			if col {
				rl.DrawPixel(int32(x), int32(y), rl.NewColor(0, 0, 0, 255))
			}
		}
	}

	rl.EndTextureMode()

	rl.DrawTexturePro(
		w.Canvas.Texture,
		rl.NewRectangle(0, 0, float32(w.Canvas.Texture.Width), float32(w.Canvas.Texture.Height)),
		rl.NewRectangle(0, 0, float32(w.Width), float32(w.Height)),
		rl.NewVector2(float32(0), float32(0)),
		0,
		rl.RayWhite,
	)

	rl.EndDrawing()
}

// Finish releases resources from window and closes it.
func (w *Window) Close() {
	rl.UnloadTexture(w.Canvas.Texture)
	rl.CloseWindow()
}

// Refresh
func (w *Window) Refresh() bool {
	return rl.WindowShouldClose()
}
