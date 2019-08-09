package digestfs_driver

import (
	"testing"
)

func TestInternalUnsupportedAlgorithmAsError(t *testing.T) {

	var err error = internalUnsupportedAlgorithm{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalUnsupportedAlgorithmAsUnsupportedAlgorithm(t *testing.T) {

	var complainer UnsupportedAlgorithm = internalUnsupportedAlgorithm{} // THIS IS WHAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
