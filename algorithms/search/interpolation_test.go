// +build all_tests algorythm_tests

package search

import (
	"testing"
)

type interpolationTest struct {
	item     int
	expected int
}

var interpolationTests = []interpolationTest{
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

func TestInterpolation(t *testing.T) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	for idx, spec := range interpolationTests {
		result := Interpolation(items, spec.item)

		if result != spec.expected {
			t.Errorf("[spec %d] expected to get %d; got %d", idx, spec.expected, result)
		}
	}
}
