package digestfs_driver

type Mounter interface {
	Mount(params ...interface{}) (MountPoint, error)
}
