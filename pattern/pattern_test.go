package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return a not all false 5 x 5 matrix When rows equals 5 & cols equals 5 ", func(t *testing.T) {
		zeroPattern := [][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		}

		randomPattern, err := Random(5, 5)

		assert.NoError(err)
		assert.NotEmpty(randomPattern)
		assert.Equal(5, len(randomPattern))
		assert.NotEqual(zeroPattern, randomPattern)
	})

}

func TestCreateEmptyMatrix(t *testing.T) {
	assert := assert.New(t)

	t.Run("Should return a 5 x 5 matrix When rows equals 5 & cols equals 5 ", func(t *testing.T) {
		actualMatrix := New(5, 5)

		assert.NotEmpty(actualMatrix)
		assert.Equal(5, len(actualMatrix))
	})
}
