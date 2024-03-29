package digestfs_driver

import (
	"fmt"
)

// BadLocation is an error that should be returned by a MountPoint's OpenLocation() method when the
// value of the ‘location’ parameter passed to OpenLocation(location string) is bad.
//
// (You the programmer will decide when the value of the ‘location’ parameter passed to
// OpenLocation(location string) is bad.)
//
// You can create this kind of error by calling the ErrBadLocation() function.
// For example:
//
//	func (receiver MyMountPoint) OpenLocation(location string) (digestfs_driver.Content, error) {
//
//		// ...
//
//		if locationIsBad {
//			nil, return digestfs_driver.ErrBadLocation(location)
//		}
//
//		// ...
//	}
//
// Note that this is a different error that ContentNotFound.
//
// The BadLocation Error means that the format of the value of ‘location’ is bad.
// Ex: it is not understandable by your code.
//
// So, for example, if you decided that your driver will accept locations in this kind of format:
//
//	"@->/path/on/the/file/system/to/the/content"
//	"@->/home/jowblow/photos/1996/vacation-vancouver/.ii/content/00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf"
//
// Then if OpenLocation() received something that didn't begin with a "@->" then it would return
// a BadLocation error.
//
// (You the programmer will decide the format that the value of the ‘location’ parameter passed to
// OpenLocation(location string) will be in.)
//
// If the format of the value of ‘location’ was good, but it just wasn't there, then you would
// return a ContentNotFound error (using ErrContentNotFound()) instead.
type BadLocation interface {
	error
	BadLocation() string
}

// ErrBadLocation returns a BadLocation error.
//
// Here is how one might use this function:
//
//	func (receiver MyMountPoint) OpenLocation(location string) (digestfs_driver.Content, error) {
//
//		// ...
//
//		if locationIsBad {
//			nil, return digestfs_driver.ErrBadLocation(location)
//		}
//
//		// ...
//	}
func ErrBadLocation(location string) error {
	var e BadLocation = &internalBadLocation{
		location: location,
	}

	return e
}

type internalBadLocation struct {
	location string
}

func (receiver internalBadLocation) Error() string {
	return fmt.Sprintf("digestfs: Bad Location: %q", receiver.location)
}

func (receiver internalBadLocation) BadLocation() string {
	return receiver.location
}
