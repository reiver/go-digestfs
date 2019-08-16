/*
Package digestfs provides a content-addressable virtual file system (VFS)
by providing a common interface to one or more content-addressable storage (CAS).

Example

Here is an example of how to use an already mounted digestfs.MountPoint to get content:

	content, err := mountpoint.Open("SHA-256", "\xc0\x53\x5e\x4b\xe2\xb7\x9f\xfd\x93\x29\x13\x05\x43\x6b\xf8\x89\x31\x4e\x4a\x3f\xae\xc0\x5e\xcf\xfc\xbb\x7d\xf3\x1a\xd9\xe5\x1a")
	if nil != err {
		return err
	}
	defer content.Close()

Note that if you had the digest encoded as hexadecimal, then you could decode it into binary using:
As in:

	var hexadecimalDigest = "c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a"

	digestBytes, err := hex.DecodeString(hexadecimalDigest)

	var digest string = string(digestBytes)

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

Some content-addressable storages (CAS) that can be used with digestfs include:

• https://godoc.org/github.com/reiver/go-memdigest
*/
package digestfs
