package digestfs

import (
	"io"
)

// Content provides a read-only interface to content.
//
//
type Content interface {
	io.Closer
	io.ReaderAt

	Len() int
}
