package window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return correctly a new window instance When New,5,5,1,'test')", func(t *testing.T) {
		expectedWindow := &Raylib{
			Title:  "test",
			Width:  int32(5),
			Height: int32(5),
			FPS:    int32(1),
		}

		actualWindow, err := New(int32(5), int32(5), int32(1), "test")

		assert.NoError(err)
		assert.Equal(expectedWindow, actualWindow)
	})

	t.Run("Should assign default title When title is empty", func(t *testing.T) {
		expectedWindow := &Raylib{
			Title:  "Game of Life",
			Width:  int32(5),
			Height: int32(5),
			FPS:    int32(1),
		}

		actualWindow, err := New(int32(5), int32(5), int32(1), "")

		assert.NoError(err)
		assert.Equal(expectedWindow, actualWindow)
	})

	t.Run("Should return `ErrScreenWidthRequired` When width is 0", func(t *testing.T) {
		_, err := New(int32(0), int32(5), int32(1), "test")

		assert.EqualError(err, ErrScreenWidthRequired.Error())
	})

	t.Run("Should return `ErrScreenHeigthRequired` When heigth is 0", func(t *testing.T) {
		_, err := New(int32(5), int32(0), int32(1), "test")

		assert.EqualError(err, ErrScreenHeigthRequired.Error())
	})

	t.Run("Should return `ErrFPSRequired` When fps is 0", func(t *testing.T) {
		_, err := New(int32(5), int32(5), int32(0), "test")

		assert.EqualError(err, ErrFPSRequired.Error())
	})
}
