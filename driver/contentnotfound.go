package digestfs_driver

import (
	"fmt"
)

type ContentNotFound interface {
	error
	ContentNotFound() (algorithm string, digest string)
}

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
