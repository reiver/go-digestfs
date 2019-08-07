package digestfs_driver_test

import (
	"github.com/reiver/go-digestfs/driver"

	"testing"
)

func TestMounterFuncAsMounter(t *testing.T) {

	var mounter digestfs_driver.Mounter = digestfs_driver.MounterFunc(nil) // THIS IS WHAT ACTUALLY MATTERS!

	if nil == mounter {
		t.Error("This should never happen.")
	}
}
