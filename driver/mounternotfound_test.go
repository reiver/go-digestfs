package digestfs_driver

import (
	"testing"
)

func TestInternalMounterNotFoundAsError(t *testing.T) {

	var err error = internalMounterNotFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalMounterNotFoundAsMounterNotFound(t *testing.T) {

	var complainer MounterNotFound = internalMounterNotFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
