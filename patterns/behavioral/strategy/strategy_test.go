// //go:build all_tests || pattern_tests

package patterns

// import (
// 	"bytes"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestNewStrategyA_ReturnsNonNil(t *testing.T) {
// 	t.Parallel()
// 	strategy := NewStrategyA()
// 	assert.NotNil(t, strategy)
// }

// func TestStrategyAExecute_ExecutesStrategyA(t *testing.T) {
// 	outputWriter := new(bytes.Buffer)
// 	bufferOutputWriter := outputWriter
// 	defer func() { outputWriter = bufferOutputWriter }()

// 	strategy := NewStrategyA()
// 	strategy.Execute()

// 	assert.Equal(t, "executing strategy A\n", outputWriter.(*bytes.Buffer).String())
// }

// func TestNewStrategyB_ReturnsNonNil(t *testing.T) {
// 	t.Parallel()
// 	strategy := NewStrategyB()
// 	assert.NotNil(t, strategy)
// }

// func TestStrategyBExecute_ExecutesStrategyB(t *testing.T) {
// 	outputWriter := new(bytes.Buffer)
// 	bufferOutputWriter := outputWriter
// 	defer func() { outputWriter = bufferOutputWriter }()

// 	strategy := NewStrategyB()
// 	strategy.Execute()

// 	assert.Equal(t, "executing strategy B\n", outputWriter.(*bytes.Buffer).String())
// }

// func TestNewContext_ReturnsNonNil(t *testing.T) {
// 	t.Parallel()
// 	context := NewContext()
// 	assert.NotNil(t, context)
// }

// func TestSetStrategy_SetsStrategy(t *testing.T) {
// 	t.Parallel()
// 	strategy := NewStrategyB()
// 	context := NewContext()
// 	context.SetStrategy(strategy)
// 	assert.Equal(t, strategy, context.strategy)
// }

// func TestContextExecute_ExecutesSetStrategy(t *testing.T) {
// 	outputWriter := new(bytes.Buffer)
// 	bufferOutputWriter := outputWriter
// 	defer func() { outputWriter = bufferOutputWriter }()

// 	strategy := NewStrategyB()
// 	context := NewContext()
// 	context.SetStrategy(strategy)
// 	context.Execute()
// 	assert.Equal(t, "executing strategy B\n", outputWriter.(*bytes.Buffer).String())
// }
