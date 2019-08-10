package digestfs_driver

import (
	"fmt"
)

type ContentNotFound interface {
	error
	ContentNotFound() (algorithm string, digest []byte)
}

func ErrContentNotFound(algorithm string, digest []byte) error {
	var e ContentNotFound = &internalContentNotFound{
		algorithm: algorithm,
		digest: append([]byte(nil), digest...),
	}

	return e
}

type internalContentNotFound struct {
	algorithm string
	digest []byte
}

func (receiver internalContentNotFound) Error() string {
	return fmt.Sprintf("digestfs: Content Not Found: algorithm=%q digest=%x", receiver.algorithm, receiver.digest)
}

func (receiver internalContentNotFound) ContentNotFound() (string, []byte) {
	return receiver.algorithm, receiver.digest
}
