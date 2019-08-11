/*
Package digestfs provides a content-addressable file system, and virtual file system (VFS)
by providing a common interface to one or more content-addressable storage (CAS).

Example

Here is an example of how to use an already mounted digestfs.MountPoint to get content:

	var mountpoint digestfs.MountPoint
	
	// ...
	
	// c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a
	var digest [64]byte = [64]byte{0xc0, 0x53, 0x5e, 0x4b, 0xe2, 0xb7, 0x9f, 0xfd, 0x93, 0x29, 0x13, 0x05, 0x43, 0x6b, 0xf8, 0x89, 0x31, 0x4e, 0x4a, 0x3f, 0xae, 0xc0, 0x5e, 0xcf, 0xfc, 0xbb, 0x7d, 0xf3, 0x1a, 0xd9, 0xe5, 0x1a}
	
	content, err := mountpoint.Open("SHA-256", digest[:])
	if nil != err {
		return err
	}
	defer content.Close()

Note that if you had the digest encoded as hexadecimal, then you could decode it using the Go built-in package "encoding/hex".
As in:

	digest, err := hex.DecodeString("c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a")

Here, ‘content’ would give you access to the data on the content mostly via an io.ReaderAt interface.
For example:

	var buffer [256]byte

	var p []byte = buffer[:]

	var offset int64 = 0
	
	n, err := content.ReadAt(p, offset)

Alternatively, if you want to use the io.Reader interface, you can upgrade ‘content’ from an io.ReaderAt to an io.Reader by using an io.SectionReader.
For example:

	r := io.NewSectionReader(content, 0, int64(content.Len()))
	
	// ...
	
	contentBytes, err := ioutil.RealAll(r)
*/
package digestfs
