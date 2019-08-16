package digestfs_driver

import (
	"fmt"
)

// UnsupportedAlgorithm is an error that should be returned by a MountPoint's Open() method
// when the algorithm specified by value of the ‘algorithm’ is not supported.
//
// You can create this kind of error by calling the ErrUnsupportedAlgorithm() function.
// For example:
//
//	func (receiver MyMountPoint) Open(algorithm string, digest string) (digestfs_driver.Content, error) {
//
//		// ...
//
//		if algorithmNotSupported {
//			nil, return digestfs_driver.ErrUnsupportedAlgorithm(algorithm)
//		}
//
//		// ...
//	}
type UnsupportedAlgorithm interface {
	error
	UnsupportedAlgorithm() string
}

// ErrUnsupportedAlgorithm returns a UnsupportedAlgorithm error.
//
// Here is how one might use this function:
//
//	func (receiver MyMountPoint) Open(algorithm string, digest string) (digestfs_driver.Content, error) {
//
//		// ...
//
//		if algorithmNotSupported {
//			nil, return digestfs_driver.ErrUnsupportedAlgorithm(algorithm)
//		}
//
//		// ...
//	}
func ErrUnsupportedAlgorithm(name string) error {
	var e UnsupportedAlgorithm = &internalUnsupportedAlgorithm{
		name: name,
	}

	return e
}

type internalUnsupportedAlgorithm struct {
	name string
}

func (receiver internalUnsupportedAlgorithm) Error() string {
	return fmt.Sprintf("digestfs: Unsupported Algorithm: %q", receiver.name)
}

func (receiver internalUnsupportedAlgorithm) UnsupportedAlgorithm() string {
	return receiver.name
}
