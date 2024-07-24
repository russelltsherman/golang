package patterns

import (
	"io"
	"os"
)

var OutputWriter io.Writer = os.Stdout // modified during testing
