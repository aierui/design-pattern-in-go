package builder

type director struct {
	builder iBuilder
}

func newDirector(c iBuilder) *director {
	return &director{
		builder: c,
	}
}

func (c *director) setBuilder(b iBuilder) {
	c.builder = b
}

func (c *director) buildHouse() house {
	c.builder.setDoorType()
	c.builder.setWindowType()
	c.builder.setNumFloor()

	return c.builder.getHouse()
}
