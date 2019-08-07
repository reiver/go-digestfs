package digestfs_driver

type MounterFunc func(string, ...interface{})(Opener, error)

func (fn MounterFunc) Mount(src string, params ...interface{}) (Opener, error) {
	return fn(src, params...)
}
