// +build all_tests algorythm_tests

package search

import (
	"testing"

	. "gopkg.in/check.v1"
)

func BenchmarkBinarySearch(b *testing.B) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	for i := 0; i < b.N; i++ {
		Binary(items, 63)
	}
}

// Hook up gocheck into the "go test" runner.
func TestBinarySearch(t *testing.T) { TestingT(t) }

type BinarySearchSuite struct{}

var _ = Suite(&BinarySearchSuite{})

func (s *BinarySearchSuite) TestBinarySearch(c *C) {
	specs := []struct {
		item     int
		expected bool
	}{
		{
			item:     63,
			expected: true,
		},
		{
			item:     200,
			expected: false,
		},
	}
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	for _, spec := range specs {
		result := Binary(items, spec.item)
		c.Assert(result, Equals, spec.expected)
	}

}
