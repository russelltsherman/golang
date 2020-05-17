// +build all_tests algorythm_tests

package math

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestMathFib(t *testing.T) { TestingT(t) }

type MathFibSuite struct{}

var _ = Suite(&MathFibSuite{})

func (s *MathFibSuite) TestMathFib(c *C) {
	specs := []struct {
		item     int // input
		expected int // expected result
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, spec := range specs {
		result := Fib(spec.item)
		c.Assert(result, Equals, spec.expected)
	}
}
