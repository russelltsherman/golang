// +build all_tests algorythm_tests

package search

import (
	"testing"
)

type linearTest struct {
	item     int
	expected bool
}

var linearTests = []linearTest{
	{
		item:     63,
		expected: true,
	},
	{
		item:     200,
		expected: false,
	},
}

func TestLinear(t *testing.T) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	for idx, spec := range linearTests {
		result := Linear(items, spec.item)

		if result != spec.expected {
			t.Errorf("[spec %d] expected to get %v; got %v", idx, spec.expected, result)
		}
	}
}
