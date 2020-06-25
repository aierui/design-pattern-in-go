package abstractfactory

type adidas struct{}

type adidasShoe struct {
	shoeItem
}

func (a adidasShoe) getLogo() string {
	return a.shoeItem.logo
}

type adidasShort struct {
	shortItem
}

func (a adidasShort) getLogo() string {
	return a.shortItem.logo
}

func (c *adidas) makeShoe() shoe {
	return &adidasShoe{
		shoeItem: shoeItem{
			logo: "adidas",
			size: 40,
		},
	}
}

func (c *adidas) makeShort() short {
	return &adidasShort{
		shortItem: shortItem{
			logo: "adidas",
			size: 170,
		},
	}
}
