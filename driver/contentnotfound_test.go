package digestfs_driver

import (
	"testing"
)

func TestInternalContentNotFoundAsError(t *testing.T) {

	var err error = internalContentNotFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalContentNotFoundAsContentNotFound(t *testing.T) {

	var complainer ContentNotFound = internalContentNotFound{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
