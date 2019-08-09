package digestfs_driver

import (
	"sync"
)

type internalRegistrar struct {
	mutex sync.RWMutex
	mounters map[string]Mounter
}

func (receiver *internalRegistrar) Fetch(name string) (Mounter, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	mounters := receiver.mounters
	if nil == mounters {
		return nil, errMounterNotFound(name)
	}

	mounter, found := mounters[name]
	if !found {
		return nil, errMounterNotFound(name)
	}
	if nil == mounter {
		return nil, errNilMounter
	}

	return mounter, nil
}

func (receiver *internalRegistrar) Register(mounter Mounter, name string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == mounter {
		return errNilMounter
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.mounters {
		receiver.mounters = map[string]Mounter{}
	}

	_, found := receiver.mounters[name]
	if found {
		return errMounterFound(name)
	}

	receiver.mounters[name]  = mounter

	return nil
}
