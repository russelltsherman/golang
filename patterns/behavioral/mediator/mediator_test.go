package patterns

// import (
// 	"bytes"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestMediator(t *testing.T) {
// 	outputWriter := new(bytes.Buffer)
// 	bufferOutputWriter := outputWriter
// 	defer func() { outputWriter = bufferOutputWriter }()

// 	mediator := NewMediator()
// 	mediator.Ted.Talk()

// 	assert.Equal(t, "Ted: Bill?\n"+
// 		"Bill: What?\n"+
// 		"Ted: Strange things are afoot at the Circle K.\n", outputWriter.(*bytes.Buffer).String())
// }
