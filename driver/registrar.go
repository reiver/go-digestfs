package digestfs_driver

type Registrar interface {
	Fetch(string) (Mounter, error)
	Register(Mounter, string) error
}
