package digestfs

import (
	"github.com/reiver/go-digestfs/driver"

	"sync"
)

type MountPoint struct {
	mutex sync.RWMutex
	mountpoint digestfs_driver.MountPoint
}

func (dest *MountPoint) Mount(fstype string, params ...interface{}) error {
	if nil == dest {
		return errNilDestination
	}

	dest.mutex.Lock()
	defer dest.mutex.Unlock()

	if nil != dest.mountpoint {
		return errMounted
	}

	mounter, err := digestfs_driver.Registry.Fetch(fstype)
	if nil != err {
		return err
	}
	if nil == mounter {
		return errNilMounter
	}

	mountpoint, err := mounter.Mount(params...)
	if nil != err {
		return err
	}
	if nil == mountpoint {
		return errNilMountPoint
	}

	dest.mountpoint = mountpoint

	return nil
}

func (receiver *MountPoint) Open(algorithm string, digest string) (Content, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	mountpoint := receiver.mountpoint
	if nil == mountpoint {
		return nil, errNilMountPoint
	}

	return mountpoint.Open(algorithm, digest)
}


func (receiver *MountPoint) OpenLocation(location string) (Content, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	mountpoint := receiver.mountpoint
	if nil == mountpoint {
		return nil, errNilMountPoint
	}

	return mountpoint.OpenLocation(location)
}

func (dest *MountPoint) Unmount() error {
	if nil == dest {
		return errNilDestination
	}

	dest.mutex.Lock()
	defer dest.mutex.Unlock()

	if nil != dest.mountpoint {
		return nil
	}

	if err := dest.mountpoint.Unmount(); nil != err {
		return err
	}

	dest.mountpoint = nil

	return nil
}
