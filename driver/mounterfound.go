package digestfs_driver

import (
	"fmt"
)

type MounterFound interface {
	error
	MounterFound() string
}

func errMounterFound(name string) error {
	var e MounterFound = &internalMounterFound{
		name: name,
	}

	return e
}

type internalMounterFound struct {
	name string
}

func (receiver internalMounterFound) Error() string {
	return fmt.Sprintf("digestfs: Mounter Found: mounter name=%q", receiver.name)
}

func (receiver internalMounterFound) MounterFound() string {
	return receiver.name
}
