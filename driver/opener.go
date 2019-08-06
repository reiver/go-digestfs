package digestfs_driver

type Opener interface {
	Open(algorithm string, digest string) (Content, error)
}
