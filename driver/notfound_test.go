package digestfs_driver

import (
	"testing"
)

func TestInternalNotFoundAsError(t *testing.T) {

	var err error = internalNotFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalNotFoundAsNotFound(t *testing.T) {

	var notfound NotFound = internalNotFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == notfound {
		t.Error("This should never happen.")
	}
}
