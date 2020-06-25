package builder

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	} else if builderType == "igloo" {
		return &iglooBuilder{}
	}

	return nil
}
