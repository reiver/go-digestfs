package digestfs_driver

import (
	"fmt"
)

type MounterNotFound interface {
	error
	MounterNotFound() string
}

func errMounterNotFound(name string) error {
	var e MounterNotFound = &internalMounterNotFound{
		name: name,
	}

	return e
}

type internalMounterNotFound struct {
	name string
}

func (receiver internalMounterNotFound) Error() string {
	return fmt.Sprintf("digestfs: Mounter Not Found: mounter name=%q", receiver.name)
}

func (receiver internalMounterNotFound) MounterNotFound() string {
	return receiver.name
}
