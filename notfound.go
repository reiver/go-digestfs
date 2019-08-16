package digestfs

// ContentNotFound may be returned by MountPoint.Open(), and MountPoint.OpenLocation().
//
// Example
//
// Here is an example usage:
//
//	var mountpoint digestfs.MountPoint
//	
//	// ...
//	
//	content, err := mountpoint.Open("SHA-256", "\x00\x00\x00\x00\x00\x19\xd6\x68\x9c\x08\x5a\xe1\x65\x83\x1e\x93\x4f\xf7\x63\xae\x46\xa2\xa6\xc1\x72\xb3\xf1\xb6\x0a\x8c\xe2\x6f")
//	
//	if nil != err {
//		switch err.(type) {
//		case digestfs.ContentNotFound:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type ContentNotFound interface {
	error
	ContentNotFound() (algorithm string, digest string)
}

// MounterNotFound may be returned by Mounter.Mount().
//
// Example
//
// Here is an example usage:
//
//	var mountpoint digestfs.MountPoint
//	
//	err := mountpoint.Mount("git", "/home/joeblow/workspaces/myproject")
//	
//	if nil != err {
//		switch err.(type) {
//		case digestfs.MounterNotFound:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type MounterNotFound interface {
	error
	MounterNotFound() string
}
