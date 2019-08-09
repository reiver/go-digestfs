/*
Package digestfs provides a content-addressable file system, and virtual file system (VFS)
by providing a common interface to one or more content-addressable storage (CAS).

Example

Here is an example of how to use an already mounted digestfs.MountPoint to get content:

	var mountpoint digestfs.MountPoint
	
	// ...
	
	content, err := mountpoint.Open("SHA-256", "c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a")
*/
package digestfs
