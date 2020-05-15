package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	sorted := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	BubbleSort(elements)
	assert.Equal(t, sorted, elements, "the slice should be sorted")
}
