package math

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestMathSum(t *testing.T) { TestingT(t) }

type MathSumSuite struct{}

var _ = Suite(&MathSumSuite{})

func (s *MathSumSuite) TestMathSum(c *C) {
	specs := []struct {
		x        int
		y        int
		expected int // expected result
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
		{5, 5, 10},
	}

	for _, spec := range specs {
		result := Sum(spec.x, spec.y)
		c.Assert(result, Equals, spec.expected)
	}
}
