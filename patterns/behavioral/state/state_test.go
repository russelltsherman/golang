package patterns

// import (
// 	"bytes"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestState(t *testing.T) {
// 	outputWriter := new(bytes.Buffer)
// 	bufferOutputWriter := outputWriter
// 	defer func() { outputWriter = bufferOutputWriter }()

// 	machine := NewMachine()
// 	machine.Off()
// 	machine.On()
// 	machine.On()
// 	machine.Off()

// 	assert.Equal(t, "Machine is ready.\n"+
// 		"   already OFF\n"+
// 		"   going from OFF to ON\n"+
// 		"   already ON\n"+
// 		"   going from ON to OFF\n", outputWriter.(*bytes.Buffer).String())
// }
