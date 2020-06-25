package builder

type normalBuilder struct {
	widowType string
	doorType  string
	floor     int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (c *normalBuilder) setWindowType() {
	c.widowType = "Wooden Window"
}

func (c *normalBuilder) setDoorType() {
	c.doorType = "Wooden Door"
}

func (c *normalBuilder) setNumFloor() {
	c.floor = 2
}

func (c *normalBuilder) getHouse() house {
	return house{
		doorType:   c.doorType,
		windowType: c.widowType,
		floor:      c.floor,
	}
}
