//go:build all_tests || pattern_tests
// +build all_tests pattern_tests

package patterns

// func TestMediator(t *testing.T) {
// 	bufferOutputWriter := outputWriter
// 	outputWriter = new(bytes.Buffer)
// 	defer func() { outputWriter = bufferOutputWriter }()

// 	mediator := NewMediator()
// 	mediator.Ted.Talk()

// 	assert.Equal(t, "Ted: Bill?\n"+
// 		"Bill: What?\n"+
// 		"Ted: Strange things are afoot at the Circle K.\n", outputWriter.(*bytes.Buffer).String())
// }
