# go-digestfs

Package **digestfs provides** a **content-addressable virtual file system** (**VFS**) 
by providing a common interface to one or more **content-addressable storage** (**CAS**).

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-digestfs

[![GoDoc](https://godoc.org/github.com/reiver/go-digestfs?status.svg)](https://godoc.org/github.com/reiver/go-digestfs)

## Example

Here is an example of how to use an already mounted `digestfs.MountPoint` to get content:
```go
content, err := mountpoint.Open("SHA-256.hexadecimal", "c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a")
if nil != err {
	return err
}
defer content.Close()
```

Note that if you had the digest encoded as binary, then you could encode it as hexadecimal using:
As in:
```go
digest := fmt.Sprintf("%x", binaryDigest)
```

## Content Addressing

A **content-addressable file system**, or **content-addressable storage** (**CAS**) may be used with **content addressing**.

Some examples of **content addressing** include:

* [Magnet URI Project](http://magnet-uri.sourceforge.net/)

* [MAGNET v0.1 (Draft Tech Overview/Spec)](http://magnet-uri.sourceforge.net/magnet-draft-overview.txt),
  by Gordon Mohr

* [Magnet URI scheme (Wikipedia)](https://en.wikipedia.org/wiki/Magnet_URI_scheme),

* [A URN Namespace For Identifiers Based on Cryptographic Hashes](https://tools.ietf.org/html/draft-thiemann-hash-urn-01)
  by Peter Thiemann

* [IETF RFC 6920: Naming Things with Hashes](https://tools.ietf.org/search/rfc6920)

* [Trusty URIs: Verifiable, Immutable, and Permanent Digital Artifacts for Linked Data](https://arxiv.org/abs/1401.5775),
  by Tobias Kuhn, Michel Dumontier

* [BIP 122](https://github.com/bitcoin/bips/blob/master/bip-0122.mediawiki)

* [HTTP Extensions for a Content-Addressable Web](http://lists.w3.org/Archives/Public/www-talk/2001NovDec/0090.html)

* [IETF RFC 2397: The "data" URL scheme](https://tools.ietf.org/html/rfc2397)

* [Hash URI Specification](https://github.com/hash-uri/hash-uri)

## See Also

* https://godoc.org/github.com/reiver/go-memdigest
