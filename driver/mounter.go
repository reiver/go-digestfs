package digestfs_driver

type Mounter interface {
	Mount(src string, params ...interface{}) error
}
