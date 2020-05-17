// +build all_tests algorythm_tests

package search

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestLinearSearch(t *testing.T) { TestingT(t) }

type LinearSearchSuite struct{}

var _ = Suite(&LinearSearchSuite{})

func (s *LinearSearchSuite) TestLinearSearch(c *C) {
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
		result := Linear(items, spec.item)
		c.Assert(result, Equals, spec.expected)
	}
}
