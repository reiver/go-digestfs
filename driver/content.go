package digestfs_driver

import (
	"io"
)

// Content represents open content.
//
// Content is returned by Open(), and OpenLocation().
//
// MountPoints can use the StringContent() helper function create a Content.
//
// Alternatively, MountPoints can create this own implementation of Content.
type Content interface {
	io.Closer
	io.ReaderAt

	Len() int
}
