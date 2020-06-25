package factorymethod

type gun struct {
	name  string
	power int
}

func (c *gun) setName(name string) {
	c.name = name
}

func (c *gun) getName() string {
	return c.name
}

func (c *gun) setPower(power int) {
	c.power = power
}

func (c *gun) getPower() int {
	return c.power
}
