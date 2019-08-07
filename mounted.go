package digestfs

import (
	"github.com/reiver/go-digestfs/driver"

	"sync"
)

type Mounted struct {
	mutex sync.RWMutex
	opener digestfs_driver.Opener
}

func Mount(dest *Mounted, fstype string, src string, params ...interface{}) error {
	if nil == dest {
		return errNilDestination
	}

	dest.mutex.Lock()
	defer dest.mutex.Unlock()

	if nil != dest.opener {
		return errMounted
	}

	mounter, err := digestfs_driver.Registry.Fetch(fstype)
	if nil != err {
		return err
	}
	if nil == mounter {
		return errNilMounter
	}

	opener, err := mounter.Mount(src, params...)
	if nil != err {
		return err
	}
	if nil == opener {
		return errNilOpener
	}

	dest.opener = opener

	return nil
}

func (receiver *Mounted) Open(algorithm string, digest string) (Content, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	opener := receiver.opener

	if nil == opener {
		return nil, errNilOpener
	}

	return opener.Open(algorithm, digest)
}
