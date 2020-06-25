package command

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}
