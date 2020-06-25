package abstractfactory

type shoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shoeItem struct {
	logo string
	size int
}

func (c *shoeItem) setLogo(logo string) {
	c.logo = logo
}

func (c *shoeItem) getLost() string {
	return c.logo
}

func (c *shoeItem) setSize(size int) {
	c.size = size
}

func (c *shoeItem) getSize() int {
	return c.size
}
