package digestfs_driver

import (
	"testing"
)

func TestInternalNotMountedAsError(t *testing.T) {

	var err error = internalNotMounted{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalNotMountedAsNotMounted(t *testing.T) {

	var complainer NotMounted = internalNotMounted{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
