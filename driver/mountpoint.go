package digestfs_driver

type MountPoint interface {

	// Create creates new content at the MoutPoint.
	//
	// The value of ‘digest’ is in binary.
	//
	// The value of ‘digest’ is NOT a binary-to-text encoding such as hexadecimal, bas64, etc.
	//
	// So the value of ‘digest’ might be something such as:
	//
	//	"\x00\x00\x00\x00\x00\x19\xd6\x68\x9c\x08\x5a\xe1\x65\x83\x1e\x93\x4f\xf7\x63\xae\x46\xa2\xa6\xc1\x72\xb3\xf1\xb6\x0a\x8c\xe2\x6f"
	//
	// Rather than any of these:
	//
	//	"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
	//
	//	"0x000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
	//
	//	"AAAAAAAZ1micCFrhZYMek0_3Y65GoqbBcrPxtgqM4m8" (not this)
	//
	//	"00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
	//
	//	"0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
	Create(content []byte) (algorithm string, digest string, err error)

	// Open returns content at the MountPoint.
	//
	// The value of ‘digest’ is in binary.
	//
	// The value of ‘digest’ is NOT a binary-to-text encoding such as hexadecimal, bas64, etc.
	//
	// So the value of ‘digest’ might be something such as:
	//
	//	"\x00\x00\x00\x00\x00\x19\xd6\x68\x9c\x08\x5a\xe1\x65\x83\x1e\x93\x4f\xf7\x63\xae\x46\xa2\xa6\xc1\x72\xb3\xf1\xb6\x0a\x8c\xe2\x6f"
	//
	// Rather than any of these:
	//
	//	"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
	//
	//	"0x000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" (not this)
	//
	//	"AAAAAAAZ1micCFrhZYMek0_3Y65GoqbBcrPxtgqM4m8" (not this)
	//
	//	"00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
	//
	//	"0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf" (not this)
	Open(algorithm string, digest string) (Content, error)

	// OpenLocation returns content based on a location, rather than based on a digest.
	//
	// The value of ‘location’ can be in a format that is specific to a MountPoint,
	// and does NOT have to be in a universally supported format.
	OpenLocation(location string) (Content, error)

	// Unmount should make the MountPoint stop working.
	//
	// This method can be used for any cleanup that needs to be done.
	Unmount() error
}
