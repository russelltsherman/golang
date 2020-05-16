package patterns

import (
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout // modified during testing
