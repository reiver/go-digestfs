package digestfs

import (
	"errors"
)

var (
	errMounted        = errors.New("digestfs: Mounted")
	errNilDestination = errors.New("digestfs: Nil Destination")
	errNilMounter     = errors.New("digestfs: Nil Mounter")
	errNilMountPoint  = errors.New("digestfs: Nil Mount Point")
	errNilReceiver    = errors.New("digestfs: Nil Receiver")
)
