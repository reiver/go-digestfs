package digestfs_driver

var (
	Registry Registrar
)

func init() {
	Registry = new(internalRegistrar)
}
