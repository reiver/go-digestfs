package digestfs

type NotMounted interface {
	error
	NotMounted()
}
