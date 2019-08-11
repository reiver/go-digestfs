package digestfs

// BadLocation may be returned by MountPoint.OpenLocation().
//
// Example
//
// Here is an example usage:
//
//	var mountpoint digestfs.MountPoint
//	
//	// ...
//	
//	content, err := mountpoint.OpenLoction("@^/apple/banana/cherry")
//	
//	if nil != err {
//		switch err.(type) {
//		case digestfs.BadLocation:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type BadLocation interface {
	error
	BadLocation() string
}
