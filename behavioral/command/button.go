package command

type button struct {
	command command
}

func (c *button) press() {
	c.command.execute()
}
