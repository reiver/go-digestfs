package digestfs_driver

import (
	"errors"
)

var (
	errClosed         = errors.New("digestfs: Closed")
	errNegativeOffset = errors.New("digestfs: Negative Offset")
	errNilMounter     = errors.New("digestfs: Nil Mounter")
	errNilReceiver    = errors.New("digestfs: Nil Receiver")
)
