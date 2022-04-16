package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCube(t *testing.T) {
	assert.Equal(t, 27, Cube(3))
}

func TestSquare(t *testing.T) {
	assert.Equal(t, 25, Cube(5))
}
