// +build all_tests algorythm_tests

package search

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestInterpolationSearch(t *testing.T) { TestingT(t) }

type InterpolationSearchSuite struct{}

var _ = Suite(&InterpolationSearchSuite{})

func (s *InterpolationSearchSuite) TestInterpolationSearch(c *C) {
	specs := []struct {
		item     int
		expected int
	}{
		{
			item:     63,
			expected: 6,
		},
		{
			item:     0,
			expected: 0,
		},
		{
			item:     200,
			expected: 9,
		},
	}
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	for _, spec := range specs {
		result := Interpolation(items, spec.item)
		c.Assert(result, Equals, spec.expected)
	}
}
