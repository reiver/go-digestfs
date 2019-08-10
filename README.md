# go-digestfs

Package **digestfs provides** a **content-addressable file system**, and **virtual file system** (**VFS**) 
by providing a common interface to one or more **content-addressable storage** (**CAS**).

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-digestfs

[![GoDoc](https://godoc.org/github.com/reiver/go-digestfs?status.svg)](https://godoc.org/github.com/reiver/go-digestfs)

## Example

Here is an example of how to use an already mounted `digestfs.MountPoint` to get content:
```go
var mountpoint digestfs.MountPoint

// ...

// c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a
var key [64]byte = [64]byte{0xc0, 0x53, 0x5e, 0x4b, 0xe2, 0xb7, 0x9f, 0xfd, 0x93, 0x29, 0x13, 0x05, 0x43, 0x6b, 0xf8, 0x89, 0x31, 0x4e, 0x4a, 0x3f, 0xae, 0xc0, 0x5e, 0xcf, 0xfc, 0xbb, 0x7d, 0xf3, 0x1a, 0xd9, 0xe5, 0x1a}

content, err := mountpoint.Open("SHA-256", key[:])
if nil != err {
	return err
}
defer content.Close()
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
