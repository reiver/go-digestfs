package digestfs_test

import (
	"github.com/reiver/go-digestfs"

	"fmt"
)

func ExampleMounterNotFound() {

	var mountpoint digestfs.MountPoint

	err := mountpoint.Mount("faux", "/path/to/stuff")

	if nil != err {
		fmt.Print(err)
		return
	}

	fmt.Printf("MOUNTED!")

	// Output:
	// digestfs: Not Found: mounter name="faux"
}
