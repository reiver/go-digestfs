package digestfs_driver

import (
	"testing"
)

func TestInternalMounterFoundAsError(t *testing.T) {

	var err error = internalMounterFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalMounterFoundAsMounterFound(t *testing.T) {

	var complainer MounterFound = internalMounterFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
