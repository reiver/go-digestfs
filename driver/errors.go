package digestfs_driver

import (
	"errors"
)

var (
	errNilMounter  = errors.New("digestfs: Nil Mounter")
	errNilReceiver = errors.New("digestfs: Nil Receiver")
)
