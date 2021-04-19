// window package serves as a wrapper to the raylib library.
// raylib is an abstraction of OpenGL to handle graphics control.
// In this way window package abstracts the interaction directly
// to hardware in a simple manner and usage.

package window

// Window interface wraps the common functionality to manage a window on the screen.
// The purpose of this interface is to be able to change the current implementation.
// So that, if the actual library gets deprecated, the layer where the library is used
// is the only one that needs to change.
type Window interface {
	Setup()
	Init(boardWidth, boardHeight int32)
	Draw(pixels [][]bool)
	Refresh() bool
	Close()
}
