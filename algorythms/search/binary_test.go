package search

import (
	"testing"
)

func BenchmarkBinary(b *testing.B) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	for i := 0; i < b.N; i++ {
		Binary(items, 63)
	}
}

type binaryTest struct {
	item     int
	expected bool
}

var binaryTests = []binaryTest{
	{
		item:     63,
		expected: true,
	},
	{
		item:     200,
		expected: false,
	},
}

func TestBinary(t *testing.T) {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	for idx, spec := range binaryTests {
		result := Binary(items, spec.item)
		if result != spec.expected {
			t.Errorf("[spec %d] expected to get %v; got %v", idx, spec.expected, result)
		}
	}
}
