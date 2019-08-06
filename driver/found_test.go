package digestfs_driver

import (
	"testing"
)

func TestInternalFoundAsError(t *testing.T) {

	var err error = internalFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalFoundAsFound(t *testing.T) {

	var notfound Found = internalFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == notfound {
		t.Error("This should never happen.")
	}
}
