package digestfs_driver

import (
	"testing"
)

func TestInternalBadLocationAsError(t *testing.T) {

	var err error = internalBadLocation{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalBadLocationAsBadLocation(t *testing.T) {

	var complainer BadLocation = internalBadLocation{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
