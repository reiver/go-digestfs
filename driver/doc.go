/*
Package digestfs_driver provides tools for registering content-addressable file system drivers for the virtual file system (VFS) for package digestfs.

Drivers

To create a driver, do something like the following:

	import "github.com/reiver/go-digestfs/driver"
	
	func init() {
		
		var name string
		
		var mounter MyMounter
		
		// ...
		
		if err := digestfs_driver.Registry.Register(mounter, name); nil != err {
			panic(err)
		}
	}
*/
package digestfs_driver
