package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinear(t *testing.T) {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	assert.Equal(t, true, Linear(items, 58), "the item should be found")
}
