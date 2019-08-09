package digestfs

type UnsupportedAlgorithm interface {
	error
	UnsupportedAlgorithm() string
}
