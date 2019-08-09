package digestfs_driver

type MounterFunc func(...interface{})(MountPoint, error)

func (fn MounterFunc) Mount(params ...interface{}) (MountPoint, error) {
	return fn(params...)
}
