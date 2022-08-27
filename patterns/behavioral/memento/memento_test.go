//go:build all_tests || pattern_tests

package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	t.Parallel()

	n := NewNumber(10)
	n.Dubble()
	assert.Equal(t, 20, n.Value())

	memento := n.CreateMemento()

	n.Half()
	assert.Equal(t, 10, n.Value())

	n.ReinstateMemento(memento)
	assert.Equal(t, 20, n.Value())
}
