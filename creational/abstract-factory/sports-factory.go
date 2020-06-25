package abstractfactory

import (
	"errors"
)

type sportFactory interface {
	makeShoe() shoe
	makeShort() short
}

func getSportFactory(brand string) (sportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	} else if brand == "nike" {
		return &nike{}, nil
	}

	return nil, errors.New("Wrong brand type passed")
}
