package window

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ErrScreenWidthMissing thrown when missing screen width is missing.
var ErrScreenWidthRequired = errors.New("screen width is required")

// ErrScreenHeigthRequired thrown when missing screen heigth is missing.
var ErrScreenHeigthRequired = errors.New("screen height is required")

// ErrFPSRequired thrown when missing screen fps is missing.
var ErrFPSRequired = errors.New("screen height is required")

const defaultTitle = "Game of Life"

// Raylib struct is an actual implementation of the Window interface.
type Raylib struct {
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

// New returns a newly created Raylib instance..
func New(width, height, fps int32, title string) (Window, error) {
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

	return &Raylib{
		Width:  width,
		Height: height,
		FPS:    fps,
		Title:  _title,
	}, nil
}

// Setup prepares the configuration for the window.
func (r *Raylib) Setup() {
	rl.InitWindow(r.Width, r.Height, r.Title)
	rl.SetTargetFPS(r.FPS)
}

// Init loads the screen pixels based on the board to be used for next drawings.
func (r *Raylib) Init(boardWidth, boardHeight int32) {
	r.Canvas = rl.LoadRenderTexture(boardWidth, boardHeight)
}

// Draw receives a set of pixels and draws them in the window.
func (r *Raylib) Draw(pixels [][]bool) {
	rl.BeginDrawing()
	rl.BeginTextureMode(r.Canvas)
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
		r.Canvas.Texture,
		rl.NewRectangle(0, 0, float32(r.Canvas.Texture.Width), float32(r.Canvas.Texture.Height)),
		rl.NewRectangle(0, 0, float32(r.Width), float32(r.Height)),
		rl.NewVector2(float32(0), float32(0)),
		0,
		rl.RayWhite,
	)

	rl.EndDrawing()
}

// Refresh creates an infinite loop of refreshment on the screen, that can only be exited
// by pressing ESC key or the application got the signal of termination.
func (r *Raylib) Refresh() bool {
	return rl.WindowShouldClose()
}

// Finish releases resources from window and closes it.
func (r *Raylib) Close() {
	rl.UnloadTexture(r.Canvas.Texture)
	rl.CloseWindow()
}
