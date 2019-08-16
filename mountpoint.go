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

// Create creates new content at the MoutPoint.
//
// The value of ‘digest’ is in binary.
//
// The value of ‘digest’ is NOT a binary-to-text encoding such as hexadecimal, bas64, etc.
//
// So the value of ‘digest’ might be something such as:
//
//	"\x00\x00\x00\x00\x00\x19\xd6\x68\x9c\x08\x5a\xe1\x65\x83\x1e\x93\x4f\xf7\x63\xae\x46\xa2\xa6\xc1\x72\xb3\xf1\xb6\x0a\x8c\xe2\x6f"
//
// Rather than any of these:
//
//	"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"0x000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"AAAAAAAZ1micCFrhZYMek0_3Y65GoqbBcrPxtgqM4m8" (not this)
//
//	"00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
//
//	"0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
func (receiver *MountPoint) Create(content []byte) (algorithm string, digest string, err error) {
	if nil == receiver {
		return "", "", errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	mountpoint := receiver.mountpoint
	if nil == mountpoint {
		return "", "", errNilMountPoint
	}

	return mountpoint.Create(content)
}

// Open returns content at the MountPoint.
//
// The value of ‘digest’ is in binary.
//
// The value of ‘digest’ is NOT a binary-to-text encoding such as hexadecimal, bas64, etc.
//
// So the value of ‘digest’ might be something such as:
//
//	"\x00\x00\x00\x00\x00\x19\xd6\x68\x9c\x08\x5a\xe1\x65\x83\x1e\x93\x4f\xf7\x63\xae\x46\xa2\xa6\xc1\x72\xb3\xf1\xb6\x0a\x8c\xe2\x6f"
//
// Rather than any of these:
//
//	"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"0x000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
//
//	"AAAAAAAZ1micCFrhZYMek0_3Y65GoqbBcrPxtgqM4m8" (not this)
//
//	"00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
//
//	"0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
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
