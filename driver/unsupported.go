package digestfs_driver

import (
	"fmt"
)

type UnsupportedAlgorithm interface {
	error
	UnsupportedAlgorithm() string
}

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
