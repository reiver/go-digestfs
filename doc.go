/*
Package digestfs provides a content-addressable virtual file system (VFS)
by providing a common interface to one or more content-addressable storage (CAS).

Example

Here is an example of how to use an already mounted digestfs.MountPoint to get content:

	var mem memdigest.SHA1

	binaryDigest, err := mem.Store([]byte("The request has been accepted for processing, but the processing has not been completed."))
	
	binaryDigest, err := mem.Store([]byte("The server has fulfilled the request but does not need to return an entity-body, and might want to return updated metainformation."))
	
	binaryDigest, err := mem.Store([]byte("The server encountered an unexpected condition which prevented it from fulfilling the request."))

	// ...

	var mountpoint digestfs.MountPoint
	
	err := mountpoint.Mount("memdigest.SHA1", &mem)
	
	// ...
	
	content, err := mountpoint.Open("SHA-256", "c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a")
	if nil != err {
		return err
	}
	defer content.Close()

Note that if you had the digest encoded as binary, then you could encode it as hexadecimal using:
As in:

	digest := fmt.Sprintf("%x", binaryDigest)

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
