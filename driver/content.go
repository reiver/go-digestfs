package digestfs_driver

import (
	"io"
)

// Content represents open content.
//
// Content is returned by Open(), and OpenLocation().
type Content interface {
	io.Closer
	io.ReaderAt

	Len() int
}
