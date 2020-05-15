package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpolation(t *testing.T) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	assert.Equal(t, 6, Interpolation(items, 63), "the item should be found")
}
