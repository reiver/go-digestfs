package digestfs_driver

type MounterFunc func(string, ...interface{})(MountPoint, error)

func (fn MounterFunc) Mount(src string, params ...interface{}) (MountPoint, error) {
	return fn(src, params...)
}
