package digestfs_driver

import (
	"io"
)

type Content interface {
	io.Closer
	io.ReaderAt

	Len() int
}
