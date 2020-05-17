// +build all_tests algorythm_tests

package sort

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestBubbleSort(t *testing.T) { TestingT(t) }

type BubbleSortSuite struct{}

var _ = Suite(&BubbleSortSuite{})

func (s *BubbleSortSuite) TestBubbleSort(c *C) {
	specs := []struct {
		items    []int
		expected []int
	}{
		{
			items:    []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, spec := range specs {
		Bubble(spec.items)
		result := spec.items
		c.Assert(result, DeepEquals, spec.expected)
	}
}
