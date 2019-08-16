package digestfs_driver

import (
	"fmt"
)

// ContentNotFound is an error that should be returned by a MountPoint's Open(), or OpenLocation() method
// when the content specified by values of the ‘algorithm’, and ‘digest’ parameters passed to
// Open(algorithm string, digest string) is not found.
//
// You can create this kind of error by calling the ErrContentNotFound() function.
// For example:
//
//	func (receiver MyMountPoint) Open(algorithm string, digest string) (digestfs_driver.Content, error) {
//
//		// ...
//
//		if notfound {
//			nil, return digestfs_driver.ErrContentNotFound(algorithm, digest)
//		}
//
//		// ...
//	}
//
// Note that this is a different error that BadLocation.
type ContentNotFound interface {
	error
	ContentNotFound() (algorithm string, digest string)
}

// ErrContentNotFound returns a ContentNotFound error.
//
// Here is how one might use this function:
//
//	func (receiver MyMountPoint) Open(algorithm string, digest string) (digestfs_driver.Content, error) {
//
//		// ...
//
//		if notfound {
//			nil, return digestfs_driver.ErrContentNotFound(algorithm, digest)
//		}
//
//		// ...
//	}
func ErrContentNotFound(algorithm string, digest string) error {
	var e ContentNotFound = &internalContentNotFound{
		algorithm: algorithm,
		digest: digest,
	}

	return e
}

type internalContentNotFound struct {
	algorithm string
	digest string
}

func (receiver internalContentNotFound) Error() string {
	return fmt.Sprintf("digestfs: Content Not Found: algorithm=%q digest=%q", receiver.algorithm, receiver.digest)
}

func (receiver internalContentNotFound) ContentNotFound() (string, string) {
	return receiver.algorithm, receiver.digest
}
