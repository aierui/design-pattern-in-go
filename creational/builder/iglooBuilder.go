package builder

type iglooBuilder struct {
	widowType string
	doorType  string
	floor     int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (c *iglooBuilder) setWindowType() {
	c.widowType = "Wooden Window"
}

func (c *iglooBuilder) setDoorType() {
	c.doorType = "Wooden Door"
}

func (c *iglooBuilder) setNumFloor() {
	c.floor = 2
}

func (c *iglooBuilder) getHouse() house {
	return house{
		doorType:   c.doorType,
		windowType: c.widowType,
		floor:      c.floor,
	}
}
