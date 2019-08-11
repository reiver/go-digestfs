package digestfs

import (
	"io"
)

// Content provides a read-only interface to content.
//
// Example
//
// Here is an example of digestfs.Content being used.
//
//	content, err := Open(algorithm, digest)
//	if nil != err {
//		return err
//	}
//	defer content.Close()
//	
//	// ...
//	
//	var buffer [128]byte
//	p := buffer[:]
//	
//	n, err := content.ReadAt(p, offset)
//
// Reader, Seeker
//
// If you want to turn Content into a io.Reader, and io.Seeker, you can
// do that by wrapping it in a io.SectionReader.
//
// For example:
//
//	r := io.NewSectionReader(content, 0, int64(content.Len()))
type Content interface {
	io.Closer
	io.ReaderAt

	Len() int
}
