package digestfs_driver

import (
	"fmt"
)

type NotFound interface {
	error
	NotFound() string
}

func errNotFound(name string) error {
	var e NotFound = &internalNotFound{
		name: name,
	}

	return e
}

type internalNotFound struct {
	name string
}

func (receiver internalNotFound) Error() string {
	return fmt.Sprintf("digestfs: Not Found: mounter name=%q", receiver.name)
}

func (receiver internalNotFound) NotFound() string {
	return receiver.name
}
