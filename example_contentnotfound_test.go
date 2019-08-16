package digestfs_test

import (
	"github.com/reiver/go-digestfs"
	"github.com/reiver/go-memdigest"

	"encoding/hex"
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
		// 0xd3486ae9136e7856bc42212385ea797094475802
		var digest string = "\xd3\x48\x6a\xe9\x13\x6e\x78\x56\xbc\x42\x21\x23\x85\xea\x79\x70\x94\x47\x58\x02"

		content, err := mountpoint.Open("SHA-1", digest)
		if nil != err {
			switch err.(type) {
			case digestfs.ContentNotFound: // <--------- This is where digestfs.ContentNotFound is being used
				fmt.Printf("ERROR: could not find content with SHA-1 digest “%s”\n", hex.EncodeToString([]byte(digest)))
			default:
				fmt.Printf("ERROR: could not open content with SHA-1 digest “%s” because: %s\n", hex.EncodeToString([]byte(digest)), err)
			}
			return
		}
		defer content.Close()

		r := io.NewSectionReader(content, 0, int64(content.Len()))
		contentBytes, err := ioutil.ReadAll(r)
		if nil != err {
			fmt.Printf("ERROR: could not read all bytes of content with SHA-1 digest “%s” because: %s\n", hex.EncodeToString([]byte(digest)), err)
			return
		}

		fmt.Printf("Content with SHA-1 digest “%s” is: %q\n", hex.EncodeToString([]byte(digest)), contentBytes)
	}()

	func(){
		// 0x80256f39a9d308650ac90d9be9a72a9562454574
		var digest string = "\x80\x25\x6f\x39\xa9\xd3\x08\x65\x0a\xc9\x0d\x9b\xe9\xa7\x2a\x95\x62\x45\x45\x74"

		content, err := mountpoint.Open("SHA-1", digest)
		if nil != err {
			switch err.(type) {
			case digestfs.ContentNotFound: // <--------- This is where digestfs.ContentNotFound is being used
				fmt.Printf("ERROR: could not find content with SHA-1 digest “%s”\n", hex.EncodeToString([]byte(digest)))
			default:
				fmt.Printf("ERROR: could not open content with SHA-1 digest “%s” because: %s\n", hex.EncodeToString([]byte(digest)), err)
			}
			return
		}
		defer content.Close()

		r := io.NewSectionReader(content, 0, int64(content.Len()))
		contentBytes, err := ioutil.ReadAll(r)
		if nil != err {
			fmt.Printf("ERROR: could not read all bytes of content with SHA-1 digest “%s” because: %s\n", hex.EncodeToString([]byte(digest)), err)
			return
		}

		fmt.Printf("Content with SHA-1 digest “%s” is: %q\n", hex.EncodeToString([]byte(digest)), contentBytes)
	}()


	// Output:
	// Content with SHA-1 digest “d3486ae9136e7856bc42212385ea797094475802” is: "Hello world!"
	// ERROR: could not find content with SHA-1 digest “80256f39a9d308650ac90d9be9a72a9562454574”
}
