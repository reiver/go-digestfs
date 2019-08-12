package digestfs_driver

type MountPoint interface {
	Create(content []byte) (algorithm string, digest string, err error)
	Open(algorithm string, digest string) (Content, error)
	OpenLocation(location string) (Content, error)
	Unmount() error
}
