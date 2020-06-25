package abstractfactory

type nike struct{}

type nikeShoe struct {
	shoeItem
}

func (n nikeShoe) getLogo() string {
	return n.shoeItem.logo
}

type nikeShort struct {
	shortItem
}

func (n nikeShort) getLogo() string {
	return n.shortItem.logo
}

func (c *nike) makeShoe() shoe {
	return &nikeShoe{
		shoeItem: shoeItem{
			logo: "nike",
			size: 40,
		},
	}
}

func (c *nike) makeShort() short {
	return &nikeShort{
		shortItem: shortItem{
			logo: "nike",
			size: 170,
		},
	}
}
