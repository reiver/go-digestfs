package digestfs_test

import (
	"github.com/reiver/go-digestfs"

	"fmt"
)

func ExampleMounterNotFound() {

	var mountpoint digestfs.MountPoint

	err := mountpoint.Mount("faux", "/path/to/stuff")

	if nil != err {
		switch casted := err.(type) {
		case digestfs.MounterNotFound:
			fmt.Printf("mounter not found for fstype = %q", casted.MounterNotFound())
			return
		default:
			fmt.Print(err)
			return
		}
	}

	fmt.Printf("MOUNTED!")

	// Output:
	// mounter not found for fstype = "faux"
}
