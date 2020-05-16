package sort

import (
	"reflect"
	"testing"
)

type bubbleTest struct {
	items    []int
	expected []int
}

var bubbleTests = []bubbleTest{
	{
		items:    []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0},
		expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	},
}

func TestBubble(t *testing.T) {
	for idx, spec := range bubbleTests {
		Bubble(spec.items)
		result := spec.items

		if !reflect.DeepEqual(result, spec.expected) {
			t.Errorf("[spec %d] expected to get %v; got %v", idx, spec.expected, result)
		}
	}
}
