package digestfs_driver

type MounterFunc func(string, ...interface{})error

func (fn MounterFunc) Mount(src string, params ...interface{}) error {
	return fn(src, params...)
}
