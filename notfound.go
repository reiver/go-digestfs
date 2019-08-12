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
//	content, err := mountpoint.Open("SHA-256", "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
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
//	err := mountpoint.Mount("git", "/home/me/workspaces/project")
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
