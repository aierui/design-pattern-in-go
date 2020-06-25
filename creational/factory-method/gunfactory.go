package factorymethod

import (
	"errors"
)

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	} else if gunType == "maverick" {
		return newMaverick(), nil
	}

	return nil, errors.New("Wrong gun type passed")
}
