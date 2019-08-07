package digestfs_driver

type MountPoint interface {
	Open(algorithm string, digest string) (Content, error)
	Unmount() error
}
