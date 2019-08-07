package digestfs

import (
	"errors"
)

var (
	errMounted        = errors.New("digestfs: Mounted")
	errNilDestination = errors.New("digestfs: Nil Destination")
	errNilMounter     = errors.New("digestfs: Nil Mounter")
	errNilOpener      = errors.New("digestfs: Nil Opener")
	errNilReceiver    = errors.New("digestfs: Nil Receiver")
)
