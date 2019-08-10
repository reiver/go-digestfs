package digestfs_driver

type MountPoint interface {
	Open(algorithm string, digest []byte) (Content, error)
	OpenLocation(location string) (Content, error)
	Unmount() error
}
