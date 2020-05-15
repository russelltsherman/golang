package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinary(t *testing.T) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	assert.Equal(t, true, Binary(items, 63), "the item should be found")
}
