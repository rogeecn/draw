package draw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRectFromSize(t *testing.T) {
	rect := NewRectFromSize(NewPoint(0, 0), NewSize(100, 200))
	assert.Equal(t, 0, rect.Left(), "left not equal")
	assert.Equal(t, 100, rect.Width, "width not equal")
}
