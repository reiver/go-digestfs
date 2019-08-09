package digestfs

import (
	"fmt"
)

// ContentNotFound may be returned by MountPoint.Open(), and MountPoint.OpenLocation().
type ContentNotFound interface {
	error
	ContentNotFound() (algorithm string, digest string)
}

// MounterNotFound may be returned by Mounter.Mount().
type MounterNotFound interface {
	error
	MounterNotFound() (algorithm string, digest string)
}
