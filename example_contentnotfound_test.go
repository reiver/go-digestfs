package digestfs_test

import (
	"github.com/reiver/go-digestfs"
	"github.com/reiver/go-memdigest"

	"fmt"
	"io"
	"io/ioutil"
)

func ExampleContentNotFound() {

	var mem memdigest.SHA1

	var mountpoint digestfs.MountPoint

	err := mountpoint.Mount("memdigest.SHA1", &mem)

	if nil != err {
		switch casted := err.(type) {
		case digestfs.MounterNotFound:
			fmt.Printf("mounter not found for fstype = %q\n", casted.MounterNotFound())
			return
		default:
			fmt.Println(err)
			return
		}
	}


	_, _, err = mountpoint.Create([]byte("Hello world!"))
	if nil != err {
		fmt.Printf("ERROR: could not create content “Hello world!” because: %s\n", err)
		return
	}
	_, _, err = mountpoint.Create([]byte("apple banana cherry"))
	if nil != err {
		fmt.Printf("ERROR: could not create content “apple banana cherry” because: %s\n", err)
		return
	}
	_, _, err = mountpoint.Create([]byte("intergalactic planetary 👾"))
	if nil != err {
		fmt.Printf("ERROR: could not create content “intergalactic planetary 👾” because: %s\n", err)
		return
	}


	func(){
		content, err := mountpoint.Open("SHA-1.hexadecimal", "d3486ae9136e7856bc42212385ea797094475802")
		if nil != err {
			switch err.(type) {
			case digestfs.ContentNotFound: // <--------- This is where digestfs.ContentNotFound is being used
				fmt.Print("ERROR: could not find content with SHA-1 digest “d3486ae9136e7856bc42212385ea797094475802”\n")
			default:
				fmt.Printf("ERROR: could not open content with SHA-1 digest “d3486ae9136e7856bc42212385ea797094475802” because: %s\n", err)
			}
			return
		}
		defer content.Close()

		r := io.NewSectionReader(content, 0, int64(content.Len()))
		contentBytes, err := ioutil.ReadAll(r)
		if nil != err {
			fmt.Printf("ERROR: could not read all bytes of content with SHA-1 digest “d3486ae9136e7856bc42212385ea797094475802” because: %s\n", err)
			return
		}

		fmt.Printf("Content with SHA-1 digest “d3486ae9136e7856bc42212385ea797094475802” is: %q\n", contentBytes)
	}()


	func(){
		content, err := mountpoint.Open("SHA-1.hexadecimal", "80256f39a9d308650ac90d9be9a72a9562454574")
		if nil != err {
			switch err.(type) {
			case digestfs.ContentNotFound: // <--------- This is where digestfs.ContentNotFound is being used
				fmt.Print("ERROR: could not find content with SHA-1 digest “80256f39a9d308650ac90d9be9a72a9562454574”\n")
			default:
				fmt.Printf("ERROR: could not open content with SHA-1 digest “80256f39a9d308650ac90d9be9a72a9562454574” because: %s\n", err)
			}
			return
		}
		defer content.Close()

		r := io.NewSectionReader(content, 0, int64(content.Len()))
		contentBytes, err := ioutil.ReadAll(r)
		if nil != err {
			fmt.Printf("ERROR: could not read all bytes of content with SHA-1 digest “80256f39a9d308650ac90d9be9a72a9562454574” because: %s\n", err)
			return
		}

		fmt.Printf("Content with SHA-1 digest “80256f39a9d308650ac90d9be9a72a9562454574” is: %q\n", contentBytes)
	}()


	// Output:
	// Content with SHA-1 digest “d3486ae9136e7856bc42212385ea797094475802” is: "Hello world!"
	// ERROR: could not find content with SHA-1 digest “80256f39a9d308650ac90d9be9a72a9562454574”
}
