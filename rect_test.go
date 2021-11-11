package draw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRectFromSize(t *testing.T) {
	rect := NewRectFromSize(ParsePontSize(0, 0, 100, 200))
	assert.Equal(t, 0, rect.Left(), "left not equal")
	assert.Equal(t, 100, rect.Width, "width not equal")
}
