package command

type device interface {
	on()
	off()
}
