package digestfs_driver

import (
	"fmt"
)

type Found interface {
	error
	Found() string
}

func errFound(name string) error {
	var e Found = &internalFound{
		name: name,
	}

	return e
}

type internalFound struct {
	name string
}

func (receiver internalFound) Error() string {
	return fmt.Sprintf("digestfs: Found: mounter name=%q", receiver.name)
}

func (receiver internalFound) Found() string {
	return receiver.name
}
