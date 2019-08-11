package digestfs_driver

import (
	"fmt"
)

type BadLocation interface {
	error
	BadLocation() string
}

func ErrBadLocation(location string) error {
	var e BadLocation = &internalBadLocation{
		location: location,
	}

	return e
}

type internalBadLocation struct {
	location string
}

func (receiver internalBadLocation) Error() string {
	return fmt.Sprintf("digestfs: Bad Location: %q", receiver.location)
}

func (receiver internalBadLocation) BadLocation() string {
	return receiver.location
}
