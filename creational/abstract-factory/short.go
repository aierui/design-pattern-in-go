package abstractfactory

type short interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shortItem struct {
	logo string
	size int
}

func (c *shortItem) setLogo(logo string) {
	c.logo = logo
}

func (c *shortItem) getLost() string {
	return c.logo
}

func (c *shortItem) setSize(size int) {
	c.size = size
}

func (c *shortItem) getSize() int {
	return c.size
}
