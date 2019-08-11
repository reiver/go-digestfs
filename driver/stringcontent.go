package digestfs_driver

import (
	"io"
)

type internalStringContent struct {
	value string
	closed bool
}

func StringContent(value string) Content {
	return &internalStringContent{
		value: value,
	}
}

func (receiver *internalStringContent) Close() error {
	if nil == receiver {
		return nil
	}

	receiver.closed = true
	return nil
}

func (receiver internalStringContent) Len() int {
	return len(receiver.value)
}

func (receiver internalStringContent) ReadAt(p []byte, offset int64) (int, error) {
	if receiver.closed {
		return 0, errClosed
	}

	if 0 == len(p) && 0 == offset {
		return 0, nil
	}

	if offset < 0 {
		return 0, errNegativeOffset
	}

	if int64(len(receiver.value)) <= offset {
		return 0, io.EOF
	}

	n := copy(p, receiver.value[offset:])

	switch {
	case n < len(p):
		return n, io.EOF
	default:
		return n, nil
	}
}
