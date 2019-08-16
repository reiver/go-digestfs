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
//
// Also note that the value of ‘digest’ is in binary.
//
// The value of ‘digest’ is NOT a binary-to-text encoding such as hexadecimal, bas64, etc.
//
// So the value of ‘digest’ might be something such as:
//
//	"\x00\x00\x00\x00\x00\x19\xd6\x68\x9c\x08\x5a\xe1\x65\x83\x1e\x93\x4f\xf7\x63\xae\x46\xa2\xa6\xc1\x72\xb3\xf1\xb6\x0a\x8c\xe2\x6f"
//
// Rather than any of these:
//
//	"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"0x000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"AAAAAAAZ1micCFrhZYMek0_3Y65GoqbBcrPxtgqM4m8" (not this)
//
//	"00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
//
//	"0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
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
//
// Note that the value of ‘digest’ is in binary.
//
// The value of ‘digest’ is NOT a binary-to-text encoding such as hexadecimal, bas64, etc.
//
// So the value of ‘digest’ might be something such as:
//
//	"\x00\x00\x00\x00\x00\x19\xd6\x68\x9c\x08\x5a\xe1\x65\x83\x1e\x93\x4f\xf7\x63\xae\x46\xa2\xa6\xc1\x72\xb3\xf1\xb6\x0a\x8c\xe2\x6f"
//
// Rather than any of these:
//
//	"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"0x000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"AAAAAAAZ1micCFrhZYMek0_3Y65GoqbBcrPxtgqM4m8" (not this)
//
//	"00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
//
//	"0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
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
