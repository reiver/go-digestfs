package digestfs_driver

type NotMounted interface {
	error
	NotMounted()
}

func ErrNotMounted() error {
	var e NotMounted = &internalNotMounted{}

	return e
}

type internalNotMounted struct{}

func (receiver internalNotMounted) Error() string {
	return "digestfs: Not Mounted"
}

func (receiver internalNotMounted) NotMounted() {
	// Nothing here
}
