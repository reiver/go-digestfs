package digestfs

import (
	"fmt"
)

type NotFound interface {
	error
	NotFound() (algorithm string, digest string)
}

func errNotFound(algorithm string, digest string) error {
	var e NotFound = &internalNotFound{
		algorithm: algorithm,
		digest:    digest,
	}

	return e
}

type internalNotFound struct {
	algorithm string
	digest string
}

func (receiver internalNotFound) Error() string {
	return fmt.Sprintf("digestfs: Not Found: algorithm=%q digest=%q", receiver.algorithm, receiver.digest)
}

func (receiver internalNotFound) NotFound() (algorithm string, digest string) {
	return receiver.algorithm, receiver.digest
}
